package repository

import (
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model"
)

type FileRepository struct{}

func NewFileRepository() *FileRepository {
	return &FileRepository{}
}

func (f *FileRepository) Create(fileModel *model.File) error {
	if err := database.DB.Create(fileModel).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}
