package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type IQuestion interface {
	GetId() uuid.UUID

	SetDescription(string)

	SetOrder(int)

	GetType() QuestionType
	SetType(QuestionType)

	SetVariantId(nullUUID uuid.NullUUID)

	SetDynamicBlockId(nullUUID uuid.NullUUID)

	SetIsQuestionFromBank(bool)
}

type Question struct {
	model.BaseModel
	Description        string
	Order              int
	Type               QuestionType
	Attachments        []Attachment
	VariantId          uuid.NullUUID `gorm:"type:uuid"`
	DynamicBlockId     uuid.NullUUID `gorm:"type:uuid"`
	Subject            model.Subject
	SubjectId          uuid.NullUUID `gorm:"type:uuid"`
	IsQuestionFromBank bool
}

func (q *Question) GetId() uuid.UUID {
	return q.Id
}

func (q *Question) SetDescription(description string) {
	q.Description = description
}

func (q *Question) SetOrder(order int) {
	q.Order = order
}

func (q *Question) GetType() QuestionType {
	return q.Type
}

func (q *Question) SetType(t QuestionType) {
	q.Type = t
}

func (q *Question) SetVariantId(id uuid.NullUUID) {
	q.VariantId = id
}

func (q *Question) SetDynamicBlockId(id uuid.NullUUID) {
	q.DynamicBlockId = id
}

func (q *Question) SetIsQuestionFromBank(isQuestionFromBank bool) {
	q.IsQuestionFromBank = isQuestionFromBank
}
