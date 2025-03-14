package factory

import (
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/question"
	"hedgehog-forms/internal/core/repository"
)

type ExistingQuestionFactory struct {
	matchingFactory       *MatchingFactory
	textInputFactory      *TextInputFactory
	singleChoiceFactory   *SingleChoiceFactory
	multipleChoiceFactory *MultipleChoiceFactory
	questionRepository    *repository.QuestionRepository
}

func NewExistingQuestionFactory() *ExistingQuestionFactory {
	return &ExistingQuestionFactory{
		matchingFactory:       NewMatchingFactory(),
		textInputFactory:      NewTextInputFactory(),
		singleChoiceFactory:   NewSingleChoiceFactory(),
		multipleChoiceFactory: NewMultipleChoiceFactory(),
		questionRepository:    repository.NewQuestionRepository(),
	}
}

func (e *ExistingQuestionFactory) BuildFromDto(existingDto *create.QuestionOnExistingDto) (*question.Question, error) {
	questionEntity, err := e.questionRepository.FindById(existingDto.QuestionId)
	if err != nil {
		return nil, err
	}

	switch existingDto.Type {
	case question.MATCHING:
		return e.matchingFactory.BuildFromEntity(questionEntity)
	case question.MULTIPLE_CHOICE:
		return e.multipleChoiceFactory.BuildFromEntity(questionEntity)
	case question.SINGLE_CHOICE:
		return e.singleChoiceFactory.BuildFromEntity(questionEntity)
	case question.TEXT_INPUT:
		return e.textInputFactory.BuildFromEntity(questionEntity)
	default:
		return nil, errs.New("invalid question type", 400)
	}
}
