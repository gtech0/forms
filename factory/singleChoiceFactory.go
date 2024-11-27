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
	questionEntity := new(question.Question)
	questionEntity.SingleChoice = new(question.SingleChoice)
	if err := s.commonMapper.MapCommonDtoFields(questionDto.NewQuestionDto, questionEntity); err != nil {
		return nil, err
	}
	questionEntity.SingleChoice.Points = questionDto.Points

	optionNames := questionDto.Options
	options := make([]question.SingleChoiceOption, 0)
	for order := 0; order < len(optionNames); order++ {
		option := s.buildOptionFromDto(questionDto, order, questionEntity.Id)
		options = append(options, option)
	}
	questionEntity.SingleChoice.Options = options
	return questionEntity, nil
}

func (s *SingleChoiceFactory) BuildFromEntity(questionEntity *question.Question) (*question.Question, error) {
	newQuestion := new(question.Question)
	newQuestion.SingleChoice = new(question.SingleChoice)
	if err := s.commonMapper.MapCommonEntityFields(*questionEntity, newQuestion); err != nil {
		return nil, err
	}
	newQuestion.SingleChoice.Points = questionEntity.SingleChoice.Points

	options := make([]question.SingleChoiceOption, 0)
	for _, option := range questionEntity.SingleChoice.Options {
		var newOption question.SingleChoiceOption
		newOption.Text = option.Text
		newOption.Order = option.Order
		newOption.IsAnswer = option.IsAnswer
	}
	newQuestion.SingleChoice.Options = options
	return newQuestion, nil
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
