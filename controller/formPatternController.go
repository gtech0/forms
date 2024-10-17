package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/service"
	"net/http"
)

type FormPatternController struct {
	formPatternFactory *factory.FormPatternFactory
	formPatternMapper  *mapper.FormPatternMapper
	attachmentService  *service.AttachmentService
}

func NewFormPatternController() *FormPatternController {
	return &FormPatternController{
		formPatternFactory: factory.NewPatternFactory(),
		formPatternMapper:  mapper.NewFormPatternMapper(),
		attachmentService:  service.NewAttachmentService(),
	}
}

func (f *FormPatternController) CreatePattern(ctx *gin.Context) {
	body := create.FormPatternDto{}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	formPattern, err := f.formPatternFactory.BuildPattern(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	attachmentIds, err := f.attachmentService.ValidatePatternAttachments(formPattern)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(attachmentIds) > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("incorrect attachment ids: %v", attachmentIds),
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

func (f *FormPatternController) GetPattern(ctx *gin.Context) {
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

	var formPattern pattern.FormPattern
	if err = database.DB.Model(&pattern.FormPattern{}).
		Preload("Subject").
		Preload("Sections.DynamicBlocks.~~~as~~~.~~~as~~~.~~~as~~~").
		Preload("Sections.StaticBlocks.Variants.~~~as~~~.~~~as~~~.~~~as~~~").
		First(&formPattern, "form_pattern.id = ?", parsedPatternId).Error; err != nil {
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
