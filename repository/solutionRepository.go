package repository

import (
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/published"
)

type SolutionRepository struct{}

func NewSolutionRepository() *SolutionRepository {
	return &SolutionRepository{}
}

func (f *SolutionRepository) Save(solution *published.Solution) error {
	if err := database.DB.Save(solution).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}
