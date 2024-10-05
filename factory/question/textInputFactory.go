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

func (t *TextInputFactory) BuildFromDto(questionDto *dto.CreateTextQuestionDto) (*question.TextInput, error) {
	var questionObj *question.TextInput
	t.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, questionObj)

	questionObj.Points = questionDto.Points
	questionObj.IsCaseSensitive = questionDto.IsCaseSensitive

	answers := make([]question.TextInputAnswer, len(questionDto.Answers))
	for _, answer := range questionDto.Answers {
		var questionObjAnswer question.TextInputAnswer
		questionObjAnswer.Answer = answer
		answers = append(answers, questionObjAnswer)
	}
	questionObj.Answers = answers
	return questionObj, nil
}

func (t *TextInputFactory) BuildFromObj(questionObj *question.TextInput) *question.TextInput {
	var newQuestionObj *question.TextInput
	t.commonMapper.MapCommonFieldsObj(questionObj.Question, newQuestionObj)

	newQuestionObj.Points = questionObj.Points
	newQuestionObj.IsCaseSensitive = questionObj.IsCaseSensitive
	newQuestionObj.Answers = questionObj.Answers
	return newQuestionObj
}
