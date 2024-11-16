// SPDX-License-Identifier: MIT

pragma solidity ^0.8.19;

import "forge-std/Test.sol";
import "../src/Simple-Storage.sol";

contract SimpleStorageTest is Test {
    SimpleStorage simpleStorage;

    function setUp() public {
        simpleStorage = new SimpleStorage();
    }

    function testStore() public {
        uint256 value = 42;
        simpleStorage.store(value);
        uint256 storedValue = simpleStorage.retrieve();
        assertEq(storedValue, value, "Stored value should be 42");
    }

    function testAddPerson() public {
        simpleStorage.addPerson("Alice", 100);
        uint256 favoriteNumber = simpleStorage.nameToFavoriteNumber("Alice");
        assertEq(favoriteNumber, 100, "Alice's favorite number should be 100");
    }
}
