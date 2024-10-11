package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
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

func (m *MultipleChoiceFactory) BuildFromDto(questionDto *create.MultipleChoiceQuestionDto) (*question.MultipleChoice, error) {
	questionObj := new(question.MultipleChoice)
	m.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, questionObj)

	optionNames := questionDto.Options
	options := make([]question.MultipleChoiceOption, 0)
	for order := 0; order < len(optionNames); order++ {
		option := m.buildOptionFromDto(questionDto, order, questionObj.Id)
		options = append(options, option)
	}

	points := make([]question.MultipleChoicePoints, 0)
	for answer, point := range questionDto.Points {
		var pointsObj question.MultipleChoicePoints
		pointsObj.CorrectAnswer = answer
		pointsObj.Points = point
		points = append(points, pointsObj)
	}

	questionObj.Options = options
	questionObj.Points = points
	return questionObj, nil
}

func (m *MultipleChoiceFactory) buildOptionFromDto(
	questionDto *create.MultipleChoiceQuestionDto,
	order int,
	questionId uuid.UUID,
) question.MultipleChoiceOption {
	var option question.MultipleChoiceOption
	option.Text = questionDto.Options[order]
	option.Order = order
	option.IsAnswer = slices.Contains(questionDto.CorrectOptions, order)
	option.MultipleChoiceId = questionId
	return option
}

func (m *MultipleChoiceFactory) BuildFromObj(questionObj *question.MultipleChoice) *question.MultipleChoice {
	newQuestionObj := new(question.MultipleChoice)
	options := make([]question.MultipleChoiceOption, 0)
	newQuestionObj.Points = questionObj.Points
	m.commonMapper.MapCommonFieldsObj(questionObj.Question, newQuestionObj)

	for _, option := range questionObj.Options {
		m.buildOptionFromEntity(&option, newQuestionObj.Id)
		options = append(options, option)
	}

	newQuestionObj.Options = options
	return newQuestionObj
}

func (m *MultipleChoiceFactory) buildOptionFromEntity(
	optionObj *question.MultipleChoiceOption,
	questionId uuid.UUID,
) {
	var option question.MultipleChoiceOption
	option.Text = optionObj.Text
	option.Order = optionObj.Order
	option.IsAnswer = optionObj.IsAnswer
	option.MultipleChoiceId = questionId
}
