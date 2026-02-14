package models

import "time"

// License represents a cached license entry
type License struct {
    UserAddress string    `json:"user_address"`
    ToolID      string    `json:"tool_id"`
    ExpiresAt   time.Time `json:"expires_at"`
    Nonce       string    `json:"nonce"`
    Price       string    `json:"price"`
    CreatedAt   time.Time `json:"created_at"`
    MaxCalls    int       `json:"max_calls"`
    CallsUsed   int       `json:"calls_used"`
    Tier        string    `json:"tier"`
}

// LicenseMetadata holds on-chain metadata for a license
type LicenseMetadata struct {
    ExpiresAt time.Time `json:"expires_at"`
}
