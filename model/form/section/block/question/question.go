package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

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

type Attachment struct {
	model.BaseModel
	Description string
	QuestionId  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
