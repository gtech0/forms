package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type IntegratedIQuestionDto interface {
	GetType() question.QuestionType
}

type IntegratedQuestionDto struct {
	Id          uuid.UUID             `json:"id"`
	Description string                `json:"description"`
	OwnerId     uuid.NullUUID         `json:"ownerId"`
	Type        question.QuestionType `json:"type"`
}

func (q *IntegratedQuestionDto) GetType() question.QuestionType {
	return q.Type
}
