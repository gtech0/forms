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

	SetVariantId(uuid.UUID)

	SetDynamicBlockId(uuid.UUID)

	SetIsQuestionFromBank(bool)
}

type Question struct {
	model.BaseModel
	Description        string
	Order              int
	Type               QuestionType
	VariantId          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	DynamicBlockId     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	IsQuestionFromBank bool
	Attachments        []Attachment
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

func (q *Question) SetVariantId(id uuid.UUID) {
	q.VariantId = id
}

func (q *Question) SetDynamicBlockId(id uuid.UUID) {
	q.DynamicBlockId = id
}

func (q *Question) SetIsQuestionFromBank(b bool) {
	q.IsQuestionFromBank = b
}

type Attachment struct {
	model.BaseModel
	Description string
	QuestionId  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
