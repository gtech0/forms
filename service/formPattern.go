package service

import (
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/model/form/section/block/question"
	"slices"
)

type PatternService struct{}

func NewPatternService() *PatternService {
	return &PatternService{}
}

func (f *PatternService) ExtractQuestionsFromPattern(pattern form.FormPattern) []question.IQuestion {
	questions := make([]question.IQuestion, 0)
	for _, currSection := range pattern.Sections {
		blocks := currSection.Blocks
		for _, iBlock := range blocks {
			switch assertedBlock := iBlock.(type) {
			case *block.DynamicBlock:
				questions = slices.Concat(questions, assertedBlock.Questions)
			case *block.StaticBlock:
				for _, variant := range assertedBlock.Variants {
					questions = slices.Concat(questions, variant.Questions)
				}
			}
		}
	}
	return questions
}
