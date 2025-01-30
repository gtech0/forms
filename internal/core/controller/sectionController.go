package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/internal/core/service"
	"net/http"
)

type SectionController struct {
	sectionService *service.SectionService
}

func NewSectionController() *SectionController {
	return &SectionController{
		sectionService: service.NewSectionService(),
	}
}

func (s *SectionController) GetSection(ctx *gin.Context) {
	sectionId := ctx.Param("sectionId")
	sectionDto, err := s.sectionService.GetSection(sectionId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *sectionDto)
}

func (s *SectionController) GetSections(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	sections, err := s.sectionService.GetSections(query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *sections)
}
