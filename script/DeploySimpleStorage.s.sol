// SPDX-License-Identifier: MIT

pragma solidity 0.8.19;

import {Script} from "../lib/forge-std/src/Script.sol";
import {SimpleStorage} from "../src/Simple-Storage.sol";

contract DeploySimpleStorage is Script {
    function run() external returns (SimpleStorage) {
        // Whatever needs to add in the transaction add in between start and stop
        // Any transaction we want to add also put in between start and stop
        vm.startBroadcast();

        SimpleStorage simpleStorage = new SimpleStorage();

        vm.stopBroadcast();
        return simpleStorage;
    }
}
