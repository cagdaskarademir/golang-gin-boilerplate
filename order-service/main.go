package main

import (
	"fmt"
	"github.com/cagdaskarademir/golang/order-service/configuration"
	"github.com/cagdaskarademir/golang/order-service/db"
	"github.com/cagdaskarademir/golang/order-service/docs"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file. Please check it!")
	}

	// Setup Connection
	connection := db.Init()

	defer func() {
		if err := connection.DB().Close(); err != nil {
			log.Print(err.Error())
		}
	}()

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("EXPOSE_PORT"))

	server := configuration.InitServer(connection)

	// Route
	configuration.Routes(server)

	if err := server.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(err.Error())
	}

}
