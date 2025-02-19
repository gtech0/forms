package factory

import (
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/model/form/pattern/question"
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
	questionEntity := new(question.Question)
	questionEntity.TextInput = new(question.TextInput)
	if err := t.commonMapper.MapCommonDtoFields(questionDto.NewQuestionDto, questionEntity); err != nil {
		return nil, err
	}

	questionEntity.TextInput.Points = questionDto.Points
	questionEntity.TextInput.IsCaseSensitive = questionDto.IsCaseSensitive

	answers := make([]question.TextInputAnswer, 0)
	for _, answer := range questionDto.Answers {
		var answerEntity question.TextInputAnswer
		answerEntity.Answer = answer
		answers = append(answers, answerEntity)
	}
	questionEntity.TextInput.Answers = answers
	return questionEntity, nil
}

func (t *TextInputFactory) BuildFromEntity(questionEntity *question.Question) (*question.Question, error) {
	newQuestion := new(question.Question)
	newQuestion.TextInput = new(question.TextInput)
	if err := t.commonMapper.MapCommonEntityFields(*questionEntity, newQuestion); err != nil {
		return nil, err
	}

	newQuestion.TextInput.Points = questionEntity.TextInput.Points
	newQuestion.TextInput.IsCaseSensitive = questionEntity.TextInput.IsCaseSensitive
	newQuestion.TextInput.Answers = questionEntity.TextInput.Answers
	return newQuestion, nil
}
