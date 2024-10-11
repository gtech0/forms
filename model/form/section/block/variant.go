package block

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/section/block/question"
	"slices"
)

type Variant struct {
	model.BaseModel
	Title          string
	Description    string
	StaticBlockID  uuid.UUID            `gorm:"type:uuid"`
	Questions      []question.IQuestion `gorm:"-"`
	MultipleChoice question.MultipleChoiceSlice
	TextInput      question.TextInputSlice
	SingleChoice   question.SingleChoiceSlice
	Matching       question.MatchingSlice
}

func (v *Variant) BeforeSave(*gorm.DB) error {
	for _, iQuestion := range v.Questions {
		switch iQuestion.GetType() {
		case question.MULTIPLE_CHOICE:
			v.MultipleChoice = append(v.MultipleChoice, iQuestion.(*question.MultipleChoice))
		case question.SINGLE_CHOICE:
			v.SingleChoice = append(v.SingleChoice, iQuestion.(*question.SingleChoice))
		case question.MATCHING:
			v.Matching = append(v.Matching, iQuestion.(*question.Matching))
		case question.TEXT_INPUT:
			v.TextInput = append(v.TextInput, iQuestion.(*question.TextInput))
		}
	}
	return nil
}

func (v *Variant) AfterFind(*gorm.DB) error {
	v.Questions = slices.Concat(
		v.MultipleChoice.ToInterface(),
		v.TextInput.ToInterface(),
		v.SingleChoice.ToInterface(),
		v.Matching.ToInterface(),
	)
	return nil
}
