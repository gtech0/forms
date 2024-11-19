package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
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

func (q *QuestionMapper) toDto(questionObj *question.Question) (get.IQuestionDto, error) {
	switch questionObj.Type {
	case question.SINGLE_CHOICE:
		return q.singleChoiceMapper.toDto(questionObj)
	case question.TEXT_INPUT:
		return q.textInputMapper.toDto(questionObj)
	case question.MULTIPLE_CHOICE:
		return q.multipleChoiceMapper.toDto(questionObj)
	case question.MATCHING:
		return q.matchingMapper.toDto(questionObj)
	default:
		return nil, errs.New("invalid question type", 400)
	}
}
