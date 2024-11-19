package factory

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/repository"
)

type BlockFactory struct {
	dynamicFactory  *DynamicBlockFactory
	staticFactory   *StaticBlockFactory
	blockRepository *repository.BlockRepository
}

func NewBlockFactory() *BlockFactory {
	return &BlockFactory{
		dynamicFactory:  NewDynamicBlockFactory(),
		staticFactory:   NewStaticBlockFactory(),
		blockRepository: repository.NewBlockRepository(),
	}
}

func (b *BlockFactory) BuildFromDto(blockDto any) (*block.Block, error) {
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

func (b *BlockFactory) buildFromObject(blockObj *block.Block) (*block.Block, error) {
	switch blockObj.Type {
	case block.DYNAMIC:
		return b.dynamicFactory.buildFromObj(blockObj)
	case block.STATIC:
		return b.staticFactory.buildFromObj(blockObj)
	default:
		return nil, errs.New("invalid block object type", 400)
	}
}

func (b *BlockFactory) buildFromExisting(dto *create.BlockOnExistingDto) (*block.Block, error) {
	blockObj, err := b.blockRepository.FindById(dto.BlockId)
	if err != nil {
		return nil, err
	}

	return b.buildFromObject(blockObj)
}
