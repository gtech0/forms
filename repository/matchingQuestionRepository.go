package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type MatchingQuestionRepository struct{}

func NewMatchingQuestionRepository() *MatchingQuestionRepository {
	return &MatchingQuestionRepository{}
}

func (m *MatchingQuestionRepository) GetById(id uuid.UUID) (*question.Matching, error) {
	var matchingQuestion *question.Matching
	if err := database.DB.Model(&question.Matching{}).
		Where("id = ?", id).
		First(matchingQuestion).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return matchingQuestion, nil
}
