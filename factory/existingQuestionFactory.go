package factory

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"hedgehog-forms/repository"
)

type ExistingQuestionFactory struct {
	matchingFactory       *MatchingFactory
	textInputFactory      *TextInputFactory
	singleChoiceFactory   *SingleChoiceFactory
	multipleChoiceFactory *MultipleChoiceFactory

	matchingRepository       *repository.MatchingQuestionRepository
	textInputRepository      *repository.TextInputRepository
	singleChoiceRepository   *repository.SingleChoiceRepository
	multipleChoiceRepository *repository.MultipleChoiceRepository
	questionRepository       *repository.QuestionRepository
}

func NewExistingQuestionFactory() *ExistingQuestionFactory {
	return &ExistingQuestionFactory{
		matchingFactory:       NewMatchingFactory(),
		textInputFactory:      NewTextInputFactory(),
		singleChoiceFactory:   NewSingleChoiceFactory(),
		multipleChoiceFactory: NewMultipleChoiceFactory(),

		matchingRepository:       repository.NewMatchingQuestionRepository(),
		textInputRepository:      repository.NewTextInputRepository(),
		singleChoiceRepository:   repository.NewSingleChoiceRepository(),
		multipleChoiceRepository: repository.NewMultipleChoiceRepository(),
		questionRepository:       repository.NewQuestionRepository(),
	}
}

func (e *ExistingQuestionFactory) BuildFromDto(existingDto *create.QuestionOnExistingDto) (*question.Question, error) {
	questionEntity, err := e.questionRepository.FindById(existingDto.QuestionId)
	if err != nil {
		return nil, err
	}

	switch existingDto.Type {
	case question.MATCHING:
		return e.matchingFactory.BuildFromObj(questionEntity)
	case question.MULTIPLE_CHOICE:
		return e.multipleChoiceFactory.BuildFromObj(questionEntity)
	case question.SINGLE_CHOICE:
		return e.singleChoiceFactory.BuildFromObj(questionEntity)
	case question.TEXT_INPUT:
		return e.textInputFactory.BuildFromObj(questionEntity)
	default:
		return nil, errs.New("invalid question type", 400)
	}
}
