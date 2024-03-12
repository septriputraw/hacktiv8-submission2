package main

import (
	"fmt"
	"log"
	"os"
	"submission2/config"
	"submission2/routers"

	_ "submission2/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/joho/godotenv"
)

// @title Hacktiv8 Golang Submission 2 - API Documentation
// @description This is a REST API documentation for Hacktiv8 Golang Submission 2.
// @version 1.0
//
// @contact.name API Support
// @contact.email septriputraw@gmail.com
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading for .env file")
	}

	// start gin server in mode release
	gin.SetMode(gin.ReleaseMode)

	// start database
	config.StartDB()

	// setup server
	serverAddress := setupServer()

	// setup router
	router := routers.SetupRouter()

	// setup swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// run server
	err := router.Run(serverAddress)
	if err != nil {
		log.Fatal("Error running server: " + err.Error())
	}

}

func setupServer() string {
	host := os.Getenv("APP_HOST")
	port := os.Getenv("APP_PORT")
	return fmt.Sprintf("%s:%s", host, port)
}
