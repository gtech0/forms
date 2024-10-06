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
	model.BaseModel
	Title          string
	Order          int
	IsAnswer       bool
	SingleChoiceId uuid.UUID `gorm:"type:uuid"`
}
