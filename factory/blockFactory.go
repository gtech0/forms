package factory

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/service"
)

type BlockFactory struct{}

func (b *BlockFactory) buildFromDto(dto dto.CreateBlockDto) block.Block {
	switch dto.Type {
	case block.DYNAMIC:
	case block.STATIC:
	case block.EXISTING:

	default:

	}
}

func (b *BlockFactory) buildFromObject(object block.Block) (block.Block, error) {
	switch object.Type {
	case block.DYNAMIC:
	case block.STATIC:

	default:

	}
}

func (b *BlockFactory) buildFromExisting(dto dto.CreateBlockOnExistingDto) (block.Block, error) {
	blockObj, err := service.NewBlockService().GetBlockObjectById(dto.BlockId)
	if err != nil {
		return block.Block{}, err
	}

	return b.buildFromObject(blockObj)
}
