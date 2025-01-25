// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

interface IWishlistFactory {
    event NewWishlistCreated(
        address indexed wishlistOwner,
        uint256 wishlistId,
        address wishlistAddr
    );
    event ProtocolFeePercentageUpdated(
        uint256 newProtocolFeePercentage,
        uint256 oldProtocolFeePercentage
    );
    event ProtocolFeeWithdrawn(uint256 feeAmount, address feeRecipient);

    error ZeroFeeRecipient();
    error ZeroFeeToWithdraw();
    error InvalidProtocolFee(uint256 protocolFee);

    function upgradeWishlistsImpl(address newWishlistsImpl_) external;

    function setProtocolFeePercentage(uint256 newFeePercentage_) external;

    function withdrawProtocolFee(address feeRecipient_) external;

    function createWishlist(uint256[] calldata initialItemPrices_) external returns (address);

    function usdcToken() external view returns (IERC20);

    function nextWishlistId() external view returns (uint256);

    function maxProtocolFeeAmount() external view returns (uint256);

    function protocolFeePercentage() external view returns (uint256);
}
