package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samcoded/go_image_resize_api/controllers"
)

func init() {
	if os.Getenv("APP_ENV") != "production" {
		// load .env file
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/", controllers.Home)
	r.GET("/images", controllers.ImagesResize)
	r.Run() // listen and serve
}
