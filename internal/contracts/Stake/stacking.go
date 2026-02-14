// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Stake

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

// StakeMetaData contains all meta data concerning the Stake contract.
var StakeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_skillTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumStake\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SlashProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"slasher\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakeWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"}],\"name\":\"ToolDelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"ToolListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"SLASH_TIMELOCK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"delistTool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"}],\"name\":\"executeSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_tokenURI\",\"type\":\"string\"}],\"name\":\"listTool\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashAmount\",\"type\":\"uint256\"}],\"name\":\"proposeSlash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"skillToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashProposalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashProposalTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"toolCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakeABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeMetaData.ABI instead.
var StakeABI = StakeMetaData.ABI

// Stake is an auto generated Go binding around an Ethereum contract.
type Stake struct {
	StakeCaller     // Read-only binding to the contract
	StakeTransactor // Write-only binding to the contract
	StakeFilterer   // Log filterer for contract events
}

// StakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeSession struct {
	Contract     *Stake            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeCallerSession struct {
	Contract *StakeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTransactorSession struct {
	Contract     *StakeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeRaw struct {
	Contract *Stake // Generic contract binding to access the raw methods on
}

// StakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeCallerRaw struct {
	Contract *StakeCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTransactorRaw struct {
	Contract *StakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStake creates a new instance of Stake, bound to a specific deployed contract.
func NewStake(address common.Address, backend bind.ContractBackend) (*Stake, error) {
	contract, err := bindStake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stake{StakeCaller: StakeCaller{contract: contract}, StakeTransactor: StakeTransactor{contract: contract}, StakeFilterer: StakeFilterer{contract: contract}}, nil
}

// NewStakeCaller creates a new read-only instance of Stake, bound to a specific deployed contract.
func NewStakeCaller(address common.Address, caller bind.ContractCaller) (*StakeCaller, error) {
	contract, err := bindStake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeCaller{contract: contract}, nil
}

// NewStakeTransactor creates a new write-only instance of Stake, bound to a specific deployed contract.
func NewStakeTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTransactor, error) {
	contract, err := bindStake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTransactor{contract: contract}, nil
}

// NewStakeFilterer creates a new log filterer instance of Stake, bound to a specific deployed contract.
func NewStakeFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeFilterer, error) {
	contract, err := bindStake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeFilterer{contract: contract}, nil
}

// bindStake binds a generic wrapper to an already deployed contract.
func bindStake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.StakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.StakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stake *StakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stake *StakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stake *StakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stake.Contract.contract.Transact(opts, method, params...)
}

// SLASHTIMELOCK is a free data retrieval call binding the contract method 0xa1e692c5.
//
// Solidity: function SLASH_TIMELOCK() view returns(uint256)
func (_Stake *StakeCaller) SLASHTIMELOCK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "SLASH_TIMELOCK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SLASHTIMELOCK is a free data retrieval call binding the contract method 0xa1e692c5.
//
// Solidity: function SLASH_TIMELOCK() view returns(uint256)
func (_Stake *StakeSession) SLASHTIMELOCK() (*big.Int, error) {
	return _Stake.Contract.SLASHTIMELOCK(&_Stake.CallOpts)
}

// SLASHTIMELOCK is a free data retrieval call binding the contract method 0xa1e692c5.
//
// Solidity: function SLASH_TIMELOCK() view returns(uint256)
func (_Stake *StakeCallerSession) SLASHTIMELOCK() (*big.Int, error) {
	return _Stake.Contract.SLASHTIMELOCK(&_Stake.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Stake *StakeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Stake *StakeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Stake.Contract.BalanceOf(&_Stake.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Stake *StakeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Stake.Contract.BalanceOf(&_Stake.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Stake *StakeCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Stake *StakeSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Stake.Contract.GetApproved(&_Stake.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Stake *StakeCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Stake.Contract.GetApproved(&_Stake.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Stake *StakeCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Stake *StakeSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Stake.Contract.IsApprovedForAll(&_Stake.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Stake *StakeCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Stake.Contract.IsApprovedForAll(&_Stake.CallOpts, owner, operator)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_Stake *StakeCaller) MinimumStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "minimumStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_Stake *StakeSession) MinimumStake() (*big.Int, error) {
	return _Stake.Contract.MinimumStake(&_Stake.CallOpts)
}

// MinimumStake is a free data retrieval call binding the contract method 0xec5ffac2.
//
// Solidity: function minimumStake() view returns(uint256)
func (_Stake *StakeCallerSession) MinimumStake() (*big.Int, error) {
	return _Stake.Contract.MinimumStake(&_Stake.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stake *StakeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stake *StakeSession) Name() (string, error) {
	return _Stake.Contract.Name(&_Stake.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stake *StakeCallerSession) Name() (string, error) {
	return _Stake.Contract.Name(&_Stake.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Stake *StakeCaller) NextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "nextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Stake *StakeSession) NextTokenId() (*big.Int, error) {
	return _Stake.Contract.NextTokenId(&_Stake.CallOpts)
}

// NextTokenId is a free data retrieval call binding the contract method 0x75794a3c.
//
// Solidity: function nextTokenId() view returns(uint256)
func (_Stake *StakeCallerSession) NextTokenId() (*big.Int, error) {
	return _Stake.Contract.NextTokenId(&_Stake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stake *StakeCallerSession) Owner() (common.Address, error) {
	return _Stake.Contract.Owner(&_Stake.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Stake *StakeCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Stake *StakeSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Stake.Contract.OwnerOf(&_Stake.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Stake *StakeCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Stake.Contract.OwnerOf(&_Stake.CallOpts, tokenId)
}

// SkillToken is a free data retrieval call binding the contract method 0xd1ebdfd3.
//
// Solidity: function skillToken() view returns(address)
func (_Stake *StakeCaller) SkillToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "skillToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SkillToken is a free data retrieval call binding the contract method 0xd1ebdfd3.
//
// Solidity: function skillToken() view returns(address)
func (_Stake *StakeSession) SkillToken() (common.Address, error) {
	return _Stake.Contract.SkillToken(&_Stake.CallOpts)
}

// SkillToken is a free data retrieval call binding the contract method 0xd1ebdfd3.
//
// Solidity: function skillToken() view returns(address)
func (_Stake *StakeCallerSession) SkillToken() (common.Address, error) {
	return _Stake.Contract.SkillToken(&_Stake.CallOpts)
}

// SlashProposalAmount is a free data retrieval call binding the contract method 0x4d523d5c.
//
// Solidity: function slashProposalAmount(uint256 ) view returns(uint256)
func (_Stake *StakeCaller) SlashProposalAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "slashProposalAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashProposalAmount is a free data retrieval call binding the contract method 0x4d523d5c.
//
// Solidity: function slashProposalAmount(uint256 ) view returns(uint256)
func (_Stake *StakeSession) SlashProposalAmount(arg0 *big.Int) (*big.Int, error) {
	return _Stake.Contract.SlashProposalAmount(&_Stake.CallOpts, arg0)
}

// SlashProposalAmount is a free data retrieval call binding the contract method 0x4d523d5c.
//
// Solidity: function slashProposalAmount(uint256 ) view returns(uint256)
func (_Stake *StakeCallerSession) SlashProposalAmount(arg0 *big.Int) (*big.Int, error) {
	return _Stake.Contract.SlashProposalAmount(&_Stake.CallOpts, arg0)
}

// SlashProposalTime is a free data retrieval call binding the contract method 0x43c710d3.
//
// Solidity: function slashProposalTime(uint256 ) view returns(uint256)
func (_Stake *StakeCaller) SlashProposalTime(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "slashProposalTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashProposalTime is a free data retrieval call binding the contract method 0x43c710d3.
//
// Solidity: function slashProposalTime(uint256 ) view returns(uint256)
func (_Stake *StakeSession) SlashProposalTime(arg0 *big.Int) (*big.Int, error) {
	return _Stake.Contract.SlashProposalTime(&_Stake.CallOpts, arg0)
}

// SlashProposalTime is a free data retrieval call binding the contract method 0x43c710d3.
//
// Solidity: function slashProposalTime(uint256 ) view returns(uint256)
func (_Stake *StakeCallerSession) SlashProposalTime(arg0 *big.Int) (*big.Int, error) {
	return _Stake.Contract.SlashProposalTime(&_Stake.CallOpts, arg0)
}

// StakeAmount is a free data retrieval call binding the contract method 0x6c57f602.
//
// Solidity: function stakeAmount(uint256 ) view returns(uint256)
func (_Stake *StakeCaller) StakeAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "stakeAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x6c57f602.
//
// Solidity: function stakeAmount(uint256 ) view returns(uint256)
func (_Stake *StakeSession) StakeAmount(arg0 *big.Int) (*big.Int, error) {
	return _Stake.Contract.StakeAmount(&_Stake.CallOpts, arg0)
}

// StakeAmount is a free data retrieval call binding the contract method 0x6c57f602.
//
// Solidity: function stakeAmount(uint256 ) view returns(uint256)
func (_Stake *StakeCallerSession) StakeAmount(arg0 *big.Int) (*big.Int, error) {
	return _Stake.Contract.StakeAmount(&_Stake.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stake *StakeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stake *StakeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stake.Contract.SupportsInterface(&_Stake.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stake *StakeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stake.Contract.SupportsInterface(&_Stake.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stake *StakeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stake *StakeSession) Symbol() (string, error) {
	return _Stake.Contract.Symbol(&_Stake.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stake *StakeCallerSession) Symbol() (string, error) {
	return _Stake.Contract.Symbol(&_Stake.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Stake *StakeCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Stake *StakeSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Stake.Contract.TokenURI(&_Stake.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Stake *StakeCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Stake.Contract.TokenURI(&_Stake.CallOpts, tokenId)
}

// ToolCreator is a free data retrieval call binding the contract method 0xacd79be7.
//
// Solidity: function toolCreator(uint256 ) view returns(address)
func (_Stake *StakeCaller) ToolCreator(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stake.contract.Call(opts, &out, "toolCreator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToolCreator is a free data retrieval call binding the contract method 0xacd79be7.
//
// Solidity: function toolCreator(uint256 ) view returns(address)
func (_Stake *StakeSession) ToolCreator(arg0 *big.Int) (common.Address, error) {
	return _Stake.Contract.ToolCreator(&_Stake.CallOpts, arg0)
}

// ToolCreator is a free data retrieval call binding the contract method 0xacd79be7.
//
// Solidity: function toolCreator(uint256 ) view returns(address)
func (_Stake *StakeCallerSession) ToolCreator(arg0 *big.Int) (common.Address, error) {
	return _Stake.Contract.ToolCreator(&_Stake.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Stake *StakeTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Stake *StakeSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Approve(&_Stake.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Stake *StakeTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.Approve(&_Stake.TransactOpts, to, tokenId)
}

// DelistTool is a paid mutator transaction binding the contract method 0xf2e2d927.
//
// Solidity: function delistTool(uint256 _tokenId) returns()
func (_Stake *StakeTransactor) DelistTool(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "delistTool", _tokenId)
}

// DelistTool is a paid mutator transaction binding the contract method 0xf2e2d927.
//
// Solidity: function delistTool(uint256 _tokenId) returns()
func (_Stake *StakeSession) DelistTool(_tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DelistTool(&_Stake.TransactOpts, _tokenId)
}

// DelistTool is a paid mutator transaction binding the contract method 0xf2e2d927.
//
// Solidity: function delistTool(uint256 _tokenId) returns()
func (_Stake *StakeTransactorSession) DelistTool(_tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.DelistTool(&_Stake.TransactOpts, _tokenId)
}

// ExecuteSlash is a paid mutator transaction binding the contract method 0x1f769230.
//
// Solidity: function executeSlash(uint256 _tokenId, address _recipient) returns()
func (_Stake *StakeTransactor) ExecuteSlash(opts *bind.TransactOpts, _tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "executeSlash", _tokenId, _recipient)
}

// ExecuteSlash is a paid mutator transaction binding the contract method 0x1f769230.
//
// Solidity: function executeSlash(uint256 _tokenId, address _recipient) returns()
func (_Stake *StakeSession) ExecuteSlash(_tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Stake.Contract.ExecuteSlash(&_Stake.TransactOpts, _tokenId, _recipient)
}

// ExecuteSlash is a paid mutator transaction binding the contract method 0x1f769230.
//
// Solidity: function executeSlash(uint256 _tokenId, address _recipient) returns()
func (_Stake *StakeTransactorSession) ExecuteSlash(_tokenId *big.Int, _recipient common.Address) (*types.Transaction, error) {
	return _Stake.Contract.ExecuteSlash(&_Stake.TransactOpts, _tokenId, _recipient)
}

// ListTool is a paid mutator transaction binding the contract method 0xa5285b95.
//
// Solidity: function listTool(uint256 _stakeAmount, string _tokenURI) returns(uint256)
func (_Stake *StakeTransactor) ListTool(opts *bind.TransactOpts, _stakeAmount *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "listTool", _stakeAmount, _tokenURI)
}

// ListTool is a paid mutator transaction binding the contract method 0xa5285b95.
//
// Solidity: function listTool(uint256 _stakeAmount, string _tokenURI) returns(uint256)
func (_Stake *StakeSession) ListTool(_stakeAmount *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Stake.Contract.ListTool(&_Stake.TransactOpts, _stakeAmount, _tokenURI)
}

// ListTool is a paid mutator transaction binding the contract method 0xa5285b95.
//
// Solidity: function listTool(uint256 _stakeAmount, string _tokenURI) returns(uint256)
func (_Stake *StakeTransactorSession) ListTool(_stakeAmount *big.Int, _tokenURI string) (*types.Transaction, error) {
	return _Stake.Contract.ListTool(&_Stake.TransactOpts, _stakeAmount, _tokenURI)
}

// ProposeSlash is a paid mutator transaction binding the contract method 0x298bd532.
//
// Solidity: function proposeSlash(uint256 _tokenId, uint256 _slashAmount) returns()
func (_Stake *StakeTransactor) ProposeSlash(opts *bind.TransactOpts, _tokenId *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "proposeSlash", _tokenId, _slashAmount)
}

// ProposeSlash is a paid mutator transaction binding the contract method 0x298bd532.
//
// Solidity: function proposeSlash(uint256 _tokenId, uint256 _slashAmount) returns()
func (_Stake *StakeSession) ProposeSlash(_tokenId *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.ProposeSlash(&_Stake.TransactOpts, _tokenId, _slashAmount)
}

// ProposeSlash is a paid mutator transaction binding the contract method 0x298bd532.
//
// Solidity: function proposeSlash(uint256 _tokenId, uint256 _slashAmount) returns()
func (_Stake *StakeTransactorSession) ProposeSlash(_tokenId *big.Int, _slashAmount *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.ProposeSlash(&_Stake.TransactOpts, _tokenId, _slashAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stake.Contract.RenounceOwnership(&_Stake.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stake *StakeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stake.Contract.RenounceOwnership(&_Stake.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Stake *StakeTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Stake *StakeSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SafeTransferFrom(&_Stake.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Stake *StakeTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.SafeTransferFrom(&_Stake.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Stake *StakeTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Stake *StakeSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Stake.Contract.SafeTransferFrom0(&_Stake.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Stake *StakeTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Stake.Contract.SafeTransferFrom0(&_Stake.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Stake *StakeTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Stake *StakeSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Stake.Contract.SetApprovalForAll(&_Stake.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Stake *StakeTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Stake.Contract.SetApprovalForAll(&_Stake.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Stake *StakeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Stake *StakeSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.TransferFrom(&_Stake.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Stake *StakeTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Stake.Contract.TransferFrom(&_Stake.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stake *StakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stake.Contract.TransferOwnership(&_Stake.TransactOpts, newOwner)
}

// StakeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Stake contract.
type StakeApprovalIterator struct {
	Event *StakeApproval // Event containing the contract specifics and raw log

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
func (it *StakeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeApproval)
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
		it.Event = new(StakeApproval)
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
func (it *StakeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeApproval represents a Approval event raised by the Stake contract.
type StakeApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Stake *StakeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*StakeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeApprovalIterator{contract: _Stake.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Stake *StakeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakeApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeApproval)
				if err := _Stake.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Stake *StakeFilterer) ParseApproval(log types.Log) (*StakeApproval, error) {
	event := new(StakeApproval)
	if err := _Stake.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Stake contract.
type StakeApprovalForAllIterator struct {
	Event *StakeApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StakeApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeApprovalForAll)
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
		it.Event = new(StakeApprovalForAll)
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
func (it *StakeApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeApprovalForAll represents a ApprovalForAll event raised by the Stake contract.
type StakeApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Stake *StakeFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StakeApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakeApprovalForAllIterator{contract: _Stake.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Stake *StakeFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StakeApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeApprovalForAll)
				if err := _Stake.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Stake *StakeFilterer) ParseApprovalForAll(log types.Log) (*StakeApprovalForAll, error) {
	event := new(StakeApprovalForAll)
	if err := _Stake.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Stake contract.
type StakeBatchMetadataUpdateIterator struct {
	Event *StakeBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *StakeBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeBatchMetadataUpdate)
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
		it.Event = new(StakeBatchMetadataUpdate)
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
func (it *StakeBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Stake contract.
type StakeBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Stake *StakeFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*StakeBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &StakeBatchMetadataUpdateIterator{contract: _Stake.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Stake *StakeFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *StakeBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeBatchMetadataUpdate)
				if err := _Stake.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Stake *StakeFilterer) ParseBatchMetadataUpdate(log types.Log) (*StakeBatchMetadataUpdate, error) {
	event := new(StakeBatchMetadataUpdate)
	if err := _Stake.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Stake contract.
type StakeMetadataUpdateIterator struct {
	Event *StakeMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *StakeMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeMetadataUpdate)
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
		it.Event = new(StakeMetadataUpdate)
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
func (it *StakeMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeMetadataUpdate represents a MetadataUpdate event raised by the Stake contract.
type StakeMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Stake *StakeFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*StakeMetadataUpdateIterator, error) {

	logs, sub, err := _Stake.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &StakeMetadataUpdateIterator{contract: _Stake.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Stake *StakeFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *StakeMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Stake.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeMetadataUpdate)
				if err := _Stake.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Stake *StakeFilterer) ParseMetadataUpdate(log types.Log) (*StakeMetadataUpdate, error) {
	event := new(StakeMetadataUpdate)
	if err := _Stake.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stake contract.
type StakeOwnershipTransferredIterator struct {
	Event *StakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeOwnershipTransferred)
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
		it.Event = new(StakeOwnershipTransferred)
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
func (it *StakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeOwnershipTransferred represents a OwnershipTransferred event raised by the Stake contract.
type StakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakeOwnershipTransferredIterator{contract: _Stake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stake *StakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeOwnershipTransferred)
				if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Stake *StakeFilterer) ParseOwnershipTransferred(log types.Log) (*StakeOwnershipTransferred, error) {
	event := new(StakeOwnershipTransferred)
	if err := _Stake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeSlashProposedIterator is returned from FilterSlashProposed and is used to iterate over the raw logs and unpacked data for SlashProposed events raised by the Stake contract.
type StakeSlashProposedIterator struct {
	Event *StakeSlashProposed // Event containing the contract specifics and raw log

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
func (it *StakeSlashProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeSlashProposed)
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
		it.Event = new(StakeSlashProposed)
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
func (it *StakeSlashProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeSlashProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeSlashProposed represents a SlashProposed event raised by the Stake contract.
type StakeSlashProposed struct {
	TokenId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlashProposed is a free log retrieval operation binding the contract event 0x813a947408c790b7b53be43749748d2162bc2d4269785624dce654993f93bacd.
//
// Solidity: event SlashProposed(uint256 indexed tokenId, uint256 amount)
func (_Stake *StakeFilterer) FilterSlashProposed(opts *bind.FilterOpts, tokenId []*big.Int) (*StakeSlashProposedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "SlashProposed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeSlashProposedIterator{contract: _Stake.contract, event: "SlashProposed", logs: logs, sub: sub}, nil
}

// WatchSlashProposed is a free log subscription operation binding the contract event 0x813a947408c790b7b53be43749748d2162bc2d4269785624dce654993f93bacd.
//
// Solidity: event SlashProposed(uint256 indexed tokenId, uint256 amount)
func (_Stake *StakeFilterer) WatchSlashProposed(opts *bind.WatchOpts, sink chan<- *StakeSlashProposed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "SlashProposed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeSlashProposed)
				if err := _Stake.contract.UnpackLog(event, "SlashProposed", log); err != nil {
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

// ParseSlashProposed is a log parse operation binding the contract event 0x813a947408c790b7b53be43749748d2162bc2d4269785624dce654993f93bacd.
//
// Solidity: event SlashProposed(uint256 indexed tokenId, uint256 amount)
func (_Stake *StakeFilterer) ParseSlashProposed(log types.Log) (*StakeSlashProposed, error) {
	event := new(StakeSlashProposed)
	if err := _Stake.contract.UnpackLog(event, "SlashProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeStakeSlashedIterator is returned from FilterStakeSlashed and is used to iterate over the raw logs and unpacked data for StakeSlashed events raised by the Stake contract.
type StakeStakeSlashedIterator struct {
	Event *StakeStakeSlashed // Event containing the contract specifics and raw log

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
func (it *StakeStakeSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStakeSlashed)
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
		it.Event = new(StakeStakeSlashed)
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
func (it *StakeStakeSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakeSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStakeSlashed represents a StakeSlashed event raised by the Stake contract.
type StakeStakeSlashed struct {
	TokenId *big.Int
	Slasher common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeSlashed is a free log retrieval operation binding the contract event 0xba00be99b4ba100b19f2c080160ab22c2b8f4d62027d6524c532e7578d7a43b4.
//
// Solidity: event StakeSlashed(uint256 indexed tokenId, address slasher, uint256 amount)
func (_Stake *StakeFilterer) FilterStakeSlashed(opts *bind.FilterOpts, tokenId []*big.Int) (*StakeStakeSlashedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "StakeSlashed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeStakeSlashedIterator{contract: _Stake.contract, event: "StakeSlashed", logs: logs, sub: sub}, nil
}

// WatchStakeSlashed is a free log subscription operation binding the contract event 0xba00be99b4ba100b19f2c080160ab22c2b8f4d62027d6524c532e7578d7a43b4.
//
// Solidity: event StakeSlashed(uint256 indexed tokenId, address slasher, uint256 amount)
func (_Stake *StakeFilterer) WatchStakeSlashed(opts *bind.WatchOpts, sink chan<- *StakeStakeSlashed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "StakeSlashed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStakeSlashed)
				if err := _Stake.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
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

// ParseStakeSlashed is a log parse operation binding the contract event 0xba00be99b4ba100b19f2c080160ab22c2b8f4d62027d6524c532e7578d7a43b4.
//
// Solidity: event StakeSlashed(uint256 indexed tokenId, address slasher, uint256 amount)
func (_Stake *StakeFilterer) ParseStakeSlashed(log types.Log) (*StakeStakeSlashed, error) {
	event := new(StakeStakeSlashed)
	if err := _Stake.contract.UnpackLog(event, "StakeSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeStakeWithdrawnIterator is returned from FilterStakeWithdrawn and is used to iterate over the raw logs and unpacked data for StakeWithdrawn events raised by the Stake contract.
type StakeStakeWithdrawnIterator struct {
	Event *StakeStakeWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakeStakeWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeStakeWithdrawn)
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
		it.Event = new(StakeStakeWithdrawn)
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
func (it *StakeStakeWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeStakeWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeStakeWithdrawn represents a StakeWithdrawn event raised by the Stake contract.
type StakeStakeWithdrawn struct {
	TokenId *big.Int
	Creator common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakeWithdrawn is a free log retrieval operation binding the contract event 0x459d5aef2c549903c1eeb1736f5728845d5ccf82537ddd5bf8035795eee89263.
//
// Solidity: event StakeWithdrawn(uint256 indexed tokenId, address creator, uint256 amount)
func (_Stake *StakeFilterer) FilterStakeWithdrawn(opts *bind.FilterOpts, tokenId []*big.Int) (*StakeStakeWithdrawnIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "StakeWithdrawn", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeStakeWithdrawnIterator{contract: _Stake.contract, event: "StakeWithdrawn", logs: logs, sub: sub}, nil
}

// WatchStakeWithdrawn is a free log subscription operation binding the contract event 0x459d5aef2c549903c1eeb1736f5728845d5ccf82537ddd5bf8035795eee89263.
//
// Solidity: event StakeWithdrawn(uint256 indexed tokenId, address creator, uint256 amount)
func (_Stake *StakeFilterer) WatchStakeWithdrawn(opts *bind.WatchOpts, sink chan<- *StakeStakeWithdrawn, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "StakeWithdrawn", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeStakeWithdrawn)
				if err := _Stake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
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

// ParseStakeWithdrawn is a log parse operation binding the contract event 0x459d5aef2c549903c1eeb1736f5728845d5ccf82537ddd5bf8035795eee89263.
//
// Solidity: event StakeWithdrawn(uint256 indexed tokenId, address creator, uint256 amount)
func (_Stake *StakeFilterer) ParseStakeWithdrawn(log types.Log) (*StakeStakeWithdrawn, error) {
	event := new(StakeStakeWithdrawn)
	if err := _Stake.contract.UnpackLog(event, "StakeWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeToolDelistedIterator is returned from FilterToolDelisted and is used to iterate over the raw logs and unpacked data for ToolDelisted events raised by the Stake contract.
type StakeToolDelistedIterator struct {
	Event *StakeToolDelisted // Event containing the contract specifics and raw log

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
func (it *StakeToolDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeToolDelisted)
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
		it.Event = new(StakeToolDelisted)
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
func (it *StakeToolDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeToolDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeToolDelisted represents a ToolDelisted event raised by the Stake contract.
type StakeToolDelisted struct {
	TokenId *big.Int
	Creator common.Address
	Stake   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterToolDelisted is a free log retrieval operation binding the contract event 0xb954ba9d67083923d9e72da3afa0a96a04499aa5831fb5ae21ebe704211d3e77.
//
// Solidity: event ToolDelisted(uint256 indexed tokenId, address creator, uint256 stake)
func (_Stake *StakeFilterer) FilterToolDelisted(opts *bind.FilterOpts, tokenId []*big.Int) (*StakeToolDelistedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "ToolDelisted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeToolDelistedIterator{contract: _Stake.contract, event: "ToolDelisted", logs: logs, sub: sub}, nil
}

// WatchToolDelisted is a free log subscription operation binding the contract event 0xb954ba9d67083923d9e72da3afa0a96a04499aa5831fb5ae21ebe704211d3e77.
//
// Solidity: event ToolDelisted(uint256 indexed tokenId, address creator, uint256 stake)
func (_Stake *StakeFilterer) WatchToolDelisted(opts *bind.WatchOpts, sink chan<- *StakeToolDelisted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "ToolDelisted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeToolDelisted)
				if err := _Stake.contract.UnpackLog(event, "ToolDelisted", log); err != nil {
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

// ParseToolDelisted is a log parse operation binding the contract event 0xb954ba9d67083923d9e72da3afa0a96a04499aa5831fb5ae21ebe704211d3e77.
//
// Solidity: event ToolDelisted(uint256 indexed tokenId, address creator, uint256 stake)
func (_Stake *StakeFilterer) ParseToolDelisted(log types.Log) (*StakeToolDelisted, error) {
	event := new(StakeToolDelisted)
	if err := _Stake.contract.UnpackLog(event, "ToolDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeToolListedIterator is returned from FilterToolListed and is used to iterate over the raw logs and unpacked data for ToolListed events raised by the Stake contract.
type StakeToolListedIterator struct {
	Event *StakeToolListed // Event containing the contract specifics and raw log

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
func (it *StakeToolListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeToolListed)
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
		it.Event = new(StakeToolListed)
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
func (it *StakeToolListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeToolListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeToolListed represents a ToolListed event raised by the Stake contract.
type StakeToolListed struct {
	TokenId     *big.Int
	Creator     common.Address
	Stake       *big.Int
	MetadataURI string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterToolListed is a free log retrieval operation binding the contract event 0x998c9b1b9d0323ef228dda772effbc9e97561f6d73cb191faa28cc7d39e7f86f.
//
// Solidity: event ToolListed(uint256 indexed tokenId, address creator, uint256 stake, string metadataURI)
func (_Stake *StakeFilterer) FilterToolListed(opts *bind.FilterOpts, tokenId []*big.Int) (*StakeToolListedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "ToolListed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeToolListedIterator{contract: _Stake.contract, event: "ToolListed", logs: logs, sub: sub}, nil
}

// WatchToolListed is a free log subscription operation binding the contract event 0x998c9b1b9d0323ef228dda772effbc9e97561f6d73cb191faa28cc7d39e7f86f.
//
// Solidity: event ToolListed(uint256 indexed tokenId, address creator, uint256 stake, string metadataURI)
func (_Stake *StakeFilterer) WatchToolListed(opts *bind.WatchOpts, sink chan<- *StakeToolListed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "ToolListed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeToolListed)
				if err := _Stake.contract.UnpackLog(event, "ToolListed", log); err != nil {
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

// ParseToolListed is a log parse operation binding the contract event 0x998c9b1b9d0323ef228dda772effbc9e97561f6d73cb191faa28cc7d39e7f86f.
//
// Solidity: event ToolListed(uint256 indexed tokenId, address creator, uint256 stake, string metadataURI)
func (_Stake *StakeFilterer) ParseToolListed(log types.Log) (*StakeToolListed, error) {
	event := new(StakeToolListed)
	if err := _Stake.contract.UnpackLog(event, "ToolListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Stake contract.
type StakeTransferIterator struct {
	Event *StakeTransfer // Event containing the contract specifics and raw log

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
func (it *StakeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTransfer)
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
		it.Event = new(StakeTransfer)
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
func (it *StakeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTransfer represents a Transfer event raised by the Stake contract.
type StakeTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Stake *StakeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*StakeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeTransferIterator{contract: _Stake.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Stake *StakeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakeTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Stake.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTransfer)
				if err := _Stake.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Stake *StakeFilterer) ParseTransfer(log types.Log) (*StakeTransfer, error) {
	event := new(StakeTransfer)
	if err := _Stake.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
