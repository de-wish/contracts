// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {Address} from "@openzeppelin/contracts/utils/Address.sol";

import {IProtocolTreasury} from "./interfaces/IProtocolTreasury.sol";

contract ProtocolTreasury is IProtocolTreasury, UUPSUpgradeable, OwnableUpgradeable {
    using SafeERC20 for IERC20;

    receive() external payable {}

    function initialize() external initializer {
        __Ownable_init();
    }

    function withdrawNative(
        address recipient_,
        uint256 amountToWithdraw_,
        bool isAllWithdraw_
    ) external onlyOwner {
        amountToWithdraw_ = isAllWithdraw_ ? getSelfNativeBalance() : amountToWithdraw_;

        Address.sendValue(payable(recipient_), amountToWithdraw_);

        emit NativeCurrencyWithdrawn(recipient_, amountToWithdraw_);
    }

    function withdrawERC20Tokens(
        address tokenAddr_,
        address recipient_,
        uint256 amountToWithdraw_,
        bool isAllWithdraw_
    ) external onlyOwner {
        amountToWithdraw_ = isAllWithdraw_
            ? getSelfERC20TokenBalance(tokenAddr_)
            : amountToWithdraw_;

        SafeERC20.safeTransfer(IERC20(tokenAddr_), recipient_, amountToWithdraw_);

        emit ERC20TokensWithdrawn(tokenAddr_, recipient_, amountToWithdraw_);
    }

    function getSelfNativeBalance() public view returns (uint256) {
        return address(this).balance;
    }

    function getSelfERC20TokenBalance(address tokenAddr_) public view returns (uint256) {
        return IERC20(tokenAddr_).balanceOf(address(this));
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
