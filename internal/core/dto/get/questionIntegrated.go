package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type IntegratedIQuestionDto interface {
	SetId(uuid.UUID)

	SetDescription(string)

	SetOwnerId(nullUUID uuid.NullUUID)

	GetType() question.QuestionType
	SetType(question.QuestionType)
}

type IntegratedQuestionDto struct {
	Id          uuid.UUID             `json:"id"`
	Description string                `json:"description"`
	OwnerId     uuid.NullUUID         `json:"ownerId"`
	Type        question.QuestionType `json:"type"`
}

func (q *IntegratedQuestionDto) SetId(id uuid.UUID) {
	q.Id = id
}

func (q *IntegratedQuestionDto) SetDescription(description string) {
	q.Description = description
}

func (q *IntegratedQuestionDto) SetOwnerId(ownerId uuid.NullUUID) {
	q.OwnerId = ownerId
}

func (q *IntegratedQuestionDto) GetType() question.QuestionType {
	return q.Type
}

func (q *IntegratedQuestionDto) SetType(t question.QuestionType) {
	q.Type = t
}
