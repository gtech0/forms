package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section"
)

type SectionRepository struct{}

func NewSectionRepository() *SectionRepository {
	return &SectionRepository{}
}

func (s *SectionRepository) FindById(id uuid.UUID) (*section.Section, error) {
	sectObj := new(section.Section)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&section.Section{}).
		Where("id = ?", id).
		First(sectObj).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return sectObj, nil
}

func (s *SectionRepository) FindByNameAndPaginate(name string, page, size int) ([]section.Section, error) {
	sections := make([]section.Section, 0)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&section.Section{}).
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Scopes(paginate(page, size)).
		Find(&sections).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return sections, nil
}
