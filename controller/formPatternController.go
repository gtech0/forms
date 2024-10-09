package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/model/form"
	"net/http"
)

type FormPatternController struct {
	formPatternFactory *factory.FormPatternFactory
	formPatternMapper  *mapper.FormPatternMapper
}

func NewFormPatternController() *FormPatternController {
	return &FormPatternController{
		formPatternFactory: factory.NewFormPatternFactory(),
		formPatternMapper:  mapper.NewFormPatternMapper(),
	}
}

func (f *FormPatternController) CreateFormPattern(ctx *gin.Context) {
	body := create.FormPatternDto{}
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

func (f *FormPatternController) GetFormPattern(ctx *gin.Context) {
	patternId := ctx.Param("patternId")
	if patternId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "patternId is required",
		})
		return
	}

	parsedPatternId, err := uuid.Parse(patternId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	var formPattern form.FormPattern
	if err = database.DB.Model(&form.FormPattern{}).
		Preload("Subject").
		Preload("Sections.DynamicBlocks.~~~as~~~.~~~as~~~").
		Preload("Sections.StaticBlocks.Variants.~~~as~~~.~~~as~~~").
		First(&formPattern, "form_patterns.id = ?", parsedPatternId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	dto, err := f.formPatternMapper.ToDto(formPattern)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, dto)
}
