package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type TextInput struct {
	model.Base
	Points          int
	IsCaseSensitive bool
	Answers         []TextInputAnswer
	QuestionId      uuid.UUID `gorm:"type:uuid"`
}

type TextInputAnswer struct {
	model.Base
	Answer      string
	TextInputId uuid.UUID `gorm:"type:uuid"`
}
