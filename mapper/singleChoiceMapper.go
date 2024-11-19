package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/pattern/section/block/question"
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

func (s *SingleChoiceMapper) singleOptionToDto(optionsObj []question.SingleChoiceOption) []get.SingleOptionDto {
	options := make([]get.SingleOptionDto, len(optionsObj))
	for _, optionObj := range optionsObj {
		var option get.SingleOptionDto
		option.Id = optionObj.Id
		option.Text = optionObj.Text
		option.IsAnswer = optionObj.IsAnswer
		options = append(options, option)
	}
	return options
}
