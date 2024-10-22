package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/service"
)

type FormGeneratedController struct {
	formGeneratedService *service.FormGeneratedService
}

func NewFormGeneratedController() *FormGeneratedController {
	return &FormGeneratedController{
		formGeneratedService: service.NewFormGeneratedService(),
	}
}

func (f *FormGeneratedController) GetMyForm(ctx *gin.Context) {
	publishedId := ctx.Param("publishedId")
	formGeneratedDto, err := f.formGeneratedService.GetMyForm(publishedId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, *formGeneratedDto)
}
