package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type SingleChoiceProcessor struct{}

func NewSingleChoiceProcessor() *SingleChoiceProcessor {
	return &SingleChoiceProcessor{}
}

func (s *SingleChoiceProcessor) calculateAndSetPoints(
	singleChoice *generated.SingleChoice,
	singleChoiceEntity *question.SingleChoice,
) (int, error) {
	correctOption, err := s.getAnswerEntity(singleChoiceEntity)
	if err != nil {
		return 0, err
	}

	if singleChoice.EnteredAnswer == correctOption.Id {
		singleChoice.Points = singleChoiceEntity.Points
	}

	return singleChoice.Points, nil
}

func (s *SingleChoiceProcessor) getAnswerEntity(
	singleChoiceEntity *question.SingleChoice,
) (*question.SingleChoiceOption, error) {
	for _, singleChoiceOption := range singleChoiceEntity.Options {
		if singleChoiceOption.IsAnswer {
			return &singleChoiceOption, nil
		}
	}
	return nil, errs.New(
		fmt.Sprintf("answer for sinlge choice question %v doesn't exist", singleChoiceEntity.QuestionId),
		500)
}

func (s *SingleChoiceProcessor) saveAnswer(singleChoice *generated.SingleChoice, optionId uuid.UUID) error {
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
