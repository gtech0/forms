package block

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/factory/question"
	block2 "hedgehog-forms/model/form/pattern/section/block"
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

func (s *StaticBlockFactory) buildFromDto(blockDto *create.StaticBlockDto) (*block2.StaticBlock, error) {
	blockObj := new(block2.StaticBlock)
	variants, err := s.variantFactory.buildFromDtos(blockDto.Variants, blockObj)
	if err != nil {
		return nil, err
	}

	blockObj.Title = blockDto.Title
	blockObj.Description = blockDto.Description
	blockObj.Type = block2.STATIC
	blockObj.Variants = variants
	return blockObj, nil
}

func (s *StaticBlockFactory) buildFromObj(blockObj *block2.StaticBlock) (*block2.StaticBlock, error) {
	newBlock := new(block2.StaticBlock)
	newVariants, err := s.variantFactory.buildFromObjs(blockObj.Variants, newBlock)
	if err != nil {
		return nil, err
	}

	newBlock.Title = blockObj.Title
	newBlock.Description = blockObj.Description
	newBlock.Type = block2.STATIC
	newBlock.Variants = newVariants
	return newBlock, nil
}
