package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hedgehog-forms/controller"
	"hedgehog-forms/database"
	"hedgehog-forms/enviroment"
	"hedgehog-forms/file"
	"log"
)

func init() {
	enviroment.Load()
	database.Connect()
	database.Sync()
	file.InitializeClient()
}

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	router.Use(cors.New(corsConfig))

	formPatternController := controller.NewFormPatternController()
	fileController := controller.NewFileController()

	formPattern := router.Group("/api/form/pattern")
	{
		formPattern.POST("/create", formPatternController.CreateFormPattern)
		formPattern.GET("/get/:patternId", formPatternController.GetFormPattern)
	}

	fileGroup := router.Group("/api/file")
	{
		fileGroup.POST("/upload", fileController.UploadFile)
	}

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
