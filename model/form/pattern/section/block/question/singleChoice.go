package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type SingleChoice struct {
	Question
	Points              int
	SingleChoiceOptions []SingleChoiceOption
}

type SingleChoiceOption struct {
	model.Base
	Text           string
	Order          int
	IsAnswer       bool
	SingleChoiceId uuid.UUID `gorm:"type:uuid"`
}

type SingleChoiceSlice []*SingleChoice

func (s *SingleChoiceSlice) ToInterface() []IQuestion {
	questions := make([]IQuestion, 0)
	for _, singleChoice := range *s {
		questions = append(questions, singleChoice)
	}
	return questions
}
