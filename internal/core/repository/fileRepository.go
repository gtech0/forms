package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model"
	"hedgehog-forms/pkg/database"
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

func (f *FileRepository) FindById(id uuid.UUID) (*model.File, error) {
	file := new(model.File)
	if err := database.DB.Model(&model.File{}).
		Where("id = ?", id).
		Find(&file).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return file, nil
}
