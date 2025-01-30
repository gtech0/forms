package web

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"hedgehog-forms/internal/core/controller"
	"hedgehog-forms/internal/core/errs"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	router.Use(cors.New(corsConfig))
	router.Use(errs.ErrorHandler)

	questionController := controller.NewQuestionController()
	sectionController := controller.NewSectionController()
	formPatternController := controller.NewFormPatternController()
	formPublishedController := controller.NewFormPublishedController()
	formGeneratedController := controller.NewFormGeneratedController()
	fileController := controller.NewFileController()
	subjectController := controller.NewSubjectController()

	questionRouter := router.Group("/api/question")
	{
		questionRouter.POST("/create/:subjectId", questionController.CreateQuestion)
		questionRouter.GET("/get/:questionId", questionController.GetQuestion)
		questionRouter.GET("/get", questionController.GetQuestions)
		questionRouter.DELETE("/delete/:questionId", questionController.DeleteQuestion)
	}

	sectionRouter := router.Group("/api/section")
	{
		sectionRouter.GET("/get/:sectionId", sectionController.GetSection)
		sectionRouter.GET("/get", sectionController.GetSections)
	}

	formPattern := router.Group("/api/form/pattern")
	{
		formPattern.POST("/create", formPatternController.CreateForm)
		formPattern.GET("/get/:patternId", formPatternController.GetForm)
		formPattern.GET("/get", formPatternController.GetForms)
	}

	formPublished := router.Group("/api/form/published")
	{
		formPublished.POST("/create", formPublishedController.PublishForm)
		formPublished.GET("/get/:publishedId", formPublishedController.GetForm)
		formPublished.GET("/get", formPublishedController.GetForms)
		formPublished.PUT("/update/:publishedId", formPublishedController.UpdateForm)
		formPublished.DELETE("/delete/:publishedId", formPublishedController.DeleteForm)
	}

	formGenerated := router.Group("/api/form/generated")
	{
		formGenerated.POST("/get/:publishedId", formGeneratedController.GetMyForm)
		formGenerated.POST("/save/:generatedId", formGeneratedController.SaveAnswers)
		formGenerated.POST("/submit/:generatedId", formGeneratedController.SubmitForm)
		formGenerated.POST("/unsubmit/:generatedId", formGeneratedController.UnSubmitForm)
		formGenerated.GET("/get/all/:subjectId/:userId", formGeneratedController.GetMyForms)
		formGenerated.GET("/get/submitted/all/:publishedId/:userId", formGeneratedController.GetSubmittedForms)
		formGenerated.GET("/get/unsubmitted/:publishedId", formGeneratedController.GetUsersWithUnsubmittedForms)
		formGenerated.GET("/get/submitted/:generatedId", formGeneratedController.GetSubmittedForm)
		formGenerated.POST("/verify/:generatedId", formGeneratedController.VerifyForm)
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

	return router
}
