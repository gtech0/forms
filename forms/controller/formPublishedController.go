package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/mapper"
	"hedgehog-forms/service"
	"net/http"
)

type FormPublishedController struct {
	formPublishedMapper  *mapper.FormPublishedMapper
	formPublishedService *service.FormPublishedService
}

func NewFormPublishedController() *FormPublishedController {
	return &FormPublishedController{
		formPublishedMapper:  mapper.NewFormPublishedMapper(),
		formPublishedService: service.NewFormPublishedService(),
	}
}

func (f *FormPublishedController) PublishForm(ctx *gin.Context) {
	var publishDto create.FormPublishDto
	if err := ctx.Bind(&publishDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	formPublished, err := f.formPublishedService.PublishForm(publishDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formPublished)
}

func (f *FormPublishedController) GetForm(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")

	formPublished, err := f.formPublishedService.GetForm(publishedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formPublished)
}

func (f *FormPublishedController) GetForms(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	forms, err := f.formPublishedService.GetForms(query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *forms)
}

func (f *FormPublishedController) UpdateForm(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	var updateFormPublishedDto create.UpdateFormPublishedDto
	if err := ctx.Bind(&updateFormPublishedDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	response, err := f.formPublishedService.UpdateForm(publishedId, updateFormPublishedDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *response)
}

func (f *FormPublishedController) DeleteForm(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	err := f.formPublishedService.DeleteForm(publishedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
