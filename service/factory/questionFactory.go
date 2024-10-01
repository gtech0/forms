package factory

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block/question"
)

type QuestionFactory struct {
}

func (q *QuestionFactory) buildQuestion(questionDto dto.CreateQuestion) (question.Question, error) {
	switch questionDto.GetType() {
	case question.EXISTING:
	case question.MATCHING:
		NewMatchingFactory().buildFromDto(questionDto.(dto.CreateMatchingQuestionDto))
	case question.TEXT_INPUT:
	case question.SINGLE_CHOICE:
	case question.MULTIPLE_CHOICE:
	default:

	}
}
