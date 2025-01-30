package factory

import (
	"hedgehog-forms/internal/core/dto/create"
	question2 "hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type TextInputFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewTextInputFactory() *TextInputFactory {
	return &TextInputFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (t *TextInputFactory) BuildFromDto(questionDto *create.TextQuestionDto) (*question2.Question, error) {
	questionEntity := new(question2.Question)
	questionEntity.TextInput = new(question2.TextInput)
	if err := t.commonMapper.MapCommonDtoFields(questionDto.NewQuestionDto, questionEntity); err != nil {
		return nil, err
	}

	questionEntity.TextInput.Points = questionDto.Points
	questionEntity.TextInput.IsCaseSensitive = questionDto.IsCaseSensitive

	answers := make([]question2.TextInputAnswer, 0)
	for _, answer := range questionDto.Answers {
		var answerEntity question2.TextInputAnswer
		answerEntity.Answer = answer
		answers = append(answers, answerEntity)
	}
	questionEntity.TextInput.Answers = answers
	return questionEntity, nil
}

func (t *TextInputFactory) BuildFromEntity(questionEntity *question2.Question) (*question2.Question, error) {
	newQuestion := new(question2.Question)
	newQuestion.TextInput = new(question2.TextInput)
	if err := t.commonMapper.MapCommonEntityFields(*questionEntity, newQuestion); err != nil {
		return nil, err
	}

	newQuestion.TextInput.Points = questionEntity.TextInput.Points
	newQuestion.TextInput.IsCaseSensitive = questionEntity.TextInput.IsCaseSensitive
	newQuestion.TextInput.Answers = questionEntity.TextInput.Answers
	return newQuestion, nil
}
