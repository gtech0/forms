package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type TextInputRepository struct{}

func NewTextInputRepository() *TextInputRepository {
	return &TextInputRepository{}
}

func (t *TextInputRepository) GetById(id uuid.UUID) (*question.TextInput, error) {
	var textInputQuestion *question.TextInput
	if err := database.DB.Model(&question.TextInput{}).
		Where("id = ?", id).
		First(textInputQuestion).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return textInputQuestion, nil
}
