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
	Bin: "0x6080604052348015600f57600080fd5b506119f98061001f6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806363579b7c116100975780638de2041c116100665780638de2041c14610204578063c3524f4d14610227578063c45a01551461023a578063ece23f431461025357600080fd5b806363579b7c146101b557806368742da6146101c85780636a868974146101db5780636abf1822146101e457600080fd5b806315604773116100d357806315604773146101635780632493c84014610178578063398ef7f81461018f57806356b83b1f146101a257600080fd5b806306effc75146100fa57806311e06d3a1461012357806311eac85514610138575b600080fd5b61010d6101083660046112c9565b610266565b60405161011a9190611396565b60405180910390f35b6101366101313660046114c6565b6103ef565b005b60015461014b906001600160a01b031681565b6040516001600160a01b03909116815260200161011a565b61016b610405565b60405161011a9190611508565b61018160025481565b60405190815260200161011a565b60035461014b906001600160a01b031681565b6101366101b036600461162f565b6104ea565b6101366101c336600461168b565b6106cb565b6101366101d63660046116ad565b61084a565b61018160045481565b6101f76101f23660046116ca565b61093f565b60405161011a91906116e3565b6102176102123660046116ca565b610a46565b604051901515815260200161011a565b6101366102353660046114c6565b610a59565b60005461014b906201000090046001600160a01b031681565b61013661026136600461172c565b610a6b565b6060815167ffffffffffffffff81111561028257610282611282565b6040519080825280602002602001820160405280156102ee57816020015b6102db6040518060c0016040528060008152602001600081526020016000815260200160008152602001606081526020016000151581525090565b8152602001906001900390816102a05790505b50905060005b82518110156103e957600060076000858481518110610315576103156117a3565b6020026020010151815260200190815260200160002090506040518060c0016040528085848151811061034a5761034a6117a3565b6020026020010151815260200182600001548152602001826001015481526020018260020154815260200161039786858151811061038a5761038a6117a3565b602002602001015161093f565b81526020016103be8685815181106103b1576103b16117a3565b6020026020010151610a46565b15158152508383815181106103d5576103d56117a3565b6020908102919091010152506001016102f4565b50919050565b6103f7610a7d565b6104018282610ab9565b5050565b6104406040518060a001604052806000815260200160006001600160a01b031681526020016000815260200160008152602001606081525090565b6040805160a08101825260025481526003546001600160a01b03908116602083015260015483516370a0823160e01b815230600482015292938401929116906370a0823190602401602060405180830381865afa1580156104a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104c991906117b9565b815260200160045481526020016104e36101086005610b80565b9052919050565b600054610100900460ff161580801561050a5750600054600160ff909116105b806105245750303b158015610524575060005460ff166001145b61058c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084015b60405180910390fd5b6000805460ff1916600117905580156105af576000805461ff0019166101001790555b6002859055600380546001600160a01b0319166001600160a01b0386161790556000805462010000600160b01b0319163362010000810291909117909155604080516311eac85560e01b8152905182916311eac8559160048083019260209291908290030181865afa158015610629573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064d91906117d2565b600180546001600160a01b0319166001600160a01b0392909216919091179055821561067d5761067d8484610ab9565b5080156106c4576000805461ff0019169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050505050565b6106d482610b94565b6000828152600760205260408120600181015481549192916106f69190611805565b905080831115610704578092505b600061070f84610bbf565b60015490915061072a906001600160a01b0316333087610c8d565b8383600101600082825461073e9190611818565b90915550506040805180820190915233815260208101859052600284018054600386019160009190826107708361182b565b9091555081526020808201929092526040908101600020835181546001600160a01b0319166001600160a01b03909116178155928201516001909301929092558151868152908101839052339181019190915285907fbe98b27ee08c5583ca284ecc07244f9b21bb05e6d15139b462bfa212de145c779060600160405180910390a28184036106c457610804600586610cfe565b50847f8f6b950b6d8e13af470e6510f6703514c9fb6fbb79ca7115724f14111ea3db13846002015460405161083b91815260200190565b60405180910390a25050505050565b610852610a7d565b6001546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa15801561089b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bf91906117b9565b9050600081116108e25760405163fd8fd35d60e01b815260040160405180910390fd5b6001546108f9906001600160a01b03168383610d0a565b604080518281526001600160a01b03841660208201527f046fd9783c1bef6489a56fede34c9c9300fdb9b8ea95ec5c1cd0828cd6033f3591015b60405180910390a15050565b600081815260076020526040902060028101546060919067ffffffffffffffff81111561096e5761096e611282565b6040519080825280602002602001820160405280156109cc57816020015b6109b960405180606001604052806000815260200160006001600160a01b03168152602001600081525090565b81526020019060019003908161098c5790505b50915060005b8251811015610a3f5760408051606081018252828152600083815260038501602081815284832080546001600160a01b031682860152928690525260010154918101919091528351849083908110610a2c57610a2c6117a3565b60209081029190910101526001016109d2565b5050919050565b6000610a53600583610d3f565b92915050565b610a61610a7d565b6104018282610d57565b610a73610a7d565b6104018282610ddd565b60035433906001600160a01b03168114610ab657604051631a18e31960e21b81526001600160a01b039091166004820152602401610583565b50565b80610ad757604051635291d13360e01b815260040160405180910390fd5b60045460005b82811015610b3857610af0600583610e38565b50610b306007600084610b028161182b565b95508152602001908152602001600020858584818110610b2457610b246117a3565b90506020020135610e44565b600101610add565b5060048190557f702ff69d0f79d3c0b29e2eea147352276c6600bc52f2e532f794fe98625c98f3610b6a600183611805565b60405190815260200160405180910390a1505050565b60606000610b8d83610e69565b9392505050565b610b9d81610a46565b8190610401576040516333a06fe760e11b815260040161058391815260200190565b600080600060029054906101000a90046001600160a01b03166001600160a01b0316637e9991066040518163ffffffff1660e01b8152600401606060405180830381865afa158015610c15573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c399190611844565b80519091506b033b2e3c9fd0803ce800000090610c5690856118a6565b610c6091906118bd565b9150610c70828260200151610ec5565b5060408101516001546103e9916001600160a01b03909116903390855b6040516001600160a01b0380851660248301528316604482015260648101829052610cf89085906323b872dd60e01b906084015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610edb565b50505050565b6000610b8d8383610fb0565b6040516001600160a01b038316602482015260448101829052610d3a90849063a9059cbb60e01b90606401610cc1565b505050565b60008181526001830160205260408120541515610b8d565b60005b81811015610dab576000838383818110610d7657610d766117a3565b905060200201359050610d8881610b94565b610d93600582610cfe565b50600090815260076020526040812055600101610d5a565b507f1d86218a4dca4e462d2c953ff74b90cebd8a27b84bf78a45a4f9614973418a4182826040516109339291906118df565b60005b81811015610d3a5736838383818110610dfb57610dfb6117a3565b9050604002019050610e108160000135610b94565b80356000908152600760209081526040909120610e2f91830135610e44565b50600101610de0565b6000610b8d83836110a3565b60008111610e6557604051639b3055a960e01b815260040160405180910390fd5b9055565b606081600001805480602002602001604051908101604052809291908181526020018280548015610eb957602002820191906000526020600020905b815481526020019060010190808311610ea5575b50505050509050919050565b6000818310610ed45781610b8d565b5090919050565b6000610f30826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166110f29092919063ffffffff16565b9050805160001480610f51575080806020019051810190610f519190611918565b610d3a5760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610583565b60008181526001830160205260408120548015611099576000610fd4600183611805565b8554909150600090610fe890600190611805565b905081811461104d576000866000018281548110611008576110086117a3565b906000526020600020015490508087600001848154811061102b5761102b6117a3565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061105e5761105e61193a565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610a53565b6000915050610a53565b60008181526001830160205260408120546110ea57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610a53565b506000610a53565b60606111018484600085611109565b949350505050565b60608247101561116a5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610583565b600080866001600160a01b031685876040516111869190611974565b60006040518083038185875af1925050503d80600081146111c3576040519150601f19603f3d011682016040523d82523d6000602084013e6111c8565b606091505b50915091506111d9878383876111e4565b979650505050505050565b6060831561125357825160000361124c576001600160a01b0385163b61124c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610583565b5081611101565b61110183838151156112685781518083602001fd5b8060405162461bcd60e51b81526004016105839190611990565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff811182821017156112c1576112c1611282565b604052919050565b6000602082840312156112db57600080fd5b813567ffffffffffffffff8111156112f257600080fd5b8201601f8101841361130357600080fd5b803567ffffffffffffffff81111561131d5761131d611282565b8060051b61132d60208201611298565b9182526020818401810192908101908784111561134957600080fd5b6020850194505b838510156111d957843580835260209586019590935090910190611350565b805182526020808201516001600160a01b0316908301526040908101519082015260600190565b6000602082016020835280845180835260408501915060408160051b86010192506020860160005b8281101561146e57603f19878603018452815160c0860181518752602082015160208801526040820151604088015260608201516060880152608082015160c0608089015281815180845260e08a019150602083019350600092505b808310156114405761142d82855161136f565b915060208401935060018301925061141a565b5060a0840151935061145660a08a018515159052565b975050506020948501949290920191506001016113be565b50929695505050505050565b60008083601f84011261148c57600080fd5b50813567ffffffffffffffff8111156114a457600080fd5b6020830191508360208260051b85010111156114bf57600080fd5b9250929050565b600080602083850312156114d957600080fd5b823567ffffffffffffffff8111156114f057600080fd5b6114fc8582860161147a565b90969095509350505050565b60208152600060c082018351602084015260018060a01b0360208501511660408401526040840151606084015260608401516080840152608084015160a08085015281815180845260e08601915060e08160051b870101935060208301925060005b8181101561146e5760df19878603018352835160c0860181518752602082015160208801526040820151604088015260608201516060880152608082015160c0608089015281815180845260e08a019150602083019350600092505b808310156115ec576115d982855161136f565b91506020840193506001830192506115c6565b5060a0840151935061160260a08a018515159052565b9750505060209485019493909301925060010161156a565b6001600160a01b0381168114610ab657600080fd5b6000806000806060858703121561164557600080fd5b8435935060208501356116578161161a565b9250604085013567ffffffffffffffff81111561167357600080fd5b61167f8782880161147a565b95989497509550505050565b6000806040838503121561169e57600080fd5b50508035926020909101359150565b6000602082840312156116bf57600080fd5b8135610b8d8161161a565b6000602082840312156116dc57600080fd5b5035919050565b602080825282518282018190526000918401906040840190835b818110156117215761171083855161136f565b6020949094019392506001016116fd565b509095945050505050565b6000806020838503121561173f57600080fd5b823567ffffffffffffffff81111561175657600080fd5b8301601f8101851361176757600080fd5b803567ffffffffffffffff81111561177e57600080fd5b8560208260061b840101111561179357600080fd5b6020919091019590945092505050565b634e487b7160e01b600052603260045260246000fd5b6000602082840312156117cb57600080fd5b5051919050565b6000602082840312156117e457600080fd5b8151610b8d8161161a565b634e487b7160e01b600052601160045260246000fd5b81810381811115610a5357610a536117ef565b80820180821115610a5357610a536117ef565b60006001820161183d5761183d6117ef565b5060010190565b6000606082840312801561185757600080fd5b506040516060810167ffffffffffffffff8111828210171561187b5761187b611282565b6040908152835182526020808501519083015283015161189a8161161a565b60408201529392505050565b8082028115828204841417610a5357610a536117ef565b6000826118da57634e487b7160e01b600052601260045260246000fd5b500490565b6020808252810182905260006001600160fb1b038311156118ff57600080fd5b8260051b80856040850137919091016040019392505050565b60006020828403121561192a57600080fd5b81518015158114610b8d57600080fd5b634e487b7160e01b600052603160045260246000fd5b60005b8381101561196b578181015183820152602001611953565b50506000910152565b60008251611986818460208701611950565b9190910192915050565b60208152600082518060208401526119af816040850160208701611950565b601f01601f1916919091016040019291505056fea26469706673582212207f7874071e3c7323238ea22f5b724c603912487e0576f5c33b5b497ac245840164736f6c634300081c0033",
}

// WishlistABI is the input ABI used to generate the binding from.
// Deprecated: Use WishlistMetaData.ABI instead.
var WishlistABI = WishlistMetaData.ABI

// WishlistBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WishlistMetaData.Bin instead.
var WishlistBin = WishlistMetaData.Bin

// DeployWishlist deploys a new Ethereum contract, binding an instance of Wishlist to it.
func DeployWishlist(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Wishlist, error) {
	parsed, err := WishlistMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WishlistBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wishlist{WishlistCaller: WishlistCaller{contract: contract}, WishlistTransactor: WishlistTransactor{contract: contract}, WishlistFilterer: WishlistFilterer{contract: contract}}, nil
}

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
