package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OrderHandler struct {
}

func (handler OrderHandler) GetOrder() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.String(http.StatusOK, "Hello Gin!")
	}
}
