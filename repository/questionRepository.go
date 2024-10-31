package repository

import (
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type QuestionRepository struct {
	textInputRepository        *TextInputRepository
	multipleChoiceRepository   *MultipleChoiceRepository
	singleChoiceRepository     *SingleChoiceRepository
	matchingQuestionRepository *MatchingQuestionRepository
}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{
		textInputRepository:        NewTextInputRepository(),
		multipleChoiceRepository:   NewMultipleChoiceRepository(),
		singleChoiceRepository:     NewSingleChoiceRepository(),
		matchingQuestionRepository: NewMatchingQuestionRepository(),
	}
}

func (q *QuestionRepository) Save(iQuestion question.IQuestion) error {
	switch iQuestion.GetType() {
	case question.MULTIPLE_CHOICE:
		if err := q.multipleChoiceRepository.Save(iQuestion.(*question.MultipleChoice)); err != nil {
			return err
		}
	case question.SINGLE_CHOICE:
		if err := q.singleChoiceRepository.Save(iQuestion.(*question.SingleChoice)); err != nil {
			return err
		}
	case question.MATCHING:
		if err := q.matchingQuestionRepository.Save(iQuestion.(*question.Matching)); err != nil {
			return err
		}
	case question.TEXT_INPUT:
		if err := q.multipleChoiceRepository.Save(iQuestion.(*question.MultipleChoice)); err != nil {
			return err
		}
	}
	return nil
}
