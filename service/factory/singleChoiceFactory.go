package factory

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block/question"
)

type SingleChoiceFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewSingleChoiceFactory() *SingleChoiceFactory {
	return &SingleChoiceFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (s *SingleChoiceFactory) BuildFromDto(questionDto dto.CreateSingleChoiceQuestionDto) (question.SingleChoice, error) {
	var questionObj question.SingleChoice
	s.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, questionObj.Question)
	questionObj.Points = questionDto.Points

	optionNames := questionDto.Options
	options := make([]question.SingleChoiceOption, len(optionNames))
	for order := 0; order < len(optionNames); order++ {
		option := s.buildOptionFromDto(questionDto, order, questionObj)
		options = append(options, option)
	}
	questionObj.SingleChoiceOptions = options
	return questionObj, nil
}

func (s *SingleChoiceFactory) BuildFromEntity(questionObj question.SingleChoice) question.SingleChoice {
	var newQuestionObj question.SingleChoice
	s.commonMapper.MapCommonFieldsObj(questionObj.Question, newQuestionObj.Question)
	newQuestionObj.Points = questionObj.Points

	options := make([]question.SingleChoiceOption, len(questionObj.SingleChoiceOptions))
	for _, option := range questionObj.SingleChoiceOptions {
		var newOption question.SingleChoiceOption
		newOption.Title = option.Title
		newOption.Order = option.Order
		newOption.IsAnswer = option.IsAnswer
		newOption.SingleChoiceId = newQuestionObj.Id
	}
	newQuestionObj.SingleChoiceOptions = options
	return newQuestionObj
}

func (s *SingleChoiceFactory) buildOptionFromDto(
	questionDto dto.CreateSingleChoiceQuestionDto,
	order int,
	questionObj question.SingleChoice,
) question.SingleChoiceOption {
	var option question.SingleChoiceOption
	option.Title = questionDto.Options[order]
	option.Order = order
	option.IsAnswer = questionDto.CorrectOption == order
	option.SingleChoiceId = questionObj.Id
	return option
}
