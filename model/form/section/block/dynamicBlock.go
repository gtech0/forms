package block

import (
	"gorm.io/gorm"
	"hedgehog-forms/model/form/section/block/question"
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

type DynamicBlockSlice []*DynamicBlock

func (d *DynamicBlockSlice) ToInterface() []IBlock {
	blocks := make([]IBlock, 0)
	for _, dynamicBlock := range *d {
		blocks = append(blocks, dynamicBlock)
	}
	return blocks
}
