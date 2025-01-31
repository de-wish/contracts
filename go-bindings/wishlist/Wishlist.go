// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wishlist

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

// IWishlistChangedItemPriceData is an auto generated low-level Go binding around an user-defined struct.
type IWishlistChangedItemPriceData struct {
	ItemId       *big.Int
	NewItemPrice *big.Int
}

// IWishlistContributionInfo is an auto generated low-level Go binding around an user-defined struct.
type IWishlistContributionInfo struct {
	ContributionId     *big.Int
	Contributor        common.Address
	ContributionAmount *big.Int
}

// IWishlistWishlistInfo is an auto generated low-level Go binding around an user-defined struct.
type IWishlistWishlistInfo struct {
	WishlistId      *big.Int
	OwnerAddr       common.Address
	OwedUsdcAmount  *big.Int
	TotalItemsCount *big.Int
	ActiveItemsInfo []IWishlistWishlistItemInfo
}

// IWishlistWishlistItemInfo is an auto generated low-level Go binding around an user-defined struct.
type IWishlistWishlistItemInfo struct {
	ItemId                   *big.Int
	ItemPrice                *big.Int
	CollectedTokensAmount    *big.Int
	TotalContributionsNumber *big.Int
	ContributionsInfo        []IWishlistContributionInfo
	IsActive                 bool
}

// WishlistMetaData contains all meta data concerning the Wishlist contract.
var WishlistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddr\",\"type\":\"address\"}],\"name\":\"NotAWishlistOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"NotAnActiveItem\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAmountToWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroItemPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroItemsToAdd\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalContributionsNumber\",\"type\":\"uint256\"}],\"name\":\"FundsCollectionFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"contributedAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"contributionAddr\",\"type\":\"address\"}],\"name\":\"FundsContributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fundsRecipient\",\"type\":\"address\"}],\"name\":\"WishlistFundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lastItemId\",\"type\":\"uint256\"}],\"name\":\"WishlistItemsAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"removedItemIds\",\"type\":\"uint256[]\"}],\"name\":\"WishlistItemsRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"itemPrices_\",\"type\":\"uint256[]\"}],\"name\":\"addWishlistItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newItemPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIWishlist.ChangedItemPriceData[]\",\"name\":\"changedItemsDataArr_\",\"type\":\"tuple[]\"}],\"name\":\"changeWishlistItemsPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountToContribute_\",\"type\":\"uint256\"}],\"name\":\"contributeFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIWishlistFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId_\",\"type\":\"uint256\"}],\"name\":\"getContributionsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"contributionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"contributionAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIWishlist.ContributionInfo[]\",\"name\":\"resultArr_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWishlistInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"wishlistId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ownerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"owedUsdcAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalItemsCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectedTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalContributionsNumber\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"contributionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"contributionAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIWishlist.ContributionInfo[]\",\"name\":\"contributionsInfo\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIWishlist.WishlistItemInfo[]\",\"name\":\"activeItemsInfo\",\"type\":\"tuple[]\"}],\"internalType\":\"structIWishlist.WishlistInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"itemIds_\",\"type\":\"uint256[]\"}],\"name\":\"getWishlistItemsInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"itemPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collectedTokensAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalContributionsNumber\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"contributionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"contributionAmount\",\"type\":\"uint256\"}],\"internalType\":\"structIWishlist.ContributionInfo[]\",\"name\":\"contributionsInfo\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIWishlist.WishlistItemInfo[]\",\"name\":\"resultArr_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wishlistId_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"wishlistOwner_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"initialItemPrices_\",\"type\":\"uint256[]\"}],\"name\":\"initializeWishlist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"itemId_\",\"type\":\"uint256\"}],\"name\":\"isItemActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextItemId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"itemIds_\",\"type\":\"uint256[]\"}],\"name\":\"removeWishlistItems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdcToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wishlistId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wishlistOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fundsRecipient_\",\"type\":\"address\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// WishlistABI is the input ABI used to generate the binding from.
// Deprecated: Use WishlistMetaData.ABI instead.
var WishlistABI = WishlistMetaData.ABI

// Wishlist is an auto generated Go binding around an Ethereum contract.
type Wishlist struct {
	WishlistCaller     // Read-only binding to the contract
	WishlistTransactor // Write-only binding to the contract
	WishlistFilterer   // Log filterer for contract events
}

// WishlistCaller is an auto generated read-only Go binding around an Ethereum contract.
type WishlistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WishlistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WishlistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WishlistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WishlistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WishlistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WishlistSession struct {
	Contract     *Wishlist         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WishlistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WishlistCallerSession struct {
	Contract *WishlistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WishlistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WishlistTransactorSession struct {
	Contract     *WishlistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WishlistRaw is an auto generated low-level Go binding around an Ethereum contract.
type WishlistRaw struct {
	Contract *Wishlist // Generic contract binding to access the raw methods on
}

// WishlistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WishlistCallerRaw struct {
	Contract *WishlistCaller // Generic read-only contract binding to access the raw methods on
}

// WishlistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WishlistTransactorRaw struct {
	Contract *WishlistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWishlist creates a new instance of Wishlist, bound to a specific deployed contract.
func NewWishlist(address common.Address, backend bind.ContractBackend) (*Wishlist, error) {
	contract, err := bindWishlist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wishlist{WishlistCaller: WishlistCaller{contract: contract}, WishlistTransactor: WishlistTransactor{contract: contract}, WishlistFilterer: WishlistFilterer{contract: contract}}, nil
}

// NewWishlistCaller creates a new read-only instance of Wishlist, bound to a specific deployed contract.
func NewWishlistCaller(address common.Address, caller bind.ContractCaller) (*WishlistCaller, error) {
	contract, err := bindWishlist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WishlistCaller{contract: contract}, nil
}

// NewWishlistTransactor creates a new write-only instance of Wishlist, bound to a specific deployed contract.
func NewWishlistTransactor(address common.Address, transactor bind.ContractTransactor) (*WishlistTransactor, error) {
	contract, err := bindWishlist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WishlistTransactor{contract: contract}, nil
}

// NewWishlistFilterer creates a new log filterer instance of Wishlist, bound to a specific deployed contract.
func NewWishlistFilterer(address common.Address, filterer bind.ContractFilterer) (*WishlistFilterer, error) {
	contract, err := bindWishlist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WishlistFilterer{contract: contract}, nil
}

// bindWishlist binds a generic wrapper to an already deployed contract.
func bindWishlist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WishlistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wishlist *WishlistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wishlist.Contract.WishlistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wishlist *WishlistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wishlist.Contract.WishlistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wishlist *WishlistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wishlist.Contract.WishlistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wishlist *WishlistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wishlist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wishlist *WishlistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wishlist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wishlist *WishlistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wishlist.Contract.contract.Transact(opts, method, params...)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Wishlist *WishlistCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wishlist.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Wishlist *WishlistSession) Factory() (common.Address, error) {
	return _Wishlist.Contract.Factory(&_Wishlist.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Wishlist *WishlistCallerSession) Factory() (common.Address, error) {
	return _Wishlist.Contract.Factory(&_Wishlist.CallOpts)
}

// GetContributionsInfo is a free data retrieval call binding the contract method 0x6abf1822.
//
// Solidity: function getContributionsInfo(uint256 itemId_) view returns((uint256,address,uint256)[] resultArr_)
func (_Wishlist *WishlistCaller) GetContributionsInfo(opts *bind.CallOpts, itemId_ *big.Int) ([]IWishlistContributionInfo, error) {
	var out []interface{}
	err := _Wishlist.contract.Call(opts, &out, "getContributionsInfo", itemId_)

	if err != nil {
		return *new([]IWishlistContributionInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWishlistContributionInfo)).(*[]IWishlistContributionInfo)

	return out0, err

}

// GetContributionsInfo is a free data retrieval call binding the contract method 0x6abf1822.
//
// Solidity: function getContributionsInfo(uint256 itemId_) view returns((uint256,address,uint256)[] resultArr_)
func (_Wishlist *WishlistSession) GetContributionsInfo(itemId_ *big.Int) ([]IWishlistContributionInfo, error) {
	return _Wishlist.Contract.GetContributionsInfo(&_Wishlist.CallOpts, itemId_)
}

// GetContributionsInfo is a free data retrieval call binding the contract method 0x6abf1822.
//
// Solidity: function getContributionsInfo(uint256 itemId_) view returns((uint256,address,uint256)[] resultArr_)
func (_Wishlist *WishlistCallerSession) GetContributionsInfo(itemId_ *big.Int) ([]IWishlistContributionInfo, error) {
	return _Wishlist.Contract.GetContributionsInfo(&_Wishlist.CallOpts, itemId_)
}

// GetWishlistInfo is a free data retrieval call binding the contract method 0x15604773.
//
// Solidity: function getWishlistInfo() view returns((uint256,address,uint256,uint256,(uint256,uint256,uint256,uint256,(uint256,address,uint256)[],bool)[]))
func (_Wishlist *WishlistCaller) GetWishlistInfo(opts *bind.CallOpts) (IWishlistWishlistInfo, error) {
	var out []interface{}
	err := _Wishlist.contract.Call(opts, &out, "getWishlistInfo")

	if err != nil {
		return *new(IWishlistWishlistInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IWishlistWishlistInfo)).(*IWishlistWishlistInfo)

	return out0, err

}

// GetWishlistInfo is a free data retrieval call binding the contract method 0x15604773.
//
// Solidity: function getWishlistInfo() view returns((uint256,address,uint256,uint256,(uint256,uint256,uint256,uint256,(uint256,address,uint256)[],bool)[]))
func (_Wishlist *WishlistSession) GetWishlistInfo() (IWishlistWishlistInfo, error) {
	return _Wishlist.Contract.GetWishlistInfo(&_Wishlist.CallOpts)
}

// GetWishlistInfo is a free data retrieval call binding the contract method 0x15604773.
//
// Solidity: function getWishlistInfo() view returns((uint256,address,uint256,uint256,(uint256,uint256,uint256,uint256,(uint256,address,uint256)[],bool)[]))
func (_Wishlist *WishlistCallerSession) GetWishlistInfo() (IWishlistWishlistInfo, error) {
	return _Wishlist.Contract.GetWishlistInfo(&_Wishlist.CallOpts)
}

// GetWishlistItemsInfo is a free data retrieval call binding the contract method 0x06effc75.
//
// Solidity: function getWishlistItemsInfo(uint256[] itemIds_) view returns((uint256,uint256,uint256,uint256,(uint256,address,uint256)[],bool)[] resultArr_)
func (_Wishlist *WishlistCaller) GetWishlistItemsInfo(opts *bind.CallOpts, itemIds_ []*big.Int) ([]IWishlistWishlistItemInfo, error) {
	var out []interface{}
	err := _Wishlist.contract.Call(opts, &out, "getWishlistItemsInfo", itemIds_)

	if err != nil {
		return *new([]IWishlistWishlistItemInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]IWishlistWishlistItemInfo)).(*[]IWishlistWishlistItemInfo)

	return out0, err

}

// GetWishlistItemsInfo is a free data retrieval call binding the contract method 0x06effc75.
//
// Solidity: function getWishlistItemsInfo(uint256[] itemIds_) view returns((uint256,uint256,uint256,uint256,(uint256,address,uint256)[],bool)[] resultArr_)
func (_Wishlist *WishlistSession) GetWishlistItemsInfo(itemIds_ []*big.Int) ([]IWishlistWishlistItemInfo, error) {
	return _Wishlist.Contract.GetWishlistItemsInfo(&_Wishlist.CallOpts, itemIds_)
}

// GetWishlistItemsInfo is a free data retrieval call binding the contract method 0x06effc75.
//
// Solidity: function getWishlistItemsInfo(uint256[] itemIds_) view returns((uint256,uint256,uint256,uint256,(uint256,address,uint256)[],bool)[] resultArr_)
func (_Wishlist *WishlistCallerSession) GetWishlistItemsInfo(itemIds_ []*big.Int) ([]IWishlistWishlistItemInfo, error) {
	return _Wishlist.Contract.GetWishlistItemsInfo(&_Wishlist.CallOpts, itemIds_)
}

// IsItemActive is a free data retrieval call binding the contract method 0x8de2041c.
//
// Solidity: function isItemActive(uint256 itemId_) view returns(bool)
func (_Wishlist *WishlistCaller) IsItemActive(opts *bind.CallOpts, itemId_ *big.Int) (bool, error) {
	var out []interface{}
	err := _Wishlist.contract.Call(opts, &out, "isItemActive", itemId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsItemActive is a free data retrieval call binding the contract method 0x8de2041c.
//
// Solidity: function isItemActive(uint256 itemId_) view returns(bool)
func (_Wishlist *WishlistSession) IsItemActive(itemId_ *big.Int) (bool, error) {
	return _Wishlist.Contract.IsItemActive(&_Wishlist.CallOpts, itemId_)
}

// IsItemActive is a free data retrieval call binding the contract method 0x8de2041c.
//
// Solidity: function isItemActive(uint256 itemId_) view returns(bool)
func (_Wishlist *WishlistCallerSession) IsItemActive(itemId_ *big.Int) (bool, error) {
	return _Wishlist.Contract.IsItemActive(&_Wishlist.CallOpts, itemId_)
}

// NextItemId is a free data retrieval call binding the contract method 0x6a868974.
//
// Solidity: function nextItemId() view returns(uint256)
func (_Wishlist *WishlistCaller) NextItemId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wishlist.contract.Call(opts, &out, "nextItemId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextItemId is a free data retrieval call binding the contract method 0x6a868974.
//
// Solidity: function nextItemId() view returns(uint256)
func (_Wishlist *WishlistSession) NextItemId() (*big.Int, error) {
	return _Wishlist.Contract.NextItemId(&_Wishlist.CallOpts)
}

// NextItemId is a free data retrieval call binding the contract method 0x6a868974.
//
// Solidity: function nextItemId() view returns(uint256)
func (_Wishlist *WishlistCallerSession) NextItemId() (*big.Int, error) {
	return _Wishlist.Contract.NextItemId(&_Wishlist.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Wishlist *WishlistCaller) UsdcToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wishlist.contract.Call(opts, &out, "usdcToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Wishlist *WishlistSession) UsdcToken() (common.Address, error) {
	return _Wishlist.Contract.UsdcToken(&_Wishlist.CallOpts)
}

// UsdcToken is a free data retrieval call binding the contract method 0x11eac855.
//
// Solidity: function usdcToken() view returns(address)
func (_Wishlist *WishlistCallerSession) UsdcToken() (common.Address, error) {
	return _Wishlist.Contract.UsdcToken(&_Wishlist.CallOpts)
}

// WishlistId is a free data retrieval call binding the contract method 0x2493c840.
//
// Solidity: function wishlistId() view returns(uint256)
func (_Wishlist *WishlistCaller) WishlistId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Wishlist.contract.Call(opts, &out, "wishlistId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WishlistId is a free data retrieval call binding the contract method 0x2493c840.
//
// Solidity: function wishlistId() view returns(uint256)
func (_Wishlist *WishlistSession) WishlistId() (*big.Int, error) {
	return _Wishlist.Contract.WishlistId(&_Wishlist.CallOpts)
}

// WishlistId is a free data retrieval call binding the contract method 0x2493c840.
//
// Solidity: function wishlistId() view returns(uint256)
func (_Wishlist *WishlistCallerSession) WishlistId() (*big.Int, error) {
	return _Wishlist.Contract.WishlistId(&_Wishlist.CallOpts)
}

// WishlistOwner is a free data retrieval call binding the contract method 0x398ef7f8.
//
// Solidity: function wishlistOwner() view returns(address)
func (_Wishlist *WishlistCaller) WishlistOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wishlist.contract.Call(opts, &out, "wishlistOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WishlistOwner is a free data retrieval call binding the contract method 0x398ef7f8.
//
// Solidity: function wishlistOwner() view returns(address)
func (_Wishlist *WishlistSession) WishlistOwner() (common.Address, error) {
	return _Wishlist.Contract.WishlistOwner(&_Wishlist.CallOpts)
}

// WishlistOwner is a free data retrieval call binding the contract method 0x398ef7f8.
//
// Solidity: function wishlistOwner() view returns(address)
func (_Wishlist *WishlistCallerSession) WishlistOwner() (common.Address, error) {
	return _Wishlist.Contract.WishlistOwner(&_Wishlist.CallOpts)
}

// AddWishlistItems is a paid mutator transaction binding the contract method 0x11e06d3a.
//
// Solidity: function addWishlistItems(uint256[] itemPrices_) returns()
func (_Wishlist *WishlistTransactor) AddWishlistItems(opts *bind.TransactOpts, itemPrices_ []*big.Int) (*types.Transaction, error) {
	return _Wishlist.contract.Transact(opts, "addWishlistItems", itemPrices_)
}

// AddWishlistItems is a paid mutator transaction binding the contract method 0x11e06d3a.
//
// Solidity: function addWishlistItems(uint256[] itemPrices_) returns()
func (_Wishlist *WishlistSession) AddWishlistItems(itemPrices_ []*big.Int) (*types.Transaction, error) {
	return _Wishlist.Contract.AddWishlistItems(&_Wishlist.TransactOpts, itemPrices_)
}

// AddWishlistItems is a paid mutator transaction binding the contract method 0x11e06d3a.
//
// Solidity: function addWishlistItems(uint256[] itemPrices_) returns()
func (_Wishlist *WishlistTransactorSession) AddWishlistItems(itemPrices_ []*big.Int) (*types.Transaction, error) {
	return _Wishlist.Contract.AddWishlistItems(&_Wishlist.TransactOpts, itemPrices_)
}

// ChangeWishlistItemsPrice is a paid mutator transaction binding the contract method 0xece23f43.
//
// Solidity: function changeWishlistItemsPrice((uint256,uint256)[] changedItemsDataArr_) returns()
func (_Wishlist *WishlistTransactor) ChangeWishlistItemsPrice(opts *bind.TransactOpts, changedItemsDataArr_ []IWishlistChangedItemPriceData) (*types.Transaction, error) {
	return _Wishlist.contract.Transact(opts, "changeWishlistItemsPrice", changedItemsDataArr_)
}

// ChangeWishlistItemsPrice is a paid mutator transaction binding the contract method 0xece23f43.
//
// Solidity: function changeWishlistItemsPrice((uint256,uint256)[] changedItemsDataArr_) returns()
func (_Wishlist *WishlistSession) ChangeWishlistItemsPrice(changedItemsDataArr_ []IWishlistChangedItemPriceData) (*types.Transaction, error) {
	return _Wishlist.Contract.ChangeWishlistItemsPrice(&_Wishlist.TransactOpts, changedItemsDataArr_)
}

// ChangeWishlistItemsPrice is a paid mutator transaction binding the contract method 0xece23f43.
//
// Solidity: function changeWishlistItemsPrice((uint256,uint256)[] changedItemsDataArr_) returns()
func (_Wishlist *WishlistTransactorSession) ChangeWishlistItemsPrice(changedItemsDataArr_ []IWishlistChangedItemPriceData) (*types.Transaction, error) {
	return _Wishlist.Contract.ChangeWishlistItemsPrice(&_Wishlist.TransactOpts, changedItemsDataArr_)
}

// ContributeFunds is a paid mutator transaction binding the contract method 0x63579b7c.
//
// Solidity: function contributeFunds(uint256 itemId_, uint256 amountToContribute_) returns()
func (_Wishlist *WishlistTransactor) ContributeFunds(opts *bind.TransactOpts, itemId_ *big.Int, amountToContribute_ *big.Int) (*types.Transaction, error) {
	return _Wishlist.contract.Transact(opts, "contributeFunds", itemId_, amountToContribute_)
}

// ContributeFunds is a paid mutator transaction binding the contract method 0x63579b7c.
//
// Solidity: function contributeFunds(uint256 itemId_, uint256 amountToContribute_) returns()
func (_Wishlist *WishlistSession) ContributeFunds(itemId_ *big.Int, amountToContribute_ *big.Int) (*types.Transaction, error) {
	return _Wishlist.Contract.ContributeFunds(&_Wishlist.TransactOpts, itemId_, amountToContribute_)
}

// ContributeFunds is a paid mutator transaction binding the contract method 0x63579b7c.
//
// Solidity: function contributeFunds(uint256 itemId_, uint256 amountToContribute_) returns()
func (_Wishlist *WishlistTransactorSession) ContributeFunds(itemId_ *big.Int, amountToContribute_ *big.Int) (*types.Transaction, error) {
	return _Wishlist.Contract.ContributeFunds(&_Wishlist.TransactOpts, itemId_, amountToContribute_)
}

// InitializeWishlist is a paid mutator transaction binding the contract method 0x56b83b1f.
//
// Solidity: function initializeWishlist(uint256 wishlistId_, address wishlistOwner_, uint256[] initialItemPrices_) returns()
func (_Wishlist *WishlistTransactor) InitializeWishlist(opts *bind.TransactOpts, wishlistId_ *big.Int, wishlistOwner_ common.Address, initialItemPrices_ []*big.Int) (*types.Transaction, error) {
	return _Wishlist.contract.Transact(opts, "initializeWishlist", wishlistId_, wishlistOwner_, initialItemPrices_)
}

// InitializeWishlist is a paid mutator transaction binding the contract method 0x56b83b1f.
//
// Solidity: function initializeWishlist(uint256 wishlistId_, address wishlistOwner_, uint256[] initialItemPrices_) returns()
func (_Wishlist *WishlistSession) InitializeWishlist(wishlistId_ *big.Int, wishlistOwner_ common.Address, initialItemPrices_ []*big.Int) (*types.Transaction, error) {
	return _Wishlist.Contract.InitializeWishlist(&_Wishlist.TransactOpts, wishlistId_, wishlistOwner_, initialItemPrices_)
}

// InitializeWishlist is a paid mutator transaction binding the contract method 0x56b83b1f.
//
// Solidity: function initializeWishlist(uint256 wishlistId_, address wishlistOwner_, uint256[] initialItemPrices_) returns()
func (_Wishlist *WishlistTransactorSession) InitializeWishlist(wishlistId_ *big.Int, wishlistOwner_ common.Address, initialItemPrices_ []*big.Int) (*types.Transaction, error) {
	return _Wishlist.Contract.InitializeWishlist(&_Wishlist.TransactOpts, wishlistId_, wishlistOwner_, initialItemPrices_)
}

// RemoveWishlistItems is a paid mutator transaction binding the contract method 0xc3524f4d.
//
// Solidity: function removeWishlistItems(uint256[] itemIds_) returns()
func (_Wishlist *WishlistTransactor) RemoveWishlistItems(opts *bind.TransactOpts, itemIds_ []*big.Int) (*types.Transaction, error) {
	return _Wishlist.contract.Transact(opts, "removeWishlistItems", itemIds_)
}

// RemoveWishlistItems is a paid mutator transaction binding the contract method 0xc3524f4d.
//
// Solidity: function removeWishlistItems(uint256[] itemIds_) returns()
func (_Wishlist *WishlistSession) RemoveWishlistItems(itemIds_ []*big.Int) (*types.Transaction, error) {
	return _Wishlist.Contract.RemoveWishlistItems(&_Wishlist.TransactOpts, itemIds_)
}

// RemoveWishlistItems is a paid mutator transaction binding the contract method 0xc3524f4d.
//
// Solidity: function removeWishlistItems(uint256[] itemIds_) returns()
func (_Wishlist *WishlistTransactorSession) RemoveWishlistItems(itemIds_ []*big.Int) (*types.Transaction, error) {
	return _Wishlist.Contract.RemoveWishlistItems(&_Wishlist.TransactOpts, itemIds_)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address fundsRecipient_) returns()
func (_Wishlist *WishlistTransactor) WithdrawFunds(opts *bind.TransactOpts, fundsRecipient_ common.Address) (*types.Transaction, error) {
	return _Wishlist.contract.Transact(opts, "withdrawFunds", fundsRecipient_)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address fundsRecipient_) returns()
func (_Wishlist *WishlistSession) WithdrawFunds(fundsRecipient_ common.Address) (*types.Transaction, error) {
	return _Wishlist.Contract.WithdrawFunds(&_Wishlist.TransactOpts, fundsRecipient_)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x68742da6.
//
// Solidity: function withdrawFunds(address fundsRecipient_) returns()
func (_Wishlist *WishlistTransactorSession) WithdrawFunds(fundsRecipient_ common.Address) (*types.Transaction, error) {
	return _Wishlist.Contract.WithdrawFunds(&_Wishlist.TransactOpts, fundsRecipient_)
}

// WishlistFundsCollectionFinishedIterator is returned from FilterFundsCollectionFinished and is used to iterate over the raw logs and unpacked data for FundsCollectionFinished events raised by the Wishlist contract.
type WishlistFundsCollectionFinishedIterator struct {
	Event *WishlistFundsCollectionFinished // Event containing the contract specifics and raw log

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
func (it *WishlistFundsCollectionFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFundsCollectionFinished)
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
		it.Event = new(WishlistFundsCollectionFinished)
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
func (it *WishlistFundsCollectionFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFundsCollectionFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFundsCollectionFinished represents a FundsCollectionFinished event raised by the Wishlist contract.
type WishlistFundsCollectionFinished struct {
	ItemId                   *big.Int
	TotalContributionsNumber *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterFundsCollectionFinished is a free log retrieval operation binding the contract event 0x8f6b950b6d8e13af470e6510f6703514c9fb6fbb79ca7115724f14111ea3db13.
//
// Solidity: event FundsCollectionFinished(uint256 indexed itemId, uint256 totalContributionsNumber)
func (_Wishlist *WishlistFilterer) FilterFundsCollectionFinished(opts *bind.FilterOpts, itemId []*big.Int) (*WishlistFundsCollectionFinishedIterator, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}

	logs, sub, err := _Wishlist.contract.FilterLogs(opts, "FundsCollectionFinished", itemIdRule)
	if err != nil {
		return nil, err
	}
	return &WishlistFundsCollectionFinishedIterator{contract: _Wishlist.contract, event: "FundsCollectionFinished", logs: logs, sub: sub}, nil
}

// WatchFundsCollectionFinished is a free log subscription operation binding the contract event 0x8f6b950b6d8e13af470e6510f6703514c9fb6fbb79ca7115724f14111ea3db13.
//
// Solidity: event FundsCollectionFinished(uint256 indexed itemId, uint256 totalContributionsNumber)
func (_Wishlist *WishlistFilterer) WatchFundsCollectionFinished(opts *bind.WatchOpts, sink chan<- *WishlistFundsCollectionFinished, itemId []*big.Int) (event.Subscription, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}

	logs, sub, err := _Wishlist.contract.WatchLogs(opts, "FundsCollectionFinished", itemIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFundsCollectionFinished)
				if err := _Wishlist.contract.UnpackLog(event, "FundsCollectionFinished", log); err != nil {
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

// ParseFundsCollectionFinished is a log parse operation binding the contract event 0x8f6b950b6d8e13af470e6510f6703514c9fb6fbb79ca7115724f14111ea3db13.
//
// Solidity: event FundsCollectionFinished(uint256 indexed itemId, uint256 totalContributionsNumber)
func (_Wishlist *WishlistFilterer) ParseFundsCollectionFinished(log types.Log) (*WishlistFundsCollectionFinished, error) {
	event := new(WishlistFundsCollectionFinished)
	if err := _Wishlist.contract.UnpackLog(event, "FundsCollectionFinished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistFundsContributedIterator is returned from FilterFundsContributed and is used to iterate over the raw logs and unpacked data for FundsContributed events raised by the Wishlist contract.
type WishlistFundsContributedIterator struct {
	Event *WishlistFundsContributed // Event containing the contract specifics and raw log

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
func (it *WishlistFundsContributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistFundsContributed)
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
		it.Event = new(WishlistFundsContributed)
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
func (it *WishlistFundsContributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistFundsContributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistFundsContributed represents a FundsContributed event raised by the Wishlist contract.
type WishlistFundsContributed struct {
	ItemId            *big.Int
	ContributedAmount *big.Int
	FeeAmount         *big.Int
	ContributionAddr  common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterFundsContributed is a free log retrieval operation binding the contract event 0xbe98b27ee08c5583ca284ecc07244f9b21bb05e6d15139b462bfa212de145c77.
//
// Solidity: event FundsContributed(uint256 indexed itemId, uint256 contributedAmount, uint256 feeAmount, address contributionAddr)
func (_Wishlist *WishlistFilterer) FilterFundsContributed(opts *bind.FilterOpts, itemId []*big.Int) (*WishlistFundsContributedIterator, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}

	logs, sub, err := _Wishlist.contract.FilterLogs(opts, "FundsContributed", itemIdRule)
	if err != nil {
		return nil, err
	}
	return &WishlistFundsContributedIterator{contract: _Wishlist.contract, event: "FundsContributed", logs: logs, sub: sub}, nil
}

// WatchFundsContributed is a free log subscription operation binding the contract event 0xbe98b27ee08c5583ca284ecc07244f9b21bb05e6d15139b462bfa212de145c77.
//
// Solidity: event FundsContributed(uint256 indexed itemId, uint256 contributedAmount, uint256 feeAmount, address contributionAddr)
func (_Wishlist *WishlistFilterer) WatchFundsContributed(opts *bind.WatchOpts, sink chan<- *WishlistFundsContributed, itemId []*big.Int) (event.Subscription, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}

	logs, sub, err := _Wishlist.contract.WatchLogs(opts, "FundsContributed", itemIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistFundsContributed)
				if err := _Wishlist.contract.UnpackLog(event, "FundsContributed", log); err != nil {
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

// ParseFundsContributed is a log parse operation binding the contract event 0xbe98b27ee08c5583ca284ecc07244f9b21bb05e6d15139b462bfa212de145c77.
//
// Solidity: event FundsContributed(uint256 indexed itemId, uint256 contributedAmount, uint256 feeAmount, address contributionAddr)
func (_Wishlist *WishlistFilterer) ParseFundsContributed(log types.Log) (*WishlistFundsContributed, error) {
	event := new(WishlistFundsContributed)
	if err := _Wishlist.contract.UnpackLog(event, "FundsContributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Wishlist contract.
type WishlistInitializedIterator struct {
	Event *WishlistInitialized // Event containing the contract specifics and raw log

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
func (it *WishlistInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistInitialized)
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
		it.Event = new(WishlistInitialized)
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
func (it *WishlistInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistInitialized represents a Initialized event raised by the Wishlist contract.
type WishlistInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Wishlist *WishlistFilterer) FilterInitialized(opts *bind.FilterOpts) (*WishlistInitializedIterator, error) {

	logs, sub, err := _Wishlist.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WishlistInitializedIterator{contract: _Wishlist.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Wishlist *WishlistFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WishlistInitialized) (event.Subscription, error) {

	logs, sub, err := _Wishlist.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistInitialized)
				if err := _Wishlist.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Wishlist *WishlistFilterer) ParseInitialized(log types.Log) (*WishlistInitialized, error) {
	event := new(WishlistInitialized)
	if err := _Wishlist.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistWishlistFundsWithdrawnIterator is returned from FilterWishlistFundsWithdrawn and is used to iterate over the raw logs and unpacked data for WishlistFundsWithdrawn events raised by the Wishlist contract.
type WishlistWishlistFundsWithdrawnIterator struct {
	Event *WishlistWishlistFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *WishlistWishlistFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistWishlistFundsWithdrawn)
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
		it.Event = new(WishlistWishlistFundsWithdrawn)
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
func (it *WishlistWishlistFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistWishlistFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistWishlistFundsWithdrawn represents a WishlistFundsWithdrawn event raised by the Wishlist contract.
type WishlistWishlistFundsWithdrawn struct {
	WithdrawAmount *big.Int
	FundsRecipient common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWishlistFundsWithdrawn is a free log retrieval operation binding the contract event 0x046fd9783c1bef6489a56fede34c9c9300fdb9b8ea95ec5c1cd0828cd6033f35.
//
// Solidity: event WishlistFundsWithdrawn(uint256 withdrawAmount, address fundsRecipient)
func (_Wishlist *WishlistFilterer) FilterWishlistFundsWithdrawn(opts *bind.FilterOpts) (*WishlistWishlistFundsWithdrawnIterator, error) {

	logs, sub, err := _Wishlist.contract.FilterLogs(opts, "WishlistFundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return &WishlistWishlistFundsWithdrawnIterator{contract: _Wishlist.contract, event: "WishlistFundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchWishlistFundsWithdrawn is a free log subscription operation binding the contract event 0x046fd9783c1bef6489a56fede34c9c9300fdb9b8ea95ec5c1cd0828cd6033f35.
//
// Solidity: event WishlistFundsWithdrawn(uint256 withdrawAmount, address fundsRecipient)
func (_Wishlist *WishlistFilterer) WatchWishlistFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *WishlistWishlistFundsWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Wishlist.contract.WatchLogs(opts, "WishlistFundsWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistWishlistFundsWithdrawn)
				if err := _Wishlist.contract.UnpackLog(event, "WishlistFundsWithdrawn", log); err != nil {
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

// ParseWishlistFundsWithdrawn is a log parse operation binding the contract event 0x046fd9783c1bef6489a56fede34c9c9300fdb9b8ea95ec5c1cd0828cd6033f35.
//
// Solidity: event WishlistFundsWithdrawn(uint256 withdrawAmount, address fundsRecipient)
func (_Wishlist *WishlistFilterer) ParseWishlistFundsWithdrawn(log types.Log) (*WishlistWishlistFundsWithdrawn, error) {
	event := new(WishlistWishlistFundsWithdrawn)
	if err := _Wishlist.contract.UnpackLog(event, "WishlistFundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistWishlistItemsAddedIterator is returned from FilterWishlistItemsAdded and is used to iterate over the raw logs and unpacked data for WishlistItemsAdded events raised by the Wishlist contract.
type WishlistWishlistItemsAddedIterator struct {
	Event *WishlistWishlistItemsAdded // Event containing the contract specifics and raw log

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
func (it *WishlistWishlistItemsAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistWishlistItemsAdded)
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
		it.Event = new(WishlistWishlistItemsAdded)
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
func (it *WishlistWishlistItemsAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistWishlistItemsAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistWishlistItemsAdded represents a WishlistItemsAdded event raised by the Wishlist contract.
type WishlistWishlistItemsAdded struct {
	LastItemId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWishlistItemsAdded is a free log retrieval operation binding the contract event 0x702ff69d0f79d3c0b29e2eea147352276c6600bc52f2e532f794fe98625c98f3.
//
// Solidity: event WishlistItemsAdded(uint256 lastItemId)
func (_Wishlist *WishlistFilterer) FilterWishlistItemsAdded(opts *bind.FilterOpts) (*WishlistWishlistItemsAddedIterator, error) {

	logs, sub, err := _Wishlist.contract.FilterLogs(opts, "WishlistItemsAdded")
	if err != nil {
		return nil, err
	}
	return &WishlistWishlistItemsAddedIterator{contract: _Wishlist.contract, event: "WishlistItemsAdded", logs: logs, sub: sub}, nil
}

// WatchWishlistItemsAdded is a free log subscription operation binding the contract event 0x702ff69d0f79d3c0b29e2eea147352276c6600bc52f2e532f794fe98625c98f3.
//
// Solidity: event WishlistItemsAdded(uint256 lastItemId)
func (_Wishlist *WishlistFilterer) WatchWishlistItemsAdded(opts *bind.WatchOpts, sink chan<- *WishlistWishlistItemsAdded) (event.Subscription, error) {

	logs, sub, err := _Wishlist.contract.WatchLogs(opts, "WishlistItemsAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistWishlistItemsAdded)
				if err := _Wishlist.contract.UnpackLog(event, "WishlistItemsAdded", log); err != nil {
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

// ParseWishlistItemsAdded is a log parse operation binding the contract event 0x702ff69d0f79d3c0b29e2eea147352276c6600bc52f2e532f794fe98625c98f3.
//
// Solidity: event WishlistItemsAdded(uint256 lastItemId)
func (_Wishlist *WishlistFilterer) ParseWishlistItemsAdded(log types.Log) (*WishlistWishlistItemsAdded, error) {
	event := new(WishlistWishlistItemsAdded)
	if err := _Wishlist.contract.UnpackLog(event, "WishlistItemsAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WishlistWishlistItemsRemovedIterator is returned from FilterWishlistItemsRemoved and is used to iterate over the raw logs and unpacked data for WishlistItemsRemoved events raised by the Wishlist contract.
type WishlistWishlistItemsRemovedIterator struct {
	Event *WishlistWishlistItemsRemoved // Event containing the contract specifics and raw log

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
func (it *WishlistWishlistItemsRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WishlistWishlistItemsRemoved)
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
		it.Event = new(WishlistWishlistItemsRemoved)
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
func (it *WishlistWishlistItemsRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WishlistWishlistItemsRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WishlistWishlistItemsRemoved represents a WishlistItemsRemoved event raised by the Wishlist contract.
type WishlistWishlistItemsRemoved struct {
	RemovedItemIds []*big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWishlistItemsRemoved is a free log retrieval operation binding the contract event 0x1d86218a4dca4e462d2c953ff74b90cebd8a27b84bf78a45a4f9614973418a41.
//
// Solidity: event WishlistItemsRemoved(uint256[] removedItemIds)
func (_Wishlist *WishlistFilterer) FilterWishlistItemsRemoved(opts *bind.FilterOpts) (*WishlistWishlistItemsRemovedIterator, error) {

	logs, sub, err := _Wishlist.contract.FilterLogs(opts, "WishlistItemsRemoved")
	if err != nil {
		return nil, err
	}
	return &WishlistWishlistItemsRemovedIterator{contract: _Wishlist.contract, event: "WishlistItemsRemoved", logs: logs, sub: sub}, nil
}

// WatchWishlistItemsRemoved is a free log subscription operation binding the contract event 0x1d86218a4dca4e462d2c953ff74b90cebd8a27b84bf78a45a4f9614973418a41.
//
// Solidity: event WishlistItemsRemoved(uint256[] removedItemIds)
func (_Wishlist *WishlistFilterer) WatchWishlistItemsRemoved(opts *bind.WatchOpts, sink chan<- *WishlistWishlistItemsRemoved) (event.Subscription, error) {

	logs, sub, err := _Wishlist.contract.WatchLogs(opts, "WishlistItemsRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WishlistWishlistItemsRemoved)
				if err := _Wishlist.contract.UnpackLog(event, "WishlistItemsRemoved", log); err != nil {
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

// ParseWishlistItemsRemoved is a log parse operation binding the contract event 0x1d86218a4dca4e462d2c953ff74b90cebd8a27b84bf78a45a4f9614973418a41.
//
// Solidity: event WishlistItemsRemoved(uint256[] removedItemIds)
func (_Wishlist *WishlistFilterer) ParseWishlistItemsRemoved(log types.Log) (*WishlistWishlistItemsRemoved, error) {
	event := new(WishlistWishlistItemsRemoved)
	if err := _Wishlist.contract.UnpackLog(event, "WishlistItemsRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
