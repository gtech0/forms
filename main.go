package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hedgehog-forms/controller"
	"hedgehog-forms/database"
	"hedgehog-forms/enviroment"
	"log"
)

func init() {
	enviroment.Load()
	database.Connect()
	database.Sync()
}

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	router.Use(cors.New(corsConfig))

	formPatternController := controller.NewFormPatternController()

	formPattern := router.Group("/api/form/pattern")
	{
		formPattern.POST("/create", formPatternController.CreateFormPattern)
	}

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
