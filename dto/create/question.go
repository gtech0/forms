package create

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form"
)

type QuestionDto struct {
	Type form.QuestionType `json:"type"`
}

type QuestionOnExistingDto struct {
	QuestionDto
	QuestionId uuid.UUID `json:"questionId"`
}

type NewQuestionDto struct {
	QuestionDto
	Description string      `json:"description"`
	Attachments []uuid.UUID `json:"attachments"`
}
