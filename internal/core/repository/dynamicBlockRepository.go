package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/section/block"
	"hedgehog-forms/pkg/database"
)

type DynamicBlockRepository struct{}

func NewDynamicBlockRepository() *DynamicBlockRepository {
	return &DynamicBlockRepository{}
}

func (b *DynamicBlockRepository) FindById(id uuid.UUID) (*block.DynamicBlock, error) {
	dynamicBlock := new(block.DynamicBlock)
	if err := database.DB.Model(&block.DynamicBlock{}).
		Where("id = ?", id).
		First(dynamicBlock).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return dynamicBlock, nil
}
