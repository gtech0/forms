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

func (s *SingleChoiceFactory) BuildFromDto(questionDto *create.SingleChoiceQuestionDto) (*question.Question, error) {
	questionObj := new(question.Question)
	questionObj.SingleChoice = new(question.SingleChoice)
	if err := s.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, questionObj); err != nil {
		return nil, err
	}
	questionObj.SingleChoice.Points = questionDto.Points

	optionNames := questionDto.Options
	options := make([]question.SingleChoiceOption, 0)
	for order := 0; order < len(optionNames); order++ {
		option := s.buildOptionFromDto(questionDto, order, questionObj.Id)
		options = append(options, option)
	}
	questionObj.SingleChoice.Options = options
	return questionObj, nil
}

func (s *SingleChoiceFactory) BuildFromObj(questionObj *question.Question) (*question.Question, error) {
	newQuestionObj := new(question.Question)
	newQuestionObj.SingleChoice = new(question.SingleChoice)
	if err := s.commonMapper.MapCommonFieldsObj(*questionObj, newQuestionObj); err != nil {
		return nil, err
	}
	newQuestionObj.SingleChoice.Points = questionObj.SingleChoice.Points

	options := make([]question.SingleChoiceOption, 0)
	for _, option := range questionObj.SingleChoice.Options {
		var newOption question.SingleChoiceOption
		newOption.Text = option.Text
		newOption.Order = option.Order
		newOption.IsAnswer = option.IsAnswer
	}
	newQuestionObj.SingleChoice.Options = options
	return newQuestionObj, nil
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
