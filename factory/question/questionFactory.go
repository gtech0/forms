package question

import (
	"errors"
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/model/form/section/block/question"
)

type QuestionFactory struct {
}

func NewQuestionFactory() *QuestionFactory {
	return &QuestionFactory{}
}

func (q *QuestionFactory) buildQuestionFromDto(questionDto any) (any, error) {
	switch quest := questionDto.(type) {
	case *dto.CreateQuestionOnExistingDto:
		return NewExistingQuestionFactory().BuildFromDto(quest)
	case *dto.CreateMatchingQuestionDto:
		return NewMatchingFactory().BuildFromDto(quest)
	case *dto.CreateTextQuestionDto:
		return NewTextInputFactory().BuildFromDto(quest)
	case *dto.CreateSingleChoiceQuestionDto:
		return NewSingleChoiceFactory().BuildFromDto(quest)
	case *dto.CreateMultipleChoiceQuestionDto:
		return NewMultipleChoiceFactory().BuildFromDto(quest)
	default:
		return nil, errors.New("unknown question type")
	}
}

func (q *QuestionFactory) BuildQuestionForDynamicBlockDto(
	questionDtos []any,
	dynamicBlock block.DynamicBlock,
) ([]question.Question, error) {
	questionObjs := make([]question.Question, len(questionDtos))
	for _, questionDto := range questionDtos {
		questionObj, err := q.buildQuestionFromDto(questionDto)
		if err != nil {
			return nil, err
		}

		questionObj.(*question.Question).DynamicBlockId = dynamicBlock.Id
		questionObjs = append(questionObjs, questionObj.(question.Question))
	}
	return questionObjs, nil
}

func (q *QuestionFactory) BuildQuestionForDynamicBlockObj(
	questionObjs []question.Question,
	dynamicBlock block.DynamicBlock,
) ([]question.Question, error) {
	newQuestionObjs := make([]question.Question, len(questionObjs))
	for _, questionDto := range questionObjs {
		newQuestionObj, err := q.buildQuestionFromObj(questionDto)
		if err != nil {
			return nil, err
		}

		newQuestionObj.DynamicBlockId = dynamicBlock.Id
		newQuestionObjs = append(newQuestionObjs, newQuestionObj)
	}

	return questionObjs, nil
}

func (q *QuestionFactory) BuildQuestionForVariantDto(
	questionDtos []any,
	variant block.Variant,
) ([]question.Question, error) {
	questionObjs := make([]question.Question, len(questionDtos))
	for order, questionDto := range questionDtos {
		questionObj, err := q.buildQuestionFromDto(questionDto)
		if err != nil {
			return nil, err
		}

		questionObj.(*question.Question).VariantId = variant.Id
		questionObj.(*question.Question).Order = order
		questionObjs = append(questionObjs, questionObj.(question.Question))
	}
	return questionObjs, nil
}

func (q *QuestionFactory) BuildQuestionForVariantObj(
	questionObjs []question.Question,
	variant block.Variant,
) ([]question.Question, error) {
	newQuestionObjs := make([]question.Question, len(questionObjs))
	for order, questionDto := range questionObjs {
		newQuestionObj, err := q.buildQuestionFromObj(questionDto)
		if err != nil {
			return nil, err
		}

		newQuestionObj.VariantId = variant.Id
		newQuestionObj.Order = order
		newQuestionObjs = append(newQuestionObjs, newQuestionObj)
	}

	return questionObjs, nil
}

func (q *QuestionFactory) buildQuestionFromObj(questionObj question.Question) (question.Question, error) {
	var questionDto dto.CreateQuestionOnExistingDto
	questionDto.QuestionId = questionObj.Id
	result, err := q.buildQuestionFromDto(questionDto)
	if err != nil {
		return question.Question{}, err
	}

	return result.(question.Question), nil
}
