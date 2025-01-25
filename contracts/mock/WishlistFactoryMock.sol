// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {WishlistFactory} from "../WishlistFactory.sol";

contract WishlistFactoryMock is WishlistFactory {
    function version() external pure returns (string memory) {
        return "1.1.1";
    }
}
