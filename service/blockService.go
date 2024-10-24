package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block"
)

type BlockService struct{}

func NewBlockService() *BlockService {
	return &BlockService{}
}

func (b *BlockService) GetBlockObjectById(id uuid.UUID) (*block.Block, error) {
	blockObj := new(block.Block)
	if err := database.DB.Model(&block.Block{}).
		Where("id = ?", id).
		First(blockObj).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}

	return blockObj, nil
}
