// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract Counter is ReentrancyGuard{
    uint256 public count;
    constructor(){
        count = 0;
    }

    function incrementAndGet() public nonReentrant returns(uint256){
        return count += 1;
    }

    function get()public view returns(uint256){
        return count;
    }
}