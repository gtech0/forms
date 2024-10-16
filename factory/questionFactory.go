package factory

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type QuestionFactory struct {
	existingQuestionFactory *ExistingQuestionFactory
	matchingFactory         *MatchingFactory
	textInputFactory        *TextInputFactory
	singleChoiceFactory     *SingleChoiceFactory
	multipleChoiceFactory   *MultipleChoiceFactory
}

func NewQuestionFactory() *QuestionFactory {
	return &QuestionFactory{
		existingQuestionFactory: NewExistingQuestionFactory(),
		matchingFactory:         NewMatchingFactory(),
		textInputFactory:        NewTextInputFactory(),
		singleChoiceFactory:     NewSingleChoiceFactory(),
		multipleChoiceFactory:   NewMultipleChoiceFactory(),
	}
}

func (q *QuestionFactory) buildQuestionFromDto(questionDto any) (question.IQuestion, error) {
	switch questionTyped := questionDto.(type) {
	case *create.QuestionOnExistingDto:
		return q.existingQuestionFactory.BuildFromDto(questionTyped)
	case *create.MatchingQuestionDto:
		return q.matchingFactory.BuildFromDto(questionTyped)
	case *create.TextQuestionDto:
		return q.textInputFactory.BuildFromDto(questionTyped)
	case *create.SingleChoiceQuestionDto:
		return q.singleChoiceFactory.BuildFromDto(questionTyped)
	case *create.MultipleChoiceQuestionDto:
		return q.multipleChoiceFactory.BuildFromDto(questionTyped)
	default:
		return nil, errs.New("invalid question type", 400)
	}
}

func (q *QuestionFactory) BuildQuestionDtoForDynamicBlock(
	questionDtos []any,
	dynamicBlock *block.DynamicBlock,
) ([]question.IQuestion, error) {
	questionObjs := make([]question.IQuestion, 0)
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
	newQuestionObjs := make([]question.IQuestion, 0)
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
	questionObjs := make([]question.IQuestion, 0)
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
	newQuestionObjs := make([]question.IQuestion, 0)
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
	questionDto := new(create.QuestionOnExistingDto)
	questionDto.QuestionId = questionObj.GetId()
	result, err := q.buildQuestionFromDto(questionDto)
	if err != nil {
		return nil, err
	}

	return result, nil
}
