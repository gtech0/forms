package question

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	model2 "hedgehog-forms/internal/core/model"
	"time"
)

type Question struct {
	model2.Base
	Description        string
	Order              int
	OwnerId            uuid.NullUUID `gorm:"type:uuid"`
	Attachments        []Attachment
	VariantId          uuid.NullUUID `gorm:"type:uuid"`
	DynamicBlockId     uuid.NullUUID `gorm:"type:uuid"`
	Subject            model2.Subject
	SubjectId          uuid.NullUUID `gorm:"type:uuid"`
	IsQuestionFromBank bool
	Type               QuestionType

	SingleChoice   *SingleChoice
	MultipleChoice *MultipleChoice
	Matching       *Matching
	TextInput      *TextInput
}

type Base struct {
	QuestionId uuid.UUID `gorm:"type:uuid;primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
