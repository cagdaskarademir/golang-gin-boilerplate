package configuration

import (
	"github.com/cagdaskarademir/golang/order-service/constant"
	docs "github.com/cagdaskarademir/golang/order-service/docs"
	"github.com/cagdaskarademir/golang/order-service/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(server *Server) {
	docs.SwaggerInfo.BasePath = constant.DefaultApiPath
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	orderHandler := handler.OrderHandler{}

	v1 := server.Engine().Group("/api/v1")
	{
		eg := v1.Group("/transactions")
		{
			eg.GET("/orders", orderHandler.GetOrder())
			eg.POST("/order", orderHandler.PostOrder())
		}
	}
}
