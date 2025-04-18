package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type TextInputMapper struct {
	commonMapper *CommonFieldQuestionDtoMapper
}

func NewTextInputMapper() *TextInputMapper {
	return &TextInputMapper{
		commonMapper: NewCommonFieldQuestionDtoMapper(),
	}
}

func (t *TextInputMapper) toDto(questionEntity *question.Question) (*get.TextInputDto, error) {
	textInputDto := new(get.TextInputDto)
	t.commonMapper.CommonFieldsToDto(questionEntity, textInputDto)
	textInputDto.Points = questionEntity.TextInput.Points
	textInputDto.Answers = t.AnswersToString(questionEntity.TextInput.Answers)
	return textInputDto, nil
}

func (t *TextInputMapper) AnswersToString(textInputAnswers []question.TextInputAnswer) []string {
	answers := make([]string, 0)
	for _, answer := range textInputAnswers {
		answers = append(answers, answer.Answer)
	}
	return answers
}
