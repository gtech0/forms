package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
	"hedgehog-forms/pkg/database"
)

type MatchingQuestionRepository struct{}

func NewMatchingQuestionRepository() *MatchingQuestionRepository {
	return &MatchingQuestionRepository{}
}

func (m *MatchingQuestionRepository) Create(matching *question.Matching) error {
	if err := database.DB.Create(matching).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (m *MatchingQuestionRepository) FindById(id uuid.UUID) (*question.Matching, error) {
	var matchingQuestion *question.Matching
	if err := database.DB.Model(&question.Matching{}).
		Where("id = ?", id).
		First(matchingQuestion).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return matchingQuestion, nil
}
