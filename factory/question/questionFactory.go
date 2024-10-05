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

func (q *QuestionFactory) buildQuestionFromDto(questionDto any) (question.IQuestion, error) {
	switch questionTyped := questionDto.(type) {
	case *dto.CreateQuestionOnExistingDto:
		return NewExistingQuestionFactory().BuildFromDto(questionTyped)
	case *dto.CreateMatchingQuestionDto:
		return NewMatchingFactory().BuildFromDto(questionTyped)
	case *dto.CreateTextQuestionDto:
		return NewTextInputFactory().BuildFromDto(questionTyped)
	case *dto.CreateSingleChoiceQuestionDto:
		return NewSingleChoiceFactory().BuildFromDto(questionTyped)
	case *dto.CreateMultipleChoiceQuestionDto:
		return NewMultipleChoiceFactory().BuildFromDto(questionTyped)
	default:
		return nil, errors.New("unknown question type")
	}
}

func (q *QuestionFactory) BuildQuestionDtoForDynamicBlock(
	questionDtos []any,
	dynamicBlock *block.DynamicBlock,
) ([]question.IQuestion, error) {
	questionObjs := make([]question.IQuestion, len(questionDtos))
	for _, questionDto := range questionDtos {
		questionObj, err := q.buildQuestionFromDto(questionDto)
		if err != nil {
			return nil, err
		}

		questionObjs = append(questionObjs, questionObj)
	}

	dynamicBlock.Questions = questionObjs
	return questionObjs, nil
}

func (q *QuestionFactory) BuildQuestionObjForDynamicBlock(
	questionObjs []question.IQuestion,
	dynamicBlock *block.DynamicBlock,
) ([]question.IQuestion, error) {
	newQuestionObjs := make([]question.IQuestion, len(questionObjs))
	for _, questionObj := range questionObjs {
		newQuestionObj, err := q.buildQuestionFromObj(questionObj)
		if err != nil {
			return nil, err
		}

		newQuestionObjs = append(newQuestionObjs, newQuestionObj)
	}

	dynamicBlock.Questions = newQuestionObjs
	return questionObjs, nil
}

func (q *QuestionFactory) BuildQuestionDtoForVariant(
	questionDtos []any,
	variant *block.Variant,
) ([]question.IQuestion, error) {
	questionObjs := make([]question.IQuestion, len(questionDtos))
	for order, questionDto := range questionDtos {
		questionObj, err := q.buildQuestionFromDto(questionDto)
		if err != nil {
			return nil, err
		}

		questionObj.SetOrder(order)
		questionObjs = append(questionObjs, questionObj)
	}

	variant.Questions = questionObjs
	return questionObjs, nil
}

func (q *QuestionFactory) BuildQuestionForVariantObj(
	questionObjs []question.IQuestion,
	variant *block.Variant,
) ([]question.IQuestion, error) {
	newQuestionObjs := make([]question.IQuestion, len(questionObjs))
	for order, questionDto := range questionObjs {
		newQuestionObj, err := q.buildQuestionFromObj(questionDto)
		if err != nil {
			return nil, err
		}

		newQuestionObj.SetOrder(order)
		newQuestionObjs = append(newQuestionObjs, newQuestionObj)
	}

	variant.Questions = questionObjs
	return newQuestionObjs, nil
}

func (q *QuestionFactory) buildQuestionFromObj(questionObj question.IQuestion) (question.IQuestion, error) {
	var questionDto *dto.CreateQuestionOnExistingDto
	questionDto.QuestionId = questionObj.GetId()
	result, err := q.buildQuestionFromDto(questionDto)
	if err != nil {
		return nil, err
	}

	return result, nil
}
