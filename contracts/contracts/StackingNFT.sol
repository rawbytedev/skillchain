// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract StakingNFT is ERC721URIStorage, ReentrancyGuard, Ownable {
    // Interface for the $SKILL ERC-20 token
    IERC20 public skillToken;

    uint256 public nextTokenId=1; // Start token IDs at 1 for better UX
    uint256 public minimumStake;
    uint256 public constant SLASH_TIMELOCK = 3 days;

    // Maps Token ID to its staking info
    mapping(uint256 => uint256) public stakeAmount;
    mapping(uint256 => address) public toolCreator;
    
    // Slash proposals for time-lock
    mapping(uint256 => uint256) public slashProposalTime;
    mapping(uint256 => uint256) public slashProposalAmount;

    event ToolListed(uint256 indexed tokenId, address creator, uint256 stake, string metadataURI);
    event ToolDelisted(uint256 indexed tokenId, address creator, uint256 stake);
    event StakeWithdrawn(uint256 indexed tokenId, address creator, uint256 amount);
    event SlashProposed(uint256 indexed tokenId, uint256 amount);
    event StakeSlashed(uint256 indexed tokenId, address slasher, uint256 amount);

    constructor(address _skillTokenAddress, uint256 _minimumStake) 
        ERC721("SkillChain Tool", "SCT") 
        Ownable(msg.sender)
    {
        skillToken = IERC20(_skillTokenAddress);
        minimumStake = _minimumStake;
    }

    function listTool(uint256 _stakeAmount, string memory _tokenURI) 
        external 
        nonReentrant 
        returns (uint256) 
    {
        require(_stakeAmount >= minimumStake, "Stake below minimum");

        // Transfer stake from creator to this contract
        require(
            skillToken.transferFrom(msg.sender, address(this), _stakeAmount),
            "Transfer failed: ensure you approved this contract"
        );

        uint256 newTokenId = nextTokenId++;
        _safeMint(msg.sender, newTokenId);
        _setTokenURI(newTokenId, _tokenURI);

        stakeAmount[newTokenId] = _stakeAmount;
        toolCreator[newTokenId] = msg.sender;

        emit ToolListed(newTokenId, msg.sender, _stakeAmount, _tokenURI);
        return newTokenId;
    }

    // Allows creator to delist tool and withdraw stake
    function delistTool(uint256 _tokenId) external nonReentrant {
        require(ownerOf(_tokenId) == msg.sender, "Not token owner");
        require(stakeAmount[_tokenId] > 0, "Already delisted");
        
        uint256 amount = stakeAmount[_tokenId];
        stakeAmount[_tokenId] = 0;
        
        // Burn the NFT
        _burn(_tokenId);
        
        // Transfer stake back to creator
        require(skillToken.transfer(msg.sender, amount), "Transfer failed");
        
        emit ToolDelisted(_tokenId, msg.sender, amount);
        emit StakeWithdrawn(_tokenId, msg.sender, amount);
    }

    // Propose slash (requires time-lock before execution)
    function proposeSlash(uint256 _tokenId, uint256 _slashAmount) external onlyOwner {
        require(_ownerOf(_tokenId) != address(0), "Token does not exist");
        require(_slashAmount <= stakeAmount[_tokenId], "Slash amount exceeds stake");
        
        slashProposalTime[_tokenId] = block.timestamp;
        slashProposalAmount[_tokenId] = _slashAmount;
        
        emit SlashProposed(_tokenId, _slashAmount);
    }

    // Execute slash after time-lock expires
    function executeSlash(uint256 _tokenId, address _recipient) external onlyOwner nonReentrant {
        require(
            block.timestamp >= slashProposalTime[_tokenId] + SLASH_TIMELOCK,
            "Timelock not expired"
        );
        require(slashProposalAmount[_tokenId] > 0, "No slash proposal");
        
        uint256 _slashAmount = slashProposalAmount[_tokenId];
        require(_slashAmount <= stakeAmount[_tokenId], "Slash amount exceeds stake");
        
        stakeAmount[_tokenId] -= _slashAmount;
        slashProposalAmount[_tokenId] = 0;
        
        require(skillToken.transfer(_recipient, _slashAmount), "Transfer failed");
        
        emit StakeSlashed(_tokenId, msg.sender, _slashAmount);
    }
}