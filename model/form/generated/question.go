package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form"
)

type IQuestion interface {
	GetType() form.QuestionType
}

type Question struct {
	Id          uuid.UUID
	Description string
	Type        form.QuestionType
	Attachments []uuid.UUID
	Points      int
}
