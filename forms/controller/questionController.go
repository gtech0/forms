package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hedgehog-forms/errs"
	"hedgehog-forms/service"
	"net/http"
)

type QuestionController struct {
	questionService *service.QuestionService
}

func NewQuestionController() *QuestionController {
	return &QuestionController{
		questionService: service.NewQuestionService(),
	}
}

func (q *QuestionController) CreateQuestion(ctx *gin.Context) {
	subjectId := ctx.Param("subjectId")
	var body json.RawMessage
	if err := ctx.Bind(&body); err != nil {
		ctx.Error(errs.New(err.Error(), 500))
		return
	}

	questionDto, err := q.questionService.CreateQuestion(subjectId, body)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *questionDto)
}

func (q *QuestionController) GetQuestion(ctx *gin.Context) {
	questionId := ctx.Param("questionId")
	questionDto, err := q.questionService.GetQuestion(questionId)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *questionDto)
}

func (q *QuestionController) GetQuestions(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	questions, err := q.questionService.GetQuestions(query)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, *questions)
}

func (q *QuestionController) DeleteQuestion(ctx *gin.Context) {
	questionId := ctx.Param("questionId")
	if err := q.questionService.DeleteQuestion(questionId); err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
