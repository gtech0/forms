package controller

import (
	"github.com/gin-gonic/gin"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/service"
	"net/http"
)

type SubjectController struct {
	subjectService *service.SubjectService
}

func NewSubjectController() *SubjectController {
	return &SubjectController{
		subjectService: service.NewSubjectService(),
	}
}

func (s *SubjectController) Create(ctx *gin.Context) {
	var subjectDto create.SubjectDto
	if err := ctx.Bind(&subjectDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	subject, err := s.subjectService.Create(subjectDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *subject)
}
