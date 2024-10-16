package question

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type TextInputFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewTextInputFactory() *TextInputFactory {
	return &TextInputFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (t *TextInputFactory) BuildFromDto(questionDto *create.TextQuestionDto) (*question.TextInput, error) {
	questionObj := new(question.TextInput)
	t.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, &questionObj.Question)

	questionObj.Points = questionDto.Points
	questionObj.IsCaseSensitive = questionDto.IsCaseSensitive

	answers := make([]question.TextInputAnswer, 0)
	for _, answer := range questionDto.Answers {
		var questionObjAnswer question.TextInputAnswer
		questionObjAnswer.Answer = answer
		answers = append(answers, questionObjAnswer)
	}
	questionObj.Answers = answers
	return questionObj, nil
}

func (t *TextInputFactory) BuildFromObj(questionObj *question.TextInput) *question.TextInput {
	newQuestionObj := new(question.TextInput)
	t.commonMapper.MapCommonFieldsObj(questionObj.Question, &newQuestionObj.Question)

	newQuestionObj.Points = questionObj.Points
	newQuestionObj.IsCaseSensitive = questionObj.IsCaseSensitive
	newQuestionObj.Answers = questionObj.Answers
	return newQuestionObj
}
