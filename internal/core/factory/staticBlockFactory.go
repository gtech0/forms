package factory

import (
	"hedgehog-forms/internal/core/dto/create"
	block2 "hedgehog-forms/internal/core/model/form/pattern/section/block"
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

func (s *StaticBlockFactory) buildFromDto(blockDto *create.StaticBlockDto) (*block2.Block, error) {
	blockEntity := new(block2.Block)
	blockEntity.StaticBlock = new(block2.StaticBlock)
	variants, err := s.variantFactory.buildFromDtos(blockDto.Variants, blockEntity.Id)
	if err != nil {
		return nil, err
	}

	blockEntity.Title = blockDto.Title
	blockEntity.Description = blockDto.Description
	blockEntity.Type = block2.STATIC
	blockEntity.StaticBlock.Variants = variants
	return blockEntity, nil
}

func (s *StaticBlockFactory) buildFromEntity(blockEntity *block2.Block) (*block2.Block, error) {
	newBlock := new(block2.Block)
	newBlock.StaticBlock = new(block2.StaticBlock)
	newVariants, err := s.variantFactory.buildFromEntities(blockEntity.StaticBlock.Variants, newBlock.Id)
	if err != nil {
		return nil, err
	}

	newBlock.Title = blockEntity.Title
	newBlock.Description = blockEntity.Description
	newBlock.Type = block2.STATIC
	newBlock.StaticBlock.Variants = newVariants
	return newBlock, nil
}
