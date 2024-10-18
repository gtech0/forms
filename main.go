package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hedgehog-forms/controller"
	"hedgehog-forms/database"
	"hedgehog-forms/enviroment"
	"hedgehog-forms/errs"
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
	router.Use(errs.ErrorHandler)

	formPatternController := controller.NewFormPatternController()
	formPublishedController := controller.NewFormPublishedController()
	fileController := controller.NewFileController()

	formPattern := router.Group("/api/form/pattern")
	{
		formPattern.POST("/create", formPatternController.CreateForm)
		formPattern.GET("/get/:patternId", formPatternController.GetForm)
	}

	formPublished := router.Group("/api/form/published")
	{
		formPublished.POST("/create", formPublishedController.PublishForm)
		formPublished.GET("/get/:formId", formPublishedController.GetForm)
	}

	fileGroup := router.Group("/api/file")
	{
		fileGroup.POST("/upload", fileController.UploadFile)
		fileGroup.GET("/download/:fileId", fileController.DownloadFile)
	}

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
