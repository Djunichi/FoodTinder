package handler

import (
	"food-tinder/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

// createVote godoc
// @Summary Creates a new Vote
// @Schemas
// @Description
// @Tags votes
// @Accept  json
// @Produce  json
// @Param input body dto.CreateVoteReq true "Request Body"
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /votes/create [post]
func (h *httpHandler) createVote() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &dto.CreateVoteReq{}
		if err := c.ShouldBindJSON(req); err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(400, gin.H{"error": "error when parsing data"})
			return
		}

		err := h.voteSvc.CreateVote(c, req)
		if err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, nil)
	}
}

// createManyVotes godoc
// @Summary Creates a new Votes
// @Schemas
// @Description
// @Tags votes
// @Accept  json
// @Produce  json
// @Param input body dto.CreateVotesReq true "Request Body"
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /votes/create-many [post]
func (h *httpHandler) createManyVotes() gin.HandlerFunc {
	return func(c *gin.Context) {

		req := &dto.CreateVotesReq{}
		if err := c.ShouldBindJSON(req); err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(400, gin.H{"error": "error when parsing data"})
			return
		}

		err := h.voteSvc.CreateVotes(c, req)
		if err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, nil)
	}
}

// updateVote godoc
// @Summary Updates a new Vote
// @Schemas
// @Description
// @Tags votes
// @Accept  json
// @Produce  json
// @Param input body dto.UpdateVoteReq true "Request Body"
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /votes/update [put]
func (h *httpHandler) updateVote() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &dto.UpdateVoteReq{}
		if err := c.ShouldBindJSON(req); err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(400, gin.H{"error": "error when parsing data"})
			return
		}

		err := h.voteSvc.UpdateVote(c, req)
		if err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, nil)
	}
}

// updateManyVotes godoc
// @Summary Updates a new Votes
// @Schemas
// @Description
// @Tags votes
// @Accept  json
// @Produce  json
// @Param input body dto.UpdateVotesReq true "Request Body"
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /votes/update-many [put]
func (h *httpHandler) updateManyVotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		req := &dto.UpdateVotesReq{}
		if err := c.ShouldBindJSON(req); err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(400, gin.H{"error": "error when parsing data"})
			return
		}

		err := h.voteSvc.UpdateVotes(c, req)
		if err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, nil)
	}
}

// getVoteBySession godoc
// @Summary Get votes by specific session id
// @Schemas
// @Description
// @Tags votes
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /votes/get-by-session/{session-id} [get]
func (h *httpHandler) getVoteBySession() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := uuid.Parse(c.Param("session-id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "session-id must be a valid uuid"})
			return
		}

		votes, err := h.voteSvc.GetVotesBySession(c, sessionId)
		if err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, votes)
	}
}

// getAggregatedScores godoc
// @Summary Gets an aggregated scores across all sessions
// @Schemas
// @Description
// @Tags votes
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /votes/get-aggregated-scores [get]
func (h *httpHandler) getAggregatedScores() gin.HandlerFunc {
	return func(c *gin.Context) {

		votes, err := h.voteSvc.GetAggregatedScoresByAllSessions(c)
		if err != nil {
			log.Printf("[Vote Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, votes)
	}
}
