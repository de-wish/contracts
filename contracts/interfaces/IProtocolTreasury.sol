// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IProtocolTreasury {
    event NativeCurrencyWithdrawn(address indexed recipient, uint256 withdrawnAmount);
    event ERC20TokensWithdrawn(
        address indexed tokenAddr,
        address indexed recipient,
        uint256 withdrawnAmount
    );

    function withdrawNative(
        address recipient_,
        uint256 amountToWithdraw_,
        bool isAllWithdraw_
    ) external;

    function withdrawERC20Tokens(
        address tokenAddr_,
        address recipient_,
        uint256 amountToWithdraw_,
        bool isAllWithdraw_
    ) external;

    function getSelfNativeBalance() external view returns (uint256);

    function getSelfERC20TokenBalance(address tokenAddr_) external view returns (uint256);
}
