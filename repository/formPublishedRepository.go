package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/published"
)

type FormPublishedRepository struct{}

func NewFormPublishedRepository() *FormPublishedRepository {
	return &FormPublishedRepository{}
}

func (f *FormPublishedRepository) Create(formPublished *published.FormPublished) error {
	if err := database.DB.Create(formPublished).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (f *FormPublishedRepository) FindById(id uuid.UUID) (*published.FormPublished, error) {
	formPublished := new(published.FormPublished)
	if err := database.DB.Preload("~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~").
		Model(&published.FormPublished{}).
		Where("id = ?", id).
		First(formPublished).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formPublished, nil
}
