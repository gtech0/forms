package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
)

type FormGeneratedRepository struct{}

func NewFormGeneratedRepository() *FormGeneratedRepository {
	return &FormGeneratedRepository{}
}

func (f *FormGeneratedRepository) Create(formGenerated *generated.FormGenerated) error {
	if err := database.DB.Create(formGenerated).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (f *FormGeneratedRepository) FindByPublishedId(publishedId uuid.UUID) (*generated.FormGenerated, error) {
	formGenerated := new(generated.FormGenerated)
	if err := database.DB.Model(&generated.FormGenerated{}).
		Where("form_published_id = ?", publishedId).
		Find(formGenerated).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formGenerated, nil
}
