package factory

import (
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/repository"
)

type BlockFactory struct {
	dynamicFactory    *DynamicBlockFactory
	dynamicRepository *repository.DynamicBlockRepository
	staticFactory     *StaticBlockFactory
	staticRepository  *repository.StaticBlockRepository
}

func NewBlockFactory() *BlockFactory {
	return &BlockFactory{
		dynamicFactory:    NewDynamicBlockFactory(),
		dynamicRepository: repository.NewDynamicBlockRepository(),
		staticFactory:     NewStaticBlockFactory(),
		staticRepository:  repository.NewStaticBlockRepository(),
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
	var blockObj block.IBlock
	switch dto.Type {
	case block.STATIC:
		staticBlock, err := b.staticRepository.GetById(dto.BlockId)
		if err != nil {
			return nil, err
		}
		blockObj = staticBlock
	case block.DYNAMIC:
		dynamicBlock, err := b.dynamicRepository.GetById(dto.BlockId)
		if err != nil {
			return nil, err
		}
		blockObj = dynamicBlock
	}

	if err := database.DB.Model(&block.Block{}).
		Where("id = ?", dto.BlockId).
		First(blockObj).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}

	return b.buildFromObject(blockObj)
}
