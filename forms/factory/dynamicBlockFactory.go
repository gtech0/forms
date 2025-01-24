package factory

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern/section/block"
)

type DynamicBlockFactory struct {
	questionFactory *QuestionFactory
}

func NewDynamicBlockFactory() *DynamicBlockFactory {
	return &DynamicBlockFactory{
		questionFactory: NewQuestionFactory(),
	}
}

func (d *DynamicBlockFactory) buildFromDto(dynamicDto *create.DynamicBlockDto) (*block.Block, error) {
	blockEntity := new(block.Block)
	blockEntity.DynamicBlock = new(block.DynamicBlock)
	questions, err := d.questionFactory.BuildQuestionDtoForDynamicBlock(dynamicDto.Questions, blockEntity)
	if err != nil {
		return nil, err
	}

	blockEntity.Title = dynamicDto.Title
	blockEntity.Description = dynamicDto.Description
	blockEntity.Type = block.DYNAMIC
	blockEntity.DynamicBlock.QuestionCount = dynamicDto.QuestionCount
	blockEntity.DynamicBlock.Questions = questions
	return blockEntity, nil
}

func (d *DynamicBlockFactory) buildFromEntity(dynamicBlock *block.Block) (*block.Block, error) {
	newBlock := new(block.Block)
	newBlock.DynamicBlock = new(block.DynamicBlock)
	newQuestions, err := d.questionFactory.BuildQuestionEntityForDynamicBlock(dynamicBlock.DynamicBlock.Questions, newBlock)
	if err != nil {
		return nil, err
	}

	newBlock.Title = dynamicBlock.Title
	newBlock.Description = dynamicBlock.Description
	newBlock.Type = block.DYNAMIC
	newBlock.DynamicBlock.QuestionCount = dynamicBlock.DynamicBlock.QuestionCount
	newBlock.DynamicBlock.Questions = newQuestions
	return newBlock, nil
}
