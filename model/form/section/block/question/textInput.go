package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type TextInput struct {
	Question
	Points          int
	IsCaseSensitive bool
	Answers         []TextInputAnswer
}

type TextInputAnswer struct {
	model.BaseModel
	Answer      string
	TextInputId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
