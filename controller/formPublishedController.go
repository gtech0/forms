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
	formId := ctx.Param("formId")

	formPublished, err := f.formPublishedService.GetForm(formId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *formPublished)
}
