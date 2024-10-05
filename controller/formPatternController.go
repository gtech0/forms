package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/database"
	"hedgehog-forms/dto"
	"hedgehog-forms/factory"
	"net/http"
)

type FormPatternController struct {
	formPatternFactory *factory.FormPatternFactory
}

func NewFormPatternController() *FormPatternController {
	return &FormPatternController{}
}

func (f *FormPatternController) CreateFormPattern(ctx *gin.Context) {
	body := dto.CreateFormPatternDto{}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	formPattern, err := f.formPatternFactory.BuildFormPattern(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err = database.DB.Create(&formPattern).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create form pattern",
		})
		return
	}

	ctx.Status(http.StatusOK)
}
