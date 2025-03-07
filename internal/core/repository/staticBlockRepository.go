package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/block"
	"hedgehog-forms/pkg/database"
)

type StaticBlockRepository struct{}

func NewStaticBlockRepository() *StaticBlockRepository {
	return &StaticBlockRepository{}
}

func (b *StaticBlockRepository) FindById(id uuid.UUID) (*block.StaticBlock, error) {
	staticBlock := new(block.StaticBlock)
	if err := database.DB.Model(&block.StaticBlock{}).
		Where("id = ?", id).
		First(staticBlock).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return staticBlock, nil
}
