package core

import (
	"testing"
	"time"

	"moltket/internal/testutils"

	"github.com/stretchr/testify/assert"
)

func TestProvenanceHashEdgeCases(t *testing.T) {
	service := &VerificationService{}

	t.Run("Hash with Empty Strings", func(t *testing.T) {
		hash1 := service.generateProvenanceHash("", "", "")
		hash2 := service.generateProvenanceHash("", "", "")
		assert.Equal(t, hash1, hash2)
		assert.Len(t, hash1, 64)
	})

	t.Run("Hash Sensitivity to Input Order", func(t *testing.T) {
		hash1 := service.generateProvenanceHash("license", "tool", "user")
		hash2 := service.generateProvenanceHash("tool", "license", "user")
		assert.NotEqual(t, hash1, hash2)
	})
}

func TestSignatureValidationEdgeCases(t *testing.T) {
	service := &VerificationService{}

	t.Run("Signature Valid at Boundary", func(t *testing.T) {
		// Exactly 5 minutes ago (the boundary)
		pk := testutils.GenerateTestPrivateKey(t, "edge-boundary")
		addr := testutils.PrivateKeyToAddress(t, pk).Hex()
		timestamp := time.Now().Unix() - 300
		sig := testutils.SignAuthMessage(t, pk, addr, timestamp)

		result := service.ValidateSignature(addr, timestamp, sig)
		assert.True(t, result)
	})

	t.Run("Signature Invalid Just Over Boundary", func(t *testing.T) {
		// Just over 5 minutes ago
		pk := testutils.GenerateTestPrivateKey(t, "edge-over")
		addr := testutils.PrivateKeyToAddress(t, pk).Hex()
		timestamp := time.Now().Unix() - 301
		sig := testutils.SignAuthMessage(t, pk, addr, timestamp)

		result := service.ValidateSignature(addr, timestamp, sig)
		assert.False(t, result)
	})

	t.Run("Future Timestamp", func(t *testing.T) {
		// Future timestamp should be invalid
		pk := testutils.GenerateTestPrivateKey(t, "edge-future")
		addr := testutils.PrivateKeyToAddress(t, pk).Hex()
		timestamp := time.Now().Unix() + 60
		sig := testutils.SignAuthMessage(t, pk, addr, timestamp)

		result := service.ValidateSignature(addr, timestamp, sig)
		assert.False(t, result)
	})
}
