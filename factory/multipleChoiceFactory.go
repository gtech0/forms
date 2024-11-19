package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern/section/block/question"
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

func (m *MultipleChoiceFactory) BuildFromDto(questionDto *create.MultipleChoiceQuestionDto) (*question.Question, error) {
	questionObj := new(question.Question)
	questionObj.MultipleChoice = new(question.MultipleChoice)
	if err := m.commonMapper.MapCommonFieldsDto(questionDto.NewQuestionDto, questionObj); err != nil {
		return nil, err
	}

	options := make([]question.MultipleChoiceOption, 0)
	for order := 0; order < len(questionDto.Options); order++ {
		option := m.buildOptionFromDto(questionDto, order)
		options = append(options, option)
	}

	points := make([]question.MultipleChoicePoints, 0)
	for answer, point := range questionDto.Points {
		pointsObj := m.buildPointFromDto(answer, point)
		points = append(points, pointsObj)
	}

	questionObj.MultipleChoice.Options = options
	questionObj.MultipleChoice.Points = points
	return questionObj, nil
}

func (m *MultipleChoiceFactory) buildOptionFromDto(
	questionDto *create.MultipleChoiceQuestionDto,
	order int,
) question.MultipleChoiceOption {
	var option question.MultipleChoiceOption
	option.Text = questionDto.Options[order]
	option.Order = order
	option.IsAnswer = slices.Contains(questionDto.CorrectOptions, order)
	return option
}

func (m *MultipleChoiceFactory) buildPointFromDto(
	answer int,
	point int,
) question.MultipleChoicePoints {
	var pointsObj question.MultipleChoicePoints
	pointsObj.CorrectAnswers = answer
	pointsObj.Points = point
	return pointsObj
}

func (m *MultipleChoiceFactory) BuildFromObj(questionObj *question.Question) (*question.Question, error) {
	newQuestionObj := new(question.Question)
	newQuestionObj.MultipleChoice = new(question.MultipleChoice)
	options := make([]question.MultipleChoiceOption, 0)
	newQuestionObj.MultipleChoice.Points = questionObj.MultipleChoice.Points
	if err := m.commonMapper.MapCommonFieldsObj(*questionObj, newQuestionObj); err != nil {
		return nil, err
	}

	for _, option := range questionObj.MultipleChoice.Options {
		newOption := m.buildOptionFromEntity(option, newQuestionObj.Id)
		options = append(options, newOption)
	}

	newQuestionObj.MultipleChoice.Options = options
	return newQuestionObj, nil
}

func (m *MultipleChoiceFactory) buildOptionFromEntity(
	optionObj question.MultipleChoiceOption,
	questionId uuid.UUID,
) question.MultipleChoiceOption {
	var option question.MultipleChoiceOption
	option.Text = optionObj.Text
	option.Order = optionObj.Order
	option.IsAnswer = optionObj.IsAnswer
	option.MultipleChoiceId = questionId
	return option
}
