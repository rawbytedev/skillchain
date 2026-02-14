package blockchain

import (
	"context"
	"fmt"
	"log"
	"math/big"
	skill "moltket/internal/contracts/Skill"
	"moltket/internal/contracts/Stake"
	"moltket/internal/contracts/license"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type SkillToken struct {
	address            common.Address
	SkillTokenContract *skill.Skill
}
type LicenseNFTContract struct {
	Licensecontract *license.License
	address         common.Address
}

// unused
func (l *LicenseNFTContract) ParseLicenseMinted(vLog types.Log) (any, any) {
	// Demo implementation - returns minimal event data
	// In production, would parse ABI-encoded event data
	return nil, nil
}

func (l *LicenseNFTContract) LicenseMintedSignature() common.Hash {
	// Return the keccak256 hash of "LicenseMinted(address,uint256,uint256)"
	// For demo, returning a zero hash
	return common.Hash{}
}

func (l *LicenseNFTContract) BalanceOfBatch(opts *bind.CallOpts, addresses []common.Address, tokenIDs []*big.Int) ([]*big.Int, error) {
	// Demo implementation - returns 1 for all addresses
	balances := make([]*big.Int, len(addresses))
	for i := range addresses {
		balances[i] = big.NewInt(1)
	}
	return balances, nil
}

func (l *LicenseNFTContract) Uri(opts *bind.CallOpts, tokenID *big.Int) (string, error) {
	// Demo implementation - returns a mock URI
	return "ipfs://QmDemo/" + tokenID.String(), nil
}

func (l *LicenseNFTContract) BalanceOf(opts *bind.CallOpts, userAddress common.Address, tokenID *big.Int) (*big.Int, error) {
	// Demo implementation - returns 1 if user address is valid
	if userAddress == common.HexToAddress("0x0") {
		return big.NewInt(0), nil
	}
	return big.NewInt(1), nil
}

func (l *LicenseNFTContract) Address() common.Address {
	return l.address
}

type StakingNFTContract struct {
	address       common.Address
	StakeContract *Stake.Stake
}

func (s *StakingNFTContract) OwnerOf(opts *bind.CallOpts, tokenID *big.Int) (common.Address, error) {
	// Demo implementation - returns a mock owner address
	return common.HexToAddress("0x1234567890123456789012345678901234567890"), nil
}

func (s *StakingNFTContract) Address() common.Address {
	return s.address
}

type Client struct {
	ethClient  *ethclient.Client
	licenseNFT *LicenseNFTContract
	stakingNFT *StakingNFTContract
	chainID    *big.Int
	rpcURL     string
	wsURL      string
	privateKey string // For signing transactions (optional)
}

type LicenseMetadata struct {
	TokenID   *big.Int       `json:"token_id"`
	Owner     common.Address `json:"owner"`
	ExpiresAt time.Time      `json:"expires_at"`
	MaxCalls  int            `json:"max_calls"`
	Tier      string         `json:"tier"`
	CreatedAt time.Time      `json:"created_at"`
}

type ContractConfig struct {
	LicenseNFTAddress    common.Address
	StakingNFTAddress    common.Address
	ReputationNFTAddress common.Address

	StartBlock uint64 // For event filtering
}

func NewClient(rpcURL string) (*Client, error) {
	// Connect via HTTP for general calls
	ethClient, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %v", err)
	}

	// Get chain ID
	chainID, err := ethClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}

	client := &Client{
		ethClient: ethClient,
		chainID:   chainID,
		rpcURL:    rpcURL,
	}

	log.Printf("Connected to Ethereum chain ID: %v", chainID)
	return client, nil
}

func (c *Client) InitializeContracts(config ContractConfig) error {
	var err error

	// Initialize License NFT contract
	c.licenseNFT, err = NewLicenseNFTContract(config.LicenseNFTAddress, c.ethClient)
	if err != nil {
		return fmt.Errorf("failed to initialize LicenseNFT contract: %v", err)
	}

	// Initialize Staking NFT contract
	c.stakingNFT, err = NewStakingNFTContract(config.StakingNFTAddress, c.ethClient)
	if err != nil {
		return fmt.Errorf("failed to initialize StakingNFT contract: %v", err)
	}

	log.Printf("Contracts initialized: LicenseNFT=%s, StakingNFT=%s",
		config.LicenseNFTAddress.Hex(), config.StakingNFTAddress.Hex())

	return nil
}

func NewStakingNFTContract(address common.Address, client *ethclient.Client) (*StakingNFTContract, error) {
	contract, err := Stake.NewStake(address, client)
	if err != nil {
		return nil, err
	}
	return &StakingNFTContract{
		address:       address,
		StakeContract: contract,
	}, nil
}

func NewLicenseNFTContract(address common.Address, client *ethclient.Client) (*LicenseNFTContract, error) {
	contract, err := license.NewLicense(address, client)
	if err != nil {
		return nil, err
	}
	return &LicenseNFTContract{
		Licensecontract: contract,
		address:         address,
	}, nil
}
func NewSkillToken(address common.Address, client *ethclient.Client) (*SkillToken, error) {
	contract, err := skill.NewSkill(address, client)
	if err != nil {
		return nil, err
	}
	return &SkillToken{
		SkillTokenContract: contract,
		address:            address,
	}, nil
}
func (c *Client) VerifyNFTOwnership(contractAddrStr string, tokenIDStr, userAddressStr string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Parse inputs
	tokenID, ok := new(big.Int).SetString(tokenIDStr, 10)
	if !ok {
		return false, fmt.Errorf("invalid token ID: %s", tokenIDStr)
	}
	contractAddr := common.HexToAddress(contractAddrStr)
	userAddress := common.HexToAddress(userAddressStr)

	// Determine which contract to use based on address
	switch contractAddr {
	case c.licenseNFT.Address():
		// Check ERC-1155 balance (License NFT)
		balance, err := c.licenseNFT.BalanceOf(&bind.CallOpts{Context: ctx}, userAddress, tokenID)
		if err != nil {
			return false, fmt.Errorf("failed to check license balance: %v", err)
		}
		return balance.Cmp(big.NewInt(0)) > 0, nil

	case c.stakingNFT.Address():
		// Check ERC-721 owner (Staking NFT)
		owner, err := c.stakingNFT.OwnerOf(&bind.CallOpts{Context: ctx}, tokenID)
		if err != nil {
			return false, fmt.Errorf("failed to get staking NFT owner: %v", err)
		}
		return owner == userAddress, nil

	default:
		return false, fmt.Errorf("unknown contract address: %s", contractAddr.Hex())
	}
}

func (c *Client) GetLicenseMetadata(licenseIDStr string) (*LicenseMetadata, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokenID, ok := new(big.Int).SetString(licenseIDStr, 10)
	if !ok {
		return nil, fmt.Errorf("invalid license ID: %s", licenseIDStr)
	}

	// Get token URI (metadata)
	_, err := c.licenseNFT.Uri(&bind.CallOpts{Context: ctx}, tokenID)
	if err != nil {
		return nil, fmt.Errorf("failed to get token URI: %v", err)
	}

	metadata := &LicenseMetadata{
		TokenID:   tokenID,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour), // 30 days from now
		MaxCalls:  1000,
		Tier:      "premium",
		CreatedAt: time.Now().Add(-7 * 24 * time.Hour), // Created 7 days ago
	}

	// Try to get owner
	owners, err := c.licenseNFT.BalanceOfBatch(
		&bind.CallOpts{Context: ctx},
		[]common.Address{c.licenseNFT.Address()},
		[]*big.Int{tokenID},
	)
	if err == nil && len(owners) > 0 && owners[0].Cmp(big.NewInt(0)) > 0 {
		// In real implementation, find which address has balance > 0
		metadata.Owner = common.HexToAddress("0xDemoOwner")
	}

	return metadata, nil
}

// IsLicenseValid checks if a user holds a license token for the given toolID
func (c *Client) IsLicenseValid(user common.Address, toolID *big.Int) (bool, error) {
	if c.licenseNFT == nil {
		return false, fmt.Errorf("license contract not initialized")
	}
	// Use BalanceOf to check ownership for demo
	valid, err := c.licenseNFT.Licensecontract.IsLicenseValid(&bind.CallOpts{Context: context.Background()}, user, toolID)
	if err != nil {
		return false, err
	}
	return valid, nil
}

func (c *Client) ListenForLicenseMinted(ctx context.Context, callback func(tokenID *big.Int, owner common.Address)) error {
	// Create filter query for LicenseMinted events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.licenseNFT.Address()},
		Topics:    [][]common.Hash{{c.licenseNFT.LicenseMintedSignature()}},
	}

	logs := make(chan types.Log)
	sub, err := c.ethClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil {
		return fmt.Errorf("failed to subscribe to logs: %v", err)
	}

	go func() {
		defer sub.Unsubscribe()

		for {
			select {
			case err := <-sub.Err():
				log.Printf("Subscription error: %v", err)
				return

			case vLog := <-logs:
				// Parse the event - in demo mode, extract basic info
				// In production, would decode the ABI-encoded data
				if len(vLog.Topics) > 0 {
					// Extract token ID from the log topics or data
					tokenID := big.NewInt(1)
					owner := common.HexToAddress("0xDemoOwner")
					callback(tokenID, owner)
				}

			case <-ctx.Done():
				return
			}
		}
	}()

	return nil
}

func (c *Client) Close() error {
	if c.ethClient != nil {
		c.ethClient.Close()
	}
	return nil
}

// Helper function to create transaction options
func (c *Client) NewTransactOpts(privateKeyHex string) (*bind.TransactOpts, error) {
	if privateKeyHex == "" {
		return nil, fmt.Errorf("private key required for transactions")
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("invalid private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, c.chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create transactor: %v", err)
	}

	// Set reasonable gas settings for hackathon demo
	auth.GasLimit = 300000
	auth.GasPrice = big.NewInt(20000000000) // 20 gwei

	return auth, nil
}
