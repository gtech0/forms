package verify

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/generated"
)

type Question struct {
	QuestionWithCorrectAnswer get.QuestionDto
	QuestionWithStudentAnswer generated.IQuestion
}
