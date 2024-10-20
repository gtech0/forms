package repository

import (
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model"
)

type SubjectRepository struct{}

func NewSubjectRepository() *SubjectRepository {
	return &SubjectRepository{}
}

func (s *SubjectRepository) Create(subject model.Subject) error {
	if err := database.DB.Create(&subject).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}
