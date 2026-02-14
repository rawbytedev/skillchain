package blockchain

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"moltket/config"
	"moltket/internal/auth"
	"moltket/internal/kvstore"
	"moltket/internal/models"

	"github.com/ethereum/go-ethereum/common"
)

type LicenseService struct {
	config     *config.Config
	cache      kvstore.Store
	signer     *auth.EIP712Signer
	blockchain BlockchainInterface
}

type BlockchainInterface interface {
	IsLicenseValid(user common.Address, toolID *big.Int) (bool, error)
	GetLicenseMetadata(toolIDStr string) (*LicenseMetadata, error)
}

type LicenseServiceInterface interface {
	RequestLicense(ctx context.Context, req *LicenseRequest) (*LicenseResponse, error)
	VerifyAccess(ctx context.Context, user common.Address, toolID *big.Int) (*AccessResult, error)
	RecordLicenseMinted(ctx context.Context, user common.Address, toolID, expiresAt, nonce *big.Int) error
}

func NewLicenseService(
	cfg *config.Config,
	cache kvstore.Store,
	signer *auth.EIP712Signer,
	bc BlockchainInterface,
) *LicenseService {
	return &LicenseService{
		config:     cfg,
		cache:      cache,
		signer:     signer,
		blockchain: bc,
	}
}

type LicenseRequest struct {
	UserAddress common.Address `json:"user_address"`
	ToolID      *big.Int       `json:"tool_id"`
}

type LicenseResponse struct {
	ToolID     *big.Int       `json:"tool_id"`
	User       common.Address `json:"user"`
	ExpiresAt  *big.Int       `json:"expires_at"`
	Nonce      *big.Int       `json:"nonce"`
	SignatureR string         `json:"signature_r"`
	SignatureS string         `json:"signature_s"`
	SignatureV string         `json:"signature_v"`
	Price      string         `json:"price"`
	Contract   string         `json:"contract_address"`
}

type AccessResult struct {
	Valid          bool       `json:"valid"`
	Tier           string     `json:"tier"`
	CallsRemaining int        `json:"calls_remaining"`
	ExpiresAt      *time.Time `json:"expires_at,omitempty"`
	ProvenanceHash string     `json:"provenance_hash,omitempty"`
	Reason         string     `json:"reason,omitempty"`
}

func (s *LicenseService) RequestLicense(ctx context.Context, req *LicenseRequest) (*LicenseResponse, error) {
	// Check if user already has a pending or active license
	licenseKey := fmt.Sprintf("license:%s:%s", req.UserAddress.Hex(), req.ToolID.String())

	if cached, found := s.cache.Get(ctx, licenseKey); found {
		if license, ok := cached.(*models.License); ok && time.Now().Before(license.ExpiresAt) {
			// Return existing license info
			return nil, fmt.Errorf("license already active until %s", license.ExpiresAt.Format(time.RFC3339))
		}
	}

	// Check for pending request to prevent double-issuance
	pendingKey := fmt.Sprintf("pending:%s", licenseKey)
	if _, found := s.cache.Get(ctx, pendingKey); found {
		return nil, fmt.Errorf("license request already pending")
	}

	// Generate license parameters
	expiresAt := big.NewInt(time.Now().Add(30 * 24 * time.Hour).Unix()) // 30 days
	nonce := big.NewInt(time.Now().UnixNano())

	// Calculate dynamic price based on tool reputation
	price := s.calculateLicensePrice(req.ToolID)

	// Create EIP-712 signature
	r, sigS, sigV, err := s.signer.CreateLicenseSignature(
		req.UserAddress,
		req.ToolID,
		expiresAt,
		nonce,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create signature: %w", err)
	}

	// Store pending license (10-minute TTL to allow user to complete transaction)
	pendingLicense := &models.License{
		UserAddress: req.UserAddress.Hex(),
		ToolID:      req.ToolID.String(),
		ExpiresAt:   time.Unix(expiresAt.Int64(), 0),
		Nonce:       nonce.String(),
		Price:       price,
		CreatedAt:   time.Now(),
	}

	if err := s.cache.Set(ctx, pendingKey, pendingLicense, 10*time.Minute); err != nil {
		return nil, fmt.Errorf("failed to cache pending license: %w", err)
	}

	// Return license response
	return &LicenseResponse{
		ToolID:     req.ToolID,
		User:       req.UserAddress,
		ExpiresAt:  expiresAt,
		Nonce:      nonce,
		SignatureR: hex.EncodeToString(r),
		SignatureS: hex.EncodeToString(sigS),
		SignatureV: hex.EncodeToString(sigV),
		Price:      price,
		Contract:   s.config.LicenseNFTAddress,
	}, nil
}

func (s *LicenseService) VerifyAccess(ctx context.Context, user common.Address, toolID *big.Int) (*AccessResult, error) {
	// TIER 1: LICENSED ACCESS (check first, before free tier)
	licenseKey := fmt.Sprintf("license:%s:%s", user.Hex(), toolID.String())

	// Check cache first
	if cached, found := s.cache.Get(ctx, licenseKey); found {
		if license, ok := cached.(*models.License); ok {
			if time.Now().Before(license.ExpiresAt) {
				// Update usage count
				license.CallsUsed++
				if err := s.cache.Set(ctx, licenseKey, license, 24*time.Hour); err != nil {
					return nil, fmt.Errorf("failed to update license usage: %w", err)
				}

				return &AccessResult{
					Valid:          true,
					Tier:           "licensed",
					CallsRemaining: license.MaxCalls - license.CallsUsed,
					ExpiresAt:      &license.ExpiresAt,
				}, nil
			}
		}
	}

	// Cache miss: Check blockchain (rare - only once per license period)
	isValid, err := s.blockchain.IsLicenseValid(user, toolID)
	if err == nil && isValid {
		// License is valid on-chain - cache it
		metadata, err := s.blockchain.GetLicenseMetadata(toolID.String())
		if err == nil {
			license := &models.License{
				UserAddress: user.Hex(),
				ToolID:      toolID.String(),
				ExpiresAt:   metadata.ExpiresAt,
				MaxCalls:    1000, // Default licensed calls
				CallsUsed:   1,
				Tier:        "licensed",
			}

			if err := s.cache.Set(ctx, licenseKey, license, 24*time.Hour); err == nil {
				return &AccessResult{
					Valid:          true,
					Tier:           "licensed",
					CallsRemaining: license.MaxCalls - license.CallsUsed,
					ExpiresAt:      &license.ExpiresAt,
				}, nil
			}
		}
	}

	// TIER 2: FREE ACCESS (100 calls/day) - fallback if no license
	freeTierKey := fmt.Sprintf("free:%s:%s", user.Hex(), toolID.String())
	freeUsage, err := s.cache.Increment(ctx, freeTierKey, 1, 24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to track free usage: %w", err)
	}

	if freeUsage <= 100 {
		provenanceHash := s.generateProvenanceHash("FREE", toolID.String(), user.Hex())
		return &AccessResult{
			Valid:          true,
			Tier:           "free",
			CallsRemaining: 100 - int(freeUsage),
			ProvenanceHash: provenanceHash,
		}, nil
	}

	return &AccessResult{
		Valid:  false,
		Tier:   "none",
		Reason: "free tier exhausted and no valid license found",
	}, nil
}

func (s *LicenseService) RecordLicenseMinted(
	ctx context.Context,
	user common.Address,
	toolID *big.Int,
	expiresAt *big.Int,
	nonce *big.Int,
) error {
	// Move from pending to active
	licenseKey := fmt.Sprintf("license:%s:%s", user.Hex(), toolID.String())
	pendingKey := fmt.Sprintf("pending:%s", licenseKey)

	// Get pending license
	cached, found := s.cache.Get(ctx, pendingKey)
	if !found {
		return fmt.Errorf("no pending license found")
	}

	pendingLicense, ok := cached.(*models.License)
	if !ok {
		return fmt.Errorf("invalid pending license format")
	}

	// Verify nonce matches
	if pendingLicense.Nonce != nonce.String() {
		return fmt.Errorf("nonce mismatch")
	}

	// Create active license
	activeLicense := &models.License{
		UserAddress: user.Hex(),
		ToolID:      toolID.String(),
		ExpiresAt:   time.Unix(expiresAt.Int64(), 0),
		MaxCalls:    1000,
		CallsUsed:   0,
		Tier:        "licensed",
		CreatedAt:   time.Now(),
	}

	// Store active license, delete pending
	if err := s.cache.Set(ctx, licenseKey, activeLicense, 30*24*time.Hour); err != nil {
		return fmt.Errorf("failed to cache active license: %w", err)
	}

	s.cache.Delete(ctx, pendingKey)

	return nil
}

func (s *LicenseService) calculateLicensePrice(toolID *big.Int) string {
	// Simple pricing for demo: base price + reputation factor
	// In production, this would query tool reputation from database
	return "10000000000000000"
}

func (s *LicenseService) generateProvenanceHash(tier, toolID, userAddress string) string {
	timestamp := time.Now().UnixNano()
	data := fmt.Sprintf("%s:%s:%s:%d:%s",
		tier, toolID, userAddress, timestamp, s.config.SignatureNonce)

	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}
