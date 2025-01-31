// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ProtocolTreasury} from "../ProtocolTreasury.sol";

contract ProtocolTreasuryMock is ProtocolTreasury {
    function version() external pure returns (string memory) {
        return "3.3.3";
    }
}
