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

    address public protocolSignerAddr;

    ProtocolFeeSettings internal _feeSettings;

    EnumerableSet.UintSet internal _wishlistIds;

    mapping(uint256 => address) internal _wishlists;

    function initialize(
        address wishlistImpl_,
        address usdcTokenAddr_,
        address protocolSignerAddr_,
        ProtocolFeeSettings calldata feeSettings_
    ) external initializer {
        __Ownable_init();
        __EIP712_init("WishlistFactory", "1");

        proxyBeacon = new ProxyBeacon();
        proxyBeacon.upgradeTo(wishlistImpl_);

        usdcToken = IERC20(usdcTokenAddr_);

        _setProtocolSignerAddr(protocolSignerAddr_);
        _setProtocolFeeSettings(feeSettings_);
    }

    function upgradeWishlistsImpl(address newWishlistsImpl_) external onlyOwner {
        proxyBeacon.upgradeTo(newWishlistsImpl_);
    }

    function setProtocolSignerAddr(address newProtocolSignerAddr_) external onlyOwner {
        _setProtocolSignerAddr(newProtocolSignerAddr_);
    }

    function setProtocolFeeSettings(ProtocolFeeSettings calldata feeSettings_) external onlyOwner {
        _setProtocolFeeSettings(feeSettings_);
    }

    function createWishlist(
        address wishlistOwner_,
        uint256 wishlistId_,
        uint256 sigDeadline_,
        uint256[] calldata initialItemPrices_,
        bytes memory signature_
    ) external returns (address newWishlistAddr_) {
        require(block.timestamp <= sigDeadline_, CreateWishlistSignatureExpired());
        require(!wishlistExists(wishlistId_), WishlistAlreadyExists(wishlistId_));

        bytes32 structHash_ = _getCreateWishlistStructHash(
            wishlistOwner_,
            wishlistId_,
            sigDeadline_,
            initialItemPrices_
        );

        _checkCreateWishlistSig(structHash_, signature_);

        newWishlistAddr_ = address(new PublicBeaconProxy(address(proxyBeacon), ""));

        IWishlist(newWishlistAddr_).initializeWishlist(
            wishlistId_,
            wishlistOwner_,
            initialItemPrices_
        );

        _wishlistIds.add(wishlistId_);
        _wishlists[wishlistId_] = newWishlistAddr_;

        emit NewWishlistCreated(wishlistOwner_, wishlistId_, newWishlistAddr_);
    }

    function getProtocolFeeSettings() external view returns (ProtocolFeeSettings memory) {
        return _feeSettings;
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

    function wishlistExists(uint256 wishlistId_) public view returns (bool) {
        return _wishlistIds.contains(wishlistId_);
    }

    function _setProtocolFeeSettings(ProtocolFeeSettings memory newProtocolFeeSettings_) internal {
        require(
            newProtocolFeeSettings_.protocolFeePercentage <= PERCENTAGE_100,
            InvalidProtocolFee(newProtocolFeeSettings_.protocolFeePercentage)
        );
        require(
            newProtocolFeeSettings_.protocolFeeRecipient != address(0),
            ZeroAddr("ProtocolFeeRecipient")
        );

        _feeSettings = newProtocolFeeSettings_;

        emit ProtocolFeeSettingsUpdated(
            newProtocolFeeSettings_.protocolFeePercentage,
            newProtocolFeeSettings_.maxProtocolFeeAmount,
            newProtocolFeeSettings_.protocolFeeRecipient
        );
    }

    function _setProtocolSignerAddr(address newProtocolSignerAddr_) internal {
        require(newProtocolSignerAddr_ != address(0), ZeroAddr("ProtocolSigner"));

        address currentProtocolSignerAddr_ = protocolSignerAddr;

        protocolSignerAddr = newProtocolSignerAddr_;

        emit ProtocolSignerUpdated(newProtocolSignerAddr_, currentProtocolSignerAddr_);
    }

    function _authorizeUpgrade(address) internal override onlyOwner {}

    function _checkCreateWishlistSig(bytes32 structHash_, bytes memory signature_) internal view {
        address signer_ = ECDSA.recover(_hashTypedDataV4(structHash_), signature_);

        require(signer_ == protocolSignerAddr, InvalidCreateWishlistSignature());
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
