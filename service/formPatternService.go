package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"slices"
)

type FormPatternService struct{}

func NewFormPatternService() *FormPatternService {
	return &FormPatternService{}
}

func (f *FormPatternService) extractQuestionsFromPattern(pattern pattern.FormPattern) []question.IQuestion {
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

func (f *FormPatternService) GetForm(id uuid.UUID) (pattern.FormPattern, error) {
	var formPattern pattern.FormPattern
	if err := database.DB.Model(&pattern.FormPattern{}).
		Where("id = ?", id).
		First(&formPattern).
		Error; err != nil {
		return pattern.FormPattern{}, err
	}
	return formPattern, nil
}
