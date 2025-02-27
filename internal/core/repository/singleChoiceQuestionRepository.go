package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/question"
	"hedgehog-forms/pkg/database"
)

type SingleChoiceRepository struct{}

func NewSingleChoiceRepository() *SingleChoiceRepository {
	return &SingleChoiceRepository{}
}

func (t *SingleChoiceRepository) Create(singleChoice *question.SingleChoice) error {
	if err := database.DB.Create(singleChoice).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (t *SingleChoiceRepository) FindById(id uuid.UUID) (*question.SingleChoice, error) {
	var singleChoiceQuestion *question.SingleChoice
	if err := database.DB.Model(&question.SingleChoice{}).
		Where("id = ?", id).
		First(singleChoiceQuestion).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return singleChoiceQuestion, nil
}
