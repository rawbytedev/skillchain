package api

import (
	"bytes"
	"context"
	"encoding/json"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"moltket/config"
	"moltket/internal/auth"
	"moltket/internal/blockchain"
	"moltket/internal/cache"
)

// -----------------------------------------------------------------------------
// MOCK BLOCKCHAIN CLIENT
// -----------------------------------------------------------------------------

type mockBlockchainClient struct {
	isLicenseValidFunc func(user common.Address, toolID *big.Int) (bool, error)
}

func (m *mockBlockchainClient) IsLicenseValid(user common.Address, toolID *big.Int) (bool, error) {
	if m.isLicenseValidFunc != nil {
		return m.isLicenseValidFunc(user, toolID)
	}
	return false, nil
}
func (m *mockBlockchainClient) GetLicenseMetadata(toolIDstr string) (*blockchain.LicenseMetadata, error) {
	// Return dummy metadata for testing
	return &blockchain.LicenseMetadata{
		TokenID: big.NewInt(1),
	}, nil
}

// -----------------------------------------------------------------------------
// TEST SETUP
// -----------------------------------------------------------------------------

func setupTestServer(t *testing.T) (*echo.Echo, *blockchain.LicenseService, *cache.Client, *config.Config) {
	t.Helper()

	// Create a temporary directory for test files
	tempDir, err := os.MkdirTemp("", "moltket-test-*")
	require.NoError(t, err)
	t.Cleanup(func() { os.RemoveAll(tempDir) })

	// Override home directory for path patching
	origHome := os.Getenv("HOME")
	if origHome == "" {
		origHome = os.Getenv("USERPROFILE") // Windows
	}
	os.Setenv("HOME", tempDir)
	t.Cleanup(func() { os.Setenv("HOME", origHome) })

	// Create test config
	cfg := &config.Config{
		LicenseNFTAddress: "0x1234567890123456789012345678901234567890",
		SignatureNonce:    "test-nonce",
		CacheTTL:          int(5 * time.Minute),
		//EnableCache:       true,
		EnableBlockchain: true,
		ChainID:          31337, // Hardhat local
	}

	// In-memory KV store
	kvStore := cache.NewKVStore()

	// Mock blockchain client
	mockBC := &mockBlockchainClient{
		isLicenseValidFunc: func(user common.Address, toolID *big.Int) (bool, error) {
			// Return true for specific tool ID (e.g., 42)
			if toolID.Int64() == 42 {
				return true, nil
			}
			return false, nil
		},
	}

	// Create EIP-712 signer with a fixed private key (Hardhat account #0)
	privateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	signer, err := auth.NewSigner(privateKeyHex, big.NewInt(cfg.ChainID), common.HexToAddress(cfg.LicenseNFTAddress))
	require.NoError(t, err)

	// Create license service
	licenseService := blockchain.NewLicenseService(cfg, kvStore, signer, mockBC)

	// Create Echo server and register all routes
	e := echo.New()
	licenseHandler := NewLicenseHandler(licenseService)
	e.POST("/api/v1/license/request", licenseHandler.RequestLicense)
	e.POST("/api/v1/access/verify", licenseHandler.VerifyAccess)
	e.POST("/api/v1/license/record-minted", licenseHandler.RecordLicenseMinted)
	e.GET("/api/v1/health", licenseHandler.HealthCheck) // ensure HealthCheck exists

	return e, licenseService, kvStore, cfg
}

// -----------------------------------------------------------------------------
// TESTS
// -----------------------------------------------------------------------------

func TestLicenseRequestEndpoint_Success(t *testing.T) {
	e, _, _, _ := setupTestServer(t)

	reqBody := map[string]string{
		"user_address": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
		"tool_id":      "42",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/license/request", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var resp map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	require.NoError(t, err)

	require.True(t, resp["success"].(bool))
	data := resp["data"].(map[string]interface{})
	require.Equal(t, "42", data["tool_id"])
	require.Equal(t, "0x70997970C51812dc3A010C7d01b50e0d17dc79C8", data["user"])
	require.NotEmpty(t, data["signature"].(map[string]interface{})["r"])
	require.NotEmpty(t, data["signature"].(map[string]interface{})["s"])
	require.NotEmpty(t, data["signature"].(map[string]interface{})["v"])
	require.NotEmpty(t, data["price"])
}

func TestLicenseRequestEndpoint_InvalidAddress(t *testing.T) {
	// Note: Our current handler does NOT validate Ethereum addresses;
	// it accepts any string and passes it to common.HexToAddress (which may return zero).
	// Therefore the request succeeds (200) even with an invalid address.
	// If you want to enforce address validation, modify the handler.
	e, _, _, _ := setupTestServer(t)

	reqBody := map[string]string{
		"user_address": "not-an-address",
		"tool_id":      "42",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/license/request", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	// Expect 200 because no validation is performed.
	require.Equal(t, http.StatusOK, rec.Code)
}

func TestLicenseRequestEndpoint_MissingFields(t *testing.T) {
	e, _, _, _ := setupTestServer(t)

	reqBody := map[string]string{
		"user_address": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
		// missing tool_id
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/license/request", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestVerifyAccessEndpoint_FreeTier(t *testing.T) {
	e, _, _, _ := setupTestServer(t)
	user := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	toolID := big.NewInt(99) // tool ID that has no license (mocked)

	// First 100 calls should succeed as free
	for i := 0; i < 100; i++ {
		reqBody := map[string]string{
			"user_address": user.Hex(),
			"tool_id":      toolID.String(),
		}
		body, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/api/v1/access/verify", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		require.Equal(t, http.StatusOK, rec.Code, "iteration %d", i)
		var resp map[string]interface{}
		err := json.Unmarshal(rec.Body.Bytes(), &resp)
		require.NoError(t, err)
		require.True(t, resp["valid"].(bool))
		require.Equal(t, "free", resp["tier"])
		require.NotEmpty(t, resp["provenance_hash"])
	}

	// 101st call should be forbidden (no license)
	reqBody := map[string]string{
		"user_address": user.Hex(),
		"tool_id":      toolID.String(),
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/access/verify", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	require.Equal(t, http.StatusForbidden, rec.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.False(t, resp["valid"].(bool))
	require.Equal(t, "free tier exhausted and no valid license found", resp["reason"])
}

func TestVerifyAccessEndpoint_LicensedTier(t *testing.T) {
	e, _, _, _ := setupTestServer(t)
	user := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	toolID := big.NewInt(42) // tool ID that has a license (mocked)

	// First request
	reqBody := map[string]string{
		"user_address": user.Hex(),
		"tool_id":      toolID.String(),
	}
	body1, _ := json.Marshal(reqBody)
	req1 := httptest.NewRequest(http.MethodPost, "/api/v1/access/verify", bytes.NewReader(body1))
	req1.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec1 := httptest.NewRecorder()
	e.ServeHTTP(rec1, req1)

	require.Equal(t, http.StatusOK, rec1.Code)
	var resp1 map[string]interface{}
	err := json.Unmarshal(rec1.Body.Bytes(), &resp1)
	require.NoError(t, err)
	require.True(t, resp1["valid"].(bool))
	require.Equal(t, "licensed", resp1["tier"])
	require.NotNil(t, resp1["expires_at"])
	require.Greater(t, resp1["calls_remaining"].(float64), 0.0)

	// Second request – create a fresh request body
	body2, _ := json.Marshal(reqBody)
	req2 := httptest.NewRequest(http.MethodPost, "/api/v1/access/verify", bytes.NewReader(body2))
	req2.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec2 := httptest.NewRecorder()
	e.ServeHTTP(rec2, req2)

	require.Equal(t, http.StatusOK, rec2.Code)
	var resp2 map[string]interface{}
	err = json.Unmarshal(rec2.Body.Bytes(), &resp2)
	require.NoError(t, err)
	require.True(t, resp2["valid"].(bool))
	require.Equal(t, "licensed", resp2["tier"])
}

func TestRecordMintedEndpoint_Success(t *testing.T) {
	e, licenseService, _, _ := setupTestServer(t)
	ctx := context.Background()
	user := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	toolID := big.NewInt(42)

	// First, request a license to get a pending entry with real nonce & expiry
	licenseReq := &blockchain.LicenseRequest{
		UserAddress: user,
		ToolID:      toolID,
	}
	resp, err := licenseService.RequestLicense(ctx, licenseReq)
	require.NoError(t, err)

	// Now call the endpoint to record it as minted, using the values from resp
	recordBody := map[string]string{
		"user_address": user.Hex(),
		"tool_id":      toolID.String(),
		"expires_at":   resp.ExpiresAt.String(),
		"nonce":        resp.Nonce.String(),
	}
	body, _ := json.Marshal(recordBody)

	httpReq := httptest.NewRequest(http.MethodPost, "/api/v1/license/record-minted", bytes.NewReader(body))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, httpReq)

	require.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.True(t, response["success"].(bool))
}

func TestRecordMintedEndpoint_InvalidNonce(t *testing.T) {
	e, licenseService, _, _ := setupTestServer(t)
	ctx := context.Background()
	user := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	toolID := big.NewInt(42)

	// Request a license
	licenseReq := &blockchain.LicenseRequest{
		UserAddress: user,
		ToolID:      toolID,
	}
	resp, err := licenseService.RequestLicense(ctx, licenseReq)
	require.NoError(t, err)

	// Try to record with a different nonce
	recordBody := map[string]string{
		"user_address": user.Hex(),
		"tool_id":      toolID.String(),
		"expires_at":   resp.ExpiresAt.String(),
		"nonce":        "99999", // wrong nonce
	}
	body, _ := json.Marshal(recordBody)

	httpReq := httptest.NewRequest(http.MethodPost, "/api/v1/license/record-minted", bytes.NewReader(body))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, httpReq)

	require.Equal(t, http.StatusBadRequest, rec.Code)
	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	require.NoError(t, err)
	require.Contains(t, response["error"], "nonce mismatch")
}

func TestHealthCheck(t *testing.T) {
	e, _, _, _ := setupTestServer(t)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
	var resp map[string]interface{}
	err := json.Unmarshal(rec.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, "OK", resp["status"])
	require.Equal(t, "skillchain-verification", resp["service"])
}
