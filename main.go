package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"website_status/controllers"
)

var router *gin.Engine

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	initializeRoutes()

	router.Run("127.0.0.1:" + os.Getenv("PORT"))
}

func initializeRoutes() {
	router.GET("/", controllers.RenderStatusPage)
}
