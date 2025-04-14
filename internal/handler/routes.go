package handler

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"net/http"
	"runtime/debug"
	"time"
)

const routePrefix = "api/v1"

func (h *httpHandler) addRoutes() {
	h.router.Use(UseRecoverMiddleware())
	h.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	voteRouter := h.router.Group(pathWithAction("votes"))
	voteRouter.POST("create", h.createVote())
	voteRouter.POST("create-many", h.createManyVotes())
	voteRouter.PUT("update", h.updateVote())
	voteRouter.PUT("update-many", h.updateManyVotes())
	voteRouter.GET("get-by-session", h.getVoteBySession())
	voteRouter.GET("get-aggregated-scores", h.getAggregatedScores())

	productRouter := h.router.Group(pathWithAction("products"))
	productRouter.GET("get-all", h.getAllProducts())
	productRouter.GET("get-unrated", h.getUnratedProducts())

	sessionRouter := h.router.Group(pathWithAction("sessions"))
	sessionRouter.POST("create-session", h.createSession())
	sessionRouter.GET("get-by-id", h.getSessionById())
	sessionRouter.GET("get-active", h.getActiveSessions())
}

func UseRecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {

					stackInfo := debug.Stack()
					log.Println("Recovered from panic", err, string(stackInfo))

					// Respond with an internal server error
					c.AbortWithStatus(http.StatusInternalServerError)
				} else {
					// If the recovered value is not an error, log it as a message
					log.Println("Recovered from panic with non-error value", r)

					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}()

		// Continue with the next middleware or route
		c.Next()
	}
}

func pathWithAction(action string) string {
	return fmt.Sprintf("%s/%s", routePrefix, action)
}
