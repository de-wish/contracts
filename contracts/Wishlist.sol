// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

import {PERCENTAGE_100} from "@solarity/solidity-lib/utils/Globals.sol";

import {IWishlist} from "./interfaces/IWishlist.sol";
import {IWishlistFactory} from "./interfaces/IWishlistFactory.sol";

contract Wishlist is IWishlist, Initializable {
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.UintSet;

    IWishlistFactory public factory;
    IERC20 public usdcToken;

    uint256 public wishlistId;
    address public wishlistOwner;

    uint256 public nextItemId;

    EnumerableSet.UintSet internal _activeItemIds;

    mapping(uint256 => WishlistItemData) internal _itemsData;

    modifier onlyWishlistOwner() {
        _onlyWishlistOwner();
        _;
    }

    function initializeWishlist(
        uint256 wishlistId_,
        address wishlistOwner_,
        uint256[] calldata initialItemPrices_
    ) external initializer {
        wishlistId = wishlistId_;
        wishlistOwner = wishlistOwner_;

        IWishlistFactory factory_ = IWishlistFactory(msg.sender);

        factory = factory_;
        usdcToken = factory_.usdcToken();

        if (initialItemPrices_.length > 0) {
            _addWishlistItems(initialItemPrices_);
        }
    }

    function addWishlistItems(uint256[] calldata itemPrices_) external onlyWishlistOwner {
        _addWishlistItems(itemPrices_);
    }

    function removeWishlistItems(uint256[] calldata itemIds_) external onlyWishlistOwner {
        _removeWishlistItems(itemIds_);
    }

    function changeWishlistItemsPrice(
        ChangedItemPriceData[] calldata changedItemsDataArr_
    ) external onlyWishlistOwner {
        _changeWishlistItemsPrice(changedItemsDataArr_);
    }

    function withdrawFunds(address fundsRecipient_) external onlyWishlistOwner {
        uint256 amountToWithdraw_ = usdcToken.balanceOf(address(this));

        require(amountToWithdraw_ > 0, ZeroAmountToWithdraw());

        usdcToken.safeTransfer(fundsRecipient_, amountToWithdraw_);

        emit WishlistFundsWithdrawn(amountToWithdraw_, fundsRecipient_);
    }

    function buyWishlistItem(uint256 itemId_) external {
        _checkActiveItem(itemId_);

        uint256 itemPrice_ = _itemsData[itemId_].itemPrice;
        uint256 protocolFee_ = _countFeeAmount(itemPrice_);

        usdcToken.safeTransferFrom(msg.sender, address(this), itemPrice_);
        usdcToken.safeTransferFrom(msg.sender, address(factory), protocolFee_);

        _activeItemIds.remove(itemId_);
        _itemsData[itemId_].buyerAddr = msg.sender;

        emit ItemBought(itemId_, itemPrice_, protocolFee_, msg.sender);
    }

    function getWishlistInfo() external view returns (WishlistInfo memory) {
        return
            WishlistInfo(
                wishlistId,
                wishlistOwner,
                usdcToken.balanceOf(address(this)),
                nextItemId,
                getWishlistItemsInfo(_activeItemIds.values())
            );
    }

    function getWishlistItemsInfo(
        uint256[] memory itemIds_
    ) public view returns (WishlistItemInfo[] memory resultArr_) {
        resultArr_ = new WishlistItemInfo[](itemIds_.length);

        for (uint256 i = 0; i < itemIds_.length; i++) {
            WishlistItemData memory itemData_ = _itemsData[itemIds_[i]];

            resultArr_[i] = WishlistItemInfo(
                itemIds_[i],
                itemData_.itemPrice,
                _countFeeAmount(itemData_.itemPrice),
                itemData_.buyerAddr,
                itemData_.buyerAddr == address(0)
            );
        }
    }

    function isItemActive(uint256 itemId_) public view returns (bool) {
        return _activeItemIds.contains(itemId_);
    }

    function _addWishlistItems(uint256[] calldata itemPrices_) internal {
        require(itemPrices_.length > 0, ZeroItemsToAdd());

        uint256 nextItemId_ = nextItemId;

        for (uint256 i = 0; i < itemPrices_.length; i++) {
            _activeItemIds.add(nextItemId_);

            _changeItemPrice(_itemsData[nextItemId_++], itemPrices_[i]);
        }

        nextItemId = nextItemId_;

        emit WishlistItemsAdded(nextItemId_ - 1);
    }

    function _removeWishlistItems(uint256[] calldata itemIds_) internal {
        for (uint256 i = 0; i < itemIds_.length; i++) {
            uint256 itemIdToRemove_ = itemIds_[i];

            _checkActiveItem(itemIdToRemove_);

            _activeItemIds.remove(itemIdToRemove_);
            delete _itemsData[itemIdToRemove_].itemPrice;
        }

        emit WishlistItemsRemoved(itemIds_);
    }

    function _changeWishlistItemsPrice(
        ChangedItemPriceData[] calldata changedItemsDataArr_
    ) internal {
        for (uint256 i = 0; i < changedItemsDataArr_.length; i++) {
            ChangedItemPriceData calldata currentData_ = changedItemsDataArr_[i];

            _checkActiveItem(currentData_.itemId);
            _changeItemPrice(_itemsData[currentData_.itemId], currentData_.newItemPrice);
        }
    }

    function _changeItemPrice(WishlistItemData storage _itemData, uint256 newPrice_) internal {
        require(newPrice_ > 0, ZeroItemPrice());

        _itemData.itemPrice = newPrice_;
    }

    function _countFeeAmount(uint256 itemPrice_) internal view returns (uint256) {
        uint256 protocolFee_ = (itemPrice_ * factory.protocolFeePercentage()) / PERCENTAGE_100;

        return Math.min(protocolFee_, factory.maxProtocolFeeAmount());
    }

    function _onlyWishlistOwner() internal view {
        require(wishlistOwner == msg.sender, NotAWishlistOwner(msg.sender));
    }

    function _checkActiveItem(uint256 itemId_) internal view {
        require(isItemActive(itemId_), NotAnActiveItem(itemId_));
    }
}
