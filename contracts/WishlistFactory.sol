// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";

import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import {ProxyBeacon} from "@solarity/solidity-lib/proxy/beacon/ProxyBeacon.sol";
import {PublicBeaconProxy} from "@solarity/solidity-lib/proxy/beacon/PublicBeaconProxy.sol";

import {PERCENTAGE_100} from "@solarity/solidity-lib/utils/Globals.sol";

import {IWishlistFactory} from "./interfaces/IWishlistFactory.sol";
import {IWishlist} from "./interfaces/IWishlist.sol";

contract WishlistFactory is IWishlistFactory, UUPSUpgradeable, OwnableUpgradeable {
    using SafeERC20 for IERC20;

    IERC20 public usdcToken;

    ProxyBeacon public proxyBeacon;

    uint256 public nextWishlistId;

    uint256 public maxProtocolFeeAmount;
    uint256 public protocolFeePercentage;

    function initialize(
        address wishlistImpl_,
        address usdcTokenAddr_,
        uint256 protocolFeePercentage_,
        uint256 maxProtocolFeeAmount_
    ) external initializer {
        __Ownable_init();

        proxyBeacon = new ProxyBeacon();
        proxyBeacon.upgradeTo(wishlistImpl_);

        usdcToken = IERC20(usdcTokenAddr_);

        maxProtocolFeeAmount = maxProtocolFeeAmount_;
        _setProtocolFeePercentage(protocolFeePercentage_);

        nextWishlistId = 1;
    }

    function upgradeWishlistsImpl(address newWishlistsImpl_) external onlyOwner {
        proxyBeacon.upgradeTo(newWishlistsImpl_);
    }

    function setProtocolFeePercentage(uint256 newFeePercentage_) external onlyOwner {
        _setProtocolFeePercentage(newFeePercentage_);
    }

    function withdrawProtocolFee(address feeRecipient_) external onlyOwner {
        require(feeRecipient_ != address(0), ZeroFeeRecipient());

        uint256 feeToWithdraw_ = usdcToken.balanceOf(address(this));

        require(feeToWithdraw_ > 0, ZeroFeeToWithdraw());

        usdcToken.safeTransfer(feeRecipient_, feeToWithdraw_);

        emit ProtocolFeeWithdrawn(feeToWithdraw_, feeRecipient_);
    }

    function createWishlist(
        uint256[] calldata initialItemPrices_
    ) external returns (address newWishlistAddr_) {
        newWishlistAddr_ = address(new PublicBeaconProxy(address(proxyBeacon), ""));

        uint256 newWishlistId_ = nextWishlistId++;

        IWishlist(newWishlistAddr_).initializeWishlist(
            newWishlistId_,
            msg.sender,
            initialItemPrices_
        );

        emit NewWishlistCreated(msg.sender, newWishlistId_, newWishlistAddr_);
    }

    function _setProtocolFeePercentage(uint256 newFeePercentage_) internal {
        require(newFeePercentage_ <= PERCENTAGE_100, InvalidProtocolFee(newFeePercentage_));

        uint256 currentFeePercentage_ = protocolFeePercentage;

        protocolFeePercentage = newFeePercentage_;

        emit ProtocolFeePercentageUpdated(newFeePercentage_, currentFeePercentage_);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}
}
