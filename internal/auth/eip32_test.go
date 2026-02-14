package auth

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEIP712Signer_Integration(t *testing.T) {
    // Generate a test private key
    privateKey, err := crypto.GenerateKey()
    require.NoError(t, err, "Should generate private key")
    
    privateKeyBytes := crypto.FromECDSA(privateKey)
    privateKeyHex := hex.EncodeToString(privateKeyBytes)
    
    // Test parameters
    chainID := big.NewInt(11155111) // Sepolia
    verifyingContract := common.HexToAddress("0x1234567890123456789012345678901234567890")
    user := common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdefabcd")
    toolId := big.NewInt(42)
    expiresAt := big.NewInt(1893456000) // 2030-01-01
    nonce := big.NewInt(123456)
    
    // Create signer
    signer, err := NewSigner(privateKeyHex, chainID, verifyingContract)
    require.NoError(t, err, "Should create signer")
    
    // Test 1: Create signature
    r, s, v, err := signer.CreateLicenseSignature(user, toolId, expiresAt, nonce)
    require.NoError(t, err, "Should create signature")
    
    assert.Len(t, r, 32, "r should be 32 bytes")
    assert.Len(t, s, 32, "s should be 32 bytes")
    assert.Len(t, v, 1, "v should be 1 byte")
    assert.GreaterOrEqual(t, v[0], uint8(27), "v should be 27 or 28")
    assert.LessOrEqual(t, v[0], uint8(28), "v should be 27 or 28")
    
    // Test 2: Verify signature
    isValid, err := signer.VerifySignature(user, toolId, expiresAt, nonce, r, s, v)
    require.NoError(t, err, "Should verify signature")
    assert.True(t, isValid, "Signature should be valid")
    
    // Test 3: Verify with wrong parameters should fail
    wrongToolId := big.NewInt(43)
    isValid, err = signer.VerifySignature(user, wrongToolId, expiresAt, nonce, r, s, v)
    require.NoError(t, err, "Should verify even with wrong params")
    assert.False(t, isValid, "Signature should be invalid with wrong toolId")
    
    // Test 4: Different signer should produce different signature
    anotherPrivateKey, err := crypto.GenerateKey()
    require.NoError(t, err)
    
    anotherPrivateKeyBytes := crypto.FromECDSA(anotherPrivateKey)
    anotherPrivateKeyHex := hex.EncodeToString(anotherPrivateKeyBytes)
    
    anotherSigner, err := NewSigner(anotherPrivateKeyHex, chainID, verifyingContract)
    require.NoError(t, err)
    
    r2, s2, v2, err := anotherSigner.CreateLicenseSignature(user, toolId, expiresAt, nonce)
    require.NoError(t, err)
    
    // Signatures should be different
    assert.NotEqual(t, r, r2, "Different signers should produce different r values")
    assert.NotEqual(t, s, s2, "Different signers should produce different s values")
    assert.NotEqual(t, v, v2, "Different signers should produce different v values")
    
    // Test 5: Address should match
    expectedAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
    assert.Equal(t, expectedAddress, signer.Address(), "Signer address should match private key")
}

func TestEIP712Signer_InvalidPrivateKey(t *testing.T) {
    chainID := big.NewInt(1)
    verifyingContract := common.HexToAddress("0x1234567890123456789012345678901234567890")
    
    // Test with invalid hex
    _, err := NewSigner("not-a-valid-hex", chainID, verifyingContract)
    assert.Error(t, err, "Should error with invalid hex")
    
    // Test with valid hex but wrong length
    _, err = NewSigner("abcd", chainID, verifyingContract)
    assert.Error(t, err, "Should error with wrong length")
}