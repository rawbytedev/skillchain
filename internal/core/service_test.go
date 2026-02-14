package core

import (
	"context"
	"errors"
	"testing"
	"time"

	"moltket/config"
	"moltket/internal/blockchain"
	"moltket/internal/testutils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockStore implements the kvstore.Store interface for testing
type MockStore struct {
	mock.Mock
}

func (m *MockStore) Get(ctx context.Context, key string) (interface{}, bool) {
	args := m.Called(ctx, key)
	return args.Get(0), args.Bool(1)
}

func (m *MockStore) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	args := m.Called(ctx, key, value, ttl)
	return args.Error(0)
}

func (m *MockStore) Increment(ctx context.Context, key string, value int64, ttl time.Duration) (int64, error) {
	args := m.Called(ctx, key, value, ttl)
	return args.Get(0).(int64), args.Error(1)
}

func (m *MockStore) Delete(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func (m *MockStore) Clear(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *MockStore) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockStore) Keys(ctx context.Context) []string {
	args := m.Called(ctx)
	if args.Get(0) == nil {
		return []string{}
	}
	return args.Get(0).([]string)
}

// MockBlockchain implements blockchain operations for testing
type MockBlockchain struct {
	mock.Mock
}

func (m *MockBlockchain) VerifyNFTOwnership(contractAddr, licenseID, userAddress string) (bool, error) {
	args := m.Called(contractAddr, licenseID, userAddress)
	return args.Bool(0), args.Error(1)
}

func (m *MockBlockchain) GetLicenseMetadata(licenseID string) (*blockchain.LicenseMetadata, error) {
	args := m.Called(licenseID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*blockchain.LicenseMetadata), args.Error(1)
}

func (m *MockBlockchain) Close() error {
	args := m.Called()
	return args.Error(0)
}

// Test suite
func TestVerificationService(t *testing.T) {
	cfg := &config.Config{
		CacheTTL: 300,
	}

	t.Run("Generate Provenance Hash", func(t *testing.T) {
		service := &VerificationService{}

		hash1 := service.generateProvenanceHash("license1", "tool1", "user1")
		hash2 := service.generateProvenanceHash("license1", "tool1", "user1")
		hash3 := service.generateProvenanceHash("license2", "tool1", "user1")

		// Hashes should be consistent for same inputs
		assert.Equal(t, hash1, hash2)

		// Different inputs should produce different hashes
		assert.NotEqual(t, hash1, hash3)

		// Hash should be valid SHA256 hex string
		assert.Len(t, hash1, 64) // 64 chars for SHA256 hex
		assert.Regexp(t, `^[a-f0-9]{64}$`, hash1)
	})

	t.Run("Validate Signature - Valid Recent", func(t *testing.T) {
		service := &VerificationService{}
		// Generate deterministic key and address for test
		pk := testutils.GenerateTestPrivateKey(t, "seed-valid")
		addr := testutils.PrivateKeyToAddress(t, pk)
		timestamp := time.Now().Unix() - 60 // 1 minute ago
		sigHex := testutils.SignAuthMessage(t, pk, addr.Hex(), timestamp)
		result := service.ValidateSignature(addr.Hex(), timestamp, sigHex)
		assert.True(t, result)
	})

	t.Run("Validate Signature - Expired", func(t *testing.T) {
		service := &VerificationService{}
		pk := testutils.GenerateTestPrivateKey(t, "seed-expired")
		addr := testutils.PrivateKeyToAddress(t, pk)
		timestamp := time.Now().Unix() - 400 // 6+ minutes ago (beyond 5 min window)
		sigHex := testutils.SignAuthMessage(t, pk, addr.Hex(), timestamp)
		result := service.ValidateSignature(addr.Hex(), timestamp, sigHex)
		assert.False(t, result)
	})

	t.Run("Check Rate Limit - Below Limit", func(t *testing.T) {
		mockStore := new(MockStore)
		service := &VerificationService{
			config: cfg,
			store:  mockStore,
		}

		ctx := context.Background()
		mockStore.On("Increment", ctx, "rate_limit:user123", int64(1), time.Minute).Return(int64(50), nil)

		allowed, err := service.CheckRateLimit("rate_limit:user123")

		assert.NoError(t, err)
		assert.True(t, allowed)
		mockStore.AssertExpectations(t)
	})

	t.Run("Check Rate Limit - Above Limit", func(t *testing.T) {
		mockStore := new(MockStore)
		service := &VerificationService{
			config: cfg,
			store:  mockStore,
		}

		ctx := context.Background()
		mockStore.On("Increment", ctx, "rate_limit:user456", int64(1), time.Minute).Return(int64(101), nil)

		allowed, err := service.CheckRateLimit("rate_limit:user456")

		assert.NoError(t, err)
		assert.False(t, allowed)
		mockStore.AssertExpectations(t)
	})

	t.Run("Check Rate Limit - Cache Error (Fail Open)", func(t *testing.T) {
		mockStore := new(MockStore)
		service := &VerificationService{
			config: cfg,
			store:  mockStore,
		}

		ctx := context.Background()
		mockStore.On("Increment", ctx, "rate_limit:user789", int64(1), time.Minute).Return(int64(0), errors.New("store error"))

		allowed, err := service.CheckRateLimit("rate_limit:user789")

		assert.Error(t, err)
		assert.True(t, allowed) // Should fail open on cache error
		mockStore.AssertExpectations(t)
	})
}

// Note: All remaining tests are implemented in TestVerificationService
