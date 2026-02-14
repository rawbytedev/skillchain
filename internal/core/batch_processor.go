// internal/core/batch_processor.go
package core

import (
	"context"
	"log"
	"time"

	"moltket/internal/kvstore"
	"moltket/internal/models"
)

type BatchProcessor struct {
    voteService  *VoteService
    cache        kvstore.Store
    interval     time.Duration
    stopChan     chan struct{}
}

func NewBatchProcessor(voteService *VoteService, cache kvstore.Store, interval time.Duration) *BatchProcessor {
    return &BatchProcessor{
        voteService: voteService,
        cache:       cache,
        interval:    interval,
        stopChan:    make(chan struct{}),
    }
}

// Start begins the batch processor in a goroutine
func (bp *BatchProcessor) Start(ctx context.Context) {
    go bp.run(ctx)
    log.Printf("Batch processor started with interval %v", bp.interval)
}

// Stop gracefully stops the batch processor
func (bp *BatchProcessor) Stop() {
    close(bp.stopChan)
    log.Println("Batch processor stopped")
}

// run is the main processing loop
func (bp *BatchProcessor) run(ctx context.Context) {
    ticker := time.NewTicker(bp.interval)
    defer ticker.Stop()
    
    for {
        select {
        case <-ticker.C:
            bp.processAllPendingBatches(ctx)
        case <-bp.stopChan:
            return
        case <-ctx.Done():
            return
        }
    }
}

// processAllPendingBatches finds and processes all tools with pending votes
func (bp *BatchProcessor) processAllPendingBatches(ctx context.Context) {
    // In production, this would query a database for tools with pending votes
    // For demo, we'll use a simplified approach with cache scanning
    
    // Get all pending vote keys (this is simplified - in production use better pattern)
    // Note: This cache scanning approach only works for small-scale demos
    keys := bp.cache.Keys(ctx)
    
    processedCount := 0
    for _, key := range keys {
        // Check if this is a pending vote key
        if len(key) > 12 && key[:12] == "pending:vote:" {
            toolID := key[12:] // Extract tool ID from key
            
            // Process batch for this tool
            batch, err := bp.voteService.ProcessBatch(ctx, toolID)
            if err != nil {
                log.Printf("Failed to process batch for tool %s: %v", toolID, err)
                continue
            }
            
            processedCount++
            log.Printf("Processed batch %s for tool %s with %d votes",
                batch.ID, toolID, batch.VotesCount)
            
            // In production, here we would:
            // 1. Sign the batch data with backend private key
            // 2. Submit to ReputationOracle.sol contract
            // 3. Update batch with transaction hash
            // 4. Notify relevant parties
            
            // For demo, we'll just log the batch info
            log.Printf("Batch ready for blockchain: %s (Merkle root: %s)",
                batch.ID, batch.MerkleRoot)
        }
    }
    
    if processedCount > 0 {
        log.Printf("Batch processing complete: %d batches processed", processedCount)
    }
}

// ManualProcessBatch manually processes batches for a specific tool
func (bp *BatchProcessor) ManualProcessBatch(ctx context.Context, toolID string) (*models.VoteBatch, error) {
    return bp.voteService.ProcessBatch(ctx, toolID)
}