package mapper

import (
	"errors"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/section/block/question"
)

type QuestionMapper struct {
	singleChoiceMapper *SingleChoiceMapper
}

func NewQuestionMapper() *QuestionMapper {
	return &QuestionMapper{
		singleChoiceMapper: NewSingleChoiceMapper(),
	}
}

func (q *QuestionMapper) toDto(questionObj question.IQuestion) (get.IQuestionDto, error) {
	switch assertedQuestion := questionObj.(type) {
	case *question.SingleChoice:
		return q.singleChoiceMapper.singleChoiceToDto(assertedQuestion)
	default:
		return nil, errors.New("invalid question type")
	}
}
