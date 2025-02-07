package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type SingleChoiceMapper struct {
	commonMapper *CommonFieldQuestionDtoMapper
}

func NewSingleChoiceMapper() *SingleChoiceMapper {
	return &SingleChoiceMapper{
		commonMapper: NewCommonFieldQuestionDtoMapper(),
	}
}

func (s *SingleChoiceMapper) toDto(questionEntity *question.Question) (*get.SingleChoiceDto, error) {
	singleChoiceDto := new(get.SingleChoiceDto)
	s.commonMapper.CommonFieldsToDto(questionEntity, singleChoiceDto)
	singleChoiceDto.Points = questionEntity.SingleChoice.Points
	singleChoiceDto.Choices = s.singleOptionToDto(questionEntity.SingleChoice.Options)
	return singleChoiceDto, nil
}

func (s *SingleChoiceMapper) singleOptionToDto(singleChoiceOptions []question.SingleChoiceOption) []get.SingleOptionDto {
	options := make([]get.SingleOptionDto, len(singleChoiceOptions))
	for _, singleChoiceOption := range singleChoiceOptions {
		var option get.SingleOptionDto
		option.Id = singleChoiceOption.Id
		option.Text = singleChoiceOption.Text
		option.IsAnswer = singleChoiceOption.IsAnswer
		options = append(options, option)
	}
	return options
}
