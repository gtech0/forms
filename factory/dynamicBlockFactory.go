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

func (d *DynamicBlockFactory) buildFromDto(dynamicDto *create.DynamicBlockDto) (*block.DynamicBlock, error) {
	blockObj := new(block.DynamicBlock)
	questions, err := d.questionFactory.BuildQuestionDtoForDynamicBlock(dynamicDto.Questions, blockObj)
	if err != nil {
		return nil, err
	}

	blockObj.Title = dynamicDto.Title
	blockObj.Description = dynamicDto.Description
	blockObj.Type = block.DYNAMIC
	blockObj.QuestionCount = dynamicDto.QuestionCount
	blockObj.Questions = questions
	return blockObj, nil
}

func (d *DynamicBlockFactory) buildFromObj(dynamicBlock *block.DynamicBlock) (*block.DynamicBlock, error) {
	newBlockObj := new(block.DynamicBlock)
	newQuestions, err := d.questionFactory.BuildQuestionObjForDynamicBlock(dynamicBlock.Questions, newBlockObj)
	if err != nil {
		return nil, err
	}

	newBlockObj.Title = dynamicBlock.Title
	newBlockObj.Description = dynamicBlock.Description
	newBlockObj.Type = block.DYNAMIC
	newBlockObj.QuestionCount = dynamicBlock.QuestionCount
	newBlockObj.Questions = newQuestions
	return newBlockObj, nil
}
