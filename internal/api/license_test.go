package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"moltket/internal/blockchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type mockLicenseService struct {
    requestResp *blockchain.LicenseResponse
    requestErr  error
    verifyResp  *blockchain.AccessResult
    verifyErr   error
    recordErr   error
}

func (m *mockLicenseService) RequestLicense(ctx context.Context, req *blockchain.LicenseRequest) (*blockchain.LicenseResponse, error) {
    return m.requestResp, m.requestErr
}

func (m *mockLicenseService) VerifyAccess(ctx context.Context, user common.Address, toolID *big.Int) (*blockchain.AccessResult, error) {
    return m.verifyResp, m.verifyErr
}

func (m *mockLicenseService) RecordLicenseMinted(ctx context.Context, user common.Address, toolID, expiresAt, nonce *big.Int) error {
    return m.recordErr
}

func TestLicenseHandler_Integration(t *testing.T) {
    e := echo.New()
    
    // Test 1: Successful license request
    t.Run("SuccessfulLicenseRequest", func(t *testing.T) {
        validAddr := common.HexToAddress("0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
        mockService := &mockLicenseService{
            requestResp: &blockchain.LicenseResponse{
                ToolID:      big.NewInt(42),
                User:        validAddr,
                ExpiresAt:   big.NewInt(1893456000),
                Nonce:       big.NewInt(123456),
                SignatureR:  "r_value",
                SignatureS:  "s_value",
                SignatureV:  "v_value",
                Price:       "0.01",
                Contract:    "0xContract",
            },
        }
        
        handler := NewLicenseHandler(mockService)
        
        reqBody := map[string]string{
            "user_address": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
            "tool_id":      "42",
        }
        
        body, _ := json.Marshal(reqBody)
        req := httptest.NewRequest(http.MethodPost, "/api/v1/license/request", bytes.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        
        err := handler.RequestLicense(c)
        require.NoError(t, err)
        assert.Equal(t, http.StatusOK, rec.Code)
        
        var response map[string]interface{}
        json.Unmarshal(rec.Body.Bytes(), &response)
        
        assert.True(t, response["success"].(bool))
        data := response["data"].(map[string]interface{})
        assert.Equal(t, "42", data["tool_id"])
        assert.Equal(t, validAddr.Hex(), data["user"])
    })
    
    // Test 2: Invalid request format
    t.Run("InvalidRequestFormat", func(t *testing.T) {
        mockService := &mockLicenseService{}
        handler := NewLicenseHandler(mockService)
        
        // Missing required fields
        reqBody := map[string]string{
            "user_address": "0xUser1",
            // Missing tool_id
        }
        
        body, _ := json.Marshal(reqBody)
        req := httptest.NewRequest(http.MethodPost, "/api/v1/license/request", bytes.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        
        err := handler.RequestLicense(c)
        require.NoError(t, err)
        assert.Equal(t, http.StatusBadRequest, rec.Code)
    })
    
    // Test 3: Service returns error
    t.Run("ServiceError", func(t *testing.T) {
        mockService := &mockLicenseService{
            requestErr: fmt.Errorf("service error"),
        }
        handler := NewLicenseHandler(mockService)
        
        reqBody := map[string]string{
            "user_address": "0x70997970C51812dc3A010C7d01b50e0d17dc79C8",
            "tool_id":      "43",
        }
        
        body, _ := json.Marshal(reqBody)
        req := httptest.NewRequest(http.MethodPost, "/api/v1/license/request", bytes.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        
        err := handler.RequestLicense(c)
        require.NoError(t, err)
        assert.Equal(t, http.StatusBadRequest, rec.Code)
        
        var response map[string]interface{}
        json.Unmarshal(rec.Body.Bytes(), &response)
        assert.Contains(t, response["error"], "service error")
    })
    
    // Test 4: Verify access - valid license
    t.Run("VerifyAccessValid", func(t *testing.T) {
        expiresAt := time.Now().Add(24 * time.Hour)
        mockService := &mockLicenseService{
            verifyResp: &blockchain.AccessResult{
                Valid:          true,
                Tier:           "licensed",
                CallsRemaining: 950,
                ExpiresAt:      &expiresAt,
                ProvenanceHash: "hash123",
            },
        }
        handler := NewLicenseHandler(mockService)
        
        reqBody := map[string]string{
            "user_address": "0xUser3",
            "tool_id":      "44",
        }
        
        body, _ := json.Marshal(reqBody)
        req := httptest.NewRequest(http.MethodPost, "/api/v1/access/verify", bytes.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        
        err := handler.VerifyAccess(c)
        require.NoError(t, err)
        assert.Equal(t, http.StatusOK, rec.Code)
        
        var response map[string]interface{}
        json.Unmarshal(rec.Body.Bytes(), &response)
        
        assert.True(t, response["valid"].(bool))
        assert.Equal(t, "licensed", response["tier"])
        assert.Equal(t, float64(950), response["calls_remaining"])
    })
    
    // Test 5: Verify access - invalid license
    t.Run("VerifyAccessInvalid", func(t *testing.T) {
        mockService := &mockLicenseService{
            verifyResp: &blockchain.AccessResult{
                Valid:  false,
                Tier:   "none",
                Reason: "No valid license",
            },
        }
        handler := NewLicenseHandler(mockService)
        
        reqBody := map[string]string{
            "user_address": "0xUser4",
            "tool_id":      "45",
        }
        
        body, _ := json.Marshal(reqBody)
        req := httptest.NewRequest(http.MethodPost, "/api/v1/access/verify", bytes.NewReader(body))
        req.Header.Set("Content-Type", "application/json")
        rec := httptest.NewRecorder()
        c := e.NewContext(req, rec)
        
        err := handler.VerifyAccess(c)
        require.NoError(t, err)
        assert.Equal(t, http.StatusForbidden, rec.Code)
        
        var response map[string]interface{}
        json.Unmarshal(rec.Body.Bytes(), &response)
        assert.False(t, response["valid"].(bool))
        assert.Equal(t, "No valid license", response["reason"])
    })
}