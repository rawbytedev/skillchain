package core

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"moltket/config"
	"moltket/internal/blockchain"
	"moltket/internal/kvstore"
	"moltket/internal/models"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

type VerificationResult struct {
	Valid          bool
	Reason         string
	ErrorCode      string
	ExpiresAt      *time.Time
	CallsRemaining int
	Tier           string
	ProvenanceHash string
}

type VerificationService struct {
	config     *config.Config
	store      kvstore.Store
	blockchain blockchain.BlockchainInterface
}

func NewVerificationService(cfg *config.Config, store kvstore.Store, bc blockchain.BlockchainInterface) *VerificationService {
	return &VerificationService{
		config:     cfg,
		store:      store,
		blockchain: bc,
	}
}

func (s *VerificationService) VerifyLicense(licenseID, ToolID, UserAddress string) (*VerificationResult, error) {
	ctx := context.Background()

	// 1. Check cache first (performance optimization)
	cacheKey := fmt.Sprintf("license:%s:%s", licenseID, ToolID)
	if cached, found := s.store.Get(ctx, cacheKey); found {
		if result, ok := cached.(*VerificationResult); ok && result.Valid {
			return result, nil
		}
	}
	// Validate tool ID
	toolID, ok := new(big.Int).SetString(ToolID, 10)
	if !ok {
		return nil, fmt.Errorf("Invalid TooldID")
	}

	userAddress := common.HexToAddress(UserAddress)

	// 2. Verify on-chain NFT ownership
	isValid, err := s.blockchain.IsLicenseValid(userAddress, toolID)
	if err != nil {
		return &VerificationResult{
			Valid:     false,
			Reason:    "Blockchain verification failed",
			ErrorCode: "BLOCKCHAIN_ERROR",
		}, err
	}

	if !isValid {
		return &VerificationResult{
			Valid:     false,
			Reason:    "User does not own this license NFT",
			ErrorCode: "OWNERSHIP_INVALID",
		}, nil
	}

	// 3. Check license expiration and metadata
	metadata, err := s.blockchain.GetLicenseMetadata(licenseID)
	if err != nil {
		return &VerificationResult{
			Valid:     false,
			Reason:    "Failed to fetch license metadata",
			ErrorCode: "METADATA_ERROR",
		}, err
	}

	if time.Now().After(metadata.ExpiresAt) {
		return &VerificationResult{
			Valid:     false,
			Reason:    "License has expired",
			ErrorCode: "EXPIRED",
		}, nil
	}

	// 4. Check usage limits (simplified for demo)
	usageKey := fmt.Sprintf("usage:%s:%s", licenseID, toolID)
	usageCount, _ := s.store.Increment(ctx, usageKey, 1, 24*time.Hour)

	if int(usageCount) > metadata.MaxCalls {
		return &VerificationResult{
			Valid:     false,
			Reason:    "Usage limit exceeded",
			ErrorCode: "LIMIT_EXCEEDED",
		}, nil
	}

	// 5. Generate provenance hash (your expertise)
	provenanceHash := s.generateProvenanceHash(licenseID, ToolID, UserAddress)

	result := &VerificationResult{
		Valid:          true,
		ExpiresAt:      &metadata.ExpiresAt,
		CallsRemaining: metadata.MaxCalls - int(usageCount),
		Tier:           metadata.Tier,
		ProvenanceHash: provenanceHash,
	}

	// 6. Cache successful verification (5 minutes TTL)
	s.store.Set(ctx, cacheKey, result, time.Duration(s.config.CacheTTL)*time.Second)

	return result, nil
}

func (s *VerificationService) generateProvenanceHash(licenseID, toolID, userAddress string) string {
	timestamp := time.Now().UnixNano()
	data := fmt.Sprintf("%s:%s:%s:%d", licenseID, toolID, userAddress, timestamp)

	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func (s *VerificationService) ValidateSignature(userAddress string, timestamp int64, signature string) bool {
	// Prevent replay attacks by checking timestamp. Use integer-second
	// comparison to avoid microsecond rounding issues in tests.
	maxAge := int64((5 * time.Minute).Seconds())
	nowSec := time.Now().Unix()

	// Reject too-old or future timestamps
	if nowSec-timestamp > maxAge || timestamp > nowSec {
		return false
	}

	// Build the message in the same format used for signing
	msg := []byte(fmt.Sprintf("%s:%d", userAddress, timestamp))
	hash := accounts.TextHash(msg)

	// Decode hex signature (accept 0x prefixed)
	sigBytes, err := hexutil.Decode(signature)
	if err != nil {
		// try raw hex without 0x
		b, err2 := hex.DecodeString(signature)
		if err2 != nil {
			return false
		}
		sigBytes = b
	}

	if len(sigBytes) != 65 {
		return false
	}

	// Convert v from 27/28 to 0/1 if necessary
	if sigBytes[64] == 27 || sigBytes[64] == 28 {
		sigBytes[64] -= 27
	}

	pubKey, err := crypto.SigToPub(hash, sigBytes)
	if err != nil {
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return recoveredAddr == common.HexToAddress(userAddress)
}

func (s *VerificationService) CheckRateLimit(key string) (bool, error) {
	ctx := context.Background()
	limit := 100 // requests per minute

	count, err := s.store.Increment(ctx, key, 1, time.Minute)
	if err != nil {
		return true, err // Fail open on cache error
	}

	return count <= int64(limit), nil
}

func (s *VerificationService) VerifyAccess(licenseID, toolID, userAddress string) (*VerificationResult, error) {
	ctx := context.Background()

	// TIER 1: FREE ACCESS CHECK (NO BLOCKCHAIN)
	freeTierKey := fmt.Sprintf("free:%s:%s", userAddress, toolID)
	freeUsage, _ := s.store.Increment(ctx, freeTierKey, 1, 24*time.Hour)

	if freeUsage <= 100 { // Free tier limit: 100 calls/day
		provenanceHash := s.generateProvenanceHash("FREE", toolID, userAddress)
		return &VerificationResult{
			Valid:          true,
			Tier:           "free",
			CallsRemaining: 100 - int(freeUsage),
			ProvenanceHash: provenanceHash,
		}, nil
	}

	// TIER 2: LICENSED ACCESS (RARE BLOCKCHAIN CHECK)
	licenseKey := fmt.Sprintf("license:%s:%s", userAddress, toolID)
	if cached, found := s.store.Get(ctx, licenseKey); found {
		if license, ok := cached.(*models.License); ok && time.Now().Before(license.ExpiresAt) {
			license.CallsUsed++
			s.store.Set(ctx, licenseKey, license, 24*time.Hour)
			return &VerificationResult{
				Valid:          true,
				Tier:           license.Tier,
				CallsRemaining: license.MaxCalls - license.CallsUsed,
				ExpiresAt:      &license.ExpiresAt,
			}, nil
		}
	}
	// Validate tool ID
	ToolID, ok := new(big.Int).SetString(toolID, 10)
	if !ok {
		return nil, fmt.Errorf("Invalid toolID")
	}

	UserAddress := common.HexToAddress(userAddress)

	// CACHE MISS: verify on-chain
	isValid, err := s.blockchain.IsLicenseValid(UserAddress, ToolID)
	if err != nil || !isValid {
		return &VerificationResult{Valid: false, Reason: "No valid license"}, nil
	}

	metadata, err := s.blockchain.GetLicenseMetadata(toolID)
	if err != nil {
		return &VerificationResult{Valid: false, Reason: "failed to fetch license metadata"}, nil
	}

	license := &models.License{
		UserAddress: userAddress,
		ToolID:      toolID,
		ExpiresAt:   metadata.ExpiresAt,
		MaxCalls:    metadata.MaxCalls,
		CallsUsed:   1,
		Tier:        metadata.Tier,
	}

	s.store.Set(ctx, licenseKey, license, 24*time.Hour)

	return &VerificationResult{
		Valid:          true,
		Tier:           license.Tier,
		CallsRemaining: license.MaxCalls - license.CallsUsed,
		ExpiresAt:      &license.ExpiresAt,
	}, nil
}
