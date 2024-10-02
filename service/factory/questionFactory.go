package factory

import (
	"errors"
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block/question"
)

type QuestionFactory struct {
}

func (q *QuestionFactory) buildQuestion(questionDto dto.CreateQuestion) (any, error) {
	switch questionDto.GetDtoType() {
	case question.EXISTING:
		return NewExistingQuestionFactory().BuildFromDto(questionDto.(dto.CreateQuestionOnExistingDto))
	case question.MATCHING:
		return NewMatchingFactory().BuildFromDto(questionDto.(dto.CreateMatchingQuestionDto))
	case question.TEXT_INPUT:
		return NewTextInputFactory().BuildFromDto(questionDto.(dto.CreateTextQuestionDto))
	case question.SINGLE_CHOICE:
		return NewSingleChoiceFactory().BuildFromDto(questionDto.(dto.CreateSingleChoiceQuestionDto))
	case question.MULTIPLE_CHOICE:
		return NewMultipleChoiceFactory().BuildFromDto(questionDto.(dto.CreateMultipleChoiceQuestionDto))
	default:
		return nil, errors.New("unknown question type")
	}
}
