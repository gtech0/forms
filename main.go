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
	formGeneratedController := controller.NewFormGeneratedController()
	fileController := controller.NewFileController()
	subjectController := controller.NewSubjectController()

	formPattern := router.Group("/api/form/pattern")
	{
		formPattern.POST("/create", formPatternController.CreateForm)
		formPattern.GET("/get/:patternId", formPatternController.GetForm)
		formPattern.GET("/get", formPatternController.GetForms)
	}

	formPublished := router.Group("/api/form/published")
	{
		formPublished.POST("/create", formPublishedController.PublishForm)
		formPublished.GET("/get/:formId", formPublishedController.GetForm)
		formPublished.GET("/get", formPublishedController.GetForms)
	}

	formGenerated := router.Group("/api/form/generated")
	{
		formGenerated.GET("/get/:publishedId", formGeneratedController.GetMyForm)
	}

	fileGroup := router.Group("/api/file")
	{
		fileGroup.POST("/upload", fileController.UploadFile)
		fileGroup.GET("/download/:fileId", fileController.DownloadFile)
	}

	subject := router.Group("/api/subject")
	{
		subject.POST("/create", subjectController.CreateSubject)
		subject.GET("/get/:subjectId", subjectController.GetSubject)
		subject.GET("/get", subjectController.GetSubjects)
		subject.PUT("/update/:subjectId", subjectController.UpdateSubject)
		subject.DELETE("/delete/:subjectId", subjectController.DeleteSubject)
	}

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}
