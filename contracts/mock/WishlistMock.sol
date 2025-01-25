// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Wishlist} from "../Wishlist.sol";

contract WishlistMock is Wishlist {
    function version() external pure returns (string memory) {
        return "2.0.0";
    }
}
