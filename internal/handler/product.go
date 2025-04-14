package handler

import "github.com/gin-gonic/gin"

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

	}
}

// getUnratedProducts godoc
// @Summary Gets unrated products
// @Schemas
// @Description
// @Tags products
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 500 {object} error "Internal Server Error"
// @Router /products/get-unrated [get]
func (h *httpHandler) getUnratedProducts() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
