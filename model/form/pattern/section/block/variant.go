package block

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form"
	question2 "hedgehog-forms/model/form/pattern/section/block/question"
	"slices"
)

type Variant struct {
	model.Base
	Title          string
	Description    string
	StaticBlockID  uuid.UUID             `gorm:"type:uuid"`
	Questions      []question2.IQuestion `gorm:"-"`
	MultipleChoice question2.MultipleChoiceSlice
	TextInput      question2.TextInputSlice
	SingleChoice   question2.SingleChoiceSlice
	Matching       question2.MatchingSlice
}

func (v *Variant) BeforeSave(*gorm.DB) error {
	for _, iQuestion := range v.Questions {
		switch iQuestion.GetType() {
		case form.MULTIPLE_CHOICE:
			v.MultipleChoice = append(v.MultipleChoice, iQuestion.(*question2.MultipleChoice))
		case form.SINGLE_CHOICE:
			v.SingleChoice = append(v.SingleChoice, iQuestion.(*question2.SingleChoice))
		case form.MATCHING:
			v.Matching = append(v.Matching, iQuestion.(*question2.Matching))
		case form.TEXT_INPUT:
			v.TextInput = append(v.TextInput, iQuestion.(*question2.TextInput))
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
