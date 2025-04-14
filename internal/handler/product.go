package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// getAllProducts godoc
// @Summary Gets all products
// @Schemas
// @Description
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /products/get-all [get]
func (h *httpHandler) getAllProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		products, err := h.productSvc.GetAllProducts(c)
		if err != nil {
			h.logger.Errorf("[Product Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, products)
	}
}

// getUnratedProducts godoc
// @Summary Gets unrated products
// @Schemas
// @Description
// @Tags products
// @Accept  json
// @Produce  json
// @Param session-id query string true "Session UUID"
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /products/get-unrated [get]
func (h *httpHandler) getUnratedProducts() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := uuid.Parse(c.Query("session-id"))
		if err != nil {
			c.JSON(400, gin.H{"error": "session-id must be a valid uuid"})
			return
		}

		products, err := h.productSvc.GetUnratedProducts(c, sessionId)
		if err != nil {
			h.logger.Errorf("[Product Handler] %v", err)
			c.JSON(500, gin.H{"error": "internal server error"})
		}

		c.JSON(200, products)
	}
}
