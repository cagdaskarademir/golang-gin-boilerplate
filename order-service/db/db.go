package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

func Init() *gorm.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	fmt.Println(dataSourceName)

	dbDriver := os.Getenv("DB_DRIVER")

	// create instance for driver
	db, err := gorm.Open(dbDriver, dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	// The parameter sets the maximum number of open connections to the database.
	if maxOpenConnections, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONNECTIONS")); err != nil {
		log.Fatal(err.Error())
	} else {
		db.DB().SetMaxOpenConns(maxOpenConnections)
	}

	// The Parameter sets the maximum number of connections in the idle connection pool.
	if maxIdleConnections, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONNECTIONS")); err != nil {
		log.Fatal(err.Error())
	} else {
		db.DB().SetMaxIdleConns(maxIdleConnections)
	}

	// The Parameter sets the maximum amount of time a connection may be reused.
	if connectionMaxLifeTime, err := strconv.Atoi(os.Getenv("DB_CONNECTION_MAX_LIFE_TIME")); err != nil {
		log.Fatal(err.Error())
	} else {
		db.DB().SetConnMaxLifetime(time.Duration(connectionMaxLifeTime) * time.Second)
	}

	return db
}
