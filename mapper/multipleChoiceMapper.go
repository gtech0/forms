package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type MultipleChoiceMapper struct {
	commonMapper *CommonFieldQuestionDtoMapper
}

func NewMultipleChoiceMapper() *MultipleChoiceMapper {
	return &MultipleChoiceMapper{
		commonMapper: NewCommonFieldQuestionDtoMapper(),
	}
}

func (m *MultipleChoiceMapper) toDto(multipleChoiceObj *question.MultipleChoice) (*get.MultipleChoiceDto, error) {
	multipleChoiceDto := new(get.MultipleChoiceDto)
	m.commonMapper.commonFieldsToDto(multipleChoiceObj, multipleChoiceDto)
	multipleChoiceDto.Points = m.pointsToDto(multipleChoiceObj.Points)
	multipleChoiceDto.Options = m.optionsToDto(multipleChoiceObj.Options)
	return multipleChoiceDto, nil
}

func (m *MultipleChoiceMapper) pointsToDto(pointsObj []question.MultipleChoicePoints) map[int]int {
	points := make(map[int]int)
	for _, pointObj := range pointsObj {
		points[pointObj.CorrectAnswer] = pointObj.Points
	}
	return points
}

func (m *MultipleChoiceMapper) optionsToDto(optionsObj []question.MultipleChoiceOption) []get.MultipleOptionDto {
	options := make([]get.MultipleOptionDto, 0)
	for _, optionObj := range optionsObj {
		var option get.MultipleOptionDto
		option.Id = optionObj.Id
		option.Text = optionObj.Text
		option.IsAnswer = optionObj.IsAnswer
		options = append(options, option)
	}
	return options
}
