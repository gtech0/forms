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

func (t *TextInputMapper) toDto(textInputObj *question.TextInput) (*get.TextInputDto, error) {
	textInputDto := new(get.TextInputDto)
	t.commonMapper.commonFieldsToDto(textInputObj, textInputDto)
	textInputDto.Points = textInputObj.Points
	textInputDto.Answers = t.AnswersToDto(textInputObj.Answers)
	return textInputDto, nil
}

func (t *TextInputMapper) AnswersToDto(answersObj []question.TextInputAnswer) []string {
	answers := make([]string, 0)
	for _, answerObj := range answersObj {
		answers = append(answers, answerObj.Answer)
	}
	return answers
}
