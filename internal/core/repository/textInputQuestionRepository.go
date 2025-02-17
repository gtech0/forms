package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/question"
	"hedgehog-forms/pkg/database"
)

type TextInputRepository struct{}

func NewTextInputRepository() *TextInputRepository {
	return &TextInputRepository{}
}

func (t *TextInputRepository) Create(textInput *question.TextInput) error {
	if err := database.DB.Create(textInput).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (t *TextInputRepository) FindById(id uuid.UUID) (*question.TextInput, error) {
	var textInputQuestion *question.TextInput
	if err := database.DB.Model(&question.TextInput{}).
		Where("id = ?", id).
		First(textInputQuestion).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return textInputQuestion, nil
}
