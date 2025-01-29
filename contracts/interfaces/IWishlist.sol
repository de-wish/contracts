// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IWishlist
 * @notice Interface for managing user wishlists, including adding items, contributing funds, and withdrawing collected amounts.
 */
interface IWishlist {
    /**
     * @notice Stores information about a contribution to a wishlist item.
     * @param contributionId Unique identifier for the contribution.
     * @param contributor Address of the contributor.
     * @param contributionAmount Amount contributed.
     */
    struct ContributionInfo {
        uint256 contributionId;
        address contributor;
        uint256 contributionAmount;
    }

    /**
     * @notice Stores information about a wishlist item.
     * @param itemId Unique identifier for the item.
     * @param itemPrice Price of the item in USDC.
     * @param collectedTokensAmount Total amount of tokens collected for the item.
     * @param totalContributionsNumber Number of contributions made to this item.
     * @param contributionsInfo List of contributions to the item.
     * @param isActive Indicates whether the item is still active.
     */
    struct WishlistItemInfo {
        uint256 itemId;
        uint256 itemPrice;
        uint256 collectedTokensAmount;
        uint256 totalContributionsNumber;
        ContributionInfo[] contributionsInfo;
        bool isActive;
    }

    /**
     * @notice Stores information about a wishlist.
     * @param wishlistId Unique identifier for the wishlist.
     * @param ownerAddr Address of the wishlist owner.
     * @param owedUsdcAmount Amount of USDC owed for the wishlist.
     * @param totalItemsCount Total number of items in the wishlist.
     * @param activeItemsInfo List of active wishlist items.
     */
    struct WishlistInfo {
        uint256 wishlistId;
        address ownerAddr;
        uint256 owedUsdcAmount;
        uint256 totalItemsCount;
        WishlistItemInfo[] activeItemsInfo;
    }

    /**
     * @notice Stores data about a contribution.
     * @param contributor Address of the contributor.
     * @param contributionAmount Amount contributed.
     */
    struct ContributionData {
        address contributor;
        uint256 contributionAmount;
    }

    /**
     * @notice Stores information about a price change for a wishlist item.
     * @param itemId Unique identifier of the item.
     * @param newItemPrice New price of the item.
     */
    struct ChangedItemPriceData {
        uint256 itemId;
        uint256 newItemPrice;
    }

    /**
     * @notice Stores internal data about a wishlist item.
     * @param itemPrice Price of the item.
     * @param collectedTokensAmount Total amount of collected tokens.
     * @param totalContributionsNumber Number of contributions made.
     * @param contributrionsData Mapping with internal contributions data.
     */
    struct WishlistItemData {
        uint256 itemPrice;
        uint256 collectedTokensAmount;
        uint256 totalContributionsNumber;
        mapping(uint256 => ContributionData) contributrionsData;
    }

    /**
     * @notice Emitted when new wishlist items are added.
     * @param lastItemId The last item ID added.
     */
    event WishlistItemsAdded(uint256 lastItemId);

    /**
     * @notice Emitted when wishlist items are removed.
     * @param removedItemIds List of removed item IDs.
     */
    event WishlistItemsRemoved(uint256[] removedItemIds);

    /**
     * @notice Emitted when funds are withdrawn from the wishlist.
     * @param withdrawAmount Amount of funds withdrawn.
     * @param fundsRecipient Address receiving the funds.
     */
    event WishlistFundsWithdrawn(uint256 withdrawAmount, address fundsRecipient);

    /**
     * @notice Emitted when funds are contributed to a wishlist item.
     * @param itemId Unique identifier of the item.
     * @param contributedAmount Amount contributed.
     * @param feeAmount Protocol fee amount.
     * @param contributionAddr Address of the contributor.
     */
    event FundsContributed(
        uint256 indexed itemId,
        uint256 contributedAmount,
        uint256 feeAmount,
        address contributionAddr
    );

    /**
     * @notice Emitted when an item reaches its funding goal.
     * @param itemId Unique identifier of the item.
     * @param totalContributionsNumber Total number of contributions received.
     */
    event FundsCollectionFinished(uint256 indexed itemId, uint256 totalContributionsNumber);

    /**
     * @notice Thrown when an item price is set to zero.
     */
    error ZeroItemPrice();

    /**
     * @notice Thrown when no items are provided for addition.
     */
    error ZeroItemsToAdd();

    /**
     * @notice Thrown when a withdrawal amount is zero.
     */
    error ZeroAmountToWithdraw();

    /**
     * @notice Thrown when trying to modify an inactive item.
     * @param itemId The item ID that is not active.
     */
    error NotAnActiveItem(uint256 itemId);

    /**
     * @notice Thrown when a non-owner attempts to modify the wishlist.
     * @param senderAddr The sender's address.
     */
    error NotAWishlistOwner(address senderAddr);

    /**
     * @notice Initializes a new wishlist with given items.
     * @param wishlistId_ Unique identifier for the wishlist.
     * @param wishlistOwner_ Address of the wishlist owner.
     * @param initialItemPrices_ List of initial item prices.
     */
    function initializeWishlist(
        uint256 wishlistId_,
        address wishlistOwner_,
        uint256[] calldata initialItemPrices_
    ) external;

    /**
     * @notice Adds new items to the wishlist.
     * @param itemPrices_ List of item prices to be added.
     */
    function addWishlistItems(uint256[] calldata itemPrices_) external;

    /**
     * @notice Removes items from the wishlist.
     * @param itemIds_ List of item IDs to remove.
     */
    function removeWishlistItems(uint256[] calldata itemIds_) external;

    /**
     * @notice Changes the price of existing wishlist items.
     * @param changedItemsDataArr_ Array containing updated item prices.
     */
    function changeWishlistItemsPrice(
        ChangedItemPriceData[] calldata changedItemsDataArr_
    ) external;

    /**
     * @notice Withdraws collected funds from the wishlist.
     * @param fundsRecipient_ Address receiving the withdrawn funds.
     */
    function withdrawFunds(address fundsRecipient_) external;

    /**
     * @notice Contributes funds to a specific wishlist item.
     * @param itemId_ The ID of the item to contribute to.
     * @param amountToContribute_ The amount to contribute.
     */
    function contributeFunds(uint256 itemId_, uint256 amountToContribute_) external;

    /**
     * @notice Returns information about the wishlist.
     * @return A struct containing wishlist details.
     */
    function getWishlistInfo() external view returns (WishlistInfo memory);

    /**
     * @notice Returns information about specific wishlist items.
     * @param itemIds_ List of item IDs to retrieve information for.
     * @return An array of structs containing item details.
     */
    function getWishlistItemsInfo(
        uint256[] memory itemIds_
    ) external view returns (WishlistItemInfo[] memory);

    /**
     * @notice Returns the list of contributions for a given item.
     * @param itemId_ The ID of the item.
     * @return An array of structs containing contribution details.
     */
    function getContributionsInfo(
        uint256 itemId_
    ) external view returns (ContributionInfo[] memory);

    /**
     * @notice Checks if an item is active.
     * @param itemId_ The ID of the item to check.
     * @return True if the item is active, false otherwise.
     */
    function isItemActive(uint256 itemId_) external view returns (bool);
}
