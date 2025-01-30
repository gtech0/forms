package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hedgehog-forms/internal/core/errs"
	generated2 "hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/pkg/database"
)

type FormGeneratedRepository struct{}

func NewFormGeneratedRepository() *FormGeneratedRepository {
	return &FormGeneratedRepository{}
}

func (f *FormGeneratedRepository) Create(formGenerated *generated2.FormGenerated) error {
	if err := database.DB.Create(formGenerated).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (f *FormGeneratedRepository) Save(formGenerated *generated2.FormGenerated) error {
	if err := database.DB.Save(formGenerated).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (f *FormGeneratedRepository) FindById(generatedId uuid.UUID) (*generated2.FormGenerated, error) {
	formGenerated := new(generated2.FormGenerated)
	if err := database.DB.Model(&generated2.FormGenerated{}).
		Where("id = ?", generatedId).
		First(formGenerated).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formGenerated, nil
}

func (f *FormGeneratedRepository) FindByPublishedId(publishedId uuid.UUID) (*generated2.FormGenerated, error) {
	formGenerated := new(generated2.FormGenerated)
	if err := database.DB.Model(&generated2.FormGenerated{}).
		Where("form_published_id = ?", publishedId).
		First(formGenerated).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formGenerated, nil
}

func (f *FormGeneratedRepository) FindBySubjectIdAndPaginate(
	userId,
	subjectId uuid.UUID,
	page int,
	size int,
) ([]generated2.FormGenerated, error) {
	formsGenerated := make([]generated2.FormGenerated, 0)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&generated2.FormGenerated{}).
		Joins("inner join form_published on form_published_id = form_published.id").
		Joins("inner join form_pattern on form_published.form_pattern_id = form_pattern.id").
		Where("form_pattern.user_id = ?", userId).
		Where("form_pattern.subject_id = ?", subjectId).
		Scopes(paginate(page, size)).
		Find(&formsGenerated).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formsGenerated, nil
}

func (f *FormGeneratedRepository) FindByPublishedIdAndStatusAndPaginate(
	userId,
	publishedId uuid.UUID,
	status generated2.FormStatus,
	page int,
	size int,
) ([]generated2.FormGenerated, error) {
	formsGenerated := make([]generated2.FormGenerated, 0)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&generated2.FormGenerated{}).
		Joins("inner join form_published on form_published_id = form_published.id").
		Joins("inner join form_pattern on form_published.form_pattern_id = form_pattern.id").
		Where("form_pattern.owner_id = ?", userId).
		Where("form_published.id = ?", publishedId).
		Where("form_generated.status = ?", status).
		Scopes(paginate(page, size)).
		Find(&formsGenerated).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return formsGenerated, nil
}

func (f *FormGeneratedRepository) FindAttemptsByUserAndPublished(
	userId,
	publishedId uuid.UUID,
) ([]*generated2.FormGenerated, error) {
	attempts := make([]*generated2.FormGenerated, 0)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&generated2.FormGenerated{}).
		Where("user_id = ?", userId).
		Where("form_published_id = ?", publishedId).
		Find(&attempts).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return attempts, nil
}
