package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type IQuestion interface {
	GetType() question.QuestionType
}

type Question struct {
	Id          uuid.UUID             `json:"id"`
	Description string                `json:"description"`
	Type        question.QuestionType `json:"type"`
	Attachments []uuid.UUID           `json:"attachments"`
	Points      int                   `json:"points"`
}

func (q *Question) GetType() question.QuestionType {
	return q.Type
}
