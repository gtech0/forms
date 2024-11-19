package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type MultipleChoice struct {
	model.Base
	Options    []MultipleChoiceOption
	Points     []MultipleChoicePoints
	QuestionId uuid.UUID `gorm:"type:uuid"`
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
