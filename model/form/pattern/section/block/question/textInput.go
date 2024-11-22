package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type TextInput struct {
	Base
	Points          int
	IsCaseSensitive bool
	Answers         []TextInputAnswer
}

type TextInputAnswer struct {
	model.Base
	Answer      string
	TextInputId uuid.UUID `gorm:"type:uuid"`
}
