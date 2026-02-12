// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract LicenseNFT is ERC1155, Ownable, EIP712 {
    using ECDSA for bytes32;

    // The backend server's address that signs off on reputation proofs
    address public trustedSigner;
    // Base URI for metadata
    string public baseURI;

    // Maps (user, toolId) to its license expiration
    mapping(address => mapping(uint256 => uint256)) public licenseExpiry;
    
    // Track used nonces to prevent replay attacks
    mapping(address => mapping(uint256 => bool)) public usedNonces;

    // EIP-712 typehash for signed minting
    bytes32 private constant _MINT_LICENSE_TYPEHASH = 
        keccak256("MintLicense(address user,uint256 toolId,uint256 expiresAt,uint256 nonce)");

    event LicenseMinted(address indexed user, uint256 toolId, uint256 expiresAt, uint256 pricePaid);
    event SignerChanged(address oldSigner, address newSigner);
    event LicenseRevoked(address indexed user, uint256 toolId);

    constructor(address _trustedSigner, string memory _uri) 
        ERC1155(_uri) 
        Ownable(msg.sender)
        EIP712("SkillChainLicense", "1")
    {
        trustedSigner = _trustedSigner;
        baseURI = _uri;
    }

    function mintLicense(
        uint256 _toolId,
        uint256 _expiresAt,   // Unix timestamp
        uint256 _nonce,
        bytes memory _signature
    ) external payable {
        // 1. Verify the backend's signature
        bytes32 digest = _hashTypedDataV4(
            keccak256(abi.encode(
                _MINT_LICENSE_TYPEHASH,
                msg.sender,
                _toolId,
                _expiresAt,
                _nonce
            ))
        );
        address signer = digest.recover(_signature);
        require(signer == trustedSigner, "Invalid signature");
        require(msg.value > 0, "Payment required");
        // 2. Verify nonce hasn't been used (replay attack protection)
        require(!usedNonces[msg.sender][_nonce], "Nonce already used");
        usedNonces[msg.sender][_nonce] = true;

        // 3. Validate expiration is in the future
        require(_expiresAt > block.timestamp, "License expiration must be in the future");

        // 4. Check license isn't already active
        require(licenseExpiry[msg.sender][_toolId] < block.timestamp, "License already active");

        // 5. Set expiration
        licenseExpiry[msg.sender][_toolId] = _expiresAt;

        // 6. Mint the license NFT (ERC-1155 with token ID = toolId)
        _mint(msg.sender, _toolId, 1, "");

        emit LicenseMinted(msg.sender, _toolId, _expiresAt, 0);
    }

    // View function for your Go backend to check license validity
    function isLicenseValid(address _user, uint256 _toolId) external view returns (bool) {
        return licenseExpiry[_user][_toolId] >= block.timestamp;
    }

    // Revoke a license (admin-only)
    function revokeLicense(address _user, uint256 _toolId) external onlyOwner {
        require(licenseExpiry[_user][_toolId] > 0, "License does not exist");
        licenseExpiry[_user][_toolId] = 0;
        emit LicenseRevoked(_user, _toolId);
    }

    function updateSigner(address _newSigner) external onlyOwner {
        emit SignerChanged(trustedSigner, _newSigner);
        trustedSigner = _newSigner;
    }

    function uri(uint256 _id) public view override returns (string memory) {
        return string(abi.encodePacked(baseURI, _toString(_id)));
    }

    function _toString(uint256 value) internal pure returns (string memory) {
        if (value == 0) return "0";
        uint256 temp = value;
        uint256 digits;
        while (temp != 0) {
            digits++;
            temp /= 10;
        }
        bytes memory buffer = new bytes(digits);
        while (value != 0) {
            digits -= 1;
            buffer[digits] = bytes1(uint8(48 + uint256(value % 10)));
            value /= 10;
        }
        return string(buffer);
    }
}