package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type SingleChoice struct {
	model.Base
	Points     int
	Options    []SingleChoiceOption
	QuestionId uuid.UUID `gorm:"type:uuid"`
}

type SingleChoiceOption struct {
	model.Base
	Text           string
	Order          int
	IsAnswer       bool
	SingleChoiceId uuid.UUID `gorm:"type:uuid"`
}
