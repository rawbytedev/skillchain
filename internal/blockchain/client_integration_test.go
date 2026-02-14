//go:build !ci

package blockchain

import (
	"context"
	//"encoding/json"
	"math/big"
	"os"
	"testing"
	"time"

	"log"

	"moltket/config"
	"moltket/internal/auth"
	"moltket/internal/cache" // your generated LicenseNFT binding

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system env")
	}
	os.Exit(m.Run())
}

func TestLiveLicenseNFTIntegration(t *testing.T) {
	t.Skip()
	// -------------------- CONNECT TO LOCAL HARDHAT NODE --------------------
	rpcURL := os.Getenv("ETH_NODE_URL")
	if rpcURL == "" {
		rpcURL = "http://localhost:8545"
	}

	client, err := ethclient.Dial(rpcURL)
	require.NoError(t, err, "failed to connect to Hardhat node")
	defer client.Close()

	chainID, err := client.ChainID(context.Background())
	require.NoError(t, err, "failed to get chain ID")

	// -------------------- LOAD CONTRACT ADDRESSES --------------------
	skillTokenAddrHex := os.Getenv("SKILL_TOKEN_ADDRESS")
	stakingAddrHex := os.Getenv("STAKING_NFT_ADDRESS")
	licenseAddrHex := os.Getenv("LICENSE_NFT_ADDRESS")

	require.NotEmpty(t, skillTokenAddrHex, "SKILL_TOKEN_ADDRESS not set")
	require.NotEmpty(t, stakingAddrHex, "STAKING_NFT_ADDRESS not set")
	require.NotEmpty(t, licenseAddrHex, "LICENSE_NFT_ADDRESS not set")

	skillTokenAddr := common.HexToAddress(skillTokenAddrHex)
	stakingAddr := common.HexToAddress(stakingAddrHex)
	licenseAddr := common.HexToAddress(licenseAddrHex)

	// Instantiate contract bindings
	skillToken, err := NewSkillToken(skillTokenAddr, client)
	require.NoError(t, err, "failed to instantiate SkillToken")

	stakingNFT, err := NewStakingNFTContract(stakingAddr, client)
	require.NoError(t, err, "failed to instantiate StakingNFT")

	licenseNFT, err := NewLicenseNFTContract(licenseAddr, client)
	require.NoError(t, err, "failed to instantiate LicenseNFT")

	// -------------------- SET UP BACKEND SIGNER (EIP-712) --------------------
	backendPrivateKeyHex := os.Getenv("SIGNER_PRIVATE_KEY")
	require.NotEmpty(t, backendPrivateKeyHex, "SIGNER_PRIVATE_KEY not set")

	signer, err := auth.NewSigner(backendPrivateKeyHex, chainID, licenseAddr)
	require.NoError(t, err)

	// -------------------- INITIALIZE BACKEND COMPONENTS --------------------
	ctx := context.Background()
	kvStore := cache.NewKVStore()
	defer kvStore.Close()

	bcClient := &Client{
		ethClient:  client,
		licenseNFT: licenseNFT, // only needed for IsLicenseValid
	}

	cfg := &config.Config{
		LicenseNFTAddress: licenseAddr.Hex(),
		SignatureNonce:    "integration-test-nonce",
		CacheTTL:          int(5 * time.Minute),
	}
	licenseService := NewLicenseService(cfg, kvStore, signer, bcClient)

	// -------------------- CREATE TEST ACCOUNTS --------------------
	// Developer: Hardhat account #0
	devPrivateKeyHex := "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"
	devPrivateKey, err := crypto.HexToECDSA(devPrivateKeyHex)
	require.NoError(t, err)
	devAddress := crypto.PubkeyToAddress(devPrivateKey.PublicKey)

	// User: Hardhat account #1
	userPrivateKeyHex := "5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a"
	userPrivateKey, err := crypto.HexToECDSA(userPrivateKeyHex)
	require.NoError(t, err)
	userAddress := crypto.PubkeyToAddress(userPrivateKey.PublicKey)

	// Create transactors
	devAuth, err := bind.NewKeyedTransactorWithChainID(devPrivateKey, chainID)
	require.NoError(t, err)
	devAuth.GasLimit = 3000000

	userAuth, err := bind.NewKeyedTransactorWithChainID(userPrivateKey, chainID)
	require.NoError(t, err)
	userAuth.GasLimit = 3000000

	// -------------------- STEP 1: DEVELOPER STAKES & LISTS A TOOL --------------------
	stakeAmount := new(big.Int).Mul(big.NewInt(10), big.NewInt(1e18)) // 10 SKILL
	// --- ENSURE TRUSTED SIGNER IS CORRECT ---
	currentSigner, err := licenseNFT.Licensecontract.TrustedSigner(&bind.CallOpts{})
	require.NoError(t, err)
	backendSignerAddr := signer.Address()
	t.Logf("Current trusted signer: %s", currentSigner.Hex())
	t.Logf("Backend signer address: %s", backendSignerAddr.Hex())

	if currentSigner != backendSignerAddr {
		t.Log("⚠️ Trusted signer mismatch  updating...")

		// Use Hardhat account #0 (deployer) – this is likely the contract owner
		ownerPrivateKeyHex := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
		ownerPrivateKey, err := crypto.HexToECDSA(ownerPrivateKeyHex)
		require.NoError(t, err)
		ownerAuth, err := bind.NewKeyedTransactorWithChainID(ownerPrivateKey, chainID)
		require.NoError(t, err)
		ownerAuth.GasLimit = 3000000

		tx, err := licenseNFT.Licensecontract.UpdateSigner(ownerAuth, backendSignerAddr)
		require.NoError(t, err)
		receipt, err := bind.WaitMined(ctx, client, tx)
		require.NoError(t, err)
		require.Equal(t, uint64(1), receipt.Status)
		t.Log("✅ Trusted signer updated successfully")
	}
	// Approve StakingNFT to spend SKILL
	tx, err := skillToken.SkillTokenContract.Approve(devAuth, stakingAddr, stakeAmount)
	require.NoError(t, err)
	t.Logf("Approved StakingNFT to spend SKILL, tx: %s", tx.Hash().Hex())

	// List the tool
	metadataURI := "ipfs://QmIntegrationTestTool"
	tx, err = stakingNFT.StakeContract.ListTool(devAuth, stakeAmount, metadataURI)
	require.NoError(t, err)
	t.Logf("Tool listed, tx: %s", tx.Hash().Hex())
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	require.NoError(t, err)
	require.Equal(t, uint64(1), receipt.Status, "ListTool failed")

	// Parse the third log (assuming it's ToolListed)
	toolListedEvent, err := stakingNFT.StakeContract.ParseToolListed(*receipt.Logs[3])
	require.NoError(t, err)
	toolID := toolListedEvent.TokenId // ← this is the REAL token ID
	t.Logf("✅ Tool listed with ID: %s", toolID.String())
	// Verify ownership
	owner, err := stakingNFT.StakeContract.OwnerOf(&bind.CallOpts{Context: ctx}, toolID)
	require.NoError(t, err)
	require.Equal(t, devAddress, owner, "developer should own the tool NFT")

	// -------------------- STEP 2: FREE TIER ACCESS --------------------
	for i := 0; i < 100; i++ {
		result, err := licenseService.VerifyAccess(ctx, userAddress, toolID)
		require.NoError(t, err)
		require.True(t, result.Valid)
		require.Equal(t, "free", result.Tier)
		require.Equal(t, 100-(i+1), result.CallsRemaining)
		require.NotEmpty(t, result.ProvenanceHash)
	}

	// 101st call: should fail (no license)
	result, err := licenseService.VerifyAccess(ctx, userAddress, toolID)
	require.NoError(t, err)
	require.False(t, result.Valid)
	require.Equal(t, "free tier exhausted and no valid license found", result.Reason)

	// -------------------- STEP 3: REQUEST LICENSE (OFF-CHAIN SIGNATURE) --------------------
	licenseReq := &LicenseRequest{
		UserAddress: userAddress,
		ToolID:      toolID,
	}
	licenseResp, err := licenseService.RequestLicense(ctx, licenseReq)
	require.NoError(t, err)
	require.NotNil(t, licenseResp)
	require.Equal(t, toolID, licenseResp.ToolID)
	require.Equal(t, userAddress, licenseResp.User)
	require.NotEmpty(t, licenseResp.SignatureR)
	require.NotEmpty(t, licenseResp.SignatureS)
	require.NotEmpty(t, licenseResp.SignatureV)

	// -------------------- STEP 4: MINT LICENSE ON-CHAIN --------------------
	expiresAt := licenseResp.ExpiresAt
	nonce := licenseResp.Nonce

	r := common.Hex2Bytes(licenseResp.SignatureR)
	s := common.Hex2Bytes(licenseResp.SignatureS)
	v := common.Hex2Bytes(licenseResp.SignatureV)
	signature := append(r, s...)
	signature = append(signature, v...)
	price, ok := new(big.Int).SetString(licenseResp.Price, 10)
	require.True(t, ok, "failed to parse price")
	userAuth.Value = price
	tx, err = licenseNFT.Licensecontract.MintLicense(
		userAuth,
		toolID,
		expiresAt,
		nonce,
		signature,
	)
	require.NoError(t, err)
	t.Logf("License minted, tx: %s", tx.Hash().Hex())

	// -------------------- STEP 5: BACKEND RECORDS THE MINTED LICENSE --------------------
	err = licenseService.RecordLicenseMinted(
		ctx,
		userAddress,
		toolID,
		expiresAt,
		nonce,
	)
	require.NoError(t, err)

	// -------------------- STEP 6: VERIFY ACCESS AFTER LICENSE --------------------
	result, err = licenseService.VerifyAccess(ctx, userAddress, toolID)
	require.NoError(t, err)
	require.True(t, result.Valid)
	require.Equal(t, "licensed", result.Tier)
	require.NotNil(t, result.ExpiresAt)
	require.Greater(t, result.CallsRemaining, 0)

	// -------------------- STEP 7: VERIFY ON-CHAIN STATE DIRECTLY --------------------
	isValid, err := licenseNFT.Licensecontract.IsLicenseValid(&bind.CallOpts{Context: ctx}, userAddress, toolID)
	require.NoError(t, err)
	require.True(t, isValid)

	t.Log("✅ Full integration test passed: Staking → Listing → Free Tier → License Mint → Verification")
}
