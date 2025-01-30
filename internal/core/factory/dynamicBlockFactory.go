package factory

import (
	"hedgehog-forms/internal/core/dto/create"
	block2 "hedgehog-forms/internal/core/model/form/pattern/section/block"
)

type DynamicBlockFactory struct {
	questionFactory *QuestionFactory
}

func NewDynamicBlockFactory() *DynamicBlockFactory {
	return &DynamicBlockFactory{
		questionFactory: NewQuestionFactory(),
	}
}

func (d *DynamicBlockFactory) buildFromDto(dynamicDto *create.DynamicBlockDto) (*block2.Block, error) {
	blockEntity := new(block2.Block)
	blockEntity.DynamicBlock = new(block2.DynamicBlock)
	questions, err := d.questionFactory.BuildQuestionDtoForDynamicBlock(dynamicDto.Questions, blockEntity)
	if err != nil {
		return nil, err
	}

	blockEntity.Title = dynamicDto.Title
	blockEntity.Description = dynamicDto.Description
	blockEntity.Type = block2.DYNAMIC
	blockEntity.DynamicBlock.QuestionCount = dynamicDto.QuestionCount
	blockEntity.DynamicBlock.Questions = questions
	return blockEntity, nil
}

func (d *DynamicBlockFactory) buildFromEntity(dynamicBlock *block2.Block) (*block2.Block, error) {
	newBlock := new(block2.Block)
	newBlock.DynamicBlock = new(block2.DynamicBlock)
	newQuestions, err := d.questionFactory.BuildQuestionEntityForDynamicBlock(dynamicBlock.DynamicBlock.Questions, newBlock)
	if err != nil {
		return nil, err
	}

	newBlock.Title = dynamicBlock.Title
	newBlock.Description = dynamicBlock.Description
	newBlock.Type = block2.DYNAMIC
	newBlock.DynamicBlock.QuestionCount = dynamicBlock.DynamicBlock.QuestionCount
	newBlock.DynamicBlock.Questions = newQuestions
	return newBlock, nil
}
