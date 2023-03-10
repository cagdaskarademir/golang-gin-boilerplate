package configuration

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"os"
)

type Server struct {
	engine   *gin.Engine
	database *gorm.DB
}

// InitServer initialized the default configurations
func InitServer(connection *gorm.DB) *Server {
	environment := os.Getenv("ENVIRONMENT")
	if len(environment) == 0 {
		environment = "debug"
	}

	gin.SetMode(environment)

	return &Server{
		engine:   gin.Default(),
		database: connection,
	}
}

func (server *Server) Run(port string) error {
	return server.engine.Run(":" + port)
}

// Engine Provides Engine Values
func (server *Server) Engine() *gin.Engine {
	return server.engine
}

// Database Gets Variable
func (server *Server) Database() *gorm.DB {
	return server.database
}
