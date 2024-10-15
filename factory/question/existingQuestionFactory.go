package question

import (
	"fmt"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/pattern/section/block/question"
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
	case form.MATCHING:
		return NewMatchingFactory().BuildFromObj(questionObj.(*question.Matching)), nil
	case form.TEXT_INPUT:
		return NewTextInputFactory().BuildFromObj(questionObj.(*question.TextInput)), nil
	case form.SINGLE_CHOICE:
		return NewSingleChoiceFactory().BuildFromObj(questionObj.(*question.SingleChoice)), nil
	case form.MULTIPLE_CHOICE:
		return NewMultipleChoiceFactory().BuildFromObj(questionObj.(*question.MultipleChoice)), nil
	default:
		return nil, fmt.Errorf("unknown question type")
	}
}
