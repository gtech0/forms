package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type SingleChoice struct {
	Base
	Points  int
	Options []SingleChoiceOption
}

type SingleChoiceOption struct {
	model.Base
	Text           string
	Order          int
	IsAnswer       bool
	SingleChoiceId uuid.UUID `gorm:"type:uuid"`
}
