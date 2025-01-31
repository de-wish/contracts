// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wishlistfactory

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

// IWishlistFactoryProtocolFeeSettings is an auto generated low-level Go binding around an user-defined struct.
type IWishlistFactoryProtocolFeeSettings struct {
	ProtocolFeePercentage *big.Int
	MaxProtocolFeeAmount  *big.Int
	ProtocolFeeRecipient  common.Address
}

// WishlistFactoryMetaData contains all meta data concerning the WishlistFactory contract.
var WishlistFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"CreateWishlistSignatureExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCreateWishlistSignature\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"}],\"name\":\"InvalidProtocolFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wishlistId\",\"type\":\"uint256\"}],\"name\":\"WishlistAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addrName\",\"type\":\"string\"}],\"name\":\"ZeroAddr\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"wishlistOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wishlistId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"wishlistAddr\",\"type\":\"address\"}],\"name\":\"NewWishlistCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProtocolFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxProtocolFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProtocolFeeRecipient\",\"type\":\"address\"}],\"name\":\"ProtocolFeeSettingsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newProtocolSignerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldProtocolSignerAddr\",\"type\":\"address\"}],\"name\":\"ProtocolSignerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wishlistOwner_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"wishlistId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sigDeadline_\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"initialItemPrices_\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"signature_\",\"type\":\"bytes\"}],\"name\":\"createWishlist\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"newWishlistAddr_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeeSettings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocolFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProtocolFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"}],\"internalType\":\"structIWishlistFactory.ProtocolFeeSettings\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wishlistId_\",\"type\":\"uint256\"}],\"name\":\"getWishlistAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to_\",\"type\":\"uint256\"}],\"name\":\"getWishlistAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"wishlistAddresses_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWishlistsTotalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wishlistImpl_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"usdcTokenAddr_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"protocolSignerAddr_\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocolFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProtocolFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"}],\"internalType\":\"structIWishlistFactory.ProtocolFeeSettings\",\"name\":\"feeSettings_\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolSignerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxyBeacon\",\"outputs\":[{\"internalType\":\"contractProxyBeacon\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"protocolFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxProtocolFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"protocolFeeRecipient\",\"type\":\"address\"}],\"internalType\":\"structIWishlistFactory.ProtocolFeeSettings\",\"name\":\"feeSettings_\",\"type\":\"tuple\"}],\"name\":\"setProtocolFeeSettings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newProtocolSignerAddr_\",\"type\":\"address\"}],\"name\":\"setProtocolSignerAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newWishlistsImpl_\",\"type\":\"address\"}],\"name\":\"upgradeWishlistsImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wishlistId_\",\"type\":\"uint256\"}],\"name\":\"wishlistExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WishlistFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use WishlistFactoryMetaData.ABI instead.
var WishlistFactoryABI = WishlistFactoryMetaData.ABI

// WishlistFactory is an auto generated Go binding around an Ethereum contract.
type WishlistFactory struct {
	WishlistFactoryCaller     // Read-only binding to the contract
	WishlistFactoryTransactor // Write-only binding to the contract
	WishlistFactoryFilterer   // Log filterer for contract events
}

// WishlistFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type WishlistFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WishlistFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WishlistFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WishlistFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WishlistFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WishlistFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WishlistFactorySession struct {
	Contract     *WishlistFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WishlistFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WishlistFactoryCallerSession struct {
	Contract *WishlistFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WishlistFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WishlistFactoryTransactorSession struct {
	Contract     *WishlistFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WishlistFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type WishlistFactoryRaw struct {
	Contract *WishlistFactory // Generic contract binding to access the raw methods on
}

// WishlistFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WishlistFactoryCallerRaw struct {
	Contract *WishlistFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// WishlistFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WishlistFactoryTransactorRaw struct {
	Contract *WishlistFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWishlistFactory creates a new instance of WishlistFactory, bound to a specific deployed contract.
func NewWishlistFactory(address common.Address, backend bind.ContractBackend) (*WishlistFactory, error) {
	contract, err := bindWishlistFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WishlistFactory{WishlistFactoryCaller: WishlistFactoryCaller{contract: contract}, WishlistFactoryTransactor: WishlistFactoryTransactor{contract: contract}, WishlistFactoryFilterer: WishlistFactoryFilterer{contract: contract}}, nil
}

// NewWishlistFactoryCaller creates a new read-only instance of WishlistFactory, bound to a specific deployed contract.
func NewWishlistFactoryCaller(address common.Address, caller bind.ContractCaller) (*WishlistFactoryCaller, error) {
	contract, err := bindWishlistFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryCaller{contract: contract}, nil
}

// NewWishlistFactoryTransactor creates a new write-only instance of WishlistFactory, bound to a specific deployed contract.
func NewWishlistFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*WishlistFactoryTransactor, error) {
	contract, err := bindWishlistFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryTransactor{contract: contract}, nil
}

// NewWishlistFactoryFilterer creates a new log filterer instance of WishlistFactory, bound to a specific deployed contract.
func NewWishlistFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*WishlistFactoryFilterer, error) {
	contract, err := bindWishlistFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryFilterer{contract: contract}, nil
}

// bindWishlistFactory binds a generic wrapper to an already deployed contract.
func bindWishlistFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WishlistFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WishlistFactory *WishlistFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WishlistFactory.Contract.WishlistFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WishlistFactory *WishlistFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WishlistFactory.Contract.WishlistFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WishlistFactory *WishlistFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WishlistFactory.Contract.WishlistFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WishlistFactory *WishlistFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WishlistFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WishlistFactory *WishlistFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WishlistFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WishlistFactory *WishlistFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WishlistFactory.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WishlistFactory *WishlistFactoryCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "eip712Domain")

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
func (_WishlistFactory *WishlistFactorySession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WishlistFactory.Contract.Eip712Domain(&_WishlistFactory.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_WishlistFactory *WishlistFactoryCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _WishlistFactory.Contract.Eip712Domain(&_WishlistFactory.CallOpts)
}

// GetProtocolFeeSettings is a free data retrieval call binding the contract method 0x7e999106.
//
// Solidity: function getProtocolFeeSettings() view returns((uint256,uint256,address))
func (_WishlistFactory *WishlistFactoryCaller) GetProtocolFeeSettings(opts *bind.CallOpts) (IWishlistFactoryProtocolFeeSettings, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "getProtocolFeeSettings")

	if err != nil {
		return *new(IWishlistFactoryProtocolFeeSettings), err
	}

	out0 := *abi.ConvertType(out[0], new(IWishlistFactoryProtocolFeeSettings)).(*IWishlistFactoryProtocolFeeSettings)

	return out0, err

}

// GetProtocolFeeSettings is a free data retrieval call binding the contract method 0x7e999106.
//
// Solidity: function getProtocolFeeSettings() view returns((uint256,uint256,address))
func (_WishlistFactory *WishlistFactorySession) GetProtocolFeeSettings() (IWishlistFactoryProtocolFeeSettings, error) {
	return _WishlistFactory.Contract.GetProtocolFeeSettings(&_WishlistFactory.CallOpts)
}

// GetProtocolFeeSettings is a free data retrieval call binding the contract method 0x7e999106.
//
// Solidity: function getProtocolFeeSettings() view returns((uint256,uint256,address))
func (_WishlistFactory *WishlistFactoryCallerSession) GetProtocolFeeSettings() (IWishlistFactoryProtocolFeeSettings, error) {
	return _WishlistFactory.Contract.GetProtocolFeeSettings(&_WishlistFactory.CallOpts)
}

// GetWishlistAddress is a free data retrieval call binding the contract method 0x1cd01349.
//
// Solidity: function getWishlistAddress(uint256 wishlistId_) view returns(address)
func (_WishlistFactory *WishlistFactoryCaller) GetWishlistAddress(opts *bind.CallOpts, wishlistId_ *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "getWishlistAddress", wishlistId_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWishlistAddress is a free data retrieval call binding the contract method 0x1cd01349.
//
// Solidity: function getWishlistAddress(uint256 wishlistId_) view returns(address)
func (_WishlistFactory *WishlistFactorySession) GetWishlistAddress(wishlistId_ *big.Int) (common.Address, error) {
	return _WishlistFactory.Contract.GetWishlistAddress(&_WishlistFactory.CallOpts, wishlistId_)
}

// GetWishlistAddress is a free data retrieval call binding the contract method 0x1cd01349.
//
// Solidity: function getWishlistAddress(uint256 wishlistId_) view returns(address)
func (_WishlistFactory *WishlistFactoryCallerSession) GetWishlistAddress(wishlistId_ *big.Int) (common.Address, error) {
	return _WishlistFactory.Contract.GetWishlistAddress(&_WishlistFactory.CallOpts, wishlistId_)
}

// GetWishlistAddresses is a free data retrieval call binding the contract method 0x8942fabb.
//
// Solidity: function getWishlistAddresses(uint256 from_, uint256 to_) view returns(address[] wishlistAddresses_)
func (_WishlistFactory *WishlistFactoryCaller) GetWishlistAddresses(opts *bind.CallOpts, from_ *big.Int, to_ *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "getWishlistAddresses", from_, to_)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWishlistAddresses is a free data retrieval call binding the contract method 0x8942fabb.
//
// Solidity: function getWishlistAddresses(uint256 from_, uint256 to_) view returns(address[] wishlistAddresses_)
func (_WishlistFactory *WishlistFactorySession) GetWishlistAddresses(from_ *big.Int, to_ *big.Int) ([]common.Address, error) {
	return _WishlistFactory.Contract.GetWishlistAddresses(&_WishlistFactory.CallOpts, from_, to_)
}

// GetWishlistAddresses is a free data retrieval call binding the contract method 0x8942fabb.
//
// Solidity: function getWishlistAddresses(uint256 from_, uint256 to_) view returns(address[] wishlistAddresses_)
func (_WishlistFactory *WishlistFactoryCallerSession) GetWishlistAddresses(from_ *big.Int, to_ *big.Int) ([]common.Address, error) {
	return _WishlistFactory.Contract.GetWishlistAddresses(&_WishlistFactory.CallOpts, from_, to_)
}

// GetWishlistsTotalCount is a free data retrieval call binding the contract method 0x84441d00.
//
// Solidity: function getWishlistsTotalCount() view returns(uint256)
func (_WishlistFactory *WishlistFactoryCaller) GetWishlistsTotalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "getWishlistsTotalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWishlistsTotalCount is a free data retrieval call binding the contract method 0x84441d00.
//
// Solidity: function getWishlistsTotalCount() view returns(uint256)
func (_WishlistFactory *WishlistFactorySession) GetWishlistsTotalCount() (*big.Int, error) {
	return _WishlistFactory.Contract.GetWishlistsTotalCount(&_WishlistFactory.CallOpts)
}

// GetWishlistsTotalCount is a free data retrieval call binding the contract method 0x84441d00.
//
// Solidity: function getWishlistsTotalCount() view returns(uint256)
func (_WishlistFactory *WishlistFactoryCallerSession) GetWishlistsTotalCount() (*big.Int, error) {
	return _WishlistFactory.Contract.GetWishlistsTotalCount(&_WishlistFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WishlistFactory *WishlistFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WishlistFactory *WishlistFactorySession) Owner() (common.Address, error) {
	return _WishlistFactory.Contract.Owner(&_WishlistFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WishlistFactory *WishlistFactoryCallerSession) Owner() (common.Address, error) {
	return _WishlistFactory.Contract.Owner(&_WishlistFactory.CallOpts)
}

// ProtocolSignerAddr is a free data retrieval call binding the contract method 0x284d7192.
//
// Solidity: function protocolSignerAddr() view returns(address)
func (_WishlistFactory *WishlistFactoryCaller) ProtocolSignerAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "protocolSignerAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolSignerAddr is a free data retrieval call binding the contract method 0x284d7192.
//
// Solidity: function protocolSignerAddr() view returns(address)
func (_WishlistFactory *WishlistFactorySession) ProtocolSignerAddr() (common.Address, error) {
	return _WishlistFactory.Contract.ProtocolSignerAddr(&_WishlistFactory.CallOpts)
}

// ProtocolSignerAddr is a free data retrieval call binding the contract method 0x284d7192.
//
// Solidity: function protocolSignerAddr() view returns(address)
func (_WishlistFactory *WishlistFactoryCallerSession) ProtocolSignerAddr() (common.Address, error) {
	return _WishlistFactory.Contract.ProtocolSignerAddr(&_WishlistFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WishlistFactory *WishlistFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WishlistFactory *WishlistFactorySession) ProxiableUUID() ([32]byte, error) {
	return _WishlistFactory.Contract.ProxiableUUID(&_WishlistFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_WishlistFactory *WishlistFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _WishlistFactory.Contract.ProxiableUUID(&_WishlistFactory.CallOpts)
}

// ProxyBeacon is a free data retrieval call binding the contract method 0xa280812d.
//
// Solidity: function proxyBeacon() view returns(address)
func (_WishlistFactory *WishlistFactoryCaller) ProxyBeacon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "proxyBeacon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyBeacon is a free data retrieval call binding the contract method 0xa280812d.
//
// Solidity: function proxyBeacon() view returns(address)
func (_WishlistFactory *WishlistFactorySession) ProxyBeacon() (common.Address, error) {
	return _WishlistFactory.Contract.ProxyBeacon(&_WishlistFactory.CallOpts)
}

// ProxyBeacon is a free data retrieval call binding the contract method 0xa280812d.
//
// Solidity: function proxyBeacon() view returns(address)
func (_WishlistFactory *WishlistFactoryCallerSession) ProxyBeacon() (common.Address, error) {
	return _WishlistFactory.Contract.ProxyBeacon(&_WishlistFactory.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_WishlistFactory *WishlistFactoryCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_WishlistFactory *WishlistFactorySession) UsdcToken() (common.Address, error) {
	return _WishlistFactory.Contract.UsdcToken(&_WishlistFactory.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_WishlistFactory *WishlistFactoryCallerSession) UsdcToken() (common.Address, error) {
	return _WishlistFactory.Contract.UsdcToken(&_WishlistFactory.CallOpts)
}

// WishlistExists is a free data retrieval call binding the contract method 0x7bd36206.
//
// Solidity: function wishlistExists(uint256 wishlistId_) view returns(bool)
func (_WishlistFactory *WishlistFactoryCaller) WishlistExists(opts *bind.CallOpts, wishlistId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _WishlistFactory.contract.Call(opts, &out, "wishlistExists", wishlistId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WishlistExists is a free data retrieval call binding the contract method 0x7bd36206.
//
// Solidity: function wishlistExists(uint256 wishlistId_) view returns(bool)
func (_WishlistFactory *WishlistFactorySession) WishlistExists(wishlistId_ *big.Int) (bool, error) {
	return _WishlistFactory.Contract.WishlistExists(&_WishlistFactory.CallOpts, wishlistId_)
}

// WishlistExists is a free data retrieval call binding the contract method 0x7bd36206.
//
// Solidity: function wishlistExists(uint256 wishlistId_) view returns(bool)
func (_WishlistFactory *WishlistFactoryCallerSession) WishlistExists(wishlistId_ *big.Int) (bool, error) {
	return _WishlistFactory.Contract.WishlistExists(&_WishlistFactory.CallOpts, wishlistId_)
}

// CreateWishlist is a paid mutator transaction binding the contract method 0xbcdf75e3.
//
// Solidity: function createWishlist(address wishlistOwner_, uint256 wishlistId_, uint256 sigDeadline_, uint256[] initialItemPrices_, bytes signature_) returns(address newWishlistAddr_)
func (_WishlistFactory *WishlistFactoryTransactor) CreateWishlist(opts *bind.TransactOpts, wishlistOwner_ common.Address, wishlistId_ *big.Int, sigDeadline_ *big.Int, initialItemPrices_ []*big.Int, signature_ []byte) (*types.Transaction, error) {
	return _WishlistFactory.contract.Transact(opts, "createWishlist", wishlistOwner_, wishlistId_, sigDeadline_, initialItemPrices_, signature_)
}

// CreateWishlist is a paid mutator transaction binding the contract method 0xbcdf75e3.
//
// Solidity: function createWishlist(address wishlistOwner_, uint256 wishlistId_, uint256 sigDeadline_, uint256[] initialItemPrices_, bytes signature_) returns(address newWishlistAddr_)
func (_WishlistFactory *WishlistFactorySession) CreateWishlist(wishlistOwner_ common.Address, wishlistId_ *big.Int, sigDeadline_ *big.Int, initialItemPrices_ []*big.Int, signature_ []byte) (*types.Transaction, error) {
	return _WishlistFactory.Contract.CreateWishlist(&_WishlistFactory.TransactOpts, wishlistOwner_, wishlistId_, sigDeadline_, initialItemPrices_, signature_)
}

// CreateWishlist is a paid mutator transaction binding the contract method 0xbcdf75e3.
//
// Solidity: function createWishlist(address wishlistOwner_, uint256 wishlistId_, uint256 sigDeadline_, uint256[] initialItemPrices_, bytes signature_) returns(address newWishlistAddr_)
func (_WishlistFactory *WishlistFactoryTransactorSession) CreateWishlist(wishlistOwner_ common.Address, wishlistId_ *big.Int, sigDeadline_ *big.Int, initialItemPrices_ []*big.Int, signature_ []byte) (*types.Transaction, error) {
	return _WishlistFactory.Contract.CreateWishlist(&_WishlistFactory.TransactOpts, wishlistOwner_, wishlistId_, sigDeadline_, initialItemPrices_, signature_)
}

// Initialize is a paid mutator transaction binding the contract method 0xb03487f7.
//
// Solidity: function initialize(address wishlistImpl_, address usdcTokenAddr_, address protocolSignerAddr_, (uint256,uint256,address) feeSettings_) returns()
func (_WishlistFactory *WishlistFactoryTransactor) Initialize(opts *bind.TransactOpts, wishlistImpl_ common.Address, usdcTokenAddr_ common.Address, protocolSignerAddr_ common.Address, feeSettings_ IWishlistFactoryProtocolFeeSettings) (*types.Transaction, error) {
	return _WishlistFactory.contract.Transact(opts, "initialize", wishlistImpl_, usdcTokenAddr_, protocolSignerAddr_, feeSettings_)
}

// Initialize is a paid mutator transaction binding the contract method 0xb03487f7.
//
// Solidity: function initialize(address wishlistImpl_, address usdcTokenAddr_, address protocolSignerAddr_, (uint256,uint256,address) feeSettings_) returns()
func (_WishlistFactory *WishlistFactorySession) Initialize(wishlistImpl_ common.Address, usdcTokenAddr_ common.Address, protocolSignerAddr_ common.Address, feeSettings_ IWishlistFactoryProtocolFeeSettings) (*types.Transaction, error) {
	return _WishlistFactory.Contract.Initialize(&_WishlistFactory.TransactOpts, wishlistImpl_, usdcTokenAddr_, protocolSignerAddr_, feeSettings_)
}

// Initialize is a paid mutator transaction binding the contract method 0xb03487f7.
//
// Solidity: function initialize(address wishlistImpl_, address usdcTokenAddr_, address protocolSignerAddr_, (uint256,uint256,address) feeSettings_) returns()
func (_WishlistFactory *WishlistFactoryTransactorSession) Initialize(wishlistImpl_ common.Address, usdcTokenAddr_ common.Address, protocolSignerAddr_ common.Address, feeSettings_ IWishlistFactoryProtocolFeeSettings) (*types.Transaction, error) {
	return _WishlistFactory.Contract.Initialize(&_WishlistFactory.TransactOpts, wishlistImpl_, usdcTokenAddr_, protocolSignerAddr_, feeSettings_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WishlistFactory *WishlistFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WishlistFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WishlistFactory *WishlistFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _WishlistFactory.Contract.RenounceOwnership(&_WishlistFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WishlistFactory *WishlistFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WishlistFactory.Contract.RenounceOwnership(&_WishlistFactory.TransactOpts)
}

// SetProtocolFeeSettings is a paid mutator transaction binding the contract method 0xf6bcade6.
//
// Solidity: function setProtocolFeeSettings((uint256,uint256,address) feeSettings_) returns()
func (_WishlistFactory *WishlistFactoryTransactor) SetProtocolFeeSettings(opts *bind.TransactOpts, feeSettings_ IWishlistFactoryProtocolFeeSettings) (*types.Transaction, error) {
	return _WishlistFactory.contract.Transact(opts, "setProtocolFeeSettings", feeSettings_)
}

// SetProtocolFeeSettings is a paid mutator transaction binding the contract method 0xf6bcade6.
//
// Solidity: function setProtocolFeeSettings((uint256,uint256,address) feeSettings_) returns()
func (_WishlistFactory *WishlistFactorySession) SetProtocolFeeSettings(feeSettings_ IWishlistFactoryProtocolFeeSettings) (*types.Transaction, error) {
	return _WishlistFactory.Contract.SetProtocolFeeSettings(&_WishlistFactory.TransactOpts, feeSettings_)
}

// SetProtocolFeeSettings is a paid mutator transaction binding the contract method 0xf6bcade6.
//
// Solidity: function setProtocolFeeSettings((uint256,uint256,address) feeSettings_) returns()
func (_WishlistFactory *WishlistFactoryTransactorSession) SetProtocolFeeSettings(feeSettings_ IWishlistFactoryProtocolFeeSettings) (*types.Transaction, error) {
	return _WishlistFactory.Contract.SetProtocolFeeSettings(&_WishlistFactory.TransactOpts, feeSettings_)
}

// SetProtocolSignerAddr is a paid mutator transaction binding the contract method 0x771bea98.
//
// Solidity: function setProtocolSignerAddr(address newProtocolSignerAddr_) returns()
func (_WishlistFactory *WishlistFactoryTransactor) SetProtocolSignerAddr(opts *bind.TransactOpts, newProtocolSignerAddr_ common.Address) (*types.Transaction, error) {
	return _WishlistFactory.contract.Transact(opts, "setProtocolSignerAddr", newProtocolSignerAddr_)
}

// SetProtocolSignerAddr is a paid mutator transaction binding the contract method 0x771bea98.
//
// Solidity: function setProtocolSignerAddr(address newProtocolSignerAddr_) returns()
func (_WishlistFactory *WishlistFactorySession) SetProtocolSignerAddr(newProtocolSignerAddr_ common.Address) (*types.Transaction, error) {
	return _WishlistFactory.Contract.SetProtocolSignerAddr(&_WishlistFactory.TransactOpts, newProtocolSignerAddr_)
}

// SetProtocolSignerAddr is a paid mutator transaction binding the contract method 0x771bea98.
//
// Solidity: function setProtocolSignerAddr(address newProtocolSignerAddr_) returns()
func (_WishlistFactory *WishlistFactoryTransactorSession) SetProtocolSignerAddr(newProtocolSignerAddr_ common.Address) (*types.Transaction, error) {
	return _WishlistFactory.Contract.SetProtocolSignerAddr(&_WishlistFactory.TransactOpts, newProtocolSignerAddr_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WishlistFactory *WishlistFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WishlistFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WishlistFactory *WishlistFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WishlistFactory.Contract.TransferOwnership(&_WishlistFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WishlistFactory *WishlistFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WishlistFactory.Contract.TransferOwnership(&_WishlistFactory.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WishlistFactory *WishlistFactoryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _WishlistFactory.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WishlistFactory *WishlistFactorySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WishlistFactory.Contract.UpgradeTo(&_WishlistFactory.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_WishlistFactory *WishlistFactoryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _WishlistFactory.Contract.UpgradeTo(&_WishlistFactory.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WishlistFactory *WishlistFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WishlistFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WishlistFactory *WishlistFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WishlistFactory.Contract.UpgradeToAndCall(&_WishlistFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_WishlistFactory *WishlistFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _WishlistFactory.Contract.UpgradeToAndCall(&_WishlistFactory.TransactOpts, newImplementation, data)
}

// UpgradeWishlistsImpl is a paid mutator transaction binding the contract method 0xd6cab1f4.
//
// Solidity: function upgradeWishlistsImpl(address newWishlistsImpl_) returns()
func (_WishlistFactory *WishlistFactoryTransactor) UpgradeWishlistsImpl(opts *bind.TransactOpts, newWishlistsImpl_ common.Address) (*types.Transaction, error) {
	return _WishlistFactory.contract.Transact(opts, "upgradeWishlistsImpl", newWishlistsImpl_)
}

// UpgradeWishlistsImpl is a paid mutator transaction binding the contract method 0xd6cab1f4.
//
// Solidity: function upgradeWishlistsImpl(address newWishlistsImpl_) returns()
func (_WishlistFactory *WishlistFactorySession) UpgradeWishlistsImpl(newWishlistsImpl_ common.Address) (*types.Transaction, error) {
	return _WishlistFactory.Contract.UpgradeWishlistsImpl(&_WishlistFactory.TransactOpts, newWishlistsImpl_)
}

// UpgradeWishlistsImpl is a paid mutator transaction binding the contract method 0xd6cab1f4.
//
// Solidity: function upgradeWishlistsImpl(address newWishlistsImpl_) returns()
func (_WishlistFactory *WishlistFactoryTransactorSession) UpgradeWishlistsImpl(newWishlistsImpl_ common.Address) (*types.Transaction, error) {
	return _WishlistFactory.Contract.UpgradeWishlistsImpl(&_WishlistFactory.TransactOpts, newWishlistsImpl_)
}

// WishlistFactoryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the WishlistFactory contract.
type WishlistFactoryAdminChangedIterator struct {
	Event *WishlistFactoryAdminChanged // Event containing the contract specifics and raw log

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
func (it *WishlistFactoryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFactoryAdminChanged)
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
		it.Event = new(WishlistFactoryAdminChanged)
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
func (it *WishlistFactoryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFactoryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFactoryAdminChanged represents a AdminChanged event raised by the WishlistFactory contract.
type WishlistFactoryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WishlistFactory *WishlistFactoryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*WishlistFactoryAdminChangedIterator, error) {

	logs, sub, err := _WishlistFactory.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryAdminChangedIterator{contract: _WishlistFactory.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_WishlistFactory *WishlistFactoryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *WishlistFactoryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _WishlistFactory.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFactoryAdminChanged)
				if err := _WishlistFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_WishlistFactory *WishlistFactoryFilterer) ParseAdminChanged(log types.Log) (*WishlistFactoryAdminChanged, error) {
	event := new(WishlistFactoryAdminChanged)
	if err := _WishlistFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistFactoryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the WishlistFactory contract.
type WishlistFactoryBeaconUpgradedIterator struct {
	Event *WishlistFactoryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *WishlistFactoryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFactoryBeaconUpgraded)
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
		it.Event = new(WishlistFactoryBeaconUpgraded)
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
func (it *WishlistFactoryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFactoryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFactoryBeaconUpgraded represents a BeaconUpgraded event raised by the WishlistFactory contract.
type WishlistFactoryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WishlistFactory *WishlistFactoryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*WishlistFactoryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WishlistFactory.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryBeaconUpgradedIterator{contract: _WishlistFactory.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_WishlistFactory *WishlistFactoryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *WishlistFactoryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _WishlistFactory.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFactoryBeaconUpgraded)
				if err := _WishlistFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_WishlistFactory *WishlistFactoryFilterer) ParseBeaconUpgraded(log types.Log) (*WishlistFactoryBeaconUpgraded, error) {
	event := new(WishlistFactoryBeaconUpgraded)
	if err := _WishlistFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistFactoryEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the WishlistFactory contract.
type WishlistFactoryEIP712DomainChangedIterator struct {
	Event *WishlistFactoryEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *WishlistFactoryEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFactoryEIP712DomainChanged)
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
		it.Event = new(WishlistFactoryEIP712DomainChanged)
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
func (it *WishlistFactoryEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFactoryEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFactoryEIP712DomainChanged represents a EIP712DomainChanged event raised by the WishlistFactory contract.
type WishlistFactoryEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WishlistFactory *WishlistFactoryFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*WishlistFactoryEIP712DomainChangedIterator, error) {

	logs, sub, err := _WishlistFactory.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryEIP712DomainChangedIterator{contract: _WishlistFactory.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_WishlistFactory *WishlistFactoryFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *WishlistFactoryEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _WishlistFactory.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFactoryEIP712DomainChanged)
				if err := _WishlistFactory.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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
func (_WishlistFactory *WishlistFactoryFilterer) ParseEIP712DomainChanged(log types.Log) (*WishlistFactoryEIP712DomainChanged, error) {
	event := new(WishlistFactoryEIP712DomainChanged)
	if err := _WishlistFactory.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WishlistFactory contract.
type WishlistFactoryInitializedIterator struct {
	Event *WishlistFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *WishlistFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFactoryInitialized)
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
		it.Event = new(WishlistFactoryInitialized)
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
func (it *WishlistFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFactoryInitialized represents a Initialized event raised by the WishlistFactory contract.
type WishlistFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WishlistFactory *WishlistFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*WishlistFactoryInitializedIterator, error) {

	logs, sub, err := _WishlistFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryInitializedIterator{contract: _WishlistFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WishlistFactory *WishlistFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WishlistFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _WishlistFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFactoryInitialized)
				if err := _WishlistFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_WishlistFactory *WishlistFactoryFilterer) ParseInitialized(log types.Log) (*WishlistFactoryInitialized, error) {
	event := new(WishlistFactoryInitialized)
	if err := _WishlistFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistFactoryNewWishlistCreatedIterator is returned from FilterNewWishlistCreated and is used to iterate over the raw logs and unpacked data for NewWishlistCreated events raised by the WishlistFactory contract.
type WishlistFactoryNewWishlistCreatedIterator struct {
	Event *WishlistFactoryNewWishlistCreated // Event containing the contract specifics and raw log

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
func (it *WishlistFactoryNewWishlistCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFactoryNewWishlistCreated)
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
		it.Event = new(WishlistFactoryNewWishlistCreated)
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
func (it *WishlistFactoryNewWishlistCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFactoryNewWishlistCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFactoryNewWishlistCreated represents a NewWishlistCreated event raised by the WishlistFactory contract.
type WishlistFactoryNewWishlistCreated struct {
	WishlistOwner common.Address
	WishlistId    *big.Int
	WishlistAddr  common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewWishlistCreated is a free log retrieval operation binding the contract event 0x7f3d4ad58c7a5e976b64408fe823ca51d3c9660eb682fb4895ddc0ae7189cb66.
//
// Solidity: event NewWishlistCreated(address indexed wishlistOwner, uint256 wishlistId, address wishlistAddr)
func (_WishlistFactory *WishlistFactoryFilterer) FilterNewWishlistCreated(opts *bind.FilterOpts, wishlistOwner []common.Address) (*WishlistFactoryNewWishlistCreatedIterator, error) {

	var wishlistOwnerRule []interface{}
	for _, wishlistOwnerItem := range wishlistOwner {
		wishlistOwnerRule = append(wishlistOwnerRule, wishlistOwnerItem)
	}

	logs, sub, err := _WishlistFactory.contract.FilterLogs(opts, "NewWishlistCreated", wishlistOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryNewWishlistCreatedIterator{contract: _WishlistFactory.contract, event: "NewWishlistCreated", logs: logs, sub: sub}, nil
}

// WatchNewWishlistCreated is a free log subscription operation binding the contract event 0x7f3d4ad58c7a5e976b64408fe823ca51d3c9660eb682fb4895ddc0ae7189cb66.
//
// Solidity: event NewWishlistCreated(address indexed wishlistOwner, uint256 wishlistId, address wishlistAddr)
func (_WishlistFactory *WishlistFactoryFilterer) WatchNewWishlistCreated(opts *bind.WatchOpts, sink chan<- *WishlistFactoryNewWishlistCreated, wishlistOwner []common.Address) (event.Subscription, error) {

	var wishlistOwnerRule []interface{}
	for _, wishlistOwnerItem := range wishlistOwner {
		wishlistOwnerRule = append(wishlistOwnerRule, wishlistOwnerItem)
	}

	logs, sub, err := _WishlistFactory.contract.WatchLogs(opts, "NewWishlistCreated", wishlistOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFactoryNewWishlistCreated)
				if err := _WishlistFactory.contract.UnpackLog(event, "NewWishlistCreated", log); err != nil {
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

// ParseNewWishlistCreated is a log parse operation binding the contract event 0x7f3d4ad58c7a5e976b64408fe823ca51d3c9660eb682fb4895ddc0ae7189cb66.
//
// Solidity: event NewWishlistCreated(address indexed wishlistOwner, uint256 wishlistId, address wishlistAddr)
func (_WishlistFactory *WishlistFactoryFilterer) ParseNewWishlistCreated(log types.Log) (*WishlistFactoryNewWishlistCreated, error) {
	event := new(WishlistFactoryNewWishlistCreated)
	if err := _WishlistFactory.contract.UnpackLog(event, "NewWishlistCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WishlistFactory contract.
type WishlistFactoryOwnershipTransferredIterator struct {
	Event *WishlistFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WishlistFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFactoryOwnershipTransferred)
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
		it.Event = new(WishlistFactoryOwnershipTransferred)
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
func (it *WishlistFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the WishlistFactory contract.
type WishlistFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WishlistFactory *WishlistFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WishlistFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WishlistFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryOwnershipTransferredIterator{contract: _WishlistFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WishlistFactory *WishlistFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WishlistFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WishlistFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFactoryOwnershipTransferred)
				if err := _WishlistFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WishlistFactory *WishlistFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*WishlistFactoryOwnershipTransferred, error) {
	event := new(WishlistFactoryOwnershipTransferred)
	if err := _WishlistFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistFactoryProtocolFeeSettingsUpdatedIterator is returned from FilterProtocolFeeSettingsUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeSettingsUpdated events raised by the WishlistFactory contract.
type WishlistFactoryProtocolFeeSettingsUpdatedIterator struct {
	Event *WishlistFactoryProtocolFeeSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *WishlistFactoryProtocolFeeSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFactoryProtocolFeeSettingsUpdated)
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
		it.Event = new(WishlistFactoryProtocolFeeSettingsUpdated)
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
func (it *WishlistFactoryProtocolFeeSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFactoryProtocolFeeSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFactoryProtocolFeeSettingsUpdated represents a ProtocolFeeSettingsUpdated event raised by the WishlistFactory contract.
type WishlistFactoryProtocolFeeSettingsUpdated struct {
	NewProtocolFeePercentage *big.Int
	NewMaxProtocolFeeAmount  *big.Int
	NewProtocolFeeRecipient  common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeSettingsUpdated is a free log retrieval operation binding the contract event 0x2ef35253d66a5745228e36cabace1ea21979dd57d91ceee43149c57730fa5cda.
//
// Solidity: event ProtocolFeeSettingsUpdated(uint256 newProtocolFeePercentage, uint256 newMaxProtocolFeeAmount, address newProtocolFeeRecipient)
func (_WishlistFactory *WishlistFactoryFilterer) FilterProtocolFeeSettingsUpdated(opts *bind.FilterOpts) (*WishlistFactoryProtocolFeeSettingsUpdatedIterator, error) {

	logs, sub, err := _WishlistFactory.contract.FilterLogs(opts, "ProtocolFeeSettingsUpdated")
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryProtocolFeeSettingsUpdatedIterator{contract: _WishlistFactory.contract, event: "ProtocolFeeSettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeSettingsUpdated is a free log subscription operation binding the contract event 0x2ef35253d66a5745228e36cabace1ea21979dd57d91ceee43149c57730fa5cda.
//
// Solidity: event ProtocolFeeSettingsUpdated(uint256 newProtocolFeePercentage, uint256 newMaxProtocolFeeAmount, address newProtocolFeeRecipient)
func (_WishlistFactory *WishlistFactoryFilterer) WatchProtocolFeeSettingsUpdated(opts *bind.WatchOpts, sink chan<- *WishlistFactoryProtocolFeeSettingsUpdated) (event.Subscription, error) {

	logs, sub, err := _WishlistFactory.contract.WatchLogs(opts, "ProtocolFeeSettingsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFactoryProtocolFeeSettingsUpdated)
				if err := _WishlistFactory.contract.UnpackLog(event, "ProtocolFeeSettingsUpdated", log); err != nil {
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

// ParseProtocolFeeSettingsUpdated is a log parse operation binding the contract event 0x2ef35253d66a5745228e36cabace1ea21979dd57d91ceee43149c57730fa5cda.
//
// Solidity: event ProtocolFeeSettingsUpdated(uint256 newProtocolFeePercentage, uint256 newMaxProtocolFeeAmount, address newProtocolFeeRecipient)
func (_WishlistFactory *WishlistFactoryFilterer) ParseProtocolFeeSettingsUpdated(log types.Log) (*WishlistFactoryProtocolFeeSettingsUpdated, error) {
	event := new(WishlistFactoryProtocolFeeSettingsUpdated)
	if err := _WishlistFactory.contract.UnpackLog(event, "ProtocolFeeSettingsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistFactoryProtocolSignerUpdatedIterator is returned from FilterProtocolSignerUpdated and is used to iterate over the raw logs and unpacked data for ProtocolSignerUpdated events raised by the WishlistFactory contract.
type WishlistFactoryProtocolSignerUpdatedIterator struct {
	Event *WishlistFactoryProtocolSignerUpdated // Event containing the contract specifics and raw log

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
func (it *WishlistFactoryProtocolSignerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFactoryProtocolSignerUpdated)
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
		it.Event = new(WishlistFactoryProtocolSignerUpdated)
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
func (it *WishlistFactoryProtocolSignerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFactoryProtocolSignerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFactoryProtocolSignerUpdated represents a ProtocolSignerUpdated event raised by the WishlistFactory contract.
type WishlistFactoryProtocolSignerUpdated struct {
	NewProtocolSignerAddr common.Address
	OldProtocolSignerAddr common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProtocolSignerUpdated is a free log retrieval operation binding the contract event 0x98a339ab527d1d2b312f1ff2ea2581f6e0454de6ed228e50944944e95e0c2c40.
//
// Solidity: event ProtocolSignerUpdated(address newProtocolSignerAddr, address oldProtocolSignerAddr)
func (_WishlistFactory *WishlistFactoryFilterer) FilterProtocolSignerUpdated(opts *bind.FilterOpts) (*WishlistFactoryProtocolSignerUpdatedIterator, error) {

	logs, sub, err := _WishlistFactory.contract.FilterLogs(opts, "ProtocolSignerUpdated")
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryProtocolSignerUpdatedIterator{contract: _WishlistFactory.contract, event: "ProtocolSignerUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolSignerUpdated is a free log subscription operation binding the contract event 0x98a339ab527d1d2b312f1ff2ea2581f6e0454de6ed228e50944944e95e0c2c40.
//
// Solidity: event ProtocolSignerUpdated(address newProtocolSignerAddr, address oldProtocolSignerAddr)
func (_WishlistFactory *WishlistFactoryFilterer) WatchProtocolSignerUpdated(opts *bind.WatchOpts, sink chan<- *WishlistFactoryProtocolSignerUpdated) (event.Subscription, error) {

	logs, sub, err := _WishlistFactory.contract.WatchLogs(opts, "ProtocolSignerUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFactoryProtocolSignerUpdated)
				if err := _WishlistFactory.contract.UnpackLog(event, "ProtocolSignerUpdated", log); err != nil {
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

// ParseProtocolSignerUpdated is a log parse operation binding the contract event 0x98a339ab527d1d2b312f1ff2ea2581f6e0454de6ed228e50944944e95e0c2c40.
//
// Solidity: event ProtocolSignerUpdated(address newProtocolSignerAddr, address oldProtocolSignerAddr)
func (_WishlistFactory *WishlistFactoryFilterer) ParseProtocolSignerUpdated(log types.Log) (*WishlistFactoryProtocolSignerUpdated, error) {
	event := new(WishlistFactoryProtocolSignerUpdated)
	if err := _WishlistFactory.contract.UnpackLog(event, "ProtocolSignerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the WishlistFactory contract.
type WishlistFactoryUpgradedIterator struct {
	Event *WishlistFactoryUpgraded // Event containing the contract specifics and raw log

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
func (it *WishlistFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFactoryUpgraded)
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
		it.Event = new(WishlistFactoryUpgraded)
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
func (it *WishlistFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFactoryUpgraded represents a Upgraded event raised by the WishlistFactory contract.
type WishlistFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WishlistFactory *WishlistFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*WishlistFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WishlistFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &WishlistFactoryUpgradedIterator{contract: _WishlistFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_WishlistFactory *WishlistFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *WishlistFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _WishlistFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFactoryUpgraded)
				if err := _WishlistFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_WishlistFactory *WishlistFactoryFilterer) ParseUpgraded(log types.Log) (*WishlistFactoryUpgraded, error) {
	event := new(WishlistFactoryUpgraded)
	if err := _WishlistFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
