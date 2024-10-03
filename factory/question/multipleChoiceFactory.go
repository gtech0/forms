package question

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block/question"
	"slices"
)

type MultipleChoiceFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewMultipleChoiceFactory() *MultipleChoiceFactory {
	return &MultipleChoiceFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (m *MultipleChoiceFactory) BuildFromDto(questionDto *dto.CreateMultipleChoiceQuestionDto) (question.MultipleChoice, error) {
	var questionObj question.MultipleChoice
	m.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, questionObj.Question)

	optionNames := questionDto.Options
	options := make([]question.MultipleChoiceOption, len(optionNames))

	for order := 0; order < len(optionNames); order++ {
		option := m.buildOptionFromDto(questionDto, order, questionObj)
		options = append(options, option)
	}

	questionObj.Options = options
	for answer, points := range questionDto.Points {
		var pointsObj question.MultipleChoicePoints
		pointsObj.CorrectAnswer = answer
		pointsObj.Points = points
		questionObj.Points = append(questionObj.Points, pointsObj)
	}

	return questionObj, nil
}

func (m *MultipleChoiceFactory) buildOptionFromDto(
	questionDto *dto.CreateMultipleChoiceQuestionDto,
	order int,
	questionObj question.MultipleChoice,
) question.MultipleChoiceOption {
	var option question.MultipleChoiceOption
	option.Title = questionDto.Options[order]
	option.Order = order
	option.IsAnswer = slices.Contains(questionDto.CorrectOptions, order)
	option.MultipleChoiceId = questionObj.Id
	return option
}

func (m *MultipleChoiceFactory) BuildFromObj(questionObj question.MultipleChoice) question.MultipleChoice {
	var newQuestionObj question.MultipleChoice
	newOptions := make([]question.MultipleChoiceOption, len(questionObj.Options))
	newQuestionObj.Points = questionObj.Points
	m.commonMapper.MapCommonFieldsObj(questionObj.Question, newQuestionObj.Question)

	for _, option := range questionObj.Options {
		newOption := m.buildOptionFromEntity(option, newQuestionObj)
		newOptions = append(newOptions, newOption)
	}

	newQuestionObj.Options = newOptions

	return newQuestionObj
}

func (m *MultipleChoiceFactory) buildOptionFromEntity(
	optionObj question.MultipleChoiceOption,
	questionObj question.MultipleChoice,
) question.MultipleChoiceOption {
	var option question.MultipleChoiceOption
	option.Title = optionObj.Title
	option.Order = optionObj.Order
	option.IsAnswer = optionObj.IsAnswer
	option.MultipleChoiceId = questionObj.Id
	return option
}
