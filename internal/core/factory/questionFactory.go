package factory

import (
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/errs"
	block "hedgehog-forms/internal/core/model/form/pattern/block"
	"hedgehog-forms/internal/core/model/form/pattern/question"
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
	questions := make([]*question.Question, 0)
	for _, questionDto := range questionDtos {
		questionEntity, err := q.BuildQuestionFromDto(questionDto)
		if err != nil {
			return nil, err
		}

		questions = append(questions, questionEntity)
	}

	dynamicBlock.DynamicBlock.Questions = questions
	return questions, nil
}

func (q *QuestionFactory) BuildQuestionEntityForDynamicBlock(
	questionEntities []*question.Question,
	dynamicBlock *block.Block,
) ([]*question.Question, error) {
	newQuestionEntities := make([]*question.Question, 0)
	for _, questionEntity := range questionEntities {
		newQuestion, err := q.buildQuestionFromEntity(questionEntity)
		if err != nil {
			return nil, err
		}

		newQuestionEntities = append(newQuestionEntities, newQuestion)
	}

	dynamicBlock.DynamicBlock.Questions = newQuestionEntities
	return questionEntities, nil
}

func (q *QuestionFactory) BuildQuestionDtoForVariant(
	questionDtos []any,
	variant *block.Variant,
) ([]*question.Question, error) {
	questionEntities := make([]*question.Question, 0)
	for order, questionDto := range questionDtos {
		questionEntity, err := q.BuildQuestionFromDto(questionDto)
		if err != nil {
			return nil, err
		}

		questionEntity.Order = order
		questionEntities = append(questionEntities, questionEntity)
	}

	variant.Questions = questionEntities
	return questionEntities, nil
}

func (q *QuestionFactory) BuildQuestionForVariantEntities(
	questionEntities []*question.Question,
	variant *block.Variant,
) ([]*question.Question, error) {
	newQuestionEntities := make([]*question.Question, 0)
	for order, questionDto := range questionEntities {
		newQuestionEntity, err := q.buildQuestionFromEntity(questionDto)
		if err != nil {
			return nil, err
		}

		newQuestionEntity.Order = order
		newQuestionEntities = append(newQuestionEntities, newQuestionEntity)
	}

	variant.Questions = questionEntities
	return newQuestionEntities, nil
}

func (q *QuestionFactory) buildQuestionFromEntity(questionEntity *question.Question) (*question.Question, error) {
	questionDto := new(create.QuestionOnExistingDto)
	questionDto.QuestionId = questionEntity.Id
	result, err := q.BuildQuestionFromDto(questionDto)
	if err != nil {
		return nil, err
	}

	return result, nil
}
