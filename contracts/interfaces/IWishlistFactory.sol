// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IWishlistFactory {
    struct ProtocolFeeSettings {
        uint256 protocolFeePercentage;
        uint256 maxProtocolFeeAmount;
        address protocolFeeRecipient;
    }

    event NewWishlistCreated(
        address indexed wishlistOwner,
        uint256 wishlistId,
        address wishlistAddr
    );
    event ProtocolSignerUpdated(address newProtocolSignerAddr, address oldProtocolSignerAddr);
    event ProtocolFeeSettingsUpdated(
        uint256 newProtocolFeePercentage,
        uint256 newMaxProtocolFeeAmount,
        address newProtocolFeeRecipient
    );

    error ZeroAddr(string addrName);
    error ZeroFeeToWithdraw();
    error CreateWishlistSignatureExpired();
    error InvalidCreateWishlistSignature();
    error WishlistAlreadyExists(uint256 wishlistId);
    error InvalidProtocolFee(uint256 protocolFee);

    function upgradeWishlistsImpl(address newWishlistsImpl_) external;

    function setProtocolSignerAddr(address newProtocolSignerAddr_) external;

    function setProtocolFeeSettings(ProtocolFeeSettings calldata feeSettings_) external;

    function createWishlist(
        address wishlistOwner_,
        uint256 wishlistId_,
        uint256 sigDeadline_,
        uint256[] calldata initialItemPrices_,
        bytes memory signature_
    ) external returns (address);

    function usdcToken() external view returns (IERC20);

    function nextWishlistId() external view returns (uint256);

    function protocolSignerAddr() external view returns (address);

    function getProtocolFeeSettings() external view returns (ProtocolFeeSettings memory);

    function getWishlistsTotalCount() external view returns (uint256);

    function getWishlistAddress(uint256 wishlistId_) external view returns (address);

    function getWishlistAddresses(
        uint256 from_,
        uint256 to_
    ) external view returns (address[] memory wishlistAddresses_);

    function wishlistExists(uint256 wishlistId_) external view returns (bool);
}
