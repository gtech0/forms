package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type IQuestion interface {
	GetType() question.QuestionType
	GetId() uuid.UUID

	GetPoints() int
	SetPoints(int)
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

func (q *Question) GetId() uuid.UUID {
	return q.Id
}

func (q *Question) GetPoints() int {
	return q.Points
}

func (q *Question) SetPoints(points int) {
	q.Points = points
}
