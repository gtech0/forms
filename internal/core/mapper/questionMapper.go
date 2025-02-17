package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	question "hedgehog-forms/internal/core/model/form/pattern/question"
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

func (q *QuestionMapper) ToDto(questionEntity *question.Question) (get.IQuestionDto, error) {
	switch questionEntity.Type {
	case question.SINGLE_CHOICE:
		return q.singleChoiceMapper.toDto(questionEntity)
	case question.TEXT_INPUT:
		return q.textInputMapper.toDto(questionEntity)
	case question.MULTIPLE_CHOICE:
		return q.multipleChoiceMapper.toDto(questionEntity)
	case question.MATCHING:
		return q.matchingMapper.toDto(questionEntity)
	default:
		return nil, errs.New("invalid question type", 400)
	}
}
