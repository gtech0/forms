package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type SingleChoiceGeneratedIntegrationMapper struct{}

func NewSingleChoiceGeneratedIntegrationMapper() *SingleChoiceGeneratedIntegrationMapper {
	return &SingleChoiceGeneratedIntegrationMapper{}
}

func (s *SingleChoiceGeneratedIntegrationMapper) toDto(
	singleChoice *generated.SingleChoice,
	questionEntity *question.Question,
	isAnswerRequired bool,
) (*get.IntegratedSingleChoiceDto, error) {
	singleChoiceDto := new(get.IntegratedSingleChoiceDto)
	singleChoiceDto.Id = singleChoice.Id
	singleChoiceDto.Description = singleChoice.Description
	singleChoiceDto.OwnerId = questionEntity.OwnerId
	singleChoiceDto.Type = singleChoice.Type
	singleChoiceDto.Options, singleChoiceDto.Answer = s.singleOptionToDto(questionEntity.SingleChoice.Options, isAnswerRequired)
	singleChoiceDto.EnteredAnswer = singleChoice.EnteredAnswer
	return singleChoiceDto, nil
}

func (s *SingleChoiceGeneratedIntegrationMapper) singleOptionToDto(
	singleChoiceOptions []question.SingleChoiceOption,
	isAnswerRequired bool,
) ([]get.IntegratedSingleOptionDto, get.IntegratedSingleOptionDto) {
	options := make([]get.IntegratedSingleOptionDto, len(singleChoiceOptions))
	var answer get.IntegratedSingleOptionDto
	for _, singleChoiceOption := range singleChoiceOptions {
		var option get.IntegratedSingleOptionDto
		option.Id = singleChoiceOption.Id
		option.Text = singleChoiceOption.Text
		if singleChoiceOption.IsAnswer && isAnswerRequired {
			answer = option
		} else if !singleChoiceOption.IsAnswer {
			options = append(options, option)
		}
	}
	return options, answer
}
