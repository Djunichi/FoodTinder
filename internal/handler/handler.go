package handler

import (
	"context"
	"food-tinder/internal/config"
	"food-tinder/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const API_VERSION = "1.0"

type HttpHandler interface {
	Init()
	Stop(ctx context.Context) error
}

type httpHandler struct {
	router     *gin.Engine
	conf       *config.Config
	server     *http.Server
	productSvc ProductSvc
	sessionSvc SessionSvc
	voteSvc    VoteSvc
}

func NewHttpHandler(svc *service.Container, conf *config.Config) HttpHandler {
	return &httpHandler{
		conf:       conf,
		productSvc: svc.ProductService,
		sessionSvc: svc.SessionService,
		voteSvc:    svc.VoteService,
	}
}

func (h *httpHandler) Init() {
	h.router = gin.New()
	h.router.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/api/v1/ping"))

	h.addRoutes()

	h.server = &http.Server{
		Addr:    h.conf.HTTPPort, // например, ":8080"
		Handler: h.router,
	}

	go func() {
		log.Printf("Starting server on %s\n", h.conf.HTTPPort)
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()
}

func (h *httpHandler) Stop(ctx context.Context) error {
	log.Println("Shutting down HTTP server...")
	return h.server.Shutdown(ctx)
}

func (h *httpHandler) healthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"data": "pong"})
	}
}

func (h *httpHandler) version() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"version": API_VERSION})
	}
}
