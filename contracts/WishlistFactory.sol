// SPDX-License-Identifier: MIT
pragma solidity 0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";

import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import {EnumerableSet} from "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

import {ProxyBeacon} from "@solarity/solidity-lib/proxy/beacon/ProxyBeacon.sol";
import {PublicBeaconProxy} from "@solarity/solidity-lib/proxy/beacon/PublicBeaconProxy.sol";
import {Paginator} from "@solarity/solidity-lib/libs/arrays/Paginator.sol";

import {PERCENTAGE_100} from "@solarity/solidity-lib/utils/Globals.sol";

import {IWishlistFactory} from "./interfaces/IWishlistFactory.sol";
import {IWishlist} from "./interfaces/IWishlist.sol";

contract WishlistFactory is
    IWishlistFactory,
    UUPSUpgradeable,
    EIP712Upgradeable,
    OwnableUpgradeable
{
    using SafeERC20 for IERC20;
    using EnumerableSet for EnumerableSet.UintSet;
    using Paginator for EnumerableSet.UintSet;

    // solhint-disable-next-line var-name-mixedcase
    bytes32 private constant _CREATE_WISHLIST_TYPEHASH =
        keccak256(
            "CreateWishlist(address wishlistOwner,uint256 wishlistId,bytes32 itemPricesHash,uint256 deadline)"
        );

    IERC20 public usdcToken;

    ProxyBeacon public proxyBeacon;

    uint256 public nextWishlistId;

    uint256 public maxProtocolFeeAmount;
    uint256 public protocolFeePercentage;

    EnumerableSet.UintSet internal _wishlistIds;

    mapping(uint256 => address) internal _wishlists;

    function initialize(
        address wishlistImpl_,
        address usdcTokenAddr_,
        uint256 protocolFeePercentage_,
        uint256 maxProtocolFeeAmount_
    ) external initializer {
        __Ownable_init();
        __EIP712_init("WishlistFactory", "1");

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

        uint256 feeToWithdraw_ = getCurrentProtocolFeeAmount();

        require(feeToWithdraw_ > 0, ZeroFeeToWithdraw());

        usdcToken.safeTransfer(feeRecipient_, feeToWithdraw_);

        emit ProtocolFeeWithdrawn(feeToWithdraw_, feeRecipient_);
    }

    function createWishlist(
        uint256 wishlistId_,
        uint256 sigDeadline_,
        uint256[] calldata initialItemPrices_,
        bytes memory signature_
    ) external returns (address newWishlistAddr_) {
        require(block.timestamp <= sigDeadline_, CreateWishlistSignatureExpired());
        require(!wishlistExists(wishlistId_), WishlistAlreadyExists(wishlistId_));

        bytes32 structHash_ = _getCreateWishlistStructHash(
            msg.sender,
            wishlistId_,
            sigDeadline_,
            initialItemPrices_
        );

        _checkCreateWishlistSig(msg.sender, structHash_, signature_);

        newWishlistAddr_ = address(new PublicBeaconProxy(address(proxyBeacon), ""));

        IWishlist(newWishlistAddr_).initializeWishlist(
            wishlistId_,
            msg.sender,
            initialItemPrices_
        );

        _wishlistIds.add(wishlistId_);
        _wishlists[wishlistId_] = newWishlistAddr_;

        emit NewWishlistCreated(msg.sender, wishlistId_, newWishlistAddr_);
    }

    function getWishlistsTotalCount() external view returns (uint256) {
        return _wishlistIds.length();
    }

    function getWishlistAddress(uint256 wishlistId_) external view returns (address) {
        return _wishlists[wishlistId_];
    }

    function getWishlistAddresses(
        uint256 from_,
        uint256 to_
    ) external view returns (address[] memory wishlistAddresses_) {
        uint256[] memory wishlistIds_ = _wishlistIds.part(from_, to_);

        wishlistAddresses_ = new address[](wishlistIds_.length);

        for (uint256 i = 0; i < wishlistIds_.length; i++) {
            wishlistAddresses_[i] = _wishlists[wishlistIds_[i]];
        }
    }

    function getCurrentProtocolFeeAmount() public view returns (uint256) {
        return usdcToken.balanceOf(address(this));
    }

    function wishlistExists(uint256 wishlistId_) public view returns (bool) {
        return _wishlistIds.contains(wishlistId_);
    }

    function _setProtocolFeePercentage(uint256 newFeePercentage_) internal {
        require(newFeePercentage_ <= PERCENTAGE_100, InvalidProtocolFee(newFeePercentage_));

        uint256 currentFeePercentage_ = protocolFeePercentage;

        protocolFeePercentage = newFeePercentage_;

        emit ProtocolFeePercentageUpdated(newFeePercentage_, currentFeePercentage_);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}

    function _checkCreateWishlistSig(
        address wishlistOwner_,
        bytes32 structHash_,
        bytes memory signature_
    ) internal view {
        bytes32 typedHash_ = _hashTypedDataV4(structHash_);

        address signer_ = ECDSA.recover(typedHash_, signature_);

        require(signer_ == wishlistOwner_, InvalidCreateWishlistSignature());
    }

    function _getCreateWishlistStructHash(
        address wishlistOwner_,
        uint256 wishlistId_,
        uint256 sigDeadline_,
        uint256[] calldata initialItemPrices_
    ) internal pure returns (bytes32) {
        bytes32 itemPricesHash_ = keccak256(abi.encode(initialItemPrices_));

        return
            keccak256(
                abi.encode(
                    _CREATE_WISHLIST_TYPEHASH,
                    wishlistOwner_,
                    wishlistId_,
                    itemPricesHash_,
                    sigDeadline_
                )
            );
    }
}
