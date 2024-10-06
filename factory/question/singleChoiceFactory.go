package question

import (
	"github.com/google/uuid"
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

func (s *SingleChoiceFactory) BuildFromDto(questionDto *dto.CreateSingleChoiceQuestionDto) (*question.SingleChoice, error) {
	questionObj := new(question.SingleChoice)
	s.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, questionObj)
	questionObj.Points = questionDto.Points

	optionNames := questionDto.Options
	options := make([]question.SingleChoiceOption, 0)
	for order := 0; order < len(optionNames); order++ {
		option := s.buildOptionFromDto(questionDto, order, questionObj.Id)
		options = append(options, option)
	}
	questionObj.SingleChoiceOptions = options
	return questionObj, nil
}

func (s *SingleChoiceFactory) BuildFromObj(questionObj *question.SingleChoice) *question.SingleChoice {
	newQuestionObj := new(question.SingleChoice)
	s.commonMapper.MapCommonFieldsObj(questionObj.Question, newQuestionObj)
	newQuestionObj.Points = questionObj.Points

	options := make([]question.SingleChoiceOption, 0)
	for _, option := range questionObj.SingleChoiceOptions {
		var newOption question.SingleChoiceOption
		newOption.Title = option.Title
		newOption.Order = option.Order
		newOption.IsAnswer = option.IsAnswer
	}
	newQuestionObj.SingleChoiceOptions = options
	return newQuestionObj
}

func (s *SingleChoiceFactory) buildOptionFromDto(
	questionDto *dto.CreateSingleChoiceQuestionDto,
	order int,
	questionId uuid.UUID,
) question.SingleChoiceOption {
	var option question.SingleChoiceOption
	option.Title = questionDto.Options[order]
	option.Order = order
	option.IsAnswer = questionDto.CorrectOption == order
	option.SingleChoiceId = questionId
	return option
}
