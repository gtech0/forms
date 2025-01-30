package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/create"
	question2 "hedgehog-forms/internal/core/model/form/pattern/section/block/question"
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

func (m *MultipleChoiceFactory) BuildFromDto(questionDto *create.MultipleChoiceQuestionDto) (*question2.Question, error) {
	questionEntity := new(question2.Question)
	questionEntity.MultipleChoice = new(question2.MultipleChoice)
	if err := m.commonMapper.MapCommonDtoFields(questionDto.NewQuestionDto, questionEntity); err != nil {
		return nil, err
	}

	options := make([]question2.MultipleChoiceOption, 0)
	for order := 0; order < len(questionDto.Options); order++ {
		option := m.buildOptionFromDto(questionDto, order)
		options = append(options, option)
	}

	points := make([]question2.MultipleChoicePoints, 0)
	for answer, point := range questionDto.Points {
		pointsEntity := m.buildPointFromDto(answer, point)
		points = append(points, pointsEntity)
	}

	questionEntity.MultipleChoice.Options = options
	questionEntity.MultipleChoice.Points = points
	return questionEntity, nil
}

func (m *MultipleChoiceFactory) buildOptionFromDto(
	questionDto *create.MultipleChoiceQuestionDto,
	order int,
) question2.MultipleChoiceOption {
	var option question2.MultipleChoiceOption
	option.Text = questionDto.Options[order]
	option.Order = order
	option.IsAnswer = slices.Contains(questionDto.CorrectOptions, order)
	return option
}

func (m *MultipleChoiceFactory) buildPointFromDto(
	answer int,
	point int,
) question2.MultipleChoicePoints {
	var points question2.MultipleChoicePoints
	points.CorrectAnswer = answer
	points.Points = point
	return points
}

func (m *MultipleChoiceFactory) BuildFromEntity(questionEntity *question2.Question) (*question2.Question, error) {
	newQuestion := new(question2.Question)
	newQuestion.MultipleChoice = new(question2.MultipleChoice)
	options := make([]question2.MultipleChoiceOption, 0)
	newQuestion.MultipleChoice.Points = questionEntity.MultipleChoice.Points
	if err := m.commonMapper.MapCommonEntityFields(*questionEntity, newQuestion); err != nil {
		return nil, err
	}

	for _, option := range questionEntity.MultipleChoice.Options {
		newOption := m.buildOptionFromEntity(option, newQuestion.Id)
		options = append(options, newOption)
	}

	newQuestion.MultipleChoice.Options = options
	return newQuestion, nil
}

func (m *MultipleChoiceFactory) buildOptionFromEntity(
	optionEntity question2.MultipleChoiceOption,
	questionId uuid.UUID,
) question2.MultipleChoiceOption {
	var option question2.MultipleChoiceOption
	option.Text = optionEntity.Text
	option.Order = optionEntity.Order
	option.IsAnswer = optionEntity.IsAnswer
	option.MultipleChoiceId = questionId
	return option
}
