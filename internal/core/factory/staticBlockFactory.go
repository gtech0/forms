package factory

import (
	"hedgehog-forms/internal/core/dto/create"
	block "hedgehog-forms/internal/core/model/form/pattern/block"
)

type StaticBlockFactory struct {
	questionFactory *QuestionFactory
	variantFactory  *VariantFactory
}

func NewStaticBlockFactory() *StaticBlockFactory {
	return &StaticBlockFactory{
		questionFactory: NewQuestionFactory(),
		variantFactory:  NewVariantFactory(),
	}
}

func (s *StaticBlockFactory) buildFromDto(blockDto *create.StaticBlockDto) (*block.Block, error) {
	blockEntity := new(block.Block)
	blockEntity.StaticBlock = new(block.StaticBlock)
	variants, err := s.variantFactory.buildFromDtos(blockDto.Variants, blockEntity.Id)
	if err != nil {
		return nil, err
	}

	blockEntity.Title = blockDto.Title
	blockEntity.Description = blockDto.Description
	blockEntity.Type = block.STATIC
	blockEntity.StaticBlock.Variants = variants
	return blockEntity, nil
}

func (s *StaticBlockFactory) buildFromEntity(blockEntity *block.Block) (*block.Block, error) {
	newBlock := new(block.Block)
	newBlock.StaticBlock = new(block.StaticBlock)
	newVariants, err := s.variantFactory.buildFromEntities(blockEntity.StaticBlock.Variants, newBlock.Id)
	if err != nil {
		return nil, err
	}

	newBlock.Title = blockEntity.Title
	newBlock.Description = blockEntity.Description
	newBlock.Type = block.STATIC
	newBlock.StaticBlock.Variants = newVariants
	return newBlock, nil
}
