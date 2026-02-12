// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/AccessControl.sol";

contract SkillToken is ERC20, AccessControl {
    // Role that can mint new tokens (e.g., for rewards, initial distribution)
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    // Role that can burn tokens (e.g., for buybacks, fee burning)
    bytes32 public constant BURNER_ROLE = keccak256("BURNER_ROLE");

    // Maximum total supply (optional, good for demo)
    uint256 public immutable maxSupply;

    event TokensMinted(address to, uint256 amount);
    event TokensBurned(address from, uint256 amount);

    /**
     * @dev Constructor sets up roles and initial supply.
     * @param _initialHolder Address receiving the initial minted supply.
     * @param _initialSupply Initial token amount (e.g., for liquidity, team).
     * @param _maxSupply Hard cap on total tokens (e.g., 1 billion).
     */
    constructor(
        address _initialHolder,
        uint256 _initialSupply,
        uint256 _maxSupply
    ) ERC20("SkillChain Token", "SKILL") {
        require(_initialHolder != address(0), "Initial holder cannot be zero");
        require(_initialSupply <= _maxSupply, "Initial supply exceeds max");

        maxSupply = _maxSupply;

        // Grant roles to the contract deployer
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        _grantRole(MINTER_ROLE, msg.sender);
        _grantRole(BURNER_ROLE, msg.sender);

        // Mint the initial supply
        _mint(_initialHolder, _initialSupply);
    }

    /**
     * @dev Mints new tokens. Can be called by backend for staking rewards.
     * @param to The address receiving the tokens.
     * @param amount The amount to mint.
     */
    function mint(address to, uint256 amount) external onlyRole(MINTER_ROLE) {
        require(totalSupply() + amount <= maxSupply, "Exceeds max supply");
        _mint(to, amount);
        emit TokensMinted(to, amount);
    }

    /**
     * @dev Burns tokens. Can be used for buybacks or fee burning.
     * @param from The address whose tokens are burned.
     * @param amount The amount to burn.
     */
    function burnFrom(address from, uint256 amount) external onlyRole(BURNER_ROLE) {
        _burn(from, amount);
        emit TokensBurned(from, amount);
    }

    // Optional: Override to add hooks if needed (e.g., for governance snapshots)
    function _update(address from, address to, uint256 value) internal virtual override {
        super._update(from, to, value);
        // Potential hook for governance/tracking
    }
}