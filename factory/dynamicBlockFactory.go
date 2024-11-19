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
	blockObj := new(block.Block)
	blockObj.DynamicBlock = new(block.DynamicBlock)
	questions, err := d.questionFactory.BuildQuestionDtoForDynamicBlock(dynamicDto.Questions, blockObj)
	if err != nil {
		return nil, err
	}

	blockObj.Title = dynamicDto.Title
	blockObj.Description = dynamicDto.Description
	blockObj.Type = block.DYNAMIC
	blockObj.DynamicBlock.QuestionCount = dynamicDto.QuestionCount
	blockObj.DynamicBlock.Questions = questions
	return blockObj, nil
}

func (d *DynamicBlockFactory) buildFromObj(dynamicBlock *block.Block) (*block.Block, error) {
	newBlockObj := new(block.Block)
	newBlockObj.DynamicBlock = new(block.DynamicBlock)
	newQuestions, err := d.questionFactory.BuildQuestionObjForDynamicBlock(dynamicBlock.DynamicBlock.Questions, newBlockObj)
	if err != nil {
		return nil, err
	}

	newBlockObj.Title = dynamicBlock.Title
	newBlockObj.Description = dynamicBlock.Description
	newBlockObj.Type = block.DYNAMIC
	newBlockObj.DynamicBlock.QuestionCount = dynamicBlock.DynamicBlock.QuestionCount
	newBlockObj.DynamicBlock.Questions = newQuestions
	return newBlockObj, nil
}
