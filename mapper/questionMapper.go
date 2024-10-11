package mapper

import (
	"errors"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/section/block/question"
)

type QuestionMapper struct {
	singleChoiceMapper   *SingleChoiceMapper
	textInputMapper      *TextInputMapper
	multipleChoiceMapper *MultipleChoiceMapper
	matchingMapper       *MatchingMapper
}

func NewQuestionMapper() *QuestionMapper {
	return &QuestionMapper{
		singleChoiceMapper:   NewSingleChoiceMapper(),
		textInputMapper:      NewTextInputMapper(),
		multipleChoiceMapper: NewMultipleChoiceMapper(),
		matchingMapper:       NewMatchingMapper(),
	}
}

func (q *QuestionMapper) toDto(questionObj question.IQuestion) (get.IQuestionDto, error) {
	switch assertedQuestion := questionObj.(type) {
	case *question.SingleChoice:
		return q.singleChoiceMapper.toDto(assertedQuestion)
	case *question.TextInput:
		return q.textInputMapper.toDto(assertedQuestion)
	case *question.MultipleChoice:
		return q.multipleChoiceMapper.toDto(assertedQuestion)
	case *question.Matching:
		return q.matchingMapper.toDto(assertedQuestion)
	default:
		return nil, errors.New("invalid question type")
	}
}
