package main

import (
	"github.com/cagdaskarademir/golang-gin-boilerplate/src"
	"github.com/cagdaskarademir/golang-gin-boilerplate/src/db"
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

	// Todo:  Swagger Integration

	app := src.InitServer(connection)

	// Route

	if err := app.Run(os.Getenv("PORT")); err != nil {
		log.Fatal(err.Error())
	}

}
