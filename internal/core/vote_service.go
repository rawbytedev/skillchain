// internal/core/vote_service.go
package core

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"time"

	"moltket/config"
	"moltket/internal/auth"
	"moltket/internal/kvstore"
	"moltket/internal/models"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type VoteService struct {
	config        *config.Config
	cache         kvstore.Store
	signer        *auth.EIP712Signer
	batchInterval time.Duration
}

// NewVoteService creates a new vote service
func NewVoteService(
	cfg *config.Config,
	cache kvstore.Store,
	signer *auth.EIP712Signer,
) *VoteService {
	return &VoteService{
		config:        cfg,
		cache:         cache,
		signer:        signer,
		batchInterval: 5 * time.Minute, // Batch votes every 5 minutes
	}
}

// SubmitVote processes and stores a new vote
func (s *VoteService) SubmitVote(ctx context.Context, submission *models.VoteSubmission) (*models.VoteVerificationResult, error) {
	// 1. Verify vote signature
	isValid, reason, err := s.verifyVoteSignature(submission)
	if err != nil {
		return nil, fmt.Errorf("signature verification failed: %w", err)
	}

	if !isValid {
		return &models.VoteVerificationResult{
			Valid:  false,
			Reason: reason,
		}, nil
	}

	// 2. Check if vote already exists (replay protection)
	voteID := s.generateVoteID(submission)
	existingKey := fmt.Sprintf("vote:%s", voteID)
	if _, found := s.cache.Get(ctx, existingKey); found {
		return &models.VoteVerificationResult{
			Valid:  false,
			Reason: "vote already submitted",
		}, nil
	}

	// 3. Check voter eligibility (must have used the tool at least once)
	eligible, err := s.isVoterEligible(ctx, submission.VoterAddress, submission.ToolID)
	if err != nil {
		return nil, fmt.Errorf("eligibility check failed: %w", err)
	}

	if !eligible {
		return &models.VoteVerificationResult{
			Valid:  false,
			Reason: "voter not eligible - must use tool before voting",
		}, nil
	}

	// 4. Create and store vote
	vote := &models.Vote{
		ID:           voteID,
		ToolID:       submission.ToolID,
		VoterAddress: submission.VoterAddress,
		Score:        submission.Score,
		Nonce:        submission.Nonce,
		Signature:    submission.Signature,
		CreatedAt:    time.Now(),
		Processed:    false,
	}

	// Store vote in pending votes list
	pendingKey := fmt.Sprintf("pending:vote:%s", submission.ToolID)
	var pendingVotes []*models.Vote
	if cached, found := s.cache.Get(ctx, pendingKey); found {
		if votes, ok := cached.([]*models.Vote); ok {
			pendingVotes = votes
		}
	}

	pendingVotes = append(pendingVotes, vote)
	if err := s.cache.Set(ctx, pendingKey, pendingVotes, s.batchInterval*2); err != nil {
		return nil, fmt.Errorf("failed to store pending vote: %w", err)
	}

	// Also store individual vote for idempotency
	if err := s.cache.Set(ctx, existingKey, vote, 24*time.Hour); err != nil {
		return nil, fmt.Errorf("failed to cache vote: %w", err)
	}

	// 5. Update real-time reputation (cached, not final)
	s.updateCachedReputation(ctx, submission.ToolID, submission.Score)

	return &models.VoteVerificationResult{
		Valid:  true,
		VoteID: voteID,
	}, nil
}

// GetToolReputation returns the current reputation of a tool
func (s *VoteService) GetToolReputation(ctx context.Context, toolID string) (*models.ToolReputation, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("reputation:%s", toolID)
	if cached, found := s.cache.Get(ctx, cacheKey); found {
		if reputation, ok := cached.(*models.ToolReputation); ok {
			return reputation, nil
		}
	}

	// Calculate from pending votes if not cached
	pendingKey := fmt.Sprintf("pending:vote:%s", toolID)
	var totalScore int64
	var totalVotes int64

	if cached, found := s.cache.Get(ctx, pendingKey); found {
		if votes, ok := cached.([]*models.Vote); ok {
			for _, vote := range votes {
				totalScore += int64(vote.Score)
				totalVotes++
			}
		}
	}

	// Also check committed batches (in a real system, this would query a database)
	// For demo, we'll use cache-only approach

	var averageScore float64
	if totalVotes > 0 {
		averageScore = float64(totalScore) / float64(totalVotes)
	}

	reputation := &models.ToolReputation{
		ToolID:           toolID,
		TotalScore:       totalScore,
		TotalVotes:       totalVotes,
		AverageScore:     averageScore,
		RecentScore:      averageScore, // Simplified for demo
		LastCalculatedAt: time.Now(),
	}

	// Cache the reputation
	s.cache.Set(ctx, cacheKey, reputation, 1*time.Minute)

	return reputation, nil
}

// ProcessBatch creates a batch of pending votes for a tool
func (s *VoteService) ProcessBatch(ctx context.Context, toolID string) (*models.VoteBatch, error) {
	pendingKey := fmt.Sprintf("pending:vote:%s", toolID)

	cached, found := s.cache.Get(ctx, pendingKey)
	if !found {
		return nil, fmt.Errorf("no pending votes for tool %s", toolID)
	}

	votes, ok := cached.([]*models.Vote)
	if !ok || len(votes) == 0 {
		return nil, fmt.Errorf("no valid votes to process")
	}

	// Calculate batch statistics
	var totalScore int64
	for _, vote := range votes {
		totalScore += int64(vote.Score)
		vote.Processed = true
	}

	// Generate Merkle root (simplified for demo - in production use proper Merkle tree)
	merkleRoot := s.generateMerkleRoot(votes)

	batch := &models.VoteBatch{
		ID:         fmt.Sprintf("batch_%s_%d", toolID, time.Now().Unix()),
		ToolID:     toolID,
		VotesCount: len(votes),
		TotalScore: totalScore,
		MerkleRoot: merkleRoot,
		CreatedAt:  time.Now(),
	}

	// Store batch
	batchKey := fmt.Sprintf("batch:%s:%s", toolID, batch.ID)
	if err := s.cache.Set(ctx, batchKey, batch, 24*time.Hour); err != nil {
		return nil, fmt.Errorf("failed to store batch: %w", err)
	}

	// Store batch reference in votes
	for _, vote := range votes {
		vote.BatchID = batch.ID
		voteKey := fmt.Sprintf("vote:%s", vote.ID)
		s.cache.Set(ctx, voteKey, vote, 24*time.Hour)
	}

	// Clear pending votes (they're now in a batch)
	s.cache.Delete(ctx, pendingKey)

	// Update reputation with batch data
	s.updateReputationFromBatch(ctx, toolID, batch)

	return batch, nil
}

// verifyVoteSignature verifies an EIP-712 vote signature
func (s *VoteService) verifyVoteSignature(submission *models.VoteSubmission) (bool, string, error) {
	// Parse signature
	signatureBytes, err := hex.DecodeString(submission.Signature)
	if err != nil {
		return false, "invalid signature format", nil
	}

	if len(signatureBytes) != 65 {
		return false, "signature must be 65 bytes", nil
	}

	// Prepare EIP-712 typed data
	toolIDBig, ok := new(big.Int).SetString(submission.ToolID, 10)
	if !ok {
		return false, "invalid tool ID", nil
	}

	voterAddress := common.HexToAddress(submission.VoterAddress)

	// Note: In production, we would use the same EIP-712 structure as the signer
	// For demo, we'll verify against the backend signer's public key

	// Check if signature matches the expected signer (backend)
	// This is a simplified check - in reality, votes are signed by users, not backend
	// For the hackathon demo, we'll accept any valid ECDSA signature

	// Recover address from signature
	messageHash := s.getVoteMessageHash(voterAddress, toolIDBig, submission.Score, submission.Nonce)

	// Extract r, s, v from signature
	sigR := signatureBytes[:32]
	sigS := signatureBytes[32:64]
	vByte := signatureBytes[64]

	// Adjust v if necessary
	if vByte < 27 {
		vByte += 27
	}

	// Recover public key
	signature := append(sigR, sigS...)
	signature = append(signature, vByte)

	pubkey, err := crypto.SigToPub(messageHash, signature)
	if err != nil {
		return false, "failed to recover public key", nil
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubkey)

	// Verify recovered address matches voter address
	if recoveredAddr != voterAddress {
		return false, "signature does not match voter address", nil
	}

	return true, "", nil
}

// getVoteMessageHash generates the message hash for vote signing
func (s *VoteService) getVoteMessageHash(voter common.Address, toolID *big.Int, score int8, nonce uint64) []byte {
	// Simplified message hash for demo
	// In production, use proper EIP-712 typed data hashing
	message := fmt.Sprintf("Vote\nTool: %s\nVoter: %s\nScore: %d\nNonce: %d\nChain: %d",
		toolID.String(),
		voter.Hex(),
		score,
		nonce,
		s.config.ChainID,
	)

	return crypto.Keccak256([]byte(message))
}

// generateVoteID creates a unique ID for a vote
func (s *VoteService) generateVoteID(submission *models.VoteSubmission) string {
	data := fmt.Sprintf("%s:%s:%d:%d",
		submission.VoterAddress,
		submission.ToolID,
		submission.Score,
		submission.Nonce,
	)

	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

// isVoterEligible checks if a voter is eligible to vote on a tool
func (s *VoteService) isVoterEligible(ctx context.Context, voterAddress, toolID string) (bool, error) {
	// Check if voter has used the tool (simplified for demo)
	// In production, this would check usage logs or license status

	usageKey := fmt.Sprintf("usage:%s:%s", voterAddress, toolID)
	if cached, found := s.cache.Get(ctx, usageKey); found {
		if usageCount, ok := cached.(int64); ok && usageCount > 0 {
			return true, nil
		}
	}

	// Also check if they have a valid license
	licenseKey := fmt.Sprintf("license:%s:%s", voterAddress, toolID)
	if _, found := s.cache.Get(ctx, licenseKey); found {
		return true, nil
	}

	return false, nil
}

// updateCachedReputation updates the cached reputation for a tool
func (s *VoteService) updateCachedReputation(ctx context.Context, toolID string, score int8) {
	cacheKey := fmt.Sprintf("reputation:%s", toolID)

	var reputation *models.ToolReputation
	if cached, found := s.cache.Get(ctx, cacheKey); found {
		if rep, ok := cached.(*models.ToolReputation); ok {
			reputation = rep
		}
	}

	if reputation == nil {
		reputation = &models.ToolReputation{
			ToolID:           toolID,
			LastCalculatedAt: time.Now(),
		}
	}

	reputation.TotalScore += int64(score)
	reputation.TotalVotes++

	if reputation.TotalVotes > 0 {
		reputation.AverageScore = float64(reputation.TotalScore) / float64(reputation.TotalVotes)
		reputation.RecentScore = reputation.AverageScore // Simplified
	}

	reputation.LastCalculatedAt = time.Now()
	s.cache.Set(ctx, cacheKey, reputation, 1*time.Minute)
}

// updateReputationFromBatch updates reputation from a processed batch
func (s *VoteService) updateReputationFromBatch(ctx context.Context, toolID string, batch *models.VoteBatch) {
	cacheKey := fmt.Sprintf("reputation:%s", toolID)

	var reputation *models.ToolReputation
	if cached, found := s.cache.Get(ctx, cacheKey); found {
		if rep, ok := cached.(*models.ToolReputation); ok {
			reputation = rep
		}
	}

	if reputation == nil {
		reputation = &models.ToolReputation{
			ToolID: toolID,
		}
	}

	// In production, this would add batch stats to cumulative stats
	// For demo, we'll just update the timestamp
	reputation.LastBatchAt = time.Now()
	reputation.LastCalculatedAt = time.Now()

	s.cache.Set(ctx, cacheKey, reputation, 1*time.Minute)
}

// generateMerkleRoot generates a Merkle root from votes (simplified for demo)
func (s *VoteService) generateMerkleRoot(votes []*models.Vote) string {
	// Sort votes by ID for consistent ordering
	sort.Slice(votes, func(i, j int) bool {
		return votes[i].ID < votes[j].ID
	})

	// Generate leaf hashes
	var leafHashes [][]byte
	for _, vote := range votes {
		data := fmt.Sprintf("%s:%s:%d:%d:%s",
			vote.ID,
			vote.VoterAddress,
			vote.Score,
			vote.Nonce,
			vote.CreatedAt.Format(time.RFC3339),
		)
		hash := sha256.Sum256([]byte(data))
		leafHashes = append(leafHashes, hash[:])
	}

	// Simplified Merkle root (in production, use proper binary tree)
	if len(leafHashes) == 0 {
		emptyHash := sha256.Sum256([]byte("empty"))
		return hex.EncodeToString(emptyHash[:])
	}

	// For demo, just hash all leaf hashes together
	combined := []byte{}
	for _, leaf := range leafHashes {
		combined = append(combined, leaf...)
	}

	rootHash := sha256.Sum256(combined)
	return hex.EncodeToString(rootHash[:])
}
