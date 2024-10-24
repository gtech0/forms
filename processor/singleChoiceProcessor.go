package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type SingleChoiceProcessor struct{}

func NewSingleChoiceProcessor() *SingleChoiceProcessor {
	return &SingleChoiceProcessor{}
}

func (s *SingleChoiceProcessor) markAnswers(singleChoiceQuestions []*generated.SingleChoice, answers get.AnswerDto) error {
	for questionId, optionId := range answers.SingleChoice {
		singleChoice, err := findQuestion[*generated.SingleChoice](singleChoiceQuestions, questionId)
		if err != nil {
			return err
		}

		if err = s.markAnswer(singleChoice, optionId); err != nil {
			return err
		}
	}
	return nil
}

func (s *SingleChoiceProcessor) markAnswersAndCalculatePoints(
	singleChoiceQuestions []*generated.SingleChoice,
	singleChoiceObjs []*question.SingleChoice,
	answers get.AnswerDto,
) (int, error) {
	var points int
	for questionId, optionId := range answers.SingleChoice {
		singleChoice, err := findQuestion[*generated.SingleChoice](singleChoiceQuestions, questionId)
		if err != nil {
			return 0, err
		}

		singleChoiceObj, err := findQuestionObj[*question.SingleChoice](singleChoiceObjs, questionId)
		if err != nil {
			return 0, err
		}

		if err = s.markAnswer(singleChoice, optionId); err != nil {
			return 0, err
		}

		calculatedPoints, err := s.calculateAndSetPoints(singleChoice, singleChoiceObj)
		if err != nil {
			return 0, err
		}

		points += calculatedPoints
	}
	return points, nil
}

func (s *SingleChoiceProcessor) calculateAndSetPoints(
	singleChoice *generated.SingleChoice,
	singleChoiceObj *question.SingleChoice,
) (int, error) {
	correctOption, err := s.getSingleOptionObj(singleChoiceObj)
	if err != nil {
		return 0, err
	}

	if singleChoice.EnteredAnswer == correctOption.Id {
		singleChoice.Points = singleChoiceObj.Points
	}

	return singleChoice.Points, nil
}

func (s *SingleChoiceProcessor) getSingleOptionObj(singleChoiceObj *question.SingleChoice) (*question.SingleChoiceOption, error) {
	for _, singleChoiceOption := range singleChoiceObj.Options {
		if singleChoiceOption.IsAnswer {
			return &singleChoiceOption, nil
		}
	}
	return nil, errs.New(fmt.Sprintf("answer for sinlge choice question %v doesn't exist", singleChoiceObj.Id), 500)
}

func (s *SingleChoiceProcessor) markAnswer(singleChoice *generated.SingleChoice, optionId uuid.UUID) error {
	isOptionExist := func() bool {
		for _, singleChoiceOption := range singleChoice.Options {
			if singleChoiceOption.Id == optionId {
				return true
			}
		}
		return false
	}

	if !isOptionExist() {
		return errs.New(fmt.Sprintf("option %v doesn't exist", optionId), 500)
	}

	singleChoice.EnteredAnswer = optionId
	return nil
}
