package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section"
)

type SectionRepository struct{}

func NewSectionRepository() *SectionRepository {
	return &SectionRepository{}
}

func (i *SectionRepository) GetById(id uuid.UUID) (*section.Section, error) {
	sectObj := new(section.Section)
	if err := database.DB.Model(&section.Section{}).
		Where("id = ?", id).
		First(sectObj).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return sectObj, nil
}
