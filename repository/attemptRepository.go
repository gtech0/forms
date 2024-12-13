package repository

import (
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
)

type AttemptRepository struct{}

func NewAttemptRepository() *AttemptRepository {
	return &AttemptRepository{}
}

func (a *AttemptRepository) Save(attempt *generated.Attempt) error {
	if err := database.DB.Save(attempt).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}
