package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/service"
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

// CreateQuestion godoc
// @Tags         Question
// @Summary      Create question
// @Description  create question
// @Produce      json
// @Param   	 subjectId path string true "Subject id"
// @Success      200 {object} get.QuestionDto
// @Failure      400 {object} errs.CustomError
// @Router       /question/create/{subjectId} [post]
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

// GetQuestion godoc
// @Tags         Question
// @Summary      Get question
// @Description  get question
// @Produce      json
// @Param   	 questionId path string true "Question id"
// @Success      200 {object} get.QuestionDto
// @Failure      400 {object} errs.CustomError
// @Router       /question/get/{questionId} [get]
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

// DeleteQuestion godoc
// @Tags         Question
// @Summary      Delete question
// @Description  delete question
// @Produce      json
// @Param   	 questionId path string true "Question id"
// @Success      200
// @Failure      400 {object} errs.CustomError
// @Router       /question/delete/{questionId} [delete]
func (q *QuestionController) DeleteQuestion(ctx *gin.Context) {
	questionId := ctx.Param("questionId")
	if err := q.questionService.DeleteQuestion(questionId); err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusOK)
}
