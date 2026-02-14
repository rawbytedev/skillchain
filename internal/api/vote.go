// internal/api/vote.go
package api

import (
	"net/http"
	"time"

	"moltket/internal/core"
	"moltket/internal/models"

	"github.com/labstack/echo/v4"
)

type voteHandler struct {
    voteService *core.VoteService
}

func NewVoteHandler(voteService *core.VoteService) *voteHandler {
    return &voteHandler{
        voteService: voteService,
    }
}

// SubmitVote handles POST /api/v1/vote/submit
func (h *voteHandler) SubmitVote(c echo.Context) error {
    var submission models.VoteSubmission
    
    if err := c.Bind(&submission); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid request format",
        })
    }
    
    // Basic validation
    if submission.Score < -1 || submission.Score > 1 {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Score must be between -1 and 1",
        })
    }
    
    if submission.Nonce == 0 {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Nonce is required",
        })
    }
    
    // Submit vote
    result, err := h.voteService.SubmitVote(c.Request().Context(), &submission)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to process vote",
        })
    }
    
    if !result.Valid {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{
            "valid":  false,
            "reason": result.Reason,
        })
    }
    
    return c.JSON(http.StatusOK, map[string]interface{}{
        "valid":    true,
        "vote_id":  result.VoteID,
        "message":  "Vote submitted successfully",
    })
}

// GetReputation handles GET /api/v1/vote/reputation/:toolId
func (h *voteHandler) GetReputation(c echo.Context) error {
    toolID := c.Param("toolId")
    if toolID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Tool ID is required",
        })
    }
    
    reputation, err := h.voteService.GetToolReputation(c.Request().Context(), toolID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to get reputation",
        })
    }
    
    return c.JSON(http.StatusOK, map[string]interface{}{
        "tool_id":    reputation.ToolID,
        "score":      reputation.TotalScore,
        "votes":      reputation.TotalVotes,
        "average":    reputation.AverageScore,
        "recent":     reputation.RecentScore,
        "updated_at": reputation.LastCalculatedAt.Format(time.RFC3339),
    })
}

// ProcessBatch handles POST /api/v1/vote/process-batch/:toolId
// This is an admin endpoint to manually trigger batch processing
func (h *voteHandler) ProcessBatch(c echo.Context) error {
    toolID := c.Param("toolId")
    if toolID == "" {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Tool ID is required",
        })
    }
    
    // In production, add authentication/authorization here
    // For demo, we'll allow anyone to trigger batch processing
    
    batch, err := h.voteService.ProcessBatch(c.Request().Context(), toolID)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": err.Error(),
        })
    }
    
    return c.JSON(http.StatusOK, map[string]interface{}{
        "batch_id":      batch.ID,
        "tool_id":       batch.ToolID,
        "votes_count":   batch.VotesCount,
        "total_score":   batch.TotalScore,
        "merkle_root":   batch.MerkleRoot,
        "created_at":    batch.CreatedAt.Format(time.RFC3339),
        "message":       "Batch processed successfully. Ready for blockchain commitment.",
    })
}

// Setup function to be called from main server
func (s *Server) setupVoteRoutes(voteService *core.VoteService) {
    voteHandler := NewVoteHandler(voteService)
    
    api := s.echo.Group("/api/v1/vote")
    api.POST("/submit", voteHandler.SubmitVote)
    api.GET("/reputation/:toolId", voteHandler.GetReputation)
    api.POST("/process-batch/:toolId", voteHandler.ProcessBatch)
}