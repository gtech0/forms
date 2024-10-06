package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type MultipleChoice struct {
	Question
	Options []MultipleChoiceOption
	Points  []MultipleChoicePoints
}

type MultipleChoicePoints struct {
	model.BaseModel
	CorrectAnswer    int
	Points           int
	MultipleChoiceId uuid.UUID `gorm:"type:uuid"`
}

type MultipleChoiceOption struct {
	model.BaseModel
	Title            string
	Order            int
	IsAnswer         bool
	MultipleChoiceId uuid.UUID `gorm:"type:uuid"`
}
