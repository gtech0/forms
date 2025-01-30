package controller

import (
	"github.com/gin-gonic/gin"
	create2 "hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/service"
	"net/http"
)

type FormGeneratedController struct {
	formGeneratedService *service.FormGeneratedService
}

func NewFormGeneratedController() *FormGeneratedController {
	return &FormGeneratedController{
		formGeneratedService: service.NewFormGeneratedService(),
	}
}

//TODO: userIds

func (f *FormGeneratedController) GetMyForm(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	var userIdDto create2.FormGeneratedUser
	if err := ctx.Bind(&userIdDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	formGeneratedDto, err := f.formGeneratedService.GetMyForm(publishedId, userIdDto.UserId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formGeneratedDto)
}

func (f *FormGeneratedController) SaveAnswers(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")
	var answerDto get.AnswerDto
	if err := ctx.Bind(&answerDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	formGeneratedDto, err := f.formGeneratedService.SaveAnswers(generatedId, answerDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formGeneratedDto)
}

func (f *FormGeneratedController) SubmitForm(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")
	var answerDto get.AnswerDto
	if err := ctx.Bind(&answerDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	myGeneratedDto, err := f.formGeneratedService.SubmitForm(generatedId, answerDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *myGeneratedDto)
}

func (f *FormGeneratedController) UnSubmitForm(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")

	myGeneratedDto, err := f.formGeneratedService.UnSubmitForm(generatedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *myGeneratedDto)
}

func (f *FormGeneratedController) GetMyForms(ctx *gin.Context) {
	subjectId := ctx.Param("subjectId")
	userId := ctx.Param("userId")
	query := ctx.Request.URL.Query()
	response, err := f.formGeneratedService.GetMyForms(userId, subjectId, query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *response)
}

func (f *FormGeneratedController) GetSubmittedForms(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	userId := ctx.Param("userId")
	query := ctx.Request.URL.Query()
	response, err := f.formGeneratedService.GetSubmittedForms(userId, publishedId, query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *response)
}

func (f *FormGeneratedController) GetUsersWithUnsubmittedForms(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	ids, err := f.formGeneratedService.GetUsersWithUnsubmittedForm(publishedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, ids)
}

func (f *FormGeneratedController) GetSubmittedForm(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")
	submittedForm, err := f.formGeneratedService.GetSubmittedForm(generatedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *submittedForm)
}

func (f *FormGeneratedController) VerifyForm(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")
	var checkDto create2.CheckDto
	if err := ctx.Bind(&checkDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	formGeneratedDto, err := f.formGeneratedService.VerifyForm(generatedId, checkDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formGeneratedDto)
}

func (f *FormGeneratedController) ReturnForm(ctx *gin.Context) {
	generatedId := ctx.Param("generatedId")

	myGeneratedDto, err := f.formGeneratedService.ReturnForm(generatedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *myGeneratedDto)
}
