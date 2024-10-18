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

func (t *TextInputFactory) BuildFromDto(questionDto *create.TextQuestionDto) (*question.TextInput, error) {
	questionObj := new(question.TextInput)
	if err := t.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, &questionObj.Question); err != nil {
		return nil, err
	}

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

func (t *TextInputFactory) BuildFromObj(questionObj *question.TextInput) (*question.TextInput, error) {
	newQuestionObj := new(question.TextInput)
	if err := t.commonMapper.MapCommonFieldsObj(questionObj.Question, &newQuestionObj.Question); err != nil {
		return nil, err
	}

	newQuestionObj.Points = questionObj.Points
	newQuestionObj.IsCaseSensitive = questionObj.IsCaseSensitive
	newQuestionObj.Answers = questionObj.Answers
	return newQuestionObj, nil
}
