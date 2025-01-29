// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

/**
 * @title IProtocolTreasury
 * @notice Interface for managing protocol treasury funds, including withdrawals of native currency and ERC-20 tokens.
 */
interface IProtocolTreasury {
    /**
     * @notice Emitted when native currency is withdrawn from the treasury.
     * @param recipient The address receiving the withdrawn funds.
     * @param withdrawnAmount The amount of native currency withdrawn.
     */
    event NativeCurrencyWithdrawn(address indexed recipient, uint256 withdrawnAmount);

    /**
     * @notice Emitted when ERC-20 tokens are withdrawn from the treasury.
     * @param tokenAddr The address of the ERC-20 token.
     * @param recipient The address receiving the withdrawn tokens.
     * @param withdrawnAmount The amount of tokens withdrawn.
     */
    event ERC20TokensWithdrawn(
        address indexed tokenAddr,
        address indexed recipient,
        uint256 withdrawnAmount
    );

    /**
     * @notice Withdraws native currency from the treasury.
     * Only the contract OWNDER can call this function.
     *
     * @param recipient_ The address receiving the withdrawn funds.
     * @param amountToWithdraw_ The amount of native currency to withdraw.
     * @param isAllWithdraw_ If true, withdraws the entire balance.
     */
    function withdrawNative(
        address recipient_,
        uint256 amountToWithdraw_,
        bool isAllWithdraw_
    ) external;

    /**
     * @notice Withdraws ERC-20 tokens from the treasury.
     * Only the contract OWNDER can call this function.
     *
     * @param tokenAddr_ The address of the ERC-20 token.
     * @param recipient_ The address receiving the withdrawn tokens.
     * @param amountToWithdraw_ The amount of tokens to withdraw.
     * @param isAllWithdraw_ If true, withdraws the entire balance of the specified token.
     */
    function withdrawERC20Tokens(
        address tokenAddr_,
        address recipient_,
        uint256 amountToWithdraw_,
        bool isAllWithdraw_
    ) external;

    /**
     * @notice Returns the treasury's native currency balance.
     * @return The amount of native currency held by the contract.
     */
    function getSelfNativeBalance() external view returns (uint256);

    /**
     * @notice Returns the balance of a specific ERC-20 token in the treasury.
     * @param tokenAddr_ The address of the ERC-20 token.
     * @return The amount of the specified token held by the contract.
     */
    function getSelfERC20TokenBalance(address tokenAddr_) external view returns (uint256);
}
