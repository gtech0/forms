package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern"
	"hedgehog-forms/pkg/database"
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
		Preload(clause.Associations, preload).
		First(formPattern, "form_pattern.id = ?", id).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formPattern, nil
}

func (f *FormPatternRepository) FindAndPaginate(
	name string,
	clauses []clause.Expression,
	page int,
	size int,
) ([]pattern.FormPattern, error) {
	formPatterns := make([]pattern.FormPattern, 0)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&pattern.FormPattern{}).
		Clauses(clauses...).
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Scopes(paginate(page, size)).
		Find(&formPatterns).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formPatterns, nil
}
