package block

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/factory/question"
	"hedgehog-forms/model/form/section/block"
)

type StaticBlockFactory struct {
	questionFactory *question.QuestionFactory
	variantFactory  *VariantFactory
}

func NewStaticBlockFactory() *StaticBlockFactory {
	return &StaticBlockFactory{
		questionFactory: question.NewQuestionFactory(),
		variantFactory:  NewVariantFactory(),
	}
}

func (s *StaticBlockFactory) buildFromDto(blockDto *create.StaticBlockDto) (*block.StaticBlock, error) {
	blockObj := new(block.StaticBlock)
	variants, err := s.variantFactory.buildFromDtos(blockDto.Variants, blockObj)
	if err != nil {
		return nil, err
	}

	blockObj.Title = blockDto.Title
	blockObj.Description = blockDto.Description
	blockObj.Type = block.STATIC
	blockObj.Variants = variants
	return blockObj, nil
}

func (s *StaticBlockFactory) buildFromObj(blockObj *block.StaticBlock) (*block.StaticBlock, error) {
	newBlock := new(block.StaticBlock)
	newVariants, err := s.variantFactory.buildFromObjs(blockObj.Variants, newBlock)
	if err != nil {
		return nil, err
	}

	newBlock.Title = blockObj.Title
	newBlock.Description = blockObj.Description
	newBlock.Type = block.STATIC
	newBlock.Variants = newVariants
	return newBlock, nil
}
