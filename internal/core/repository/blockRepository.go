package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/section/block"
	"hedgehog-forms/pkg/database"
)

type BlockRepository struct{}

func NewBlockRepository() *BlockRepository {
	return &BlockRepository{}
}

func (b *BlockRepository) FindById(id uuid.UUID) (*block.Block, error) {
	newBlock := new(block.Block)
	if err := database.DB.Model(&block.Block{}).
		Where("id = ?", id).
		First(newBlock).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return newBlock, nil
}
