package factory

import (
	"fmt"
	"hedgehog-forms/database"
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block/question"
)

type ExistingQuestionFactory struct {
}

func NewExistingQuestionFactory() *ExistingQuestionFactory {
	return &ExistingQuestionFactory{}
}

func (e *ExistingQuestionFactory) BuildFromDto(existingDto dto.CreateQuestionOnExistingDto) (any, error) {
	var questionObj any
	if err := database.DB.Model(&question.Question{}).
		Where("id = ?", existingDto.QuestionId).
		First(&questionObj).Error; err != nil {
		return question.Question{}, err
	}

	switch questionObj.(question.Question).Type {
	case question.MATCHING:
		return NewMatchingFactory().buildFromObj(questionObj.(question.Matching)), nil
	case question.TEXT_INPUT:
		return NewTextInputFactory().BuildFromDto(questionObj.(dto.CreateTextQuestionDto)), nil
	case question.SINGLE_CHOICE:
		return NewSingleChoiceFactory().BuildFromDto(questionObj.(dto.CreateSingleChoiceQuestionDto)), nil
	case question.MULTIPLE_CHOICE:
		return NewMultipleChoiceFactory().BuildFromDto(questionObj.(dto.CreateMultipleChoiceQuestionDto)), nil
	default:
		return question.Question{}, fmt.Errorf("unknown question type")
	}
}