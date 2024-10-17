package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type SingleChoiceFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewSingleChoiceFactory() *SingleChoiceFactory {
	return &SingleChoiceFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (s *SingleChoiceFactory) BuildFromDto(questionDto *create.SingleChoiceQuestionDto) (*question.SingleChoice, error) {
	questionObj := new(question.SingleChoice)
	s.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, &questionObj.Question)
	questionObj.Points = questionDto.Points

	optionNames := questionDto.Options
	options := make([]question.SingleChoiceOption, 0)
	for order := 0; order < len(optionNames); order++ {
		option := s.buildOptionFromDto(questionDto, order, questionObj.Id)
		options = append(options, option)
	}
	questionObj.Options = options
	return questionObj, nil
}

func (s *SingleChoiceFactory) BuildFromObj(questionObj *question.SingleChoice) *question.SingleChoice {
	newQuestionObj := new(question.SingleChoice)
	s.commonMapper.MapCommonFieldsObj(questionObj.Question, &newQuestionObj.Question)
	newQuestionObj.Points = questionObj.Points

	options := make([]question.SingleChoiceOption, 0)
	for _, option := range questionObj.Options {
		var newOption question.SingleChoiceOption
		newOption.Text = option.Text
		newOption.Order = option.Order
		newOption.IsAnswer = option.IsAnswer
	}
	newQuestionObj.Options = options
	return newQuestionObj
}

func (s *SingleChoiceFactory) buildOptionFromDto(
	questionDto *create.SingleChoiceQuestionDto,
	order int,
	questionId uuid.UUID,
) question.SingleChoiceOption {
	var option question.SingleChoiceOption
	option.Text = questionDto.Options[order]
	option.Order = order
	option.IsAnswer = questionDto.CorrectOption == order
	option.SingleChoiceId = questionId
	return option
}
