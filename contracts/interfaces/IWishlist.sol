// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

interface IWishlist {
    struct WishlistItemInfo {
        uint256 itemId;
        uint256 itemPrice;
        uint256 collectedTokensAmount;
        address[] buyersAddresses;
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
        uint256 collectedTokensAmount;
        EnumerableSet.AddressSet buyersAddresses;
    }

    event WishlistItemsAdded(uint256 lastItemId);
    event WishlistItemsRemoved(uint256[] removedItemIds);
    event WishlistFundsWithdrawn(uint256 withdrawAmount, address fundsRecipient);
    event ItemBought(
        uint256 itemId,
        uint256 boughtAmount,
        uint256 feeAmount,
        address buyerAddr,
        bool isItemFullyPurchased
    );

    error ZeroItemPrice();
    error ZeroItemsToAdd();
    error ZeroAmountToWithdraw();
    error NotAnActiveItem(uint256 itemId);
    error NotAWishlistOwner(address senderAddr);

    function initializeWishlist(
        uint256 wishlistId_,
        address wishlistOwner_,
        uint256[] calldata initialItemPrices_
    ) external;

    function addWishlistItems(uint256[] calldata itemPrices_) external;

    function removeWishlistItems(uint256[] calldata itemIds_) external;

    function changeWishlistItemsPrice(
        ChangedItemPriceData[] calldata changedItemsDataArr_
    ) external;

    function withdrawFunds(address fundsRecipient_) external;

    function buyWishlistItem(uint256 itemId_, uint256 amountToBuy_) external;

    function getWishlistInfo() external view returns (WishlistInfo memory);

    function getWishlistItemsInfo(
        uint256[] memory itemIds_
    ) external view returns (WishlistItemInfo[] memory);

    function isItemActive(uint256 itemId_) external view returns (bool);
}
