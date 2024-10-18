package factory

import (
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type ExistingQuestionFactory struct {
	matchingFactory       *MatchingFactory
	textInputFactory      *TextInputFactory
	singleChoiceFactory   *SingleChoiceFactory
	multipleChoiceFactory *MultipleChoiceFactory
}

func NewExistingQuestionFactory() *ExistingQuestionFactory {
	return &ExistingQuestionFactory{
		matchingFactory:       NewMatchingFactory(),
		textInputFactory:      NewTextInputFactory(),
		singleChoiceFactory:   NewSingleChoiceFactory(),
		multipleChoiceFactory: NewMultipleChoiceFactory(),
	}
}

func (e *ExistingQuestionFactory) BuildFromDto(existingDto *create.QuestionOnExistingDto) (question.IQuestion, error) {
	var questionObj question.IQuestion
	if err := database.DB.Model(&question.Question{}).
		Where("id = ?", existingDto.QuestionId).
		First(&questionObj).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}

	switch questionObj.GetType() {
	case question.MATCHING:
		return e.matchingFactory.BuildFromObj(questionObj.(*question.Matching))
	case question.TEXT_INPUT:
		return e.textInputFactory.BuildFromObj(questionObj.(*question.TextInput))
	case question.SINGLE_CHOICE:
		return e.singleChoiceFactory.BuildFromObj(questionObj.(*question.SingleChoice))
	case question.MULTIPLE_CHOICE:
		return e.multipleChoiceFactory.BuildFromObj(questionObj.(*question.MultipleChoice))
	default:
		return nil, errs.New("invalid question type", 400)
	}
}
