package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
)

type FormGeneratedRepository struct{}

func NewFormGeneratedRepository() *FormGeneratedRepository {
	return &FormGeneratedRepository{}
}

func (f *FormGeneratedRepository) Save(formGenerated *generated.FormGenerated) error {
	if err := database.DB.Save(formGenerated).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (f *FormGeneratedRepository) FindById(generatedId uuid.UUID) (*generated.FormGenerated, error) {
	formGenerated := new(generated.FormGenerated)
	if err := database.DB.Model(&generated.FormGenerated{}).
		Where("id = ?", generatedId).
		First(formGenerated).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formGenerated, nil
}

func (f *FormGeneratedRepository) FindByPublishedId(publishedId uuid.UUID) (*generated.FormGenerated, error) {
	formGenerated := new(generated.FormGenerated)
	if err := database.DB.Model(&generated.FormGenerated{}).
		Where("form_published_id = ?", publishedId).
		First(formGenerated).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formGenerated, nil
}

func (f *FormGeneratedRepository) FindBySubjectIdAndPaginate(
	subjectId uuid.UUID,
	page int,
	size int,
) ([]generated.FormGenerated, error) {
	formsGenerated := make([]generated.FormGenerated, 0)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&generated.FormGenerated{}).
		Joins("inner join form_published on form_published_id = form_published.id").
		Joins("inner join form_pattern on form_published.form_pattern_id = form_pattern.id").
		Where("form_pattern.subject_id = ?", subjectId).
		Scopes(paginate(page, size)).
		Find(&formsGenerated).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formsGenerated, nil
}

func (f *FormGeneratedRepository) FindByPublishedIdAndStatusAndPaginate(
	publishedId uuid.UUID,
	status generated.FormStatus,
	page int,
	size int,
) ([]generated.FormGenerated, error) {
	formsGenerated := make([]generated.FormGenerated, 0)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&generated.FormGenerated{}).
		Joins("inner join form_published on form_published_id = form_published.id").
		Where("form_published.id = ?", publishedId).
		Where("status = ?", status).
		Scopes(paginate(page, size)).
		Find(&formsGenerated).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formsGenerated, nil
}
