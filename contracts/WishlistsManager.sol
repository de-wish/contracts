// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {PRECISION, PERCENTAGE_100} from "@solarity/solidity-lib/utils/Globals.sol";

import {IWishlistManager} from "./interfaces/IWishlistManager.sol";

contract WishlistManager is IWishlistManager, Ownable {
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.UintSet;

    uint256 public constant MAX_PROTOCOL_FEE = 5 * PRECISION;

    IERC20 public immutable usdcToken;

    uint256 public nextWishlistId;

    uint256 public protocolFee;
    address public protocolFeeRecipient;

    mapping(uint256 => WishlistData) internal _wishlistsData;
    mapping(address => EnumerableSet.UintSet) internal _userWishlists;

    modifier onlyWishlistOwner(uint256 wishlistId_) {
        _onlyWishlistOwner(wishlistId_);
        _;
    }

    constructor(
        address ownerAddr_,
        address usdcTokenAddr_,
        uint256 protocolFee_
    ) Ownable(ownerAddr_) {
        usdcToken = IERC20(usdcTokenAddr_);

        _setProtocolFee(protocolFee_);

        nextWishlistId = 1;
    }

    function setProtocolFee(uint256 newProtocolFee_) external onlyOwner {
        _setProtocolFee(newProtocolFee_);
    }

    function setProtocolFeeRecipient(address newProtocolFeeRecipient_) external onlyOwner {
        _setProtocolFeeRecipient(newProtocolFeeRecipient_);
    }

    function createWishlist(uint256[] calldata itemPrices_) external {
        uint256 newWishlistId_ = nextWishlistId++;

        WishlistData storage _newWishlistData = _wishlistsData[newWishlistId_];

        _newWishlistData.ownerAddr = msg.sender;

        if (itemPrices_.length > 0) {
            _addWishlistItems(newWishlistId_, itemPrices_);
        }

        _userWishlists[msg.sender].add(newWishlistId_);
    }

    function addWishlistItems(
        uint256 wishlistId_,
        uint256[] calldata itemPrices_
    ) external onlyWishlistOwner(wishlistId_) {
        _addWishlistItems(wishlistId_, itemPrices_);
    }

    function removeWishlistItems(
        uint256 wishlistId_,
        uint256[] calldata itemIds_
    ) external onlyWishlistOwner(wishlistId_) {
        _removeWishlistItems(wishlistId_, itemIds_);
    }

    function changeWishlistItemsPrice(
        uint256 wishlistId_,
        ChangedItemPriceData[] calldata changedItemsDataArr_
    ) external onlyWishlistOwner(wishlistId_) {
        _changeWishlistItemsPrice(wishlistId_, changedItemsDataArr_);
    }

    function withdrawFunds(
        uint256 wishlistId_,
        address fundsRecipient_
    ) external onlyWishlistOwner(wishlistId_) {
        WishlistData storage _wishlistData = _wishlistsData[wishlistId_];

        uint256 amountToWithdraw_ = _wishlistData.owedUsdcAmount;

        delete _wishlistData.owedUsdcAmount;

        usdcToken.safeTransfer(fundsRecipient_, amountToWithdraw_);

        emit WishlistFundsWithdrawn(wishlistId_, amountToWithdraw_, fundsRecipient_);
    }

    function buyWishlistItem(uint256 wishlistId_, uint256 itemId_) external {
        WishlistData storage _wishlistData = _wishlistsData[wishlistId_];

        _checkActiveItem(wishlistId_, itemId_);

        uint256 itemPrice_ = _wishlistData.itemsData[itemId_].itemPrice;

        usdcToken.safeTransferFrom(msg.sender, protocolFeeRecipient, _countFeeAmount(itemPrice_));
        usdcToken.safeTransferFrom(msg.sender, address(this), itemPrice_);

        _wishlistData.owedUsdcAmount += itemPrice_;
        _wishlistData.activeItemIds.remove(itemId_);
        _wishlistData.itemsData[itemId_].buyerAddr = msg.sender;

        emit ItemBought(wishlistId_, itemId_, itemPrice_, msg.sender);
    }

    function getWishlistInfo(uint256 wishlistId_) external view returns (WishlistInfo memory) {
        WishlistData storage _wishlistData = _wishlistsData[wishlistId_];

        return
            WishlistInfo(
                wishlistId_,
                _wishlistData.ownerAddr,
                _wishlistData.owedUsdcAmount,
                _wishlistData.nextItemId,
                getWishlistItemsInfo(wishlistId_, _wishlistData.activeItemIds.values())
            );
    }

    function getWishlistItemsInfo(
        uint256 wishlistId_,
        uint256[] memory itemIds_
    ) public view returns (WishlistItemInfo[] memory resultArr_) {
        WishlistData storage _wishlistData = _wishlistsData[wishlistId_];

        resultArr_ = new WishlistItemInfo[](itemIds_.length);

        for (uint256 i = 0; i < itemIds_.length; i++) {
            WishlistItemData memory itemData_ = _wishlistData.itemsData[itemIds_[i]];

            resultArr_[i] = WishlistItemInfo(
                itemIds_[i],
                itemData_.itemPrice,
                _countFeeAmount(itemData_.itemPrice),
                itemData_.buyerAddr,
                itemData_.buyerAddr == address(0)
            );
        }
    }

    function isItemActive(uint256 wishlistId_, uint256 itemId_) public view returns (bool) {
        return _wishlistsData[wishlistId_].activeItemIds.contains(itemId_);
    }

    function _setProtocolFee(uint256 newProtocolFee_) internal {
        require(newProtocolFee_ <= MAX_PROTOCOL_FEE, InvalidProtocolFee(newProtocolFee_));

        uint256 currentProtocolFee_ = protocolFee;

        protocolFee = newProtocolFee_;

        emit ProtocolFeeUpdated(newProtocolFee_, currentProtocolFee_);
    }

    function _setProtocolFeeRecipient(address newProtocolFeeRecipient_) internal {
        require(newProtocolFeeRecipient_ != address(0), ZeroProtocolFeeRecipient());

        address currentFeeRecipient_ = protocolFeeRecipient;

        protocolFeeRecipient = newProtocolFeeRecipient_;

        emit ProtocolFeeRecipientUpdated(newProtocolFeeRecipient_, currentFeeRecipient_);
    }

    function _addWishlistItems(uint256 wishlistId_, uint256[] calldata itemPrices_) internal {
        require(itemPrices_.length > 0, ZeroItemsToAdd());

        WishlistData storage _wishlistData = _wishlistsData[wishlistId_];

        uint256 nextItemId_ = _wishlistData.nextItemId;

        for (uint256 i = 0; i < itemPrices_.length; i++) {
            _wishlistData.activeItemIds.add(nextItemId_);

            _changeItemPrice(_wishlistData.itemsData[nextItemId_++], itemPrices_[i]);
        }

        _wishlistData.nextItemId = nextItemId_;

        emit WishlistItemsAdded(wishlistId_, nextItemId_ - 1);
    }

    function _removeWishlistItems(uint256 wishlistId_, uint256[] calldata itemIds_) internal {
        WishlistData storage _wishlistData = _wishlistsData[wishlistId_];

        for (uint256 i = 0; i < itemIds_.length; i++) {
            uint256 itemIdToRemove_ = itemIds_[i];

            _checkActiveItem(wishlistId_, itemIdToRemove_);

            _wishlistData.activeItemIds.remove(itemIdToRemove_);
            delete _wishlistData.itemsData[itemIdToRemove_].itemPrice;
        }

        emit WishlistItemsRemoved(wishlistId_, itemIds_);
    }

    function _changeWishlistItemsPrice(
        uint256 wishlistId_,
        ChangedItemPriceData[] calldata changedItemsDataArr_
    ) internal {
        WishlistData storage _wishlistData = _wishlistsData[wishlistId_];

        for (uint256 i = 0; i < changedItemsDataArr_.length; i++) {
            ChangedItemPriceData calldata currentData_ = changedItemsDataArr_[i];

            _checkActiveItem(wishlistId_, currentData_.itemId);
            _changeItemPrice(
                _wishlistData.itemsData[currentData_.itemId],
                currentData_.newItemPrice
            );
        }
    }

    function _changeItemPrice(WishlistItemData storage _itemData, uint256 newPrice_) internal {
        require(newPrice_ > 0, ZeroItemPrice());

        _itemData.itemPrice = newPrice_;
    }

    function _countFeeAmount(uint256 itemPrice_) internal view returns (uint256) {
        return (itemPrice_ * protocolFee) / PERCENTAGE_100;
    }

    function _onlyWishlistOwner(uint256 wishlistId_) internal view {
        require(
            _wishlistsData[wishlistId_].ownerAddr == msg.sender,
            NotAWishlistOwner(wishlistId_, msg.sender)
        );
    }

    function _checkActiveItem(uint256 wishlistId_, uint256 itemId_) internal view {
        require(isItemActive(wishlistId_, itemId_), NotAnActiveItem(wishlistId_, itemId_));
    }
}
