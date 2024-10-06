package block

import (
	"errors"
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/service"
)

type BlockFactory struct {
	dynamicFactory *DynamicBlockFactory
	staticFactory  *StaticBlockFactory
}

func NewBlockFactory() *BlockFactory {
	return &BlockFactory{
		dynamicFactory: NewDynamicBlockFactory(),
		staticFactory:  NewStaticBlockFactory(),
	}
}

func (b *BlockFactory) BuildFromDto(blockDto any) (block.IBlock, error) {
	switch bl := blockDto.(type) {
	case *dto.CreateDynamicBlockDto:
		return b.dynamicFactory.buildFromDto(bl)
	case *dto.CreateStaticBlockDto:
		return b.staticFactory.buildFromDto(bl)
	case *dto.CreateBlockOnExistingDto:
		return b.buildFromExisting(bl)
	default:
		return nil, errors.New("unidentified block dto type")
	}
}

func (b *BlockFactory) buildFromObject(blockObj block.IBlock) (block.IBlock, error) {
	switch bl := blockObj.(type) {
	case *block.DynamicBlock:
		return b.dynamicFactory.buildFromObj(bl)
	case *block.StaticBlock:
		return b.staticFactory.buildFromObj(bl)
	default:
		return nil, errors.New("unidentified block object type")
	}
}

func (b *BlockFactory) buildFromExisting(dto *dto.CreateBlockOnExistingDto) (block.IBlock, error) {
	blockObj, err := service.NewBlockService().GetBlockObjectById(dto.BlockId)
	if err != nil {
		return nil, err
	}

	return b.buildFromObject(blockObj)
}