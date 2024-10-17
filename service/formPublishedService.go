package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/model/form/published"
)

type FormPublishedService struct{}

func NewFormPublishedService() *FormPublishedService {
	return &FormPublishedService{}
}

func (i *FormPublishedService) GetForm(id uuid.UUID) (published.FormPublished, error) {
	var publishedForm published.FormPublished
	if err := database.DB.Preload("~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~").
		Model(&published.FormPublished{}).
		Where("id = ?", id).
		First(&publishedForm).
		Error; err != nil {
		return publishedForm, err
	}

	return publishedForm, nil
}