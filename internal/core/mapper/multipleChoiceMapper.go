package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	question2 "hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type MultipleChoiceMapper struct {
	commonMapper *CommonFieldQuestionDtoMapper
}

func NewMultipleChoiceMapper() *MultipleChoiceMapper {
	return &MultipleChoiceMapper{
		commonMapper: NewCommonFieldQuestionDtoMapper(),
	}
}

func (m *MultipleChoiceMapper) toDto(questionEntity *question2.Question) (*get.MultipleChoiceDto, error) {
	multipleChoiceDto := new(get.MultipleChoiceDto)
	m.commonMapper.CommonFieldsToDto(questionEntity, multipleChoiceDto)
	multipleChoiceDto.Points = m.pointsToDto(questionEntity.MultipleChoice.Points)
	multipleChoiceDto.Options = m.optionsToDto(questionEntity.MultipleChoice.Options)
	return multipleChoiceDto, nil
}

func (m *MultipleChoiceMapper) pointsToDto(multipleChoicePoints []question2.MultipleChoicePoints) map[int]int {
	points := make(map[int]int)
	for _, multipleChoicePoint := range multipleChoicePoints {
		points[multipleChoicePoint.CorrectAnswer] = multipleChoicePoint.Points
	}
	return points
}

func (m *MultipleChoiceMapper) optionsToDto(multipleChoiceOptions []question2.MultipleChoiceOption) []get.MultipleOptionDto {
	options := make([]get.MultipleOptionDto, 0)
	for _, multipleChoiceOption := range multipleChoiceOptions {
		var option get.MultipleOptionDto
		option.Id = multipleChoiceOption.Id
		option.Text = multipleChoiceOption.Text
		option.IsAnswer = multipleChoiceOption.IsAnswer
		options = append(options, option)
	}
	return options
}
