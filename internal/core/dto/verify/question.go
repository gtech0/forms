package verify

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/generated"
)

type Question struct {
	QuestionWithCorrectAnswer get.IQuestionDto
	QuestionWithStudentAnswer generated.IQuestion
}
