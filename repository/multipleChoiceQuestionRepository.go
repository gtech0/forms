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

func (m *MultipleChoiceRepository) GetById(id uuid.UUID) (*question.MultipleChoice, error) {
	var multipleChoiceQuestion *question.MultipleChoice
	if err := database.DB.Model(&question.MultipleChoice{}).
		Where("id = ?", id).
		First(multipleChoiceQuestion).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return multipleChoiceQuestion, nil
}
