package handler

import (
	"context"
	"errors"
	"food-tinder/internal/config"
	"food-tinder/internal/service"
	"food-tinder/internal/writer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	logger     *zap.SugaredLogger
}

func NewHttpHandler(svc *service.Container, conf *config.Config, logger *zap.SugaredLogger) HttpHandler {
	return &httpHandler{
		conf:       conf,
		productSvc: svc.ProductService,
		sessionSvc: svc.SessionService,
		voteSvc:    svc.VoteService,
		logger:     logger,
	}
}

func (h *httpHandler) Init() {
	h.router = gin.New()

	logWriter := writer.NewZapWriter(h.logger)
	h.router.Use(gin.LoggerWithWriter(logWriter, "/api/v1/ping"))

	h.addRoutes()

	h.server = &http.Server{
		Addr:    h.conf.HTTPPort, // например, ":8080"
		Handler: h.router,
	}

	go func() {
		h.logger.Infof("Starting server on %s\n", h.conf.HTTPPort)
		if err := h.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			h.logger.Fatalf("HTTP server error: %v", err)
		}
	}()
}

func (h *httpHandler) Stop(ctx context.Context) error {
	h.logger.Infof("Shutting down HTTP server...")
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
