package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type TextInputGeneratedIntegrationMapper struct{}

func NewTextInputGeneratedIntegrationMapper() *TextInputGeneratedIntegrationMapper {
	return &TextInputGeneratedIntegrationMapper{}
}

func (m *TextInputGeneratedIntegrationMapper) toDto(
	singleChoice *generated.TextInput,
	questionEntity *question.Question,
) (*get.IntegratedTextInputDto, error) {
	singleChoiceDto := new(get.IntegratedTextInputDto)
	singleChoiceDto.Id = singleChoice.Id
	singleChoiceDto.Description = singleChoice.Description
	singleChoiceDto.OwnerId = questionEntity.OwnerId
	singleChoiceDto.Type = singleChoice.Type
	singleChoiceDto.Answers = m.answersToString(questionEntity.TextInput.Answers)
	singleChoiceDto.EnteredAnswer = singleChoice.EnteredAnswer
	return singleChoiceDto, nil
}

func (m *TextInputGeneratedIntegrationMapper) answersToString(textInputAnswers []question.TextInputAnswer) []string {
	answers := make([]string, 0)
	for _, answer := range textInputAnswers {
		answers = append(answers, answer.Answer)
	}
	return answers
}
