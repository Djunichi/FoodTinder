package handler

import "github.com/gin-gonic/gin"

// createSession godoc
// @Summary Creates a new Session
// @Schemas
// @Description
// @Tags sessions
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /sessions/create-session [post]
func (h *httpHandler) createSession() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// getSessionById godoc
// @Summary Gets a session by specific id
// @Schemas
// @Description
// @Tags sessions
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /sessions/get-by-id/{session-id} [get]
func (h *httpHandler) getSessionById() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

// getActiveSessions godoc
// @Summary Gets all active sessions
// @Schemas
// @Description
// @Tags sessions
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /sessions/get-active [get]
func (h *httpHandler) getActiveSessions() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
