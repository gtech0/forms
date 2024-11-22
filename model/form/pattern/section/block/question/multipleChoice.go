package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type MultipleChoice struct {
	Base
	Options []MultipleChoiceOption
	Points  []MultipleChoicePoints
}

type MultipleChoicePoints struct {
	model.Base
	CorrectAnswers   int
	Points           int
	MultipleChoiceId uuid.UUID `gorm:"type:uuid"`
}

type MultipleChoiceOption struct {
	model.Base
	Text             string
	Order            int
	IsAnswer         bool
	MultipleChoiceId uuid.UUID `gorm:"type:uuid"`
}
