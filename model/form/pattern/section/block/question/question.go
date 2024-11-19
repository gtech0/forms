package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type Question struct {
	model.Base
	Description        string
	Order              int
	OwnerId            uuid.NullUUID `gorm:"type:uuid"`
	Attachments        []Attachment
	VariantId          uuid.NullUUID `gorm:"type:uuid"`
	DynamicBlockId     uuid.NullUUID `gorm:"type:uuid"`
	Subject            model.Subject
	SubjectId          uuid.NullUUID `gorm:"type:uuid"`
	IsQuestionFromBank bool
	Type               QuestionType

	SingleChoice   *SingleChoice
	MultipleChoice *MultipleChoice
	Matching       *Matching
	TextInput      *TextInput
}
