package block

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"slices"
)

type DynamicBlock struct {
	model.Base
	QuestionCount  int
	Questions      []question.IQuestion `gorm:"-"`
	MultipleChoice question.MultipleChoiceSlice
	TextInput      question.TextInputSlice
	SingleChoice   question.SingleChoiceSlice
	Matching       question.MatchingSlice
	BlockId        uuid.UUID `gorm:"type:uuid"`
}

func (d *DynamicBlock) BeforeSave(*gorm.DB) error {
	for _, iQuestion := range d.Questions {
		switch iQuestion.GetType() {
		case question.MULTIPLE_CHOICE:
			d.MultipleChoice = append(d.MultipleChoice, iQuestion.(*question.MultipleChoice))
		case question.SINGLE_CHOICE:
			d.SingleChoice = append(d.SingleChoice, iQuestion.(*question.SingleChoice))
		case question.MATCHING:
			d.Matching = append(d.Matching, iQuestion.(*question.Matching))
		case question.TEXT_INPUT:
			d.TextInput = append(d.TextInput, iQuestion.(*question.TextInput))
		}
	}
	return nil
}

func (d *DynamicBlock) AfterFind(*gorm.DB) error {
	d.Questions = slices.Concat(
		d.MultipleChoice.ToInterface(),
		d.TextInput.ToInterface(),
		d.SingleChoice.ToInterface(),
		d.Matching.ToInterface(),
	)
	return nil
}
