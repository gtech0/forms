package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/pkg/database"
)

type SubmissionRepository struct{}

func NewSubmissionRepository() *SubmissionRepository {
	return &SubmissionRepository{}
}

func (s *SubmissionRepository) Create(submission *generated.Submission) error {
	if err := database.DB.Create(submission).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (s *SubmissionRepository) Save(submission *generated.Submission) error {
	if err := database.DB.Save(submission).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (s *SubmissionRepository) FindById(id uuid.UUID) (*generated.Submission, error) {
	submission := new(generated.Submission)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&generated.Submission{}).
		Where("id = ?", id).
		First(submission).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return submission, nil
}

func (s *SubmissionRepository) FindByFormGeneratedId(generatedId uuid.UUID) (*generated.Submission, error) {
	submission := new(generated.Submission)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&generated.Submission{}).
		Where("form_generated_id = ?", generatedId).
		First(submission).
		Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.New(err.Error(), 500)
	}
	return submission, nil
}
