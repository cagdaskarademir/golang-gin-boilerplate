package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
}

// GetOrder get order sample
// @Summary Get Order Transactions With No Filter
// @Schemes
// @Description details summary
// @Tags Order Operations
// @Accept json
// @Produce json
// @Success 200 {string} Hello gin!
// @Router /orders [get]
func (handler OrderHandler) GetOrder() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, "Hello Gin!")
	}
}

func (handler OrderHandler) PostOrder() gin.HandlerFunc {

	return func(context *gin.Context) {
		context.JSON(http.StatusOK, "{'id':1}")
	}
}
