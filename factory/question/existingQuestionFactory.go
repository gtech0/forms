package question

import (
	"fmt"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form"
	question2 "hedgehog-forms/model/form/pattern/section/block/question"
)

type ExistingQuestionFactory struct {
}

func NewExistingQuestionFactory() *ExistingQuestionFactory {
	return &ExistingQuestionFactory{}
}

func (e *ExistingQuestionFactory) BuildFromDto(existingDto *create.QuestionOnExistingDto) (question2.IQuestion, error) {
	var questionObj question2.IQuestion
	if err := database.DB.Model(&question2.Question{}).
		Where("id = ?", existingDto.QuestionId).
		First(&questionObj).Error; err != nil {
		return nil, err
	}

	switch questionObj.GetType() {
	case form.MATCHING:
		return NewMatchingFactory().BuildFromObj(questionObj.(*question2.Matching)), nil
	case form.TEXT_INPUT:
		return NewTextInputFactory().BuildFromObj(questionObj.(*question2.TextInput)), nil
	case form.SINGLE_CHOICE:
		return NewSingleChoiceFactory().BuildFromObj(questionObj.(*question2.SingleChoice)), nil
	case form.MULTIPLE_CHOICE:
		return NewMultipleChoiceFactory().BuildFromObj(questionObj.(*question2.MultipleChoice)), nil
	default:
		return nil, fmt.Errorf("unknown question type")
	}
}
