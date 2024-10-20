package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern"
)

type FormPatternRepository struct{}

func NewFormPatternRepository() *FormPatternRepository {
	return &FormPatternRepository{}
}

func (f *FormPatternRepository) Create(formPattern *pattern.FormPattern) error {
	if err := database.DB.Create(formPattern).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (f *FormPatternRepository) FindById(id uuid.UUID) (*pattern.FormPattern, error) {
	formPattern := new(pattern.FormPattern)
	if err := database.DB.Model(&pattern.FormPattern{}).
		Preload("Subject").
		Preload("Sections.DynamicBlocks.~~~as~~~.~~~as~~~.~~~as~~~").
		Preload("Sections.StaticBlocks.Variants.~~~as~~~.~~~as~~~.~~~as~~~").
		First(formPattern, "form_pattern.id = ?", id).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formPattern, nil
}
