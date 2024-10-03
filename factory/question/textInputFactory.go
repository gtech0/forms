package question

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block/question"
)

type TextInputFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewTextInputFactory() *TextInputFactory {
	return &TextInputFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (t *TextInputFactory) BuildFromDto(questionDto *dto.CreateTextQuestionDto) (question.TextInput, error) {
	var questionObj question.TextInput
	t.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, questionObj.Question)

	questionObj.Points = questionDto.Points
	questionObj.IsCaseSensitive = questionDto.IsCaseSensitive
	for _, answer := range questionDto.Answers {
		var questionObjAnswer question.TextInputAnswer
		questionObjAnswer.Answer = answer
		questionObj.Answer = append(questionObj.Answer, questionObjAnswer)
	}
	return questionObj, nil
}

func (t *TextInputFactory) BuildFromObj(questionObj question.TextInput) question.TextInput {
	var newQuestionObj question.TextInput
	t.commonMapper.MapCommonFieldsObj(questionObj.Question, newQuestionObj.Question)

	newQuestionObj.Points = questionObj.Points
	newQuestionObj.IsCaseSensitive = questionObj.IsCaseSensitive
	newQuestionObj.Answer = questionObj.Answer
	return newQuestionObj
}
