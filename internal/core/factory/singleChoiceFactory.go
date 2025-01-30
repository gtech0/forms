package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/create"
	question2 "hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type SingleChoiceFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewSingleChoiceFactory() *SingleChoiceFactory {
	return &SingleChoiceFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (s *SingleChoiceFactory) BuildFromDto(questionDto *create.SingleChoiceQuestionDto) (*question2.Question, error) {
	questionEntity := new(question2.Question)
	questionEntity.SingleChoice = new(question2.SingleChoice)
	if err := s.commonMapper.MapCommonDtoFields(questionDto.NewQuestionDto, questionEntity); err != nil {
		return nil, err
	}
	questionEntity.SingleChoice.Points = questionDto.Points

	optionNames := questionDto.Options
	options := make([]question2.SingleChoiceOption, 0)
	for order := 0; order < len(optionNames); order++ {
		option := s.buildOptionFromDto(questionDto, order, questionEntity.Id)
		options = append(options, option)
	}
	questionEntity.SingleChoice.Options = options
	return questionEntity, nil
}

func (s *SingleChoiceFactory) BuildFromEntity(questionEntity *question2.Question) (*question2.Question, error) {
	newQuestion := new(question2.Question)
	newQuestion.SingleChoice = new(question2.SingleChoice)
	if err := s.commonMapper.MapCommonEntityFields(*questionEntity, newQuestion); err != nil {
		return nil, err
	}
	newQuestion.SingleChoice.Points = questionEntity.SingleChoice.Points

	options := make([]question2.SingleChoiceOption, 0)
	for _, option := range questionEntity.SingleChoice.Options {
		var newOption question2.SingleChoiceOption
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
) question2.SingleChoiceOption {
	var option question2.SingleChoiceOption
	option.Text = questionDto.Options[order]
	option.Order = order
	option.IsAnswer = questionDto.CorrectOption == order
	option.SingleChoiceId = questionId
	return option
}
