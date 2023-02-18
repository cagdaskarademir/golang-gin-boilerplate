package configuration

import "github.com/cagdaskarademir/golang/order-service/handler"

func Routes(server *Server) {
	orderHandler := handler.OrderHandler{}
	api := server.Engine().Group("/api")
	api.GET("/orders", orderHandler.GetOrder())
}
