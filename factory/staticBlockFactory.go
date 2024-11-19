package factory

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern/section/block"
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
	blockObj := new(block.Block)
	blockObj.StaticBlock = new(block.StaticBlock)
	variants, err := s.variantFactory.buildFromDtos(blockDto.Variants, blockObj.Id)
	if err != nil {
		return nil, err
	}

	blockObj.Title = blockDto.Title
	blockObj.Description = blockDto.Description
	blockObj.Type = block.STATIC
	blockObj.StaticBlock.Variants = variants
	return blockObj, nil
}

func (s *StaticBlockFactory) buildFromObj(blockObj *block.Block) (*block.Block, error) {
	newBlock := new(block.Block)
	newBlock.StaticBlock = new(block.StaticBlock)
	newVariants, err := s.variantFactory.buildFromObjs(blockObj.StaticBlock.Variants, newBlock.Id)
	if err != nil {
		return nil, err
	}

	newBlock.Title = blockObj.Title
	newBlock.Description = blockObj.Description
	newBlock.Type = block.STATIC
	newBlock.StaticBlock.Variants = newVariants
	return newBlock, nil
}
