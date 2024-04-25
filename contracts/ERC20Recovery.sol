// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

/**
 * @title ERC20Recovery
 * @dev Minimal version of ERC20 interface required to allow ERC20 locked tokens recovery
 */

interface ERC20Recovery {
  function balanceOf(address account) external returns (uint256);
  function transfer(address to, uint256 value) external returns (bool);
}
