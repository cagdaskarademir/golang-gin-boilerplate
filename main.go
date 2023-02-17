package main

import (
	"fmt"
	"github.com/cagdaskarademir/golang-gin-boilerplate/src/db"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file. Please check it!")
	}

	fmt.Println(os.Getenv("ProjectName"))

	// Setup Connection
	connection := db.Init()

	defer func() {
		if err := connection.DB().Close(); err != nil {
			log.Print(err.Error())
		}
	}()

}
