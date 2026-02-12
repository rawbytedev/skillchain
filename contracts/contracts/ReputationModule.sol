// SPDX-License-Identifier: MIT
// Minimal contract for storing reputation checkpoints
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

contract ReputationOracle is Ownable, EIP712{
    using ECDSA for bytes32;

    address public immutable backendSigner;
    
    // toolId -> nonce (for replay protection)
    mapping(uint256 => uint256) public toolNonce;
    
    // toolId -> timestamp -> reputationMerkleRoot
    mapping(uint256 => mapping(uint256 => bytes32)) public reputationCheckpoints;
    
    // EIP-712 typehash for reputation updates
    bytes32 private constant _UPDATE_ROOT_TYPEHASH = 
        keccak256("UpdateReputationRoot(uint256 toolId,uint256 timestamp,bytes32 rootHash,uint256 nonce)");
    
    event ReputationUpdated(uint256 toolId, uint256 timestamp, bytes32 rootHash);

    constructor(address _trustedSigner) 
        Ownable(msg.sender)
        EIP712("SkillChainLicense", "1")
    {
        backendSigner = _trustedSigner;
    }

    // Backend calls this ONCE per period with a batch proof
    function updateReputationRoot(
        uint256 _toolId,
        uint256 _timestamp,
        bytes32 _rootHash,
        uint256 _nonce,
        bytes memory _signature
    ) external {
        // 1. Verify nonce (replay attack protection)
        require(_nonce == toolNonce[_toolId], "Invalid nonce");
        
        // 2. Verify timestamp is not in future and reasonably recent
        require(_timestamp <= block.timestamp, "Timestamp cannot be in the future");
        require(_timestamp > block.timestamp - 7 days, "Signature too old");
        
        // 3. Verify backend signed this update using EIP-712
        bytes32 digest = _hashTypedDataV4(
            keccak256(abi.encode(
                _UPDATE_ROOT_TYPEHASH,
                _toolId,
                _timestamp,
                _rootHash,
                _nonce
            ))
        );
        address signer = digest.recover(_signature);
        require(signer == backendSigner, "Invalid signature");
        
        // 4. Increment nonce for next update
        toolNonce[_toolId]++;
        
        // 5. Store the reputation root
        reputationCheckpoints[_toolId][_timestamp] = _rootHash;
        emit ReputationUpdated(_toolId, _timestamp, _rootHash);
    }

    // View function to get a reputation root for a given time
    function getReputationRoot(uint256 _toolId, uint256 _timestamp) external view returns (bytes32) {
        return reputationCheckpoints[_toolId][_timestamp];
    }
}