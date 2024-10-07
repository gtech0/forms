package question

import (
	"fmt"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/section/block/question"
)

type ExistingQuestionFactory struct {
}

func NewExistingQuestionFactory() *ExistingQuestionFactory {
	return &ExistingQuestionFactory{}
}

func (e *ExistingQuestionFactory) BuildFromDto(existingDto *create.QuestionOnExistingDto) (question.IQuestion, error) {
	var questionObj question.IQuestion
	if err := database.DB.Model(&question.Question{}).
		Where("id = ?", existingDto.QuestionId).
		First(&questionObj).Error; err != nil {
		return nil, err
	}

	switch questionObj.GetType() {
	case question.MATCHING:
		return NewMatchingFactory().BuildFromObj(questionObj.(*question.Matching)), nil
	case question.TEXT_INPUT:
		return NewTextInputFactory().BuildFromObj(questionObj.(*question.TextInput)), nil
	case question.SINGLE_CHOICE:
		return NewSingleChoiceFactory().BuildFromObj(questionObj.(*question.SingleChoice)), nil
	case question.MULTIPLE_CHOICE:
		return NewMultipleChoiceFactory().BuildFromObj(questionObj.(*question.MultipleChoice)), nil
	default:
		return nil, fmt.Errorf("unknown question type")
	}
}
