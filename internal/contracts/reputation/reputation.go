// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reputation

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

// ReputationMetaData contains all meta data concerning the Reputation contract.
var ReputationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_trustedSigner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ECDSAInvalidSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"ECDSAInvalidSignatureLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"ECDSAInvalidSignatureS\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toolId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"rootHash\",\"type\":\"bytes32\"}],\"name\":\"ReputationUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"backendSigner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_toolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getReputationRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reputationCheckpoints\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"toolNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_toolId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_rootHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"updateReputationRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ReputationABI is the input ABI used to generate the binding from.
// Deprecated: Use ReputationMetaData.ABI instead.
var ReputationABI = ReputationMetaData.ABI

// Reputation is an auto generated Go binding around an Ethereum contract.
type Reputation struct {
	ReputationCaller     // Read-only binding to the contract
	ReputationTransactor // Write-only binding to the contract
	ReputationFilterer   // Log filterer for contract events
}

// ReputationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReputationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReputationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReputationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReputationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReputationSession struct {
	Contract     *Reputation       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReputationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReputationCallerSession struct {
	Contract *ReputationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReputationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReputationTransactorSession struct {
	Contract     *ReputationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReputationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReputationRaw struct {
	Contract *Reputation // Generic contract binding to access the raw methods on
}

// ReputationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReputationCallerRaw struct {
	Contract *ReputationCaller // Generic read-only contract binding to access the raw methods on
}

// ReputationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReputationTransactorRaw struct {
	Contract *ReputationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReputation creates a new instance of Reputation, bound to a specific deployed contract.
func NewReputation(address common.Address, backend bind.ContractBackend) (*Reputation, error) {
	contract, err := bindReputation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reputation{ReputationCaller: ReputationCaller{contract: contract}, ReputationTransactor: ReputationTransactor{contract: contract}, ReputationFilterer: ReputationFilterer{contract: contract}}, nil
}

// NewReputationCaller creates a new read-only instance of Reputation, bound to a specific deployed contract.
func NewReputationCaller(address common.Address, caller bind.ContractCaller) (*ReputationCaller, error) {
	contract, err := bindReputation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationCaller{contract: contract}, nil
}

// NewReputationTransactor creates a new write-only instance of Reputation, bound to a specific deployed contract.
func NewReputationTransactor(address common.Address, transactor bind.ContractTransactor) (*ReputationTransactor, error) {
	contract, err := bindReputation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReputationTransactor{contract: contract}, nil
}

// NewReputationFilterer creates a new log filterer instance of Reputation, bound to a specific deployed contract.
func NewReputationFilterer(address common.Address, filterer bind.ContractFilterer) (*ReputationFilterer, error) {
	contract, err := bindReputation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReputationFilterer{contract: contract}, nil
}

// bindReputation binds a generic wrapper to an already deployed contract.
func bindReputation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReputationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputation *ReputationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputation.Contract.ReputationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputation *ReputationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.Contract.ReputationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputation *ReputationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputation.Contract.ReputationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reputation *ReputationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reputation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reputation *ReputationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reputation *ReputationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reputation.Contract.contract.Transact(opts, method, params...)
}

// BackendSigner is a free data retrieval call binding the contract method 0x65d65e86.
//
// Solidity: function backendSigner() view returns(address)
func (_Reputation *ReputationCaller) BackendSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "backendSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BackendSigner is a free data retrieval call binding the contract method 0x65d65e86.
//
// Solidity: function backendSigner() view returns(address)
func (_Reputation *ReputationSession) BackendSigner() (common.Address, error) {
	return _Reputation.Contract.BackendSigner(&_Reputation.CallOpts)
}

// BackendSigner is a free data retrieval call binding the contract method 0x65d65e86.
//
// Solidity: function backendSigner() view returns(address)
func (_Reputation *ReputationCallerSession) BackendSigner() (common.Address, error) {
	return _Reputation.Contract.BackendSigner(&_Reputation.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Reputation *ReputationCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "eip712Domain")

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
func (_Reputation *ReputationSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Reputation.Contract.Eip712Domain(&_Reputation.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Reputation *ReputationCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Reputation.Contract.Eip712Domain(&_Reputation.CallOpts)
}

// GetReputationRoot is a free data retrieval call binding the contract method 0x1c83119d.
//
// Solidity: function getReputationRoot(uint256 _toolId, uint256 _timestamp) view returns(bytes32)
func (_Reputation *ReputationCaller) GetReputationRoot(opts *bind.CallOpts, _toolId *big.Int, _timestamp *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "getReputationRoot", _toolId, _timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetReputationRoot is a free data retrieval call binding the contract method 0x1c83119d.
//
// Solidity: function getReputationRoot(uint256 _toolId, uint256 _timestamp) view returns(bytes32)
func (_Reputation *ReputationSession) GetReputationRoot(_toolId *big.Int, _timestamp *big.Int) ([32]byte, error) {
	return _Reputation.Contract.GetReputationRoot(&_Reputation.CallOpts, _toolId, _timestamp)
}

// GetReputationRoot is a free data retrieval call binding the contract method 0x1c83119d.
//
// Solidity: function getReputationRoot(uint256 _toolId, uint256 _timestamp) view returns(bytes32)
func (_Reputation *ReputationCallerSession) GetReputationRoot(_toolId *big.Int, _timestamp *big.Int) ([32]byte, error) {
	return _Reputation.Contract.GetReputationRoot(&_Reputation.CallOpts, _toolId, _timestamp)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputation *ReputationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputation *ReputationSession) Owner() (common.Address, error) {
	return _Reputation.Contract.Owner(&_Reputation.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reputation *ReputationCallerSession) Owner() (common.Address, error) {
	return _Reputation.Contract.Owner(&_Reputation.CallOpts)
}

// ReputationCheckpoints is a free data retrieval call binding the contract method 0xb3ce0885.
//
// Solidity: function reputationCheckpoints(uint256 , uint256 ) view returns(bytes32)
func (_Reputation *ReputationCaller) ReputationCheckpoints(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "reputationCheckpoints", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ReputationCheckpoints is a free data retrieval call binding the contract method 0xb3ce0885.
//
// Solidity: function reputationCheckpoints(uint256 , uint256 ) view returns(bytes32)
func (_Reputation *ReputationSession) ReputationCheckpoints(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _Reputation.Contract.ReputationCheckpoints(&_Reputation.CallOpts, arg0, arg1)
}

// ReputationCheckpoints is a free data retrieval call binding the contract method 0xb3ce0885.
//
// Solidity: function reputationCheckpoints(uint256 , uint256 ) view returns(bytes32)
func (_Reputation *ReputationCallerSession) ReputationCheckpoints(arg0 *big.Int, arg1 *big.Int) ([32]byte, error) {
	return _Reputation.Contract.ReputationCheckpoints(&_Reputation.CallOpts, arg0, arg1)
}

// ToolNonce is a free data retrieval call binding the contract method 0x090aa68c.
//
// Solidity: function toolNonce(uint256 ) view returns(uint256)
func (_Reputation *ReputationCaller) ToolNonce(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Reputation.contract.Call(opts, &out, "toolNonce", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToolNonce is a free data retrieval call binding the contract method 0x090aa68c.
//
// Solidity: function toolNonce(uint256 ) view returns(uint256)
func (_Reputation *ReputationSession) ToolNonce(arg0 *big.Int) (*big.Int, error) {
	return _Reputation.Contract.ToolNonce(&_Reputation.CallOpts, arg0)
}

// ToolNonce is a free data retrieval call binding the contract method 0x090aa68c.
//
// Solidity: function toolNonce(uint256 ) view returns(uint256)
func (_Reputation *ReputationCallerSession) ToolNonce(arg0 *big.Int) (*big.Int, error) {
	return _Reputation.Contract.ToolNonce(&_Reputation.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reputation *ReputationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reputation *ReputationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Reputation.Contract.RenounceOwnership(&_Reputation.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reputation *ReputationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Reputation.Contract.RenounceOwnership(&_Reputation.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reputation *ReputationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reputation *ReputationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Reputation.Contract.TransferOwnership(&_Reputation.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reputation *ReputationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Reputation.Contract.TransferOwnership(&_Reputation.TransactOpts, newOwner)
}

// UpdateReputationRoot is a paid mutator transaction binding the contract method 0x99f515c4.
//
// Solidity: function updateReputationRoot(uint256 _toolId, uint256 _timestamp, bytes32 _rootHash, uint256 _nonce, bytes _signature) returns()
func (_Reputation *ReputationTransactor) UpdateReputationRoot(opts *bind.TransactOpts, _toolId *big.Int, _timestamp *big.Int, _rootHash [32]byte, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Reputation.contract.Transact(opts, "updateReputationRoot", _toolId, _timestamp, _rootHash, _nonce, _signature)
}

// UpdateReputationRoot is a paid mutator transaction binding the contract method 0x99f515c4.
//
// Solidity: function updateReputationRoot(uint256 _toolId, uint256 _timestamp, bytes32 _rootHash, uint256 _nonce, bytes _signature) returns()
func (_Reputation *ReputationSession) UpdateReputationRoot(_toolId *big.Int, _timestamp *big.Int, _rootHash [32]byte, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Reputation.Contract.UpdateReputationRoot(&_Reputation.TransactOpts, _toolId, _timestamp, _rootHash, _nonce, _signature)
}

// UpdateReputationRoot is a paid mutator transaction binding the contract method 0x99f515c4.
//
// Solidity: function updateReputationRoot(uint256 _toolId, uint256 _timestamp, bytes32 _rootHash, uint256 _nonce, bytes _signature) returns()
func (_Reputation *ReputationTransactorSession) UpdateReputationRoot(_toolId *big.Int, _timestamp *big.Int, _rootHash [32]byte, _nonce *big.Int, _signature []byte) (*types.Transaction, error) {
	return _Reputation.Contract.UpdateReputationRoot(&_Reputation.TransactOpts, _toolId, _timestamp, _rootHash, _nonce, _signature)
}

// ReputationEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Reputation contract.
type ReputationEIP712DomainChangedIterator struct {
	Event *ReputationEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *ReputationEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationEIP712DomainChanged)
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
		it.Event = new(ReputationEIP712DomainChanged)
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
func (it *ReputationEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationEIP712DomainChanged represents a EIP712DomainChanged event raised by the Reputation contract.
type ReputationEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Reputation *ReputationFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*ReputationEIP712DomainChangedIterator, error) {

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &ReputationEIP712DomainChangedIterator{contract: _Reputation.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Reputation *ReputationFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *ReputationEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationEIP712DomainChanged)
				if err := _Reputation.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_Reputation *ReputationFilterer) ParseEIP712DomainChanged(log types.Log) (*ReputationEIP712DomainChanged, error) {
	event := new(ReputationEIP712DomainChanged)
	if err := _Reputation.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Reputation contract.
type ReputationOwnershipTransferredIterator struct {
	Event *ReputationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReputationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationOwnershipTransferred)
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
		it.Event = new(ReputationOwnershipTransferred)
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
func (it *ReputationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationOwnershipTransferred represents a OwnershipTransferred event raised by the Reputation contract.
type ReputationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Reputation *ReputationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReputationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReputationOwnershipTransferredIterator{contract: _Reputation.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Reputation *ReputationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReputationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationOwnershipTransferred)
				if err := _Reputation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Reputation *ReputationFilterer) ParseOwnershipTransferred(log types.Log) (*ReputationOwnershipTransferred, error) {
	event := new(ReputationOwnershipTransferred)
	if err := _Reputation.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReputationReputationUpdatedIterator is returned from FilterReputationUpdated and is used to iterate over the raw logs and unpacked data for ReputationUpdated events raised by the Reputation contract.
type ReputationReputationUpdatedIterator struct {
	Event *ReputationReputationUpdated // Event containing the contract specifics and raw log

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
func (it *ReputationReputationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReputationReputationUpdated)
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
		it.Event = new(ReputationReputationUpdated)
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
func (it *ReputationReputationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReputationReputationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReputationReputationUpdated represents a ReputationUpdated event raised by the Reputation contract.
type ReputationReputationUpdated struct {
	ToolId    *big.Int
	Timestamp *big.Int
	RootHash  [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReputationUpdated is a free log retrieval operation binding the contract event 0xa039401c6ee8c95ffccc82a701d2d9f667d9899e690c863d6dfa851c9da60165.
//
// Solidity: event ReputationUpdated(uint256 toolId, uint256 timestamp, bytes32 rootHash)
func (_Reputation *ReputationFilterer) FilterReputationUpdated(opts *bind.FilterOpts) (*ReputationReputationUpdatedIterator, error) {

	logs, sub, err := _Reputation.contract.FilterLogs(opts, "ReputationUpdated")
	if err != nil {
		return nil, err
	}
	return &ReputationReputationUpdatedIterator{contract: _Reputation.contract, event: "ReputationUpdated", logs: logs, sub: sub}, nil
}

// WatchReputationUpdated is a free log subscription operation binding the contract event 0xa039401c6ee8c95ffccc82a701d2d9f667d9899e690c863d6dfa851c9da60165.
//
// Solidity: event ReputationUpdated(uint256 toolId, uint256 timestamp, bytes32 rootHash)
func (_Reputation *ReputationFilterer) WatchReputationUpdated(opts *bind.WatchOpts, sink chan<- *ReputationReputationUpdated) (event.Subscription, error) {

	logs, sub, err := _Reputation.contract.WatchLogs(opts, "ReputationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReputationReputationUpdated)
				if err := _Reputation.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
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

// ParseReputationUpdated is a log parse operation binding the contract event 0xa039401c6ee8c95ffccc82a701d2d9f667d9899e690c863d6dfa851c9da60165.
//
// Solidity: event ReputationUpdated(uint256 toolId, uint256 timestamp, bytes32 rootHash)
func (_Reputation *ReputationFilterer) ParseReputationUpdated(log types.Log) (*ReputationReputationUpdated, error) {
	event := new(ReputationReputationUpdated)
	if err := _Reputation.contract.UnpackLog(event, "ReputationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
