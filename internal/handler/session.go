package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
)

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
		session, err := h.sessionSvc.CreateSession(c)
		if err != nil {
			log.Printf("[Session Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, session)
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
		sessionId, err := uuid.Parse(c.Param("session-id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "session-id must be a valid uuid"})
			return
		}

		session, err := h.sessionSvc.GetSessionById(c, sessionId)
		if err != nil {
			log.Printf("[Session Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, session)
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
		sessions, err := h.sessionSvc.GetActiveSessions(c)
		if err != nil {
			log.Printf("[Session Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, sessions)
	}
}
