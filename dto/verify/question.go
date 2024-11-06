package verify

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/generated"
)

type Question struct {
	QuestionWithCorrectAnswer get.IQuestionDto
	QuestionWithStudentAnswer generated.IQuestion
}
