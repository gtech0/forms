package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type SingleChoiceRepository struct{}

func NewSingleChoiceRepository() *SingleChoiceRepository {
	return &SingleChoiceRepository{}
}

func (t *SingleChoiceRepository) GetById(id uuid.UUID) (*question.SingleChoice, error) {
	var singleChoiceQuestion *question.SingleChoice
	if err := database.DB.Model(&question.SingleChoice{}).
		Where("id = ?", id).
		First(singleChoiceQuestion).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return singleChoiceQuestion, nil
}
