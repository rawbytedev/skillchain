package blockchain

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	"moltket/config"
	"moltket/internal/auth"
	"moltket/internal/cache"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockBlockchain struct {
    isValid        bool
    metadata       *LicenseMetadata
    shouldError    bool
}

func (m *mockBlockchain) IsLicenseValid(user common.Address, toolID *big.Int) (bool, error) {
    if m.shouldError {
        return false, assert.AnError
    }
    return m.isValid, nil
}

func (m *mockBlockchain) GetLicenseMetadata(toolIDStr string) (*LicenseMetadata, error) {
    if m.shouldError {
        return nil, assert.AnError
    }
    return m.metadata, nil
}

func TestLicenseService_Integration(t *testing.T) {
    ctx := context.Background()
    
    // Setup
    cfg := &config.Config{
        LicenseNFTAddress: "0x1234567890123456789012345678901234567890",
        SignatureNonce:    "test-nonce",
    }
    
    kvStore := cache.NewKVStore()
    defer kvStore.Close()
    
    // Create test signer
    privateKey, err := auth.GenerateTestPrivateKey(t, "test-seed")
    require.NoError(t, err)
    
    privateKeyBytes := crypto.FromECDSA(privateKey)
    privateKeyHex := hex.EncodeToString(privateKeyBytes)
    
    chainID := big.NewInt(1337)
    verifyingContract := common.HexToAddress(cfg.LicenseNFTAddress)
    
    signer, err := auth.NewSigner(privateKeyHex, chainID, verifyingContract)
    require.NoError(t, err)
    
    // Test 1: Free tier access
    t.Run("FreeTierAccess", func(t *testing.T) {
        bc := &mockBlockchain{isValid: false}
        service := NewLicenseService(cfg, kvStore, signer, bc)
        
        user := common.HexToAddress("0xUser1")
        toolID := big.NewInt(1)
        
        // First 100 calls should succeed
        for i := 0; i < 100; i++ {
            result, err := service.VerifyAccess(ctx, user, toolID)
            require.NoError(t, err)
            assert.True(t, result.Valid)
            assert.Equal(t, "free", result.Tier)
            assert.Equal(t, 100-(i+1), result.CallsRemaining)
            assert.NotEmpty(t, result.ProvenanceHash)
        }
        
        // 101st call should fail (free tier exhausted, no license)
        result, err := service.VerifyAccess(ctx, user, toolID)
        require.NoError(t, err)
        assert.False(t, result.Valid)
        assert.Contains(t, result.Reason, "free tier exhausted")
    })
    
    // Test 2: License request and verification
    t.Run("LicenseRequestFlow", func(t *testing.T) {
        bc := &mockBlockchain{
            isValid: true,
            metadata: &LicenseMetadata{
                ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
            },
        }
        service := NewLicenseService(cfg, kvStore, signer, bc)
        
        user := common.HexToAddress("0xUser2")
        toolID := big.NewInt(2)
        
        // Request license
        req := &LicenseRequest{
            UserAddress: user,
            ToolID:      toolID,
        }
        
        resp, err := service.RequestLicense(ctx, req)
        require.NoError(t, err)
        
        assert.Equal(t, toolID, resp.ToolID)
        assert.Equal(t, user, resp.User)
        assert.Equal(t, cfg.LicenseNFTAddress, resp.Contract)
        assert.NotEmpty(t, resp.SignatureR)
        assert.NotEmpty(t, resp.SignatureS)
        assert.NotEmpty(t, resp.SignatureV)
        
        // Simulate license minted (would be called by blockchain event listener)
        err = service.RecordLicenseMinted(ctx, user, toolID, resp.ExpiresAt, resp.Nonce)
        require.NoError(t, err)
        
        // Now verify access with licensed tier
        result, err := service.VerifyAccess(ctx, user, toolID)
        require.NoError(t, err)
        assert.True(t, result.Valid)
        assert.Equal(t, "licensed", result.Tier)
        assert.NotNil(t, result.ExpiresAt)
    })
    
    // Test 3: Prevent duplicate license requests
    t.Run("DuplicateLicenseRequest", func(t *testing.T) {
        bc := &mockBlockchain{isValid: false}
        service := NewLicenseService(cfg, kvStore, signer, bc)
        
        user := common.HexToAddress("0xUser3")
        toolID := big.NewInt(3)
        
        // First request should succeed
        req := &LicenseRequest{UserAddress: user, ToolID: toolID}
        resp1, err := service.RequestLicense(ctx, req)
        require.NoError(t, err)
        
        // Second request while pending should fail
        _, err = service.RequestLicense(ctx, req)
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "license request already pending")
        
        // Record minting to clear pending
        err = service.RecordLicenseMinted(ctx, user, toolID, resp1.ExpiresAt, resp1.Nonce)
        require.NoError(t, err)
        
        // Third request after minting should fail (license already active)
        _, err = service.RequestLicense(ctx, req)
        assert.Error(t, err)
        assert.Contains(t, err.Error(), "license already active")
    })
    
    // Test 4: Blockchain error handling
    t.Run("BlockchainError", func(t *testing.T) {
        bc := &mockBlockchain{shouldError: true}
        service := NewLicenseService(cfg, kvStore, signer, bc)
        
        user := common.HexToAddress("0xUser4")
        toolID := big.NewInt(4)
        
        // Exhaust free tier first
        for i := 0; i < 100; i++ {
            service.VerifyAccess(ctx, user, toolID)
        }
        
        // Next call should handle free tier exhausted with no blockchain license
        result, err := service.VerifyAccess(ctx, user, toolID)
        require.NoError(t, err)
        assert.False(t, result.Valid)
        assert.Contains(t, result.Reason, "free tier exhausted")
    })
}