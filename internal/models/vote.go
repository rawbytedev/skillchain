// internal/models/vote.go
package models

import (
    "time"
)

// Vote represents a single agent vote on a tool
type Vote struct {
    ID           string    `json:"id"`           // Unique vote ID (hash of signature components)
    ToolID       string    `json:"tool_id"`      // ID of the tool being voted on
    VoterAddress string    `json:"voter_address"` // Ethereum address of the voter
    Score        int8      `json:"score"`        // Vote score: -1 (downvote), 0 (neutral), +1 (upvote)
    Nonce        uint64    `json:"nonce"`        // Unique nonce to prevent replay
    Signature    string    `json:"signature"`    // EIP-712 signature of the vote
    CreatedAt    time.Time `json:"created_at"`   // When the vote was submitted
    Processed    bool      `json:"processed"`    // Whether vote has been included in a batch
    BatchID      string    `json:"batch_id"`     // ID of the batch this vote was included in
}

// VoteBatch represents a collection of votes committed to blockchain
type VoteBatch struct {
    ID           string    `json:"id"`           // Batch ID (Merkle root hash)
    ToolID       string    `json:"tool_id"`      // Tool this batch is for
    VotesCount   int       `json:"votes_count"`  // Number of votes in this batch
    TotalScore   int64     `json:"total_score"`  // Sum of all vote scores in batch
    MerkleRoot   string    `json:"merkle_root"`  // Merkle root of votes in this batch
    BlockchainTx string    `json:"blockchain_tx"` // Transaction hash of on-chain commitment
    CommittedAt  time.Time `json:"committed_at"` // When batch was committed to blockchain
    CreatedAt    time.Time `json:"created_at"`   // When batch was created
}

// ToolReputation represents the current reputation state of a tool
type ToolReputation struct {
    ToolID           string    `json:"tool_id"`
    TotalScore       int64     `json:"total_score"`       // Sum of all vote scores
    TotalVotes       int64     `json:"total_votes"`       // Total number of votes
    AverageScore     float64   `json:"average_score"`     // Average score (total_score / total_votes)
    RecentScore      float64   `json:"recent_score"`      // Weighted average of recent votes
    LastCalculatedAt time.Time `json:"last_calculated_at"`
    LastBatchAt      time.Time `json:"last_batch_at"`     // When last batch was committed
}

// VoteSubmission is the request structure for submitting a vote
type VoteSubmission struct {
    ToolID       string `json:"tool_id" validate:"required"`
    VoterAddress string `json:"voter_address" validate:"required,eth_addr"`
    Score        int8   `json:"score" validate:"required,min=-1,max=1"`
    Nonce        uint64 `json:"nonce" validate:"required"`
    Signature    string `json:"signature" validate:"required"` // EIP-712 signature
}

// VoteVerificationResult result of vote signature verification
type VoteVerificationResult struct {
    Valid     bool   `json:"valid"`
    Reason    string `json:"reason,omitempty"`
    VoteID    string `json:"vote_id,omitempty"`
}