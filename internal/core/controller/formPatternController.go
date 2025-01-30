package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/service"
	"net/http"
)

type FormPatternController struct {
	formPatternService *service.FormPatternService
}

func NewFormPatternController() *FormPatternController {
	return &FormPatternController{
		formPatternService: service.NewFormPatternService(),
	}
}

func (f *FormPatternController) CreateForm(ctx *gin.Context) {
	body := create.FormPatternDto{}
	if err := ctx.Bind(&body); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	dto, err := f.formPatternService.CreatePattern(body)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *dto)
}

func (f *FormPatternController) GetForm(ctx *gin.Context) {
	patternId := ctx.Param("patternId")
	dto, err := f.formPatternService.GetForm(patternId)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, *dto)
}

func (f *FormPatternController) GetForms(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	forms, err := f.formPatternService.GetForms(query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *forms)
}
