package auth

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type EIP712Signer struct {
    privateKey    *ecdsa.PrivateKey
    publicAddress common.Address
    domain        apitypes.TypedDataDomain
    chainID       *big.Int
}

func NewSigner(privateKeyHex string, chainID *big.Int, verifyingContract common.Address) (*EIP712Signer, error) {
    privateKey, err := crypto.HexToECDSA(privateKeyHex)
    if err != nil {
        return nil, fmt.Errorf("invalid private key: %w", err)
    }
    
    publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        return nil, fmt.Errorf("error casting public key to ECDSA")
    }
    
    publicAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    
    return &EIP712Signer{
        privateKey:    privateKey,
        publicAddress: publicAddress,
        chainID:       chainID,
        domain: apitypes.TypedDataDomain{
            Name:              "SkillChainLicense",
            Version:           "1",
            ChainId:           math.NewHexOrDecimal256(chainID.Int64()),
            VerifyingContract: verifyingContract.Hex(),
        },
    }, nil
}

func (s *EIP712Signer) CreateLicenseSignature(
    user common.Address,
    toolId *big.Int,
    expiresAt *big.Int,
    nonce *big.Int,
) ([]byte, []byte, []byte, error) {
    
    types := apitypes.Types{
        "EIP712Domain": []apitypes.Type{
            {Name: "name", Type: "string"},
            {Name: "version", Type: "string"},
            {Name: "chainId", Type: "uint256"},
            {Name: "verifyingContract", Type: "address"},
        },
        "MintLicense": []apitypes.Type{
            {Name: "user", Type: "address"},
            {Name: "toolId", Type: "uint256"},
            {Name: "expiresAt", Type: "uint256"},
            {Name: "nonce", Type: "uint256"},
        },
    }
    
    message := apitypes.TypedDataMessage{
        "user":      user.Hex(),
        "toolId":    toolId.String(),
        "expiresAt": expiresAt.String(),
        "nonce":     nonce.String(),
    }
    
    typedData := apitypes.TypedData{
        Types:       types,
        PrimaryType: "MintLicense",
        Domain:      s.domain,
        Message:     message,
    }
    
    // Hash and sign the typed data
    domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
    if err != nil {
        return nil, nil, nil, fmt.Errorf("failed to hash domain: %w", err)
    }
    
    hash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
    if err != nil {
        return nil, nil, nil, fmt.Errorf("failed to hash message: %w", err)
    }
    
    rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(hash)))
    sighash := crypto.Keccak256(rawData)
    
    signature, err := crypto.Sign(sighash, s.privateKey)
    if err != nil {
        return nil, nil, nil, fmt.Errorf("failed to sign: %w", err)
    }
    
    // Split signature into r, s, v
    r := signature[:32]
    sigS := signature[32:64]
    sigV := signature[64:]

    // Ensure v is 27 or 28
    if sigV[0] < 27 {
        sigV[0] += 27
    }

    return r, sigS, sigV, nil
}

func (s *EIP712Signer) VerifySignature(
    user common.Address,
    toolId *big.Int,
    expiresAt *big.Int,
    nonce *big.Int,
    sigR, sigS, sigV []byte,
) (bool, error) {
    
    // Recreate the typed data
    types := apitypes.Types{
        "EIP712Domain": []apitypes.Type{
            {Name: "name", Type: "string"},
            {Name: "version", Type: "string"},
            {Name: "chainId", Type: "uint256"},
            {Name: "verifyingContract", Type: "address"},
        },
        "MintLicense": []apitypes.Type{
            {Name: "user", Type: "address"},
            {Name: "toolId", Type: "uint256"},
            {Name: "expiresAt", Type: "uint256"},
            {Name: "nonce", Type: "uint256"},
        },
    }
    
    message := apitypes.TypedDataMessage{
        "user":      user.Hex(),
        "toolId":    toolId.String(),
        "expiresAt": expiresAt.String(),
        "nonce":     nonce.String(),
    }
    
    typedData := apitypes.TypedData{
        Types:       types,
        PrimaryType: "MintLicense",
        Domain:      s.domain,
        Message:     message,
    }
    
    // Hash the typed data
    domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
    if err != nil {
        return false, fmt.Errorf("failed to hash domain: %w", err)
    }
    
    hash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
    if err != nil {
        return false, fmt.Errorf("failed to hash message: %w", err)
    }
    
    rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(hash)))
    sighash := crypto.Keccak256(rawData)
    
    // Recover the public key
    signature := append(sigR, sigS...)
    signature = append(signature, sigV...)
    
    // Adjust v if necessary
    if signature[64] >= 27 {
        signature[64] -= 27
    }
    
    pubkey, err := crypto.SigToPub(sighash, signature)
    if err != nil {
        return false, fmt.Errorf("failed to recover public key: %w", err)
    }
    
    recoveredAddr := crypto.PubkeyToAddress(*pubkey)
    return recoveredAddr == s.publicAddress, nil
}

func (s *EIP712Signer) Address() common.Address {
    return s.publicAddress
}

// GenerateTestPrivateKey is a test helper that returns a freshly generated private key.
// Provided for tests that expect a test key generation helper.
func GenerateTestPrivateKey(t interface{}, seed string) (*ecdsa.PrivateKey, error) {
    // For now ignore seed and generate a random key; deterministic seeding isn't required for these tests.
    return crypto.GenerateKey()
}