package question

import (
	"errors"
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/model/form/section/block/question"
)

type QuestionFactory struct {
}

func (q *QuestionFactory) buildQuestionFromDto(questionDto dto.CreateQuestion) (any, error) {
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

func (q *QuestionFactory) buildQuestionForDynamicBlockDto(
	questionDtos []dto.CreateQuestionDto,
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

func (q *QuestionFactory) buildQuestionForDynamicBlockObj(
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

func (q *QuestionFactory) buildQuestionForVariantDto(
	questionDtos []dto.CreateQuestionDto,
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

func (q *QuestionFactory) buildQuestionForVariantObj(
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
