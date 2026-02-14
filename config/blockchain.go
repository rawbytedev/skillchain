package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type BlockchainConfig struct {
	RPCEndpoint string `mapstructure:"RPC_ENDPOINT"`
	WSEndpoint  string `mapstructure:"WS_ENDPOINT"`

	// Contract addresses (set after deployment)
	LicenseNFTAddress common.Address `mapstructure:"LICENSE_NFT_ADDRESS"`
	StakingNFTAddress common.Address `mapstructure:"STAKING_NFT_ADDRESS"`

	// Wallet for transactions (optional for hackathon)
	AdminPrivateKey string         `mapstructure:"ADMIN_PRIVATE_KEY"`
	AdminAddress    common.Address `mapstructure:"ADMIN_ADDRESS"`

	// Gas settings
	GasLimit uint64 `mapstructure:"GAS_LIMIT"`
	GasPrice int64  `mapstructure:"GAS_PRICE"` // in gwei

	// Network
	ChainID int64  `mapstructure:"CHAIN_ID"`
	Network string `mapstructure:"NETWORK"` // sepolia, localhost, etc.
}

func DefaultBlockchainConfig() *BlockchainConfig {
	return &BlockchainConfig{
		RPCEndpoint: "http://localhost:8545",
		WSEndpoint:  "ws://localhost:8546",
		GasLimit:    300000,
		GasPrice:    20,       // gwei
		ChainID:     11155111, // Sepolia
		Network:     "sepolia",
	}
}

func (c *BlockchainConfig) Validate() error {
	if c.RPCEndpoint == "" {
		return fmt.Errorf("RPC endpoint is required")
	}

	if c.LicenseNFTAddress == (common.Address{}) {
		return fmt.Errorf("license NFT address is required")
	}

	if c.StakingNFTAddress == (common.Address{}) {
		return fmt.Errorf("staking NFT address is required")
	}

	return nil
}
