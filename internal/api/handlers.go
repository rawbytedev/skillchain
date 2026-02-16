package api

import (
	"context"
	"moltket/config"
	"moltket/internal/blockchain"
	"moltket/internal/cache"
	"moltket/internal/core"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

type Server struct {
	echo        *echo.Echo
	config      *config.Config
	service     *core.VerificationService
	cache       *cache.Client
	blockchain  blockchain.BlockchainInterface
	voteService *core.VoteService
}

func NewServer(cfg *config.Config, cacheClient *cache.Client, ethClient blockchain.BlockchainInterface, voteService *core.VoteService) *Server {
	e := echo.New()

	// Middleware for security and observability
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(rate.Limit(cfg.RateLimit))))
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	// Initialize service layer with the KVStore from cache client
	// We need to pass the store directly, so we'll need to update the approach
	service := core.NewVerificationService(cfg, cacheClient.GetStore(), ethClient)

	server := &Server{
		echo:        e,
		config:      cfg,
		service:     service,
		cache:       cacheClient,
		blockchain:  ethClient,
		voteService: voteService,
	}

	server.setupRoutes()
	server.setupLicenseRoutes()
	server.setupVoteRoutes(voteService)
	return server
}

func (s *Server) setupRoutes() {
	api := s.echo.Group("/api/v1")

	// Public health check
	api.GET("/health", s.healthCheck)

	// License verification endpoint (used by tool hosts)
	api.POST("/verify", s.verifyLicense)

	// Admin endpoints (optional for demo)
	//admin := api.Group("/admin", s.authenticateAdmin)
	//admin.POST("/cache/clear", s.clearCache)
}

func (s *Server) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"service": "skillchain-verification",
		"time":    time.Now().UTC(),
	})
}

// Main verification handler
func (s *Server) verifyLicense(c echo.Context) error {
	var req struct {
		LicenseNFTID string `json:"license_nft_id" validate:"required,alphanum"`
		ToolID       string `json:"tool_id" validate:"required,alphanum"`
		UserAddress  string `json:"user_address" validate:"required,eth_addr"`
		Timestamp    int64  `json:"timestamp" validate:"required"`
		Signature    string `json:"signature" validate:"required"` // EIP-712 signature
	}

	// Bind and validate input
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request format",
		})
	}

	// Validate signature (prevent replay attacks)
	if !s.service.ValidateSignature(req.UserAddress, req.Timestamp, req.Signature) {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Invalid or expired signature",
		})
	}

	// Check rate limit using Redis
	key := "rate_limit:" + req.UserAddress
	if allowed, err := s.service.CheckRateLimit(key); err != nil || !allowed {
		return c.JSON(http.StatusTooManyRequests, map[string]string{
			"error": "Rate limit exceeded",
		})
	}

	// Perform verification
	result, err := s.service.VerifyLicense(
		req.LicenseNFTID,
		req.ToolID,
		req.UserAddress,
	)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Verification failed",
		})
	}

	if !result.Valid {
		return c.JSON(http.StatusForbidden, map[string]interface{}{
			"valid":  false,
			"reason": result.Reason,
			"code":   result.ErrorCode,
		})
	}

	// Success response
	return c.JSON(http.StatusOK, map[string]interface{}{
		"valid":           true,
		"license_id":      req.LicenseNFTID,
		"tool_id":         req.ToolID,
		"user_address":    req.UserAddress,
		"expires_at":      result.ExpiresAt,
		"calls_remaining": result.CallsRemaining,
		"tier":            result.Tier,
		"verified_at":     time.Now().UTC(),
		"provenance_hash": result.ProvenanceHash, // Your tamper-proof hash
	})
}

func (s *Server) Start(addr string) error {
	return s.echo.Start(addr)
}

// Handler returns the underlying http.Handler (Echo) for testing or embedding
func (s *Server) Handler() http.Handler {
	return s.echo
}

// Shutdown gracefully shuts down the server
func (s *Server) Shutdown(ctx context.Context) error {
	return s.echo.Shutdown(ctx)
}
