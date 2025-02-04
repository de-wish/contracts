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
    using EnumerableSet for *;

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

    function contributeFunds(uint256 itemId_, uint256 amountToContribute_) external {
        _checkActiveItem(itemId_);

        WishlistItemData storage _itemData = _itemsData[itemId_];

        uint256 remainingAmount_ = _itemData.itemPrice - _itemData.collectedTokensAmount;

        if (amountToContribute_ > remainingAmount_) {
            amountToContribute_ = remainingAmount_;
        }

        uint256 protocolFee_ = _transferProtocolFee(amountToContribute_);

        usdcToken.safeTransferFrom(msg.sender, address(this), amountToContribute_);

        _itemData.collectedTokensAmount += amountToContribute_;

        _itemData.contributrionsData[_itemData.totalContributionsNumber++] = ContributionData(
            msg.sender,
            amountToContribute_
        );

        emit FundsContributed(itemId_, amountToContribute_, protocolFee_, msg.sender);

        if (amountToContribute_ == remainingAmount_) {
            _activeItemIds.remove(itemId_);

            emit FundsCollectionFinished(itemId_, _itemData.totalContributionsNumber);
        }
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
            WishlistItemData storage _itemData = _itemsData[itemIds_[i]];

            resultArr_[i] = WishlistItemInfo(
                itemIds_[i],
                _itemData.itemPrice,
                _itemData.collectedTokensAmount,
                _itemData.totalContributionsNumber,
                getContributionsInfo(itemIds_[i]),
                isItemActive(itemIds_[i])
            );
        }
    }

    function getContributionsInfo(
        uint256 itemId_
    ) public view returns (ContributionInfo[] memory resultArr_) {
        WishlistItemData storage _itemData = _itemsData[itemId_];

        resultArr_ = new ContributionInfo[](_itemData.totalContributionsNumber);

        for (uint256 i = 0; i < resultArr_.length; i++) {
            resultArr_[i] = ContributionInfo(
                i,
                _itemData.contributrionsData[i].contributor,
                _itemData.contributrionsData[i].contributionAmount
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

    function _transferProtocolFee(
        uint256 contributeAmount_
    ) internal returns (uint256 protocolFee_) {
        IWishlistFactory.ProtocolFeeSettings memory feeSettings_ = factory
            .getProtocolFeeSettings();

        protocolFee_ = (contributeAmount_ * feeSettings_.protocolFeePercentage) / PERCENTAGE_100;

        Math.min(protocolFee_, feeSettings_.maxProtocolFeeAmount);

        usdcToken.safeTransferFrom(msg.sender, feeSettings_.protocolFeeRecipient, protocolFee_);
    }

    function _onlyWishlistOwner() internal view {
        require(wishlistOwner == msg.sender, NotAWishlistOwner(msg.sender));
    }

    function _checkActiveItem(uint256 itemId_) internal view {
        require(isItemActive(itemId_), NotAnActiveItem(itemId_));
    }
}
