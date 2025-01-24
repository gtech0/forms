package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/published"
)

type SolutionRepository struct{}

func NewSolutionRepository() *SolutionRepository {
	return &SolutionRepository{}
}

func (f *SolutionRepository) Create(solution *published.Solution) error {
	if err := database.DB.Create(solution).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (f *SolutionRepository) Save(solution *published.Solution) error {
	if err := database.DB.Save(solution).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (f *SolutionRepository) FindById(id uuid.UUID) (*published.Solution, error) {
	solution := new(published.Solution)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&published.Solution{}).
		Where("id = ?", id).
		First(solution).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return solution, nil
}

func (f *SolutionRepository) FindByTaskIdAndUserId(taskId uuid.UUID, userId uuid.UUID) (*published.Solution, error) {
	solution := new(published.Solution)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&published.Solution{}).
		Where("class_task_id = ?", taskId).
		Where("user_owner_id = ?", userId).
		First(solution).
		Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errs.New(err.Error(), 500)
	}
	return solution, nil
}
