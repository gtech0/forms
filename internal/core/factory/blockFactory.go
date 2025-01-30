package factory

import (
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/errs"
	block2 "hedgehog-forms/internal/core/model/form/pattern/section/block"
	"hedgehog-forms/internal/core/repository"
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

func (b *BlockFactory) BuildFromDto(blockDto any) (*block2.Block, error) {
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

func (b *BlockFactory) buildFromEntity(blockEntity *block2.Block) (*block2.Block, error) {
	switch blockEntity.Type {
	case block2.DYNAMIC:
		return b.dynamicFactory.buildFromEntity(blockEntity)
	case block2.STATIC:
		return b.staticFactory.buildFromEntity(blockEntity)
	default:
		return nil, errs.New("invalid block entity type", 400)
	}
}

func (b *BlockFactory) buildFromExisting(dto *create.BlockOnExistingDto) (*block2.Block, error) {
	blockEntity, err := b.blockRepository.FindById(dto.BlockId)
	if err != nil {
		return nil, err
	}

	return b.buildFromEntity(blockEntity)
}
