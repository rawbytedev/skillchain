// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package license

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LicenseMetaData contains all meta data concerning the License contract.
var LicenseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedSigner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC1155InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idsLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valuesLength\",\"type\":\"uint256\"}],\"name\":\"ERC1155InvalidArrayLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC1155InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC1155MissingApprovalForAll\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePaid\",\"type\":\"uint256\"}],\"name\":\"LicenseMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toolId\",\"type\":\"uint256\"}],\"name\":\"LicenseRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSigner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSigner\",\"type\":\"address\"}],\"name\":\"SignerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_toolId\",\"type\":\"uint256\"}],\"name\":\"isLicenseValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"licenseExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_toolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_expiresAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"mintLicense\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_toolId\",\"type\":\"uint256\"}],\"name\":\"revokeLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"trustedSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newSigner\",\"type\":\"address\"}],\"name\":\"updateSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"usedNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LicenseABI is the input ABI used to generate the binding from.
// Deprecated: Use LicenseMetaData.ABI instead.
var LicenseABI = LicenseMetaData.ABI

// License is an auto generated Go binding around an Ethereum contract.
type License struct {
	LicenseCaller     // Read-only binding to the contract
	LicenseTransactor // Write-only binding to the contract
	LicenseFilterer   // Log filterer for contract events
}

// LicenseCaller is an auto generated read-only Go binding around an Ethereum contract.
type LicenseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LicenseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LicenseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LicenseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LicenseSession struct {
	Contract     *License          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LicenseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LicenseCallerSession struct {
	Contract *LicenseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LicenseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LicenseTransactorSession struct {
	Contract     *LicenseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LicenseRaw is an auto generated low-level Go binding around an Ethereum contract.
type LicenseRaw struct {
	Contract *License // Generic contract binding to access the raw methods on
}

// LicenseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LicenseCallerRaw struct {
	Contract *LicenseCaller // Generic read-only contract binding to access the raw methods on
}

// LicenseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LicenseTransactorRaw struct {
	Contract *LicenseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLicense creates a new instance of License, bound to a specific deployed contract.
func NewLicense(address common.Address, backend bind.ContractBackend) (*License, error) {
	contract, err := bindLicense(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &License{LicenseCaller: LicenseCaller{contract: contract}, LicenseTransactor: LicenseTransactor{contract: contract}, LicenseFilterer: LicenseFilterer{contract: contract}}, nil
}

// NewLicenseCaller creates a new read-only instance of License, bound to a specific deployed contract.
func NewLicenseCaller(address common.Address, caller bind.ContractCaller) (*LicenseCaller, error) {
	contract, err := bindLicense(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LicenseCaller{contract: contract}, nil
}

// NewLicenseTransactor creates a new write-only instance of License, bound to a specific deployed contract.
func NewLicenseTransactor(address common.Address, transactor bind.ContractTransactor) (*LicenseTransactor, error) {
	contract, err := bindLicense(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LicenseTransactor{contract: contract}, nil
}

// NewLicenseFilterer creates a new log filterer instance of License, bound to a specific deployed contract.
func NewLicenseFilterer(address common.Address, filterer bind.ContractFilterer) (*LicenseFilterer, error) {
	contract, err := bindLicense(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LicenseFilterer{contract: contract}, nil
}

// bindLicense binds a generic wrapper to an already deployed contract.
func bindLicense(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LicenseMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_License *LicenseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _License.Contract.LicenseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_License *LicenseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _License.Contract.LicenseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_License *LicenseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _License.Contract.LicenseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_License *LicenseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _License.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_License *LicenseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _License.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_License *LicenseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _License.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_License *LicenseCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_License *LicenseSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _License.Contract.BalanceOf(&_License.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_License *LicenseCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _License.Contract.BalanceOf(&_License.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_License *LicenseCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_License *LicenseSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _License.Contract.BalanceOfBatch(&_License.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_License *LicenseCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _License.Contract.BalanceOfBatch(&_License.CallOpts, accounts, ids)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_License *LicenseCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_License *LicenseSession) BaseURI() (string, error) {
	return _License.Contract.BaseURI(&_License.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_License *LicenseCallerSession) BaseURI() (string, error) {
	return _License.Contract.BaseURI(&_License.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_License *LicenseCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_License *LicenseSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _License.Contract.Eip712Domain(&_License.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_License *LicenseCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _License.Contract.Eip712Domain(&_License.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_License *LicenseCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_License *LicenseSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _License.Contract.IsApprovedForAll(&_License.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_License *LicenseCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _License.Contract.IsApprovedForAll(&_License.CallOpts, account, operator)
}

// IsLicenseValid is a free data retrieval call binding the contract method 0x197c4619.
//
// Solidity: function isLicenseValid(address _user, uint256 _toolId) view returns(bool)
func (_License *LicenseCaller) IsLicenseValid(opts *bind.CallOpts, _user common.Address, _toolId *big.Int) (bool, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "isLicenseValid", _user, _toolId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLicenseValid is a free data retrieval call binding the contract method 0x197c4619.
//
// Solidity: function isLicenseValid(address _user, uint256 _toolId) view returns(bool)
func (_License *LicenseSession) IsLicenseValid(_user common.Address, _toolId *big.Int) (bool, error) {
	return _License.Contract.IsLicenseValid(&_License.CallOpts, _user, _toolId)
}

// IsLicenseValid is a free data retrieval call binding the contract method 0x197c4619.
//
// Solidity: function isLicenseValid(address _user, uint256 _toolId) view returns(bool)
func (_License *LicenseCallerSession) IsLicenseValid(_user common.Address, _toolId *big.Int) (bool, error) {
	return _License.Contract.IsLicenseValid(&_License.CallOpts, _user, _toolId)
}

// LicenseExpiry is a free data retrieval call binding the contract method 0xd8305caf.
//
// Solidity: function licenseExpiry(address , uint256 ) view returns(uint256)
func (_License *LicenseCaller) LicenseExpiry(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "licenseExpiry", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LicenseExpiry is a free data retrieval call binding the contract method 0xd8305caf.
//
// Solidity: function licenseExpiry(address , uint256 ) view returns(uint256)
func (_License *LicenseSession) LicenseExpiry(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _License.Contract.LicenseExpiry(&_License.CallOpts, arg0, arg1)
}

// LicenseExpiry is a free data retrieval call binding the contract method 0xd8305caf.
//
// Solidity: function licenseExpiry(address , uint256 ) view returns(uint256)
func (_License *LicenseCallerSession) LicenseExpiry(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _License.Contract.LicenseExpiry(&_License.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_License *LicenseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_License *LicenseSession) Owner() (common.Address, error) {
	return _License.Contract.Owner(&_License.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_License *LicenseCallerSession) Owner() (common.Address, error) {
	return _License.Contract.Owner(&_License.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_License *LicenseCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_License *LicenseSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _License.Contract.SupportsInterface(&_License.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_License *LicenseCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _License.Contract.SupportsInterface(&_License.CallOpts, interfaceId)
}

// TrustedSigner is a free data retrieval call binding the contract method 0xf74d5480.
//
// Solidity: function trustedSigner() view returns(address)
func (_License *LicenseCaller) TrustedSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "trustedSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TrustedSigner is a free data retrieval call binding the contract method 0xf74d5480.
//
// Solidity: function trustedSigner() view returns(address)
func (_License *LicenseSession) TrustedSigner() (common.Address, error) {
	return _License.Contract.TrustedSigner(&_License.CallOpts)
}

// TrustedSigner is a free data retrieval call binding the contract method 0xf74d5480.
//
// Solidity: function trustedSigner() view returns(address)
func (_License *LicenseCallerSession) TrustedSigner() (common.Address, error) {
	return _License.Contract.TrustedSigner(&_License.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_License *LicenseCaller) Uri(opts *bind.CallOpts, _id *big.Int) (string, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "uri", _id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_License *LicenseSession) Uri(_id *big.Int) (string, error) {
	return _License.Contract.Uri(&_License.CallOpts, _id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 _id) view returns(string)
func (_License *LicenseCallerSession) Uri(_id *big.Int) (string, error) {
	return _License.Contract.Uri(&_License.CallOpts, _id)
}

// UsedNonces is a free data retrieval call binding the contract method 0x6a8a6894.
//
// Solidity: function usedNonces(address , uint256 ) view returns(bool)
func (_License *LicenseCaller) UsedNonces(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _License.contract.Call(opts, &out, "usedNonces", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedNonces is a free data retrieval call binding the contract method 0x6a8a6894.
//
// Solidity: function usedNonces(address , uint256 ) view returns(bool)
func (_License *LicenseSession) UsedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _License.Contract.UsedNonces(&_License.CallOpts, arg0, arg1)
}

// UsedNonces is a free data retrieval call binding the contract method 0x6a8a6894.
//
// Solidity: function usedNonces(address , uint256 ) view returns(bool)
func (_License *LicenseCallerSession) UsedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _License.Contract.UsedNonces(&_License.CallOpts, arg0, arg1)
}

// MintLicense is a paid mutator transaction binding the contract method 0x33386e84.
//
// Solidity: function mintLicense(uint256 _toolId, uint256 _expiresAt, uint256 _nonce, bytes _signature) payable returns()
func (_License *LicenseTransactor) MintLicense(opts *bind.TransactOpts, _toolId *big.Int, _expiresAt *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _License.contract.Transact(opts, "mintLicense", _toolId, _expiresAt, _nonce, _signature)
}

// MintLicense is a paid mutator transaction binding the contract method 0x33386e84.
//
// Solidity: function mintLicense(uint256 _toolId, uint256 _expiresAt, uint256 _nonce, bytes _signature) payable returns()
func (_License *LicenseSession) MintLicense(_toolId *big.Int, _expiresAt *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _License.Contract.MintLicense(&_License.TransactOpts, _toolId, _expiresAt, _nonce, _signature)
}

// MintLicense is a paid mutator transaction binding the contract method 0x33386e84.
//
// Solidity: function mintLicense(uint256 _toolId, uint256 _expiresAt, uint256 _nonce, bytes _signature) payable returns()
func (_License *LicenseTransactorSession) MintLicense(_toolId *big.Int, _expiresAt *big.Int, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _License.Contract.MintLicense(&_License.TransactOpts, _toolId, _expiresAt, _nonce, _signature)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_License *LicenseTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _License.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_License *LicenseSession) RenounceOwnership() (*types.Transaction, error) {
	return _License.Contract.RenounceOwnership(&_License.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_License *LicenseTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _License.Contract.RenounceOwnership(&_License.TransactOpts)
}

// RevokeLicense is a paid mutator transaction binding the contract method 0xd4961486.
//
// Solidity: function revokeLicense(address _user, uint256 _toolId) returns()
func (_License *LicenseTransactor) RevokeLicense(opts *bind.TransactOpts, _user common.Address, _toolId *big.Int) (*types.Transaction, error) {
	return _License.contract.Transact(opts, "revokeLicense", _user, _toolId)
}

// RevokeLicense is a paid mutator transaction binding the contract method 0xd4961486.
//
// Solidity: function revokeLicense(address _user, uint256 _toolId) returns()
func (_License *LicenseSession) RevokeLicense(_user common.Address, _toolId *big.Int) (*types.Transaction, error) {
	return _License.Contract.RevokeLicense(&_License.TransactOpts, _user, _toolId)
}

// RevokeLicense is a paid mutator transaction binding the contract method 0xd4961486.
//
// Solidity: function revokeLicense(address _user, uint256 _toolId) returns()
func (_License *LicenseTransactorSession) RevokeLicense(_user common.Address, _toolId *big.Int) (*types.Transaction, error) {
	return _License.Contract.RevokeLicense(&_License.TransactOpts, _user, _toolId)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_License *LicenseTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _License.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_License *LicenseSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _License.Contract.SafeBatchTransferFrom(&_License.TransactOpts, from, to, ids, values, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] values, bytes data) returns()
func (_License *LicenseTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _License.Contract.SafeBatchTransferFrom(&_License.TransactOpts, from, to, ids, values, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_License *LicenseTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _License.contract.Transact(opts, "safeTransferFrom", from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_License *LicenseSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _License.Contract.SafeTransferFrom(&_License.TransactOpts, from, to, id, value, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes data) returns()
func (_License *LicenseTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _License.Contract.SafeTransferFrom(&_License.TransactOpts, from, to, id, value, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_License *LicenseTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _License.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_License *LicenseSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _License.Contract.SetApprovalForAll(&_License.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_License *LicenseTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _License.Contract.SetApprovalForAll(&_License.TransactOpts, operator, approved)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_License *LicenseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _License.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_License *LicenseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _License.Contract.TransferOwnership(&_License.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_License *LicenseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _License.Contract.TransferOwnership(&_License.TransactOpts, newOwner)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address _newSigner) returns()
func (_License *LicenseTransactor) UpdateSigner(opts *bind.TransactOpts, _newSigner common.Address) (*types.Transaction, error) {
	return _License.contract.Transact(opts, "updateSigner", _newSigner)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address _newSigner) returns()
func (_License *LicenseSession) UpdateSigner(_newSigner common.Address) (*types.Transaction, error) {
	return _License.Contract.UpdateSigner(&_License.TransactOpts, _newSigner)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xa7ecd37e.
//
// Solidity: function updateSigner(address _newSigner) returns()
func (_License *LicenseTransactorSession) UpdateSigner(_newSigner common.Address) (*types.Transaction, error) {
	return _License.Contract.UpdateSigner(&_License.TransactOpts, _newSigner)
}

// LicenseApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the License contract.
type LicenseApprovalForAllIterator struct {
	Event *LicenseApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LicenseApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LicenseApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LicenseApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseApprovalForAll represents a ApprovalForAll event raised by the License contract.
type LicenseApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_License *LicenseFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*LicenseApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _License.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &LicenseApprovalForAllIterator{contract: _License.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_License *LicenseFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *LicenseApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _License.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseApprovalForAll)
				if err := _License.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_License *LicenseFilterer) ParseApprovalForAll(log types.Log) (*LicenseApprovalForAll, error) {
	event := new(LicenseApprovalForAll)
	if err := _License.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the License contract.
type LicenseEIP712DomainChangedIterator struct {
	Event *LicenseEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LicenseEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LicenseEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LicenseEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseEIP712DomainChanged represents a EIP712DomainChanged event raised by the License contract.
type LicenseEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_License *LicenseFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*LicenseEIP712DomainChangedIterator, error) {

	logs, sub, err := _License.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &LicenseEIP712DomainChangedIterator{contract: _License.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_License *LicenseFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *LicenseEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _License.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseEIP712DomainChanged)
				if err := _License.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_License *LicenseFilterer) ParseEIP712DomainChanged(log types.Log) (*LicenseEIP712DomainChanged, error) {
	event := new(LicenseEIP712DomainChanged)
	if err := _License.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseLicenseMintedIterator is returned from FilterLicenseMinted and is used to iterate over the raw logs and unpacked data for LicenseMinted events raised by the License contract.
type LicenseLicenseMintedIterator struct {
	Event *LicenseLicenseMinted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LicenseLicenseMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseLicenseMinted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LicenseLicenseMinted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LicenseLicenseMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseLicenseMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseLicenseMinted represents a LicenseMinted event raised by the License contract.
type LicenseLicenseMinted struct {
	User      common.Address
	ToolId    *big.Int
	ExpiresAt *big.Int
	PricePaid *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLicenseMinted is a free log retrieval operation binding the contract event 0x5f96e085fbb09368a41f4df54cf37ba4c1a4b06945d387741412dd300fd7fbae.
//
// Solidity: event LicenseMinted(address indexed user, uint256 toolId, uint256 expiresAt, uint256 pricePaid)
func (_License *LicenseFilterer) FilterLicenseMinted(opts *bind.FilterOpts, user []common.Address) (*LicenseLicenseMintedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _License.contract.FilterLogs(opts, "LicenseMinted", userRule)
	if err != nil {
		return nil, err
	}
	return &LicenseLicenseMintedIterator{contract: _License.contract, event: "LicenseMinted", logs: logs, sub: sub}, nil
}

// WatchLicenseMinted is a free log subscription operation binding the contract event 0x5f96e085fbb09368a41f4df54cf37ba4c1a4b06945d387741412dd300fd7fbae.
//
// Solidity: event LicenseMinted(address indexed user, uint256 toolId, uint256 expiresAt, uint256 pricePaid)
func (_License *LicenseFilterer) WatchLicenseMinted(opts *bind.WatchOpts, sink chan<- *LicenseLicenseMinted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _License.contract.WatchLogs(opts, "LicenseMinted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseLicenseMinted)
				if err := _License.contract.UnpackLog(event, "LicenseMinted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLicenseMinted is a log parse operation binding the contract event 0x5f96e085fbb09368a41f4df54cf37ba4c1a4b06945d387741412dd300fd7fbae.
//
// Solidity: event LicenseMinted(address indexed user, uint256 toolId, uint256 expiresAt, uint256 pricePaid)
func (_License *LicenseFilterer) ParseLicenseMinted(log types.Log) (*LicenseLicenseMinted, error) {
	event := new(LicenseLicenseMinted)
	if err := _License.contract.UnpackLog(event, "LicenseMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseLicenseRevokedIterator is returned from FilterLicenseRevoked and is used to iterate over the raw logs and unpacked data for LicenseRevoked events raised by the License contract.
type LicenseLicenseRevokedIterator struct {
	Event *LicenseLicenseRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LicenseLicenseRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseLicenseRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LicenseLicenseRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LicenseLicenseRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseLicenseRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseLicenseRevoked represents a LicenseRevoked event raised by the License contract.
type LicenseLicenseRevoked struct {
	User   common.Address
	ToolId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLicenseRevoked is a free log retrieval operation binding the contract event 0x839c38f52a1bd9ce7835e5b98999a47186ad03ada55f8f827814adb622949837.
//
// Solidity: event LicenseRevoked(address indexed user, uint256 toolId)
func (_License *LicenseFilterer) FilterLicenseRevoked(opts *bind.FilterOpts, user []common.Address) (*LicenseLicenseRevokedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _License.contract.FilterLogs(opts, "LicenseRevoked", userRule)
	if err != nil {
		return nil, err
	}
	return &LicenseLicenseRevokedIterator{contract: _License.contract, event: "LicenseRevoked", logs: logs, sub: sub}, nil
}

// WatchLicenseRevoked is a free log subscription operation binding the contract event 0x839c38f52a1bd9ce7835e5b98999a47186ad03ada55f8f827814adb622949837.
//
// Solidity: event LicenseRevoked(address indexed user, uint256 toolId)
func (_License *LicenseFilterer) WatchLicenseRevoked(opts *bind.WatchOpts, sink chan<- *LicenseLicenseRevoked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _License.contract.WatchLogs(opts, "LicenseRevoked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseLicenseRevoked)
				if err := _License.contract.UnpackLog(event, "LicenseRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLicenseRevoked is a log parse operation binding the contract event 0x839c38f52a1bd9ce7835e5b98999a47186ad03ada55f8f827814adb622949837.
//
// Solidity: event LicenseRevoked(address indexed user, uint256 toolId)
func (_License *LicenseFilterer) ParseLicenseRevoked(log types.Log) (*LicenseLicenseRevoked, error) {
	event := new(LicenseLicenseRevoked)
	if err := _License.contract.UnpackLog(event, "LicenseRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the License contract.
type LicenseOwnershipTransferredIterator struct {
	Event *LicenseOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LicenseOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LicenseOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LicenseOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseOwnershipTransferred represents a OwnershipTransferred event raised by the License contract.
type LicenseOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_License *LicenseFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LicenseOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _License.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LicenseOwnershipTransferredIterator{contract: _License.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_License *LicenseFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LicenseOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _License.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseOwnershipTransferred)
				if err := _License.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_License *LicenseFilterer) ParseOwnershipTransferred(log types.Log) (*LicenseOwnershipTransferred, error) {
	event := new(LicenseOwnershipTransferred)
	if err := _License.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseSignerChangedIterator is returned from FilterSignerChanged and is used to iterate over the raw logs and unpacked data for SignerChanged events raised by the License contract.
type LicenseSignerChangedIterator struct {
	Event *LicenseSignerChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LicenseSignerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseSignerChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LicenseSignerChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LicenseSignerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseSignerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseSignerChanged represents a SignerChanged event raised by the License contract.
type LicenseSignerChanged struct {
	OldSigner common.Address
	NewSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSignerChanged is a free log retrieval operation binding the contract event 0xeeb293e1f8f3a9db91ade748726387ed1352ca78f5430c5f06fe3d1e1ad50579.
//
// Solidity: event SignerChanged(address oldSigner, address newSigner)
func (_License *LicenseFilterer) FilterSignerChanged(opts *bind.FilterOpts) (*LicenseSignerChangedIterator, error) {

	logs, sub, err := _License.contract.FilterLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return &LicenseSignerChangedIterator{contract: _License.contract, event: "SignerChanged", logs: logs, sub: sub}, nil
}

// WatchSignerChanged is a free log subscription operation binding the contract event 0xeeb293e1f8f3a9db91ade748726387ed1352ca78f5430c5f06fe3d1e1ad50579.
//
// Solidity: event SignerChanged(address oldSigner, address newSigner)
func (_License *LicenseFilterer) WatchSignerChanged(opts *bind.WatchOpts, sink chan<- *LicenseSignerChanged) (event.Subscription, error) {

	logs, sub, err := _License.contract.WatchLogs(opts, "SignerChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseSignerChanged)
				if err := _License.contract.UnpackLog(event, "SignerChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSignerChanged is a log parse operation binding the contract event 0xeeb293e1f8f3a9db91ade748726387ed1352ca78f5430c5f06fe3d1e1ad50579.
//
// Solidity: event SignerChanged(address oldSigner, address newSigner)
func (_License *LicenseFilterer) ParseSignerChanged(log types.Log) (*LicenseSignerChanged, error) {
	event := new(LicenseSignerChanged)
	if err := _License.contract.UnpackLog(event, "SignerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the License contract.
type LicenseTransferBatchIterator struct {
	Event *LicenseTransferBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LicenseTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseTransferBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LicenseTransferBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LicenseTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseTransferBatch represents a TransferBatch event raised by the License contract.
type LicenseTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_License *LicenseFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*LicenseTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _License.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LicenseTransferBatchIterator{contract: _License.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_License *LicenseFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *LicenseTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _License.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseTransferBatch)
				if err := _License.contract.UnpackLog(event, "TransferBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_License *LicenseFilterer) ParseTransferBatch(log types.Log) (*LicenseTransferBatch, error) {
	event := new(LicenseTransferBatch)
	if err := _License.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the License contract.
type LicenseTransferSingleIterator struct {
	Event *LicenseTransferSingle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LicenseTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseTransferSingle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LicenseTransferSingle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LicenseTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseTransferSingle represents a TransferSingle event raised by the License contract.
type LicenseTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_License *LicenseFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*LicenseTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _License.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LicenseTransferSingleIterator{contract: _License.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_License *LicenseFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *LicenseTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _License.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseTransferSingle)
				if err := _License.contract.UnpackLog(event, "TransferSingle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_License *LicenseFilterer) ParseTransferSingle(log types.Log) (*LicenseTransferSingle, error) {
	event := new(LicenseTransferSingle)
	if err := _License.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LicenseURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the License contract.
type LicenseURIIterator struct {
	Event *LicenseURI // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LicenseURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LicenseURI)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LicenseURI)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LicenseURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LicenseURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LicenseURI represents a URI event raised by the License contract.
type LicenseURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_License *LicenseFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*LicenseURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _License.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &LicenseURIIterator{contract: _License.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_License *LicenseFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *LicenseURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _License.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LicenseURI)
				if err := _License.contract.UnpackLog(event, "URI", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_License *LicenseFilterer) ParseURI(log types.Log) (*LicenseURI, error) {
	event := new(LicenseURI)
	if err := _License.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
