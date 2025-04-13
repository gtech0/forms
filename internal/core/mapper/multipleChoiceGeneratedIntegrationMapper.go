package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type MultipleChoiceGeneratedIntegrationMapper struct{}

func NewMultipleChoiceGeneratedIntegrationMapper() *MultipleChoiceGeneratedIntegrationMapper {
	return &MultipleChoiceGeneratedIntegrationMapper{}
}

func (m *MultipleChoiceGeneratedIntegrationMapper) toDto(
	multipleChoice *generated.MultipleChoice,
	questionEntity *question.Question,
	isAnswerRequired bool,
) (*get.IntegratedMultipleChoiceDto, error) {
	multipleChoiceDto := new(get.IntegratedMultipleChoiceDto)
	multipleChoiceDto.Id = multipleChoice.Id
	multipleChoiceDto.Description = multipleChoice.Description
	multipleChoiceDto.OwnerId = questionEntity.OwnerId
	multipleChoiceDto.Type = multipleChoice.Type
	multipleChoiceDto.Options, multipleChoiceDto.CorrectOptions = m.optionsToDto(questionEntity.MultipleChoice.Options, isAnswerRequired)
	multipleChoiceDto.EnteredAnswers = multipleChoice.EnteredAnswers
	return multipleChoiceDto, nil
}

func (m *MultipleChoiceGeneratedIntegrationMapper) optionsToDto(
	multipleChoiceOptions []question.MultipleChoiceOption,
	isAnswerRequired bool,
) ([]get.IntegratedMultipleOptionDto, []get.IntegratedMultipleOptionDto) {
	options := make([]get.IntegratedMultipleOptionDto, 0)
	correctOptions := make([]get.IntegratedMultipleOptionDto, 0)
	for _, multipleChoiceOption := range multipleChoiceOptions {
		var option get.IntegratedMultipleOptionDto
		option.Id = multipleChoiceOption.Id
		option.Text = multipleChoiceOption.Text
		if multipleChoiceOption.IsAnswer && isAnswerRequired {
			correctOptions = append(correctOptions, option)
		} else if !isAnswerRequired {
			options = append(options, option)
		}
	}
	return options, correctOptions
}
