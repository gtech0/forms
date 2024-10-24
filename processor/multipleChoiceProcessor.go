package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"slices"
)

type MultipleChoiceProcessor struct{}

func NewMultipleChoiceProcessor() *MultipleChoiceProcessor {
	return &MultipleChoiceProcessor{}
}

func (m *MultipleChoiceProcessor) markAnswers(
	multipleChoiceQuestions []*generated.MultipleChoice,
	answersDto get.AnswerDto,
) error {
	for questionId, enteredAnswers := range answersDto.MultipleChoice {
		multipleChoice, err := findQuestion[*generated.MultipleChoice](multipleChoiceQuestions, questionId)
		if err != nil {
			return err
		}

		if err = m.markAnswer(multipleChoice, enteredAnswers); err != nil {
			return err
		}
	}
	return nil
}

func (m *MultipleChoiceProcessor) markAnswersAndCalculatePoints(
	multipleChoiceQuestions []*generated.MultipleChoice,
	multipleChoiceObjs []*question.MultipleChoice,
	answersDto get.AnswerDto,
) (int, error) {
	var points int
	for questionId, enteredAnswers := range answersDto.MultipleChoice {
		multipleChoice, err := findQuestion[*generated.MultipleChoice](multipleChoiceQuestions, questionId)
		if err != nil {
			return 0, err
		}

		multipleChoiceObj, err := findQuestionObj[*question.MultipleChoice](multipleChoiceObjs, questionId)
		if err != nil {
			return 0, err
		}

		if err = m.markAnswer(multipleChoice, enteredAnswers); err != nil {
			return 0, err
		}

		calculatedPoints, err := m.calculateAndSetPoints(multipleChoice, multipleChoiceObj)
		if err != nil {
			return 0, err
		}

		points += calculatedPoints
	}
	return points, nil
}

func (m *MultipleChoiceProcessor) calculateAndSetPoints(
	multipleChoice *generated.MultipleChoice,
	multipleChoiceObj *question.MultipleChoice,
) (int, error) {
	var correctAnswers int
	for _, enteredAnswer := range multipleChoice.EnteredAnswers {
		option, err := m.getMultipleOptionObj(multipleChoiceObj, enteredAnswer)
		if err != nil {
			return 0, err
		}

		if option.IsAnswer {
			correctAnswers++
		}
	}

	incorrectAnswers := len(multipleChoice.EnteredAnswers) - correctAnswers
	correctAnswers = correctAnswers - incorrectAnswers

	points := m.calculatePoints(multipleChoiceObj.Points, correctAnswers)
	return points, nil
}

func (m *MultipleChoiceProcessor) getMultipleOptionObj(
	multipleChoiceObj *question.MultipleChoice,
	optionId uuid.UUID,
) (*question.MultipleChoiceOption, error) {
	for _, multipleChoiceOption := range multipleChoiceObj.Options {
		if multipleChoiceOption.Id == optionId {
			return &multipleChoiceOption, nil
		}
	}
	return nil, errs.New(fmt.Sprintf("answer for multiple choice question %v doesn't exist", multipleChoiceObj.Id), 500)
}

func (m *MultipleChoiceProcessor) calculatePoints(matchingPoints []question.MultipleChoicePoints, correctAnswers int) int {
	var points int
	for _, matchingPoint := range matchingPoints {
		if matchingPoint.CorrectAnswers > points && matchingPoint.CorrectAnswers <= correctAnswers {
			points = matchingPoint.CorrectAnswers
		}
	}

	return points
}

func (m *MultipleChoiceProcessor) markAnswer(multipleChoice *generated.MultipleChoice, optionIds []uuid.UUID) error {
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
