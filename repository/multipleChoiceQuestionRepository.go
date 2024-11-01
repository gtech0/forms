package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type MultipleChoiceRepository struct{}

func NewMultipleChoiceRepository() *MultipleChoiceRepository {
	return &MultipleChoiceRepository{}
}

func (m *MultipleChoiceRepository) Create(multipleChoice *question.MultipleChoice) error {
	if err := database.DB.Create(multipleChoice).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (m *MultipleChoiceRepository) FindById(id uuid.UUID) (*question.MultipleChoice, error) {
	var multipleChoiceQuestion *question.MultipleChoice
	if err := database.DB.Model(&question.MultipleChoice{}).
		Where("id = ?", id).
		First(multipleChoiceQuestion).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return multipleChoiceQuestion, nil
}
