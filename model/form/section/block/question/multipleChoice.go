package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type MultipleChoice struct {
	Question
	MultipleChoiceOptions []MultipleChoiceOption
}

type MultipleChoicePoints struct {
	model.BaseModel
	CorrectAnswers   int
	Points           int
	MultipleChoiceId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

type MultipleChoiceOption struct {
	model.BaseModel
	Title            string
	Order            int
	IsAnswer         bool
	MultipleChoiceId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
