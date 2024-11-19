package factory

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

func (t *TextInputFactory) BuildFromDto(questionDto *create.TextQuestionDto) (*question.Question, error) {
	questionObj := new(question.Question)
	questionObj.TextInput = new(question.TextInput)
	if err := t.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, questionObj); err != nil {
		return nil, err
	}

	questionObj.TextInput.Points = questionDto.Points
	questionObj.TextInput.IsCaseSensitive = questionDto.IsCaseSensitive

	answers := make([]question.TextInputAnswer, 0)
	for _, answer := range questionDto.Answers {
		var questionObjAnswer question.TextInputAnswer
		questionObjAnswer.Answer = answer
		answers = append(answers, questionObjAnswer)
	}
	questionObj.TextInput.Answers = answers
	return questionObj, nil
}

func (t *TextInputFactory) BuildFromObj(questionObj *question.Question) (*question.Question, error) {
	newQuestionObj := new(question.Question)
	newQuestionObj.TextInput = new(question.TextInput)
	if err := t.commonMapper.MapCommonFieldsObj(*questionObj, newQuestionObj); err != nil {
		return nil, err
	}

	newQuestionObj.TextInput.Points = questionObj.TextInput.Points
	newQuestionObj.TextInput.IsCaseSensitive = questionObj.TextInput.IsCaseSensitive
	newQuestionObj.TextInput.Answers = questionObj.TextInput.Answers
	return newQuestionObj, nil
}
