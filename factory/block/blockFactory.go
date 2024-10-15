package block

import (
	"errors"
	"hedgehog-forms/dto/create"
	block2 "hedgehog-forms/model/form/pattern/section/block"
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

func (b *BlockFactory) BuildFromDto(blockDto any) (block2.IBlock, error) {
	switch bl := blockDto.(type) {
	case *create.DynamicBlockDto:
		return b.dynamicFactory.buildFromDto(bl)
	case *create.StaticBlockDto:
		return b.staticFactory.buildFromDto(bl)
	case *create.BlockOnExistingDto:
		return b.buildFromExisting(bl)
	default:
		return nil, errors.New("unidentified block dto type")
	}
}

func (b *BlockFactory) buildFromObject(blockObj block2.IBlock) (block2.IBlock, error) {
	switch bl := blockObj.(type) {
	case *block2.DynamicBlock:
		return b.dynamicFactory.buildFromObj(bl)
	case *block2.StaticBlock:
		return b.staticFactory.buildFromObj(bl)
	default:
		return nil, errors.New("unidentified block object type")
	}
}

func (b *BlockFactory) buildFromExisting(dto *create.BlockOnExistingDto) (block2.IBlock, error) {
	blockObj, err := service.NewBlockService().GetBlockObjectById(dto.BlockId)
	if err != nil {
		return nil, err
	}

	return b.buildFromObject(blockObj)
}
