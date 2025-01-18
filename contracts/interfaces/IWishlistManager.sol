// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

interface IWishlistManager {
    struct WishlistItemInfo {
        uint256 itemId;
        uint256 itemPrice;
        uint256 itemBuyingFeeAmount;
        address buyerAddr;
        bool isActive;
    }

    struct WishlistInfo {
        uint256 wishlistId;
        address ownerAddr;
        uint256 owedUsdcAmount;
        uint256 totalItemsCount;
        WishlistItemInfo[] activeItemsInfo;
    }

    struct ChangedItemPriceData {
        uint256 itemId;
        uint256 newItemPrice;
    }

    struct WishlistItemData {
        uint256 itemPrice;
        address buyerAddr;
    }

    struct WishlistData {
        address ownerAddr;
        uint256 nextItemId;
        uint256 owedUsdcAmount;
        EnumerableSet.UintSet activeItemIds;
        mapping(uint256 => WishlistItemData) itemsData;
    }

    event ProtocolFeeUpdated(uint256 newProtocolFee, uint256 oldProtocolFee);
    event ProtocolFeeRecipientUpdated(
        address newProtocolFeeRecipient,
        address oldProtocolFeeRecipient
    );
    event WishlistCreated(address indexed owner, uint256 wishlistId);
    event WishlistItemsAdded(uint256 wishlistId, uint256 lastItemId);
    event WishlistItemsRemoved(uint256 wishlistId, uint256[] removedItemIds);
    event WishlistFundsWithdrawn(
        uint256 wishlistId,
        uint256 withdrawAmount,
        address fundsRecipient
    );
    event ItemBought(uint256 wishlistId, uint256 itemId, uint256 itemPrice, address buyerAddr);

    error ZeroItemPrice();
    error ZeroItemsToAdd();
    error ZeroProtocolFeeRecipient();
    error InvalidProtocolFee(uint256 protocolFee);
    error NotAnActiveItem(uint256 wishlistId, uint256 itemId);
    error NotAWishlistOwner(uint256 wishlistId, address senderAddr);

    function setProtocolFee(uint256 newProtocolFee_) external;

    function setProtocolFeeRecipient(address newProtocolFeeRecipient_) external;

    function createWishlist(uint256[] calldata itemPrices_) external;

    function addWishlistItems(uint256 wishlistId_, uint256[] calldata itemPrices_) external;

    function removeWishlistItems(uint256 wishlistId_, uint256[] calldata itemIds_) external;

    function changeWishlistItemsPrice(
        uint256 wishlistId_,
        ChangedItemPriceData[] calldata changedItemsDataArr_
    ) external;

    function withdrawFunds(uint256 wishlistId_, address fundsRecipient_) external;

    function buyWishlistItem(uint256 wishlistId_, uint256 itemId_) external;

    function getWishlistInfo(uint256 wishlistId_) external view returns (WishlistInfo memory);

    function getWishlistItemsInfo(
        uint256 wishlistId_,
        uint256[] memory itemIds_
    ) external view returns (WishlistItemInfo[] memory);

    function isItemActive(uint256 wishlistId_, uint256 itemId_) external view returns (bool);
}
