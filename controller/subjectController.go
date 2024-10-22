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

func (s *SubjectController) CreateSubject(ctx *gin.Context) {
	var subjectDto create.SubjectDto
	if err := ctx.Bind(&subjectDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	subject, err := s.subjectService.CreateSubject(subjectDto)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *subject)
}

func (s *SubjectController) GetSubject(ctx *gin.Context) {
	subjectId := ctx.Param("subjectId")
	subject, err := s.subjectService.GetSubject(subjectId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *subject)
}

func (s *SubjectController) GetSubjects(ctx *gin.Context) {
	name := ctx.Request.URL.Query().Get("name")
	subjects, err := s.subjectService.GetSubjects(name)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, subjects)
}

func (s *SubjectController) UpdateSubject(ctx *gin.Context) {
	var subjectDto create.SubjectDto
	if err := ctx.Bind(&subjectDto); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	subjectId := ctx.Param("subjectId")
	subject, err := s.subjectService.UpdateSubject(subjectId, subjectDto.Name)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *subject)
}

func (s *SubjectController) DeleteSubject(ctx *gin.Context) {
	subjectId := ctx.Param("subjectId")
	err := s.subjectService.DeleteSubject(subjectId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
