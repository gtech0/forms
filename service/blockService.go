package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/model/form/section/block"
)

type BlockService struct{}

func NewBlockService() *BlockService {
	return &BlockService{}
}

func (b *BlockService) GetBlockObjectById(id uuid.UUID) (*block.Block, error) {
	var obj *block.Block
	if err := database.DB.Model(block.Block{}).
		Where("id = ?", id).
		First(obj).Error; err != nil {
		return nil, err
	}

	return obj, nil
}
