package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/question"
	"slices"
)

type MultipleChoiceProcessor struct{}

func NewMultipleChoiceProcessor() *MultipleChoiceProcessor {
	return &MultipleChoiceProcessor{}
}

func (m *MultipleChoiceProcessor) saveAnswerAndCalculatePoints(
	multipleChoice *generated.MultipleChoice,
	multipleChoiceEntity *question.MultipleChoice,
	optionIds []uuid.UUID,
) (int, error) {
	if err := m.saveAnswer(multipleChoice, optionIds); err != nil {
		return 0, err
	}

	points, err := m.calculateAndSetPoints(multipleChoice, multipleChoiceEntity)
	if err != nil {
		return 0, err
	}

	return points, nil
}

func (m *MultipleChoiceProcessor) calculateAndSetPoints(
	multipleChoice *generated.MultipleChoice,
	multipleChoiceEntity *question.MultipleChoice,
) (int, error) {
	var correctAnswers int
	for _, enteredAnswer := range multipleChoice.EnteredAnswers {
		option, err := m.getAnswerEntity(multipleChoiceEntity, enteredAnswer)
		if err != nil {
			return 0, err
		}

		if option.IsAnswer {
			correctAnswers++
		}
	}

	incorrectAnswers := len(multipleChoice.EnteredAnswers) - correctAnswers
	correctAnswers = correctAnswers - incorrectAnswers

	multipleChoice.Points = m.calculatePoints(multipleChoiceEntity.Points, correctAnswers)
	return multipleChoice.Points, nil
}

func (m *MultipleChoiceProcessor) getAnswerEntity(
	multipleChoiceEntity *question.MultipleChoice,
	optionId uuid.UUID,
) (*question.MultipleChoiceOption, error) {
	for _, multipleChoiceOption := range multipleChoiceEntity.Options {
		if multipleChoiceOption.Id == optionId {
			return &multipleChoiceOption, nil
		}
	}

	return nil, errs.New(
		fmt.Sprintf("answer for multiple choice question %v doesn't exist", multipleChoiceEntity.QuestionId),
		500)
}

func (m *MultipleChoiceProcessor) calculatePoints(matchingPoints []question.MultipleChoicePoints, correctAnswers int) int {
	var points int
	for _, matchingPoint := range matchingPoints {
		if matchingPoint.CorrectAnswer > points && matchingPoint.CorrectAnswer <= correctAnswers {
			points = matchingPoint.CorrectAnswer
		}
	}

	return points
}

func (m *MultipleChoiceProcessor) saveAnswer(multipleChoice *generated.MultipleChoice, optionIds []uuid.UUID) error {
	questionOptionIds := make([]uuid.UUID, 0)
	for _, optionId := range multipleChoice.Options {
		questionOptionIds = append(questionOptionIds, optionId.Id)
	}

	missingOptionIds := make([]uuid.UUID, 0)
	for _, optionId := range optionIds {
		if !slices.Contains(questionOptionIds, optionId) {
			missingOptionIds = append(missingOptionIds, optionId)
		}
	}

	if len(missingOptionIds) > 0 {
		return errs.New(fmt.Sprintf("options %v not found", missingOptionIds), 404)
	}

	multipleChoice.EnteredAnswers = optionIds
	return nil
}
