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
	model.Base
	Answer      string
	TextInputId uuid.UUID `gorm:"type:uuid"`
}

type TextInputSlice []*TextInput

func (t *TextInputSlice) ToInterface() []IQuestion {
	questions := make([]IQuestion, 0)
	for _, textInput := range *t {
		questions = append(questions, textInput)
	}
	return questions
}
