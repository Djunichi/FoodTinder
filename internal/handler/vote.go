package handler

import "github.com/gin-gonic/gin"

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

	}
}
