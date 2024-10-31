package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/published"
)

type FormPublishedRepository struct{}

func NewFormPublishedRepository() *FormPublishedRepository {
	return &FormPublishedRepository{}
}

func (f *FormPublishedRepository) Save(formPublished *published.FormPublished) error {
	if err := database.DB.Save(formPublished).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (f *FormPublishedRepository) FindById(id uuid.UUID) (*published.FormPublished, error) {
	formPublished := new(published.FormPublished)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&published.FormPublished{}).
		Where("id = ?", id).
		First(formPublished).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formPublished, nil
}

func (f *FormPublishedRepository) FindByNameAndPaginate(name string, page int, size int) ([]published.FormPublished, error) {
	formsPublished := make([]published.FormPublished, 0)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&published.FormPublished{}).
		Joins("inner join form_pattern on form_pattern_id = form_pattern.id").
		Where("form_pattern.title LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Scopes(paginate(page, size)).
		Find(&formsPublished).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formsPublished, nil
}

func (f *FormPublishedRepository) DeleteById(publishedId uuid.UUID) error {
	if err := database.DB.Delete(&published.FormPublished{}, publishedId).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}
