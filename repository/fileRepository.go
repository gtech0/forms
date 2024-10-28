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

func (f *FileRepository) Save(fileModel *model.File) error {
	if err := database.DB.Save(fileModel).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}
