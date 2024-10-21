package repository

import (
	"github.com/google/uuid"
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

func (s *SubjectRepository) GetById(id uuid.UUID) (*model.Subject, error) {
	subject := new(model.Subject)
	if err := database.DB.Model(&model.Subject{}).
		First(subject, "id = ?", id).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return subject, nil
}

func (s *SubjectRepository) Update(id uuid.UUID, name string) error {
	if err := database.DB.Model(&model.Subject{}).
		Where("id = ?", id).
		Update("name", name).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (s *SubjectRepository) Delete(id uuid.UUID) error {
	if err := database.DB.Delete(&model.Subject{}, id).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}
