// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

/**
 * @title IWishlistFactory
 * @notice Interface for creating and managing wishlists, handling protocol fees, and maintaining a registry of wishlists.
 */
interface IWishlistFactory {
    /**
     * @notice Stores protocol fee settings.
     * @param protocolFeePercentage Percentage fee charged by the protocol.
     * @param maxProtocolFeeAmount Maximum fee amount that can be charged.
     * @param protocolFeeRecipient Address that receives the protocol fees.
     */
    struct ProtocolFeeSettings {
        uint256 protocolFeePercentage;
        uint256 maxProtocolFeeAmount;
        address protocolFeeRecipient;
    }

    /**
     * @notice Emitted when a new wishlist is created.
     * @param wishlistOwner Address of the wishlist owner.
     * @param wishlistId Unique identifier of the wishlist.
     * @param wishlistAddr Address of the newly created wishlist contract.
     */
    event NewWishlistCreated(
        address indexed wishlistOwner,
        uint256 wishlistId,
        address wishlistAddr
    );

    /**
     * @notice Emitted when the protocol signer address is updated.
     * @param newProtocolSignerAddr New protocol signer address.
     * @param oldProtocolSignerAddr Previous protocol signer address.
     */
    event ProtocolSignerUpdated(address newProtocolSignerAddr, address oldProtocolSignerAddr);

    /**
     * @notice Emitted when the protocol fee settings are updated.
     * @param newProtocolFeePercentage Updated fee percentage.
     * @param newMaxProtocolFeeAmount Updated maximum fee amount.
     * @param newProtocolFeeRecipient Updated protocol fee recipient address.
     */
    event ProtocolFeeSettingsUpdated(
        uint256 newProtocolFeePercentage,
        uint256 newMaxProtocolFeeAmount,
        address newProtocolFeeRecipient
    );

    /**
     * @notice Thrown when a zero address is provided where a valid address is required.
     * @param addrName The name of the parameter that received a zero address.
     */
    error ZeroAddr(string addrName);

    /**
     * @notice Thrown when a wishlist creation signature has expired.
     */
    error CreateWishlistSignatureExpired();

    /**
     * @notice Thrown when a wishlist creation signature is invalid.
     */
    error InvalidCreateWishlistSignature();

    /**
     * @notice Thrown when a wishlist with the given ID already exists.
     * @param wishlistId The ID of the existing wishlist.
     */
    error WishlistAlreadyExists(uint256 wishlistId);

    /**
     * @notice Thrown when the provided protocol fee is invalid.
     * @param protocolFee The invalid fee value.
     */
    error InvalidProtocolFee(uint256 protocolFee);

    /**
     * @notice Upgrades the implementation of all wishlists created by this factory.
     * @param newWishlistsImpl_ Address of the new wishlist implementation contract.
     */
    function upgradeWishlistsImpl(address newWishlistsImpl_) external;

    /**
     * @notice Sets a new protocol signer address.
     * @param newProtocolSignerAddr_ Address of the new protocol signer.
     */
    function setProtocolSignerAddr(address newProtocolSignerAddr_) external;

    /**
     * @notice Updates the protocol fee settings.
     * @param feeSettings_ New protocol fee settings.
     */
    function setProtocolFeeSettings(ProtocolFeeSettings calldata feeSettings_) external;

    /**
     * @notice Creates a new wishlist.
     * @param wishlistOwner_ Address of the wishlist owner.
     * @param wishlistId_ Unique identifier for the wishlist.
     * @param sigDeadline_ Deadline for the wishlist creation signature.
     * @param initialItemPrices_ Initial prices of items in the wishlist.
     * @param signature_ Signature authorizing wishlist creation.
     * @return The address of the newly created wishlist contract.
     */
    function createWishlist(
        address wishlistOwner_,
        uint256 wishlistId_,
        uint256 sigDeadline_,
        uint256[] calldata initialItemPrices_,
        bytes memory signature_
    ) external returns (address);

    /**
     * @notice Returns the address of the USDC token used in the protocol.
     * @return The USDC token contract address.
     */
    function usdcToken() external view returns (IERC20);

    /**
     * @notice Returns the address of the protocol signer.
     * @return The protocol signer address.
     */
    function protocolSignerAddr() external view returns (address);

    /**
     * @notice Returns the current protocol fee settings.
     * @return The protocol fee settings.
     */
    function getProtocolFeeSettings() external view returns (ProtocolFeeSettings memory);

    /**
     * @notice Returns the total number of wishlists created by the factory.
     * @return The total count of wishlists.
     */
    function getWishlistsTotalCount() external view returns (uint256);

    /**
     * @notice Returns the address of a wishlist by its ID.
     * @param wishlistId_ The ID of the wishlist.
     * @return The wishlist contract address.
     */
    function getWishlistAddress(uint256 wishlistId_) external view returns (address);

    /**
     * @notice Returns the addresses of wishlists within a given range.
     * @param from_ Starting pagination index.
     * @param to_ Ending pagination index.
     * @return wishlistAddresses_ Array of wishlist addresses within the given range.
     */
    function getWishlistAddresses(
        uint256 from_,
        uint256 to_
    ) external view returns (address[] memory wishlistAddresses_);

    /**
     * @notice Checks if a wishlist with the given ID exists.
     * @param wishlistId_ The ID of the wishlist.
     * @return True if the wishlist exists, false otherwise.
     */
    function wishlistExists(uint256 wishlistId_) external view returns (bool);

    /**
     * @notice Computes the EIP712 signature hash for creating a wishlist.
     * @param wishlistOwner_ The address of the wishlist owner.
     * @param wishlistId_ The unique identifier of the wishlist.
     * @param sigDeadline_ The timestamp until which the signature is valid.
     * @param initialItemPrices_ The initial prices of items in the wishlist.
     * @return The computed EIP712 signature hash.
     */
    function getCreateWishlistSigHash(
        address wishlistOwner_,
        uint256 wishlistId_,
        uint256 sigDeadline_,
        uint256[] calldata initialItemPrices_
    ) external view returns (bytes32);
}
