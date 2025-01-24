package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/pattern/section/block/question"
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
	textInputDto.Answers = t.AnswersToDto(questionEntity.TextInput.Answers)
	return textInputDto, nil
}

func (t *TextInputMapper) AnswersToDto(textInputAnswers []question.TextInputAnswer) []string {
	answers := make([]string, 0)
	for _, answer := range textInputAnswers {
		answers = append(answers, answer.Answer)
	}
	return answers
}
