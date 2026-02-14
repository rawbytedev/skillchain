package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	ServerPort         string
	EthNodeURL         string
	ContractAddr       string
	JWTSecret          string
	CacheTTL           int           // seconds
	RateLimit          int           // requests per minute
	CleanupInterval    time.Duration // KVStore cleanup interval
	DemoMode           bool          // Use demo mode for testing
	// Blockchain / signing config
	LicenseNFTAddress  string
	SignatureNonce     string
	ChainID            int64
	SignerPrivateKey   string
	EnableBlockchain   bool
	WSEndpoint         string
	Env                string
}

func Load() *Config {
	cleanupInterval := getEnvAsDuration("CLEANUP_INTERVAL", 5*time.Minute)
	demoMode := getEnvAsBool("DEMO_MODE", false)

	cfg := &Config{
		ServerPort:      getEnv("PORT", "8080"),
		EthNodeURL:      getEnv("ETH_NODE_URL", "wss://sepolia.infura.io/ws/v3/YOUR_KEY"),
		ContractAddr:    getEnv("CONTRACT_ADDR", "0x..."),
		JWTSecret:       getEnv("JWT_SECRET", "your-secret-key-change-in-production"),
		CacheTTL:        getEnvAsInt("CACHE_TTL", 300),
		RateLimit:       getEnvAsInt("RATE_LIMIT", 100),
		CleanupInterval: cleanupInterval,
		DemoMode:        demoMode,
		LicenseNFTAddress: getEnv("LICENSE_NFT_ADDRESS", "0x..."),
		SignatureNonce:    getEnv("SIGNATURE_NONCE", "default-nonce"),
		ChainID:           int64(getEnvAsInt("CHAIN_ID", 11155111)),
		SignerPrivateKey:  getEnv("SIGNER_PRIVATE_KEY", ""),
		EnableBlockchain:  getEnvAsBool("ENABLE_BLOCKCHAIN", false),
		WSEndpoint:        getEnv("WS_ENDPOINT", ""),
		Env:               getEnv("ENV", "development"),
	}

	// Validate critical config in non-demo mode
	if !cfg.DemoMode {
		if cfg.ContractAddr == "0x..." || cfg.ContractAddr == "" {
			fmt.Println("warning: CONTRACT_ADDR not configured, using placeholder")
		}
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if duration, err := time.ParseDuration(value); err == nil {
			return duration
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return defaultValue
}
