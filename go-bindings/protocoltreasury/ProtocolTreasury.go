// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package protocoltreasury

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

// ProtocolTreasuryMetaData contains all meta data concerning the ProtocolTreasury contract.
var ProtocolTreasuryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"name\":\"ERC20TokensWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnAmount\",\"type\":\"uint256\"}],\"name\":\"NativeCurrencyWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr_\",\"type\":\"address\"}],\"name\":\"getSelfERC20TokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSelfNativeBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAllWithdraw_\",\"type\":\"bool\"}],\"name\":\"withdrawERC20Tokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToWithdraw_\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAllWithdraw_\",\"type\":\"bool\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// ProtocolTreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtocolTreasuryMetaData.ABI instead.
var ProtocolTreasuryABI = ProtocolTreasuryMetaData.ABI

// ProtocolTreasury is an auto generated Go binding around an Ethereum contract.
type ProtocolTreasury struct {
	ProtocolTreasuryCaller     // Read-only binding to the contract
	ProtocolTreasuryTransactor // Write-only binding to the contract
	ProtocolTreasuryFilterer   // Log filterer for contract events
}

// ProtocolTreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProtocolTreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProtocolTreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProtocolTreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProtocolTreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProtocolTreasurySession struct {
	Contract     *ProtocolTreasury // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProtocolTreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProtocolTreasuryCallerSession struct {
	Contract *ProtocolTreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ProtocolTreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProtocolTreasuryTransactorSession struct {
	Contract     *ProtocolTreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ProtocolTreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProtocolTreasuryRaw struct {
	Contract *ProtocolTreasury // Generic contract binding to access the raw methods on
}

// ProtocolTreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProtocolTreasuryCallerRaw struct {
	Contract *ProtocolTreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// ProtocolTreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProtocolTreasuryTransactorRaw struct {
	Contract *ProtocolTreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProtocolTreasury creates a new instance of ProtocolTreasury, bound to a specific deployed contract.
func NewProtocolTreasury(address common.Address, backend bind.ContractBackend) (*ProtocolTreasury, error) {
	contract, err := bindProtocolTreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasury{ProtocolTreasuryCaller: ProtocolTreasuryCaller{contract: contract}, ProtocolTreasuryTransactor: ProtocolTreasuryTransactor{contract: contract}, ProtocolTreasuryFilterer: ProtocolTreasuryFilterer{contract: contract}}, nil
}

// NewProtocolTreasuryCaller creates a new read-only instance of ProtocolTreasury, bound to a specific deployed contract.
func NewProtocolTreasuryCaller(address common.Address, caller bind.ContractCaller) (*ProtocolTreasuryCaller, error) {
	contract, err := bindProtocolTreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryCaller{contract: contract}, nil
}

// NewProtocolTreasuryTransactor creates a new write-only instance of ProtocolTreasury, bound to a specific deployed contract.
func NewProtocolTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*ProtocolTreasuryTransactor, error) {
	contract, err := bindProtocolTreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryTransactor{contract: contract}, nil
}

// NewProtocolTreasuryFilterer creates a new log filterer instance of ProtocolTreasury, bound to a specific deployed contract.
func NewProtocolTreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*ProtocolTreasuryFilterer, error) {
	contract, err := bindProtocolTreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryFilterer{contract: contract}, nil
}

// bindProtocolTreasury binds a generic wrapper to an already deployed contract.
func bindProtocolTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProtocolTreasuryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolTreasury *ProtocolTreasuryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolTreasury.Contract.ProtocolTreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolTreasury *ProtocolTreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.ProtocolTreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolTreasury *ProtocolTreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.ProtocolTreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProtocolTreasury *ProtocolTreasuryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProtocolTreasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProtocolTreasury *ProtocolTreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProtocolTreasury *ProtocolTreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.contract.Transact(opts, method, params...)
}

// GetSelfERC20TokenBalance is a free data retrieval call binding the contract method 0xfc88a027.
//
// Solidity: function getSelfERC20TokenBalance(address tokenAddr_) view returns(uint256)
func (_ProtocolTreasury *ProtocolTreasuryCaller) GetSelfERC20TokenBalance(opts *bind.CallOpts, tokenAddr_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolTreasury.contract.Call(opts, &out, "getSelfERC20TokenBalance", tokenAddr_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSelfERC20TokenBalance is a free data retrieval call binding the contract method 0xfc88a027.
//
// Solidity: function getSelfERC20TokenBalance(address tokenAddr_) view returns(uint256)
func (_ProtocolTreasury *ProtocolTreasurySession) GetSelfERC20TokenBalance(tokenAddr_ common.Address) (*big.Int, error) {
	return _ProtocolTreasury.Contract.GetSelfERC20TokenBalance(&_ProtocolTreasury.CallOpts, tokenAddr_)
}

// GetSelfERC20TokenBalance is a free data retrieval call binding the contract method 0xfc88a027.
//
// Solidity: function getSelfERC20TokenBalance(address tokenAddr_) view returns(uint256)
func (_ProtocolTreasury *ProtocolTreasuryCallerSession) GetSelfERC20TokenBalance(tokenAddr_ common.Address) (*big.Int, error) {
	return _ProtocolTreasury.Contract.GetSelfERC20TokenBalance(&_ProtocolTreasury.CallOpts, tokenAddr_)
}

// GetSelfNativeBalance is a free data retrieval call binding the contract method 0xcaa2f7a0.
//
// Solidity: function getSelfNativeBalance() view returns(uint256)
func (_ProtocolTreasury *ProtocolTreasuryCaller) GetSelfNativeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProtocolTreasury.contract.Call(opts, &out, "getSelfNativeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSelfNativeBalance is a free data retrieval call binding the contract method 0xcaa2f7a0.
//
// Solidity: function getSelfNativeBalance() view returns(uint256)
func (_ProtocolTreasury *ProtocolTreasurySession) GetSelfNativeBalance() (*big.Int, error) {
	return _ProtocolTreasury.Contract.GetSelfNativeBalance(&_ProtocolTreasury.CallOpts)
}

// GetSelfNativeBalance is a free data retrieval call binding the contract method 0xcaa2f7a0.
//
// Solidity: function getSelfNativeBalance() view returns(uint256)
func (_ProtocolTreasury *ProtocolTreasuryCallerSession) GetSelfNativeBalance() (*big.Int, error) {
	return _ProtocolTreasury.Contract.GetSelfNativeBalance(&_ProtocolTreasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProtocolTreasury *ProtocolTreasuryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProtocolTreasury.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProtocolTreasury *ProtocolTreasurySession) Owner() (common.Address, error) {
	return _ProtocolTreasury.Contract.Owner(&_ProtocolTreasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ProtocolTreasury *ProtocolTreasuryCallerSession) Owner() (common.Address, error) {
	return _ProtocolTreasury.Contract.Owner(&_ProtocolTreasury.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProtocolTreasury *ProtocolTreasuryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProtocolTreasury.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProtocolTreasury *ProtocolTreasurySession) ProxiableUUID() ([32]byte, error) {
	return _ProtocolTreasury.Contract.ProxiableUUID(&_ProtocolTreasury.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ProtocolTreasury *ProtocolTreasuryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ProtocolTreasury.Contract.ProxiableUUID(&_ProtocolTreasury.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolTreasury.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ProtocolTreasury *ProtocolTreasurySession) Initialize() (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.Initialize(&_ProtocolTreasury.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactorSession) Initialize() (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.Initialize(&_ProtocolTreasury.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolTreasury.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProtocolTreasury *ProtocolTreasurySession) RenounceOwnership() (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.RenounceOwnership(&_ProtocolTreasury.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.RenounceOwnership(&_ProtocolTreasury.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ProtocolTreasury.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProtocolTreasury *ProtocolTreasurySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.TransferOwnership(&_ProtocolTreasury.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.TransferOwnership(&_ProtocolTreasury.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ProtocolTreasury.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ProtocolTreasury *ProtocolTreasurySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.UpgradeTo(&_ProtocolTreasury.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.UpgradeTo(&_ProtocolTreasury.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProtocolTreasury.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProtocolTreasury *ProtocolTreasurySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.UpgradeToAndCall(&_ProtocolTreasury.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.UpgradeToAndCall(&_ProtocolTreasury.TransactOpts, newImplementation, data)
}

// WithdrawERC20Tokens is a paid mutator transaction binding the contract method 0xc499bee6.
//
// Solidity: function withdrawERC20Tokens(address tokenAddr_, address recipient_, uint256 amountToWithdraw_, bool isAllWithdraw_) returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactor) WithdrawERC20Tokens(opts *bind.TransactOpts, tokenAddr_ common.Address, recipient_ common.Address, amountToWithdraw_ *big.Int, isAllWithdraw_ bool) (*types.Transaction, error) {
	return _ProtocolTreasury.contract.Transact(opts, "withdrawERC20Tokens", tokenAddr_, recipient_, amountToWithdraw_, isAllWithdraw_)
}

// WithdrawERC20Tokens is a paid mutator transaction binding the contract method 0xc499bee6.
//
// Solidity: function withdrawERC20Tokens(address tokenAddr_, address recipient_, uint256 amountToWithdraw_, bool isAllWithdraw_) returns()
func (_ProtocolTreasury *ProtocolTreasurySession) WithdrawERC20Tokens(tokenAddr_ common.Address, recipient_ common.Address, amountToWithdraw_ *big.Int, isAllWithdraw_ bool) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.WithdrawERC20Tokens(&_ProtocolTreasury.TransactOpts, tokenAddr_, recipient_, amountToWithdraw_, isAllWithdraw_)
}

// WithdrawERC20Tokens is a paid mutator transaction binding the contract method 0xc499bee6.
//
// Solidity: function withdrawERC20Tokens(address tokenAddr_, address recipient_, uint256 amountToWithdraw_, bool isAllWithdraw_) returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactorSession) WithdrawERC20Tokens(tokenAddr_ common.Address, recipient_ common.Address, amountToWithdraw_ *big.Int, isAllWithdraw_ bool) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.WithdrawERC20Tokens(&_ProtocolTreasury.TransactOpts, tokenAddr_, recipient_, amountToWithdraw_, isAllWithdraw_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0xa5662f85.
//
// Solidity: function withdrawNative(address recipient_, uint256 amountToWithdraw_, bool isAllWithdraw_) returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactor) WithdrawNative(opts *bind.TransactOpts, recipient_ common.Address, amountToWithdraw_ *big.Int, isAllWithdraw_ bool) (*types.Transaction, error) {
	return _ProtocolTreasury.contract.Transact(opts, "withdrawNative", recipient_, amountToWithdraw_, isAllWithdraw_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0xa5662f85.
//
// Solidity: function withdrawNative(address recipient_, uint256 amountToWithdraw_, bool isAllWithdraw_) returns()
func (_ProtocolTreasury *ProtocolTreasurySession) WithdrawNative(recipient_ common.Address, amountToWithdraw_ *big.Int, isAllWithdraw_ bool) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.WithdrawNative(&_ProtocolTreasury.TransactOpts, recipient_, amountToWithdraw_, isAllWithdraw_)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0xa5662f85.
//
// Solidity: function withdrawNative(address recipient_, uint256 amountToWithdraw_, bool isAllWithdraw_) returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactorSession) WithdrawNative(recipient_ common.Address, amountToWithdraw_ *big.Int, isAllWithdraw_ bool) (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.WithdrawNative(&_ProtocolTreasury.TransactOpts, recipient_, amountToWithdraw_, isAllWithdraw_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProtocolTreasury.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProtocolTreasury *ProtocolTreasurySession) Receive() (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.Receive(&_ProtocolTreasury.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ProtocolTreasury *ProtocolTreasuryTransactorSession) Receive() (*types.Transaction, error) {
	return _ProtocolTreasury.Contract.Receive(&_ProtocolTreasury.TransactOpts)
}

// ProtocolTreasuryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ProtocolTreasury contract.
type ProtocolTreasuryAdminChangedIterator struct {
	Event *ProtocolTreasuryAdminChanged // Event containing the contract specifics and raw log

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
func (it *ProtocolTreasuryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolTreasuryAdminChanged)
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
		it.Event = new(ProtocolTreasuryAdminChanged)
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
func (it *ProtocolTreasuryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolTreasuryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolTreasuryAdminChanged represents a AdminChanged event raised by the ProtocolTreasury contract.
type ProtocolTreasuryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ProtocolTreasuryAdminChangedIterator, error) {

	logs, sub, err := _ProtocolTreasury.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryAdminChangedIterator{contract: _ProtocolTreasury.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ProtocolTreasuryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ProtocolTreasury.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolTreasuryAdminChanged)
				if err := _ProtocolTreasury.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) ParseAdminChanged(log types.Log) (*ProtocolTreasuryAdminChanged, error) {
	event := new(ProtocolTreasuryAdminChanged)
	if err := _ProtocolTreasury.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolTreasuryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ProtocolTreasury contract.
type ProtocolTreasuryBeaconUpgradedIterator struct {
	Event *ProtocolTreasuryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ProtocolTreasuryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolTreasuryBeaconUpgraded)
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
		it.Event = new(ProtocolTreasuryBeaconUpgraded)
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
func (it *ProtocolTreasuryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolTreasuryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolTreasuryBeaconUpgraded represents a BeaconUpgraded event raised by the ProtocolTreasury contract.
type ProtocolTreasuryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ProtocolTreasuryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryBeaconUpgradedIterator{contract: _ProtocolTreasury.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ProtocolTreasuryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolTreasuryBeaconUpgraded)
				if err := _ProtocolTreasury.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) ParseBeaconUpgraded(log types.Log) (*ProtocolTreasuryBeaconUpgraded, error) {
	event := new(ProtocolTreasuryBeaconUpgraded)
	if err := _ProtocolTreasury.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolTreasuryERC20TokensWithdrawnIterator is returned from FilterERC20TokensWithdrawn and is used to iterate over the raw logs and unpacked data for ERC20TokensWithdrawn events raised by the ProtocolTreasury contract.
type ProtocolTreasuryERC20TokensWithdrawnIterator struct {
	Event *ProtocolTreasuryERC20TokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *ProtocolTreasuryERC20TokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolTreasuryERC20TokensWithdrawn)
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
		it.Event = new(ProtocolTreasuryERC20TokensWithdrawn)
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
func (it *ProtocolTreasuryERC20TokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolTreasuryERC20TokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolTreasuryERC20TokensWithdrawn represents a ERC20TokensWithdrawn event raised by the ProtocolTreasury contract.
type ProtocolTreasuryERC20TokensWithdrawn struct {
	TokenAddr       common.Address
	Recipient       common.Address
	WithdrawnAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterERC20TokensWithdrawn is a free log retrieval operation binding the contract event 0x00e763f7778b8ceef7270c89b7d1df1008b0e482da39c43831417733af96fb0d.
//
// Solidity: event ERC20TokensWithdrawn(address indexed tokenAddr, address indexed recipient, uint256 withdrawnAmount)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) FilterERC20TokensWithdrawn(opts *bind.FilterOpts, tokenAddr []common.Address, recipient []common.Address) (*ProtocolTreasuryERC20TokensWithdrawnIterator, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.FilterLogs(opts, "ERC20TokensWithdrawn", tokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryERC20TokensWithdrawnIterator{contract: _ProtocolTreasury.contract, event: "ERC20TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchERC20TokensWithdrawn is a free log subscription operation binding the contract event 0x00e763f7778b8ceef7270c89b7d1df1008b0e482da39c43831417733af96fb0d.
//
// Solidity: event ERC20TokensWithdrawn(address indexed tokenAddr, address indexed recipient, uint256 withdrawnAmount)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) WatchERC20TokensWithdrawn(opts *bind.WatchOpts, sink chan<- *ProtocolTreasuryERC20TokensWithdrawn, tokenAddr []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenAddrRule []interface{}
	for _, tokenAddrItem := range tokenAddr {
		tokenAddrRule = append(tokenAddrRule, tokenAddrItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.WatchLogs(opts, "ERC20TokensWithdrawn", tokenAddrRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolTreasuryERC20TokensWithdrawn)
				if err := _ProtocolTreasury.contract.UnpackLog(event, "ERC20TokensWithdrawn", log); err != nil {
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

// ParseERC20TokensWithdrawn is a log parse operation binding the contract event 0x00e763f7778b8ceef7270c89b7d1df1008b0e482da39c43831417733af96fb0d.
//
// Solidity: event ERC20TokensWithdrawn(address indexed tokenAddr, address indexed recipient, uint256 withdrawnAmount)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) ParseERC20TokensWithdrawn(log types.Log) (*ProtocolTreasuryERC20TokensWithdrawn, error) {
	event := new(ProtocolTreasuryERC20TokensWithdrawn)
	if err := _ProtocolTreasury.contract.UnpackLog(event, "ERC20TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolTreasuryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ProtocolTreasury contract.
type ProtocolTreasuryInitializedIterator struct {
	Event *ProtocolTreasuryInitialized // Event containing the contract specifics and raw log

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
func (it *ProtocolTreasuryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolTreasuryInitialized)
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
		it.Event = new(ProtocolTreasuryInitialized)
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
func (it *ProtocolTreasuryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolTreasuryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolTreasuryInitialized represents a Initialized event raised by the ProtocolTreasury contract.
type ProtocolTreasuryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProtocolTreasuryInitializedIterator, error) {

	logs, sub, err := _ProtocolTreasury.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryInitializedIterator{contract: _ProtocolTreasury.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProtocolTreasuryInitialized) (event.Subscription, error) {

	logs, sub, err := _ProtocolTreasury.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolTreasuryInitialized)
				if err := _ProtocolTreasury.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) ParseInitialized(log types.Log) (*ProtocolTreasuryInitialized, error) {
	event := new(ProtocolTreasuryInitialized)
	if err := _ProtocolTreasury.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolTreasuryNativeCurrencyWithdrawnIterator is returned from FilterNativeCurrencyWithdrawn and is used to iterate over the raw logs and unpacked data for NativeCurrencyWithdrawn events raised by the ProtocolTreasury contract.
type ProtocolTreasuryNativeCurrencyWithdrawnIterator struct {
	Event *ProtocolTreasuryNativeCurrencyWithdrawn // Event containing the contract specifics and raw log

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
func (it *ProtocolTreasuryNativeCurrencyWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolTreasuryNativeCurrencyWithdrawn)
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
		it.Event = new(ProtocolTreasuryNativeCurrencyWithdrawn)
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
func (it *ProtocolTreasuryNativeCurrencyWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolTreasuryNativeCurrencyWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolTreasuryNativeCurrencyWithdrawn represents a NativeCurrencyWithdrawn event raised by the ProtocolTreasury contract.
type ProtocolTreasuryNativeCurrencyWithdrawn struct {
	Recipient       common.Address
	WithdrawnAmount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNativeCurrencyWithdrawn is a free log retrieval operation binding the contract event 0x5317b81e67f8cd60e038f9dc090c51c9401b1fbae411b50e2524794180844835.
//
// Solidity: event NativeCurrencyWithdrawn(address indexed recipient, uint256 withdrawnAmount)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) FilterNativeCurrencyWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*ProtocolTreasuryNativeCurrencyWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.FilterLogs(opts, "NativeCurrencyWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryNativeCurrencyWithdrawnIterator{contract: _ProtocolTreasury.contract, event: "NativeCurrencyWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNativeCurrencyWithdrawn is a free log subscription operation binding the contract event 0x5317b81e67f8cd60e038f9dc090c51c9401b1fbae411b50e2524794180844835.
//
// Solidity: event NativeCurrencyWithdrawn(address indexed recipient, uint256 withdrawnAmount)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) WatchNativeCurrencyWithdrawn(opts *bind.WatchOpts, sink chan<- *ProtocolTreasuryNativeCurrencyWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.WatchLogs(opts, "NativeCurrencyWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolTreasuryNativeCurrencyWithdrawn)
				if err := _ProtocolTreasury.contract.UnpackLog(event, "NativeCurrencyWithdrawn", log); err != nil {
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

// ParseNativeCurrencyWithdrawn is a log parse operation binding the contract event 0x5317b81e67f8cd60e038f9dc090c51c9401b1fbae411b50e2524794180844835.
//
// Solidity: event NativeCurrencyWithdrawn(address indexed recipient, uint256 withdrawnAmount)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) ParseNativeCurrencyWithdrawn(log types.Log) (*ProtocolTreasuryNativeCurrencyWithdrawn, error) {
	event := new(ProtocolTreasuryNativeCurrencyWithdrawn)
	if err := _ProtocolTreasury.contract.UnpackLog(event, "NativeCurrencyWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolTreasuryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ProtocolTreasury contract.
type ProtocolTreasuryOwnershipTransferredIterator struct {
	Event *ProtocolTreasuryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ProtocolTreasuryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolTreasuryOwnershipTransferred)
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
		it.Event = new(ProtocolTreasuryOwnershipTransferred)
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
func (it *ProtocolTreasuryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolTreasuryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolTreasuryOwnershipTransferred represents a OwnershipTransferred event raised by the ProtocolTreasury contract.
type ProtocolTreasuryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ProtocolTreasuryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryOwnershipTransferredIterator{contract: _ProtocolTreasury.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ProtocolTreasuryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolTreasuryOwnershipTransferred)
				if err := _ProtocolTreasury.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ProtocolTreasury *ProtocolTreasuryFilterer) ParseOwnershipTransferred(log types.Log) (*ProtocolTreasuryOwnershipTransferred, error) {
	event := new(ProtocolTreasuryOwnershipTransferred)
	if err := _ProtocolTreasury.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProtocolTreasuryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ProtocolTreasury contract.
type ProtocolTreasuryUpgradedIterator struct {
	Event *ProtocolTreasuryUpgraded // Event containing the contract specifics and raw log

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
func (it *ProtocolTreasuryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProtocolTreasuryUpgraded)
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
		it.Event = new(ProtocolTreasuryUpgraded)
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
func (it *ProtocolTreasuryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProtocolTreasuryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProtocolTreasuryUpgraded represents a Upgraded event raised by the ProtocolTreasury contract.
type ProtocolTreasuryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ProtocolTreasuryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ProtocolTreasuryUpgradedIterator{contract: _ProtocolTreasury.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ProtocolTreasuryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ProtocolTreasury.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProtocolTreasuryUpgraded)
				if err := _ProtocolTreasury.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ProtocolTreasury *ProtocolTreasuryFilterer) ParseUpgraded(log types.Log) (*ProtocolTreasuryUpgraded, error) {
	event := new(ProtocolTreasuryUpgraded)
	if err := _ProtocolTreasury.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
