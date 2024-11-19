package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
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

func (q *QuestionRepository) Create(questionEntity *question.Question) error {
	if err := database.DB.Create(questionEntity).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (q *QuestionRepository) FindById(id uuid.UUID) (*question.Question, error) {
	var questionEntity *question.Question
	if err := database.DB.Model(&question.Question{}).
		Where("id = ?", id).
		First(questionEntity).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return questionEntity, nil
}

//TODO
//func (q *QuestionRepository) FindById(id uuid.UUID) (question.IQuestion, error) {
//	switch iQuestion.GetType() {
//	case question.MULTIPLE_CHOICE:
//		question, err := q.multipleChoiceRepository.FindById(id)
//		if err != nil {
//
//		}
//	case question.SINGLE_CHOICE:
//		if err := q.singleChoiceRepository.Create(iQuestion.(*question.SingleChoice)); err != nil {
//			return err
//		}
//	case question.MATCHING:
//		if err := q.matchingQuestionRepository.Create(iQuestion.(*question.Matching)); err != nil {
//			return err
//		}
//	case question.TEXT_INPUT:
//		if err := q.textInputRepository.Create(iQuestion.(*question.TextInput)); err != nil {
//			return err
//		}
//	}
//	return nil
//}
