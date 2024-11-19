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

func (q *QuestionFactory) BuildQuestionFromDto(questionDto any) (*question.Question, error) {
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
	dynamicBlock *block.Block,
) ([]*question.Question, error) {
	questionObjs := make([]*question.Question, 0)
	for _, questionDto := range questionDtos {
		questionObj, err := q.BuildQuestionFromDto(questionDto)
		if err != nil {
			return nil, err
		}

		questionObjs = append(questionObjs, questionObj)
	}

	dynamicBlock.DynamicBlock.Questions = questionObjs
	return questionObjs, nil
}

func (q *QuestionFactory) BuildQuestionObjForDynamicBlock(
	questionObjs []*question.Question,
	dynamicBlock *block.Block,
) ([]*question.Question, error) {
	newQuestionObjs := make([]*question.Question, 0)
	for _, questionObj := range questionObjs {
		newQuestionObj, err := q.buildQuestionFromObj(questionObj)
		if err != nil {
			return nil, err
		}

		newQuestionObjs = append(newQuestionObjs, newQuestionObj)
	}

	dynamicBlock.DynamicBlock.Questions = newQuestionObjs
	return questionObjs, nil
}

func (q *QuestionFactory) BuildQuestionDtoForVariant(
	questionDtos []any,
	variant *block.Variant,
) ([]*question.Question, error) {
	questionObjs := make([]*question.Question, 0)
	for order, questionDto := range questionDtos {
		questionObj, err := q.BuildQuestionFromDto(questionDto)
		if err != nil {
			return nil, err
		}

		questionObj.Order = order
		questionObjs = append(questionObjs, questionObj)
	}

	variant.Questions = questionObjs
	return questionObjs, nil
}

func (q *QuestionFactory) BuildQuestionForVariantObj(
	questionObjs []*question.Question,
	variant *block.Variant,
) ([]*question.Question, error) {
	newQuestionObjs := make([]*question.Question, 0)
	for order, questionDto := range questionObjs {
		newQuestionObj, err := q.buildQuestionFromObj(questionDto)
		if err != nil {
			return nil, err
		}

		newQuestionObj.Order = order
		newQuestionObjs = append(newQuestionObjs, newQuestionObj)
	}

	variant.Questions = questionObjs
	return newQuestionObjs, nil
}

func (q *QuestionFactory) buildQuestionFromObj(questionObj *question.Question) (*question.Question, error) {
	questionDto := new(create.QuestionOnExistingDto)
	questionDto.QuestionId = questionObj.Id
	result, err := q.BuildQuestionFromDto(questionDto)
	if err != nil {
		return nil, err
	}

	return result, nil
}
