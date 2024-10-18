package factory

import (
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block"
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
	case *create.DynamicBlockDto:
		return b.dynamicFactory.buildFromDto(bl)
	case *create.StaticBlockDto:
		return b.staticFactory.buildFromDto(bl)
	case *create.BlockOnExistingDto:
		return b.buildFromExisting(bl)
	default:
		return nil, errs.New("invalid block dto type", 400)
	}
}

func (b *BlockFactory) buildFromObject(blockObj block.IBlock) (block.IBlock, error) {
	switch bl := blockObj.(type) {
	case *block.DynamicBlock:
		return b.dynamicFactory.buildFromObj(bl)
	case *block.StaticBlock:
		return b.staticFactory.buildFromObj(bl)
	default:
		return nil, errs.New("invalid block object type", 400)
	}
}

func (b *BlockFactory) buildFromExisting(dto *create.BlockOnExistingDto) (block.IBlock, error) {
	blockObj := new(block.Block)
	if err := database.DB.Model(&block.Block{}).
		Where("id = ?", dto.BlockId).
		First(blockObj).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}

	return b.buildFromObject(blockObj)
}
