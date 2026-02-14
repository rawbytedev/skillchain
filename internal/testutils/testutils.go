package testutils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// GenerateTestLicenseID creates a deterministic test license ID
func GenerateTestLicenseID(t *testing.T, index int) string {
	t.Helper()
	data := []byte{byte(index), byte(index >> 8), byte(index >> 16)}
	hash := crypto.Keccak256Hash(data)
	return hexutil.Encode(hash[:])[2:] // Remove 0x prefix
}

// GenerateTestAddress creates a deterministic test Ethereum address
func GenerateTestAddress(t *testing.T, index int) common.Address {
	t.Helper()
	data := []byte{byte(index), byte(index >> 8), byte(index >> 16)}
	hash := crypto.Keccak256Hash(data)
	return common.BytesToAddress(hash[12:])
}

// GenerateTestSignature creates a valid EIP-712 signature for testing
func GenerateTestSignature(t *testing.T, privateKey *ecdsa.PrivateKey, message []byte) (hexutil.Bytes, error) {
	t.Helper()

	hash := accounts.TextHash(message)
	signature, err := crypto.Sign(hash, privateKey)
	if err != nil {
		return nil, err
	}

	// Ethereum requires v to be 27 or 28
	signature[64] += 27

	return signature, nil
}

// GenerateTestPrivateKey creates a deterministic test private key
func GenerateTestPrivateKey(t *testing.T, seed string) *ecdsa.PrivateKey {
	t.Helper()

	hash := crypto.Keccak256Hash([]byte(seed))
	privateKey, err := crypto.ToECDSA(hash[:])
	require.NoError(t, err)

	return privateKey
}

// RandomHex generates a random hex string of specified length
func RandomHex(t *testing.T, n int) string {
	t.Helper()

	bytes := make([]byte, n)
	_, err := rand.Read(bytes)
	require.NoError(t, err)

	return hex.EncodeToString(bytes)
}

// FutureTime returns a time in the future for testing
func FutureTime(duration time.Duration) time.Time {
	return time.Now().Add(duration)
}

// PastTime returns a time in the past for testing
func PastTime(duration time.Duration) time.Time {
	return time.Now().Add(-duration)
}

// PrivateKeyToAddress returns the Ethereum address for a given private key
func PrivateKeyToAddress(t *testing.T, pk *ecdsa.PrivateKey) common.Address {
	t.Helper()
	return crypto.PubkeyToAddress(pk.PublicKey)
}

// SignAuthMessage signs the authentication message for address+timestamp
// Message format: "<address>:<timestamp>" and uses the same prefixing as
// accounts.TextHash (ethereum personal sign) to be compatible with verification.
func SignAuthMessage(t *testing.T, pk *ecdsa.PrivateKey, address string, timestamp int64) string {
	t.Helper()
	msg := []byte(fmt.Sprintf("%s:%d", address, timestamp))
	sig, err := GenerateTestSignature(t, pk, msg)
	require.NoError(t, err)
	return hexutil.Encode(sig)
}
