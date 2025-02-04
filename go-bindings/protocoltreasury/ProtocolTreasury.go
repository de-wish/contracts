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
	Bin: "0x60a060405230608052348015601357600080fd5b5060805161130361004b600039600081816101f801528181610241015281816102e00152818161032001526103b301526113036000f3fe6080604052600436106100a05760003560e01c80638da5cb5b116100645780638da5cb5b14610133578063a5662f851461015b578063c499bee61461017b578063caa2f7a01461019b578063f2fde38b146101ae578063fc88a027146101ce57600080fd5b80633659cfe6146100ac5780634f1ef286146100ce57806352d1902d146100e1578063715018a6146101095780638129fc1c1461011e57600080fd5b366100a757005b600080fd5b3480156100b857600080fd5b506100cc6100c7366004610f62565b6101ee565b005b6100cc6100dc366004610f93565b6102d6565b3480156100ed57600080fd5b506100f66103a6565b6040519081526020015b60405180910390f35b34801561011557600080fd5b506100cc610459565b34801561012a57600080fd5b506100cc61046d565b34801561013f57600080fd5b506097546040516001600160a01b039091168152602001610100565b34801561016757600080fd5b506100cc61017636600461106b565b61057d565b34801561018757600080fd5b506100cc6101963660046110ab565b6105e6565b3480156101a757600080fd5b50476100f6565b3480156101ba57600080fd5b506100cc6101c9366004610f62565b610661565b3480156101da57600080fd5b506100f66101e9366004610f62565b6106d7565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361023f5760405162461bcd60e51b8152600401610236906110fa565b60405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610288600080516020611287833981519152546001600160a01b031690565b6001600160a01b0316146102ae5760405162461bcd60e51b815260040161023690611146565b6102b781610748565b604080516000808252602082019092526102d391839190610750565b50565b6001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016300361031e5760405162461bcd60e51b8152600401610236906110fa565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610367600080516020611287833981519152546001600160a01b031690565b6001600160a01b03161461038d5760405162461bcd60e51b815260040161023690611146565b61039682610748565b6103a282826001610750565b5050565b6000306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104465760405162461bcd60e51b815260206004820152603860248201527f555550535570677261646561626c653a206d757374206e6f742062652063616c60448201527f6c6564207468726f7567682064656c656761746563616c6c00000000000000006064820152608401610236565b5060008051602061128783398151915290565b6104616108c0565b61046b600061091a565b565b600054610100900460ff161580801561048d5750600054600160ff909116105b806104a75750303b1580156104a7575060005460ff166001145b61050a5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608401610236565b6000805460ff19166001179055801561052d576000805461ff0019166101001790555b61053561096c565b80156102d3576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a150565b6105856108c0565b806105905781610592565b475b915061059e838361099b565b826001600160a01b03167f5317b81e67f8cd60e038f9dc090c51c9401b1fbae411b50e2524794180844835836040516105d991815260200190565b60405180910390a2505050565b6105ee6108c0565b806105f95781610602565b610602846106d7565b915061060f848484610ab4565b826001600160a01b0316846001600160a01b03167ee763f7778b8ceef7270c89b7d1df1008b0e482da39c43831417733af96fb0d8460405161065391815260200190565b60405180910390a350505050565b6106696108c0565b6001600160a01b0381166106ce5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610236565b6102d38161091a565b6040516370a0823160e01b81523060048201526000906001600160a01b038316906370a0823190602401602060405180830381865afa15801561071e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107429190611192565b92915050565b6102d36108c0565b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd91435460ff16156107885761078383610b06565b505050565b826001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156107e2575060408051601f3d908101601f191682019092526107df91810190611192565b60015b6108455760405162461bcd60e51b815260206004820152602e60248201527f45524331393637557067726164653a206e657720696d706c656d656e7461746960448201526d6f6e206973206e6f74205555505360901b6064820152608401610236565b60008051602061128783398151915281146108b45760405162461bcd60e51b815260206004820152602960248201527f45524331393637557067726164653a20756e737570706f727465642070726f786044820152681a58589b195555525160ba1b6064820152608401610236565b50610783838383610ba2565b6097546001600160a01b0316331461046b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610236565b609780546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff166109935760405162461bcd60e51b8152600401610236906111ab565b61046b610bcd565b804710156109eb5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610236565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114610a38576040519150601f19603f3d011682016040523d82523d6000602084013e610a3d565b606091505b50509050806107835760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610236565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b179052610783908490610bfd565b6001600160a01b0381163b610b735760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610236565b60008051602061128783398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b610bab83610cd2565b600082511180610bb85750805b1561078357610bc78383610d12565b50505050565b600054610100900460ff16610bf45760405162461bcd60e51b8152600401610236906111ab565b61046b3361091a565b6000610c52826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316610d3e9092919063ffffffff16565b9050805160001480610c73575080806020019051810190610c7391906111f6565b6107835760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610236565b610cdb81610b06565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b6060610d3783836040518060600160405280602781526020016112a760279139610d55565b9392505050565b6060610d4d8484600085610dcd565b949350505050565b6060600080856001600160a01b031685604051610d729190611237565b600060405180830381855af49150503d8060008114610dad576040519150601f19603f3d011682016040523d82523d6000602084013e610db2565b606091505b5091509150610dc386838387610ea8565b9695505050505050565b606082471015610e2e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610236565b600080866001600160a01b03168587604051610e4a9190611237565b60006040518083038185875af1925050503d8060008114610e87576040519150601f19603f3d011682016040523d82523d6000602084013e610e8c565b606091505b5091509150610e9d87838387610ea8565b979650505050505050565b60608315610f17578251600003610f10576001600160a01b0385163b610f105760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610236565b5081610d4d565b610d4d8383815115610f2c5781518083602001fd5b8060405162461bcd60e51b81526004016102369190611253565b80356001600160a01b0381168114610f5d57600080fd5b919050565b600060208284031215610f7457600080fd5b610d3782610f46565b634e487b7160e01b600052604160045260246000fd5b60008060408385031215610fa657600080fd5b610faf83610f46565b9150602083013567ffffffffffffffff811115610fcb57600080fd5b8301601f81018513610fdc57600080fd5b803567ffffffffffffffff811115610ff657610ff6610f7d565b604051601f8201601f19908116603f0116810167ffffffffffffffff8111828210171561102557611025610f7d565b60405281815282820160200187101561103d57600080fd5b816020840160208301376000602083830101528093505050509250929050565b80151581146102d357600080fd5b60008060006060848603121561108057600080fd5b61108984610f46565b92506020840135915060408401356110a08161105d565b809150509250925092565b600080600080608085870312156110c157600080fd5b6110ca85610f46565b93506110d860208601610f46565b92506040850135915060608501356110ef8161105d565b939692955090935050565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b6000602082840312156111a457600080fd5b5051919050565b6020808252602b908201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960408201526a6e697469616c697a696e6760a81b606082015260800190565b60006020828403121561120857600080fd5b8151610d378161105d565b60005b8381101561122e578181015183820152602001611216565b50506000910152565b60008251611249818460208701611213565b9190910192915050565b6020815260008251806020840152611272816040850160208701611213565b601f01601f1916919091016040019291505056fe360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220ce057079165a78aab712a88545d772fd5a7e5ef669b826f41014a7199f34f6a764736f6c634300081c0033",
}

// ProtocolTreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use ProtocolTreasuryMetaData.ABI instead.
var ProtocolTreasuryABI = ProtocolTreasuryMetaData.ABI

// ProtocolTreasuryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProtocolTreasuryMetaData.Bin instead.
var ProtocolTreasuryBin = ProtocolTreasuryMetaData.Bin

// DeployProtocolTreasury deploys a new Ethereum contract, binding an instance of ProtocolTreasury to it.
func DeployProtocolTreasury(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProtocolTreasury, error) {
	parsed, err := ProtocolTreasuryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProtocolTreasuryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProtocolTreasury{ProtocolTreasuryCaller: ProtocolTreasuryCaller{contract: contract}, ProtocolTreasuryTransactor: ProtocolTreasuryTransactor{contract: contract}, ProtocolTreasuryFilterer: ProtocolTreasuryFilterer{contract: contract}}, nil
}

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
