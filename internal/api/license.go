package api

import (
	"log"
	"math/big"
	"net/http"
	"time"

	"moltket/internal/auth"
	"moltket/internal/blockchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
)

type licenseHandler struct {
	licenseService blockchain.LicenseServiceInterface
}

func NewLicenseHandler(licenseService blockchain.LicenseServiceInterface) *licenseHandler {
	return &licenseHandler{
		licenseService: licenseService,
	}
}
func (h *licenseHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "healthy",
		"service": "skillchain-verification",
		"time":    time.Now().UTC(),
	})
}

// RequestLicense handles POST /api/v1/license/request
func (h *licenseHandler) RequestLicense(c echo.Context) error {
	var req struct {
		UserAddress string `json:"user_address" validate:"required,eth_addr"`
		ToolID      string `json:"tool_id" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	// Validate tool ID
	toolID, ok := new(big.Int).SetString(req.ToolID, 10)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid tool ID format",
		})
	}

	userAddress := common.HexToAddress(req.UserAddress)

	// Create license request
	licenseReq := &blockchain.LicenseRequest{
		UserAddress: userAddress,
		ToolID:      toolID,
	}

	// Request license from service
	resp, err := h.licenseService.RequestLicense(c.Request().Context(), licenseReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"tool_id":    resp.ToolID.String(),
			"user":       resp.User.Hex(),
			"expires_at": resp.ExpiresAt.String(),
			"nonce":      resp.Nonce.String(),
			"signature": map[string]string{
				"r": resp.SignatureR,
				"s": resp.SignatureS,
				"v": resp.SignatureV,
			},
			"price":            resp.Price,
			"contract_address": resp.Contract,
			"instructions":     "Call mintLicense() on the contract with these parameters",
		},
	})
}

// VerifyAccess handles POST /api/v1/access/verify
func (h *licenseHandler) VerifyAccess(c echo.Context) error {
	var req struct {
		UserAddress string `json:"user_address" validate:"required,eth_addr"`
		ToolID      string `json:"tool_id" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	toolID, ok := new(big.Int).SetString(req.ToolID, 10)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid tool ID format",
		})
	}

	userAddress := common.HexToAddress(req.UserAddress)

	// Verify access
	result, err := h.licenseService.VerifyAccess(c.Request().Context(), userAddress, toolID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal server error",
		})
	}

	if !result.Valid {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"valid":  false,
			"tier":   result.Tier,
			"reason": result.Reason,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"valid":           result.Valid,
		"tier":            result.Tier,
		"calls_remaining": result.CallsRemaining,
		"expires_at":      result.ExpiresAt,
		"provenance_hash": result.ProvenanceHash,
	})
}

// RecordLicenseMinted handles POST /api/v1/license/record-minted
// This is called by a blockchain event listener when a license is minted
func (h *licenseHandler) RecordLicenseMinted(c echo.Context) error {
	var req struct {
		UserAddress string `json:"user_address" validate:"required,eth_addr"`
		ToolID      string `json:"tool_id" validate:"required"`
		ExpiresAt   string `json:"expires_at" validate:"required"`
		Nonce       string `json:"nonce" validate:"required"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	// Parse inputs
	userAddress := common.HexToAddress(req.UserAddress)

	toolID, ok := new(big.Int).SetString(req.ToolID, 10)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid tool ID",
		})
	}

	expiresAt, ok := new(big.Int).SetString(req.ExpiresAt, 10)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid expires_at",
		})
	}

	nonce, ok := new(big.Int).SetString(req.Nonce, 10)
	if !ok {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid nonce",
		})
	}

	// Record the minted license
	err := h.licenseService.RecordLicenseMinted(
		c.Request().Context(),
		userAddress,
		toolID,
		expiresAt,
		nonce,
	)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"message": "License recorded successfully",
	})
}

// Update server setup to include license routes
func (s *Server) setupLicenseRoutes() {
	// Initialize license service dependencies
	chainID, _ := new(big.Int).SetString("11155111", 10) // Sepolia
	verifyingContract := common.HexToAddress(s.config.LicenseNFTAddress)

	signer, err := auth.NewSigner(
		s.config.JWTSecret, // Using JWT secret as private key for demo
		chainID,
		verifyingContract,
	)
	if err != nil {
		log.Fatalf("Failed to create signer: %v", err)
	}

	// Create license service
	licenseService := blockchain.NewLicenseService(
		s.config,
		s.cache,
		signer,
		s.blockchain,
	)

	// Create handler
	licenseHandler := NewLicenseHandler(licenseService)

	// Register routes
	api := s.echo.Group("/api/v1")
	api.POST("/license/request", licenseHandler.RequestLicense)
	api.POST("/access/verify", licenseHandler.VerifyAccess)
	api.POST("/license/record-minted", licenseHandler.RecordLicenseMinted)
}
