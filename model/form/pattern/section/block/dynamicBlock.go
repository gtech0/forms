package block

import (
	"gorm.io/gorm"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"slices"
)

type DynamicBlock struct {
	Block
	Questions      []question.IQuestion `gorm:"-"`
	MultipleChoice question.MultipleChoiceSlice
	TextInput      question.TextInputSlice
	SingleChoice   question.SingleChoiceSlice
	Matching       question.MatchingSlice
}

func (d *DynamicBlock) BeforeSave(*gorm.DB) error {
	for _, iQuestion := range d.Questions {
		switch iQuestion.GetType() {
		case form.MULTIPLE_CHOICE:
			d.MultipleChoice = append(d.MultipleChoice, iQuestion.(*question.MultipleChoice))
		case form.SINGLE_CHOICE:
			d.SingleChoice = append(d.SingleChoice, iQuestion.(*question.SingleChoice))
		case form.MATCHING:
			d.Matching = append(d.Matching, iQuestion.(*question.Matching))
		case form.TEXT_INPUT:
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

type DynamicBlockSlice []*DynamicBlock

func (d *DynamicBlockSlice) ToInterface() []IBlock {
	blocks := make([]IBlock, 0)
	for _, dynamicBlock := range *d {
		blocks = append(blocks, dynamicBlock)
	}
	return blocks
}
