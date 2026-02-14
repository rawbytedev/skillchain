package main

import (
	"bytes"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"moltket/config"
	"moltket/internal/api"
	"moltket/internal/auth"
	"moltket/internal/blockchain"
	"moltket/internal/cache"
	"moltket/internal/core"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockBlockchain implements the blockchain.Interface for testing
type mockBlockchainClient struct{}

func (m *mockBlockchainClient) IsLicenseValid(user common.Address, toolID *big.Int) (bool, error) {
	// For integration test: licenses are only valid once recorded/minted on-chain
	// This mock doesn't have minting tracking, so we just return false
	// (In real scenario, this would check on-chain contract state)
	return false, nil
}

func (m *mockBlockchainClient) GetLicenseMetadata(toolIDStr string) (*blockchain.LicenseMetadata, error) {
	return &blockchain.LicenseMetadata{
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}, nil
}

func TestFullIntegration(t *testing.T) {
    // Setup test environment
    cfg := &config.Config{
        ServerPort:        "8081",
        Env:               "test",
        LicenseNFTAddress: "0x1234567890123456789012345678901234567890",
        JWTSecret:         "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", // Hardhat #0 private key for testing
        CacheTTL:          int(5 * time.Minute),
        ChainID:           31337, // Hardhat chain ID
        RateLimit:         10000, // High rate limit for testing
        SignatureNonce:    "test-nonce",
    }
    
    // Create real instances
    kvStore := cache.NewKVStore()
    defer kvStore.Close()
    
    // Create signer for services
    chainID := big.NewInt(cfg.ChainID)
    verifyingContract := common.HexToAddress(cfg.LicenseNFTAddress)
    signer, err := auth.NewSigner(cfg.JWTSecret, chainID, verifyingContract)
    require.NoError(t, err)
    
    // Create vote service for server
    voteService := core.NewVoteService(cfg, kvStore, signer)
    
    // Create server with all components
    server := api.NewServer(cfg, kvStore, &mockBlockchainClient{}, voteService)
    
    // Start test server
    testServer := httptest.NewServer(server.Handler())
    defer testServer.Close()
    
    // Test 1: Request license
    t.Run("EndToEndLicenseFlow", func(t *testing.T) {
        client := testServer.Client()
        
        // Step 1: Request license
        licenseReq := map[string]string{
            "user_address": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", // Hardhat #1
            "tool_id":      "42",
        }
        
        body, _ := json.Marshal(licenseReq)
        resp, err := client.Post(testServer.URL+"/api/v1/license/request", "application/json", bytes.NewReader(body))
        require.NoError(t, err)
        assert.Equal(t, http.StatusOK, resp.StatusCode)
        
        var licenseResp map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&licenseResp)
        resp.Body.Close()
        
        assert.True(t, licenseResp["success"].(bool))
        data := licenseResp["data"].(map[string]interface{})
        assert.Equal(t, "42", data["tool_id"])
        
        // Step 2: Verify access (should fail before license is recorded as minted)
        verifyReq := map[string]string{
            "user_address": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
            "tool_id":      "42",
        }
        
        body, _ = json.Marshal(verifyReq)
        resp, err = client.Post(testServer.URL+"/api/v1/access/verify", "application/json", bytes.NewReader(body))
        require.NoError(t, err)
        
        // First 100 calls should be free tier
        if resp.StatusCode == http.StatusOK {
            var verifyResp map[string]interface{}
            json.NewDecoder(resp.Body).Decode(&verifyResp)
            resp.Body.Close()
            
            assert.Equal(t, "free", verifyResp["tier"])
        }
        
        // Step 3: Record license minted (simulating blockchain event)
        recordReq := map[string]string{
            "user_address": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
            "tool_id":      "42",
            "expires_at":   data["expires_at"].(string),
            "nonce":        data["nonce"].(string),
        }
        
        body, _ = json.Marshal(recordReq)
        resp, err = client.Post(testServer.URL+"/api/v1/license/record-minted", "application/json", bytes.NewReader(body))
        require.NoError(t, err)
        assert.Equal(t, http.StatusOK, resp.StatusCode)
        
        // Step 4: Verify access again (should now show licensed tier)
        body, _ = json.Marshal(verifyReq)
        resp, err = client.Post(testServer.URL+"/api/v1/access/verify", "application/json", bytes.NewReader(body))
        require.NoError(t, err)
        assert.Equal(t, http.StatusOK, resp.StatusCode)
        
        var verifyResp map[string]interface{}
        json.NewDecoder(resp.Body).Decode(&verifyResp)
        resp.Body.Close()
        
        // Should be licensed tier now
        assert.True(t, verifyResp["valid"].(bool))
        assert.Equal(t, "licensed", verifyResp["tier"])
    })
    
    // Test 2: Free tier exhaustion
    t.Run("FreeTierExhaustion", func(t *testing.T) {
        client := testServer.Client()
        
        // Make 100 free calls
        for i := 0; i < 100; i++ {
            req := map[string]string{
                "user_address": "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC", // Hardhat #2
                "tool_id":      "99",
            }
            
            body, _ := json.Marshal(req)
            resp, err := client.Post(testServer.URL+"/api/v1/access/verify", "application/json", bytes.NewReader(body))
            require.NoError(t, err)
            
            if i < 100 {
                assert.Equal(t, http.StatusOK, resp.StatusCode)
            }
            resp.Body.Close()
        }
        
        // 101st call should fail
        req := map[string]string{
            "user_address": "0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC",
            "tool_id":      "99",
        }
        
        body, _ := json.Marshal(req)
        resp, err := client.Post(testServer.URL+"/api/v1/access/verify", "application/json", bytes.NewReader(body))
        require.NoError(t, err)
        assert.Equal(t, http.StatusForbidden, resp.StatusCode)
    })
}