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
}

func NewExistingQuestionFactory() *ExistingQuestionFactory {
	return &ExistingQuestionFactory{
		matchingFactory:          NewMatchingFactory(),
		textInputFactory:         NewTextInputFactory(),
		singleChoiceFactory:      NewSingleChoiceFactory(),
		multipleChoiceFactory:    NewMultipleChoiceFactory(),
		matchingRepository:       repository.NewMatchingQuestionRepository(),
		textInputRepository:      repository.NewTextInputRepository(),
		singleChoiceRepository:   repository.NewSingleChoiceRepository(),
		multipleChoiceRepository: repository.NewMultipleChoiceRepository(),
	}
}

func (e *ExistingQuestionFactory) BuildFromDto(existingDto *create.QuestionOnExistingDto) (question.IQuestion, error) {
	switch existingDto.Type {
	case question.MATCHING:
		matchingQuestion, err := e.matchingRepository.GetById(existingDto.QuestionId)
		if err != nil {
			return nil, err
		}
		return e.matchingFactory.BuildFromObj(matchingQuestion)
	case question.MULTIPLE_CHOICE:
		multipleChoiceQuestion, err := e.multipleChoiceRepository.GetById(existingDto.QuestionId)
		if err != nil {
			return nil, err
		}
		return e.multipleChoiceFactory.BuildFromObj(multipleChoiceQuestion)
	case question.SINGLE_CHOICE:
		singleChoiceQuestion, err := e.singleChoiceRepository.GetById(existingDto.QuestionId)
		if err != nil {
			return nil, err
		}
		return e.singleChoiceFactory.BuildFromObj(singleChoiceQuestion)
	case question.TEXT_INPUT:
		textInputQuestion, err := e.textInputRepository.GetById(existingDto.QuestionId)
		if err != nil {
			return nil, err
		}
		return e.textInputFactory.BuildFromObj(textInputQuestion)
	default:
		return nil, errs.New("invalid question type", 400)
	}
}
