package block

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/factory/question"
	block2 "hedgehog-forms/model/form/pattern/section/block"
)

type DynamicBlockFactory struct {
	questionFactory *question.QuestionFactory
}

func NewDynamicBlockFactory() *DynamicBlockFactory {
	return &DynamicBlockFactory{
		questionFactory: question.NewQuestionFactory(),
	}
}

func (d *DynamicBlockFactory) buildFromDto(dynamicDto *create.DynamicBlockDto) (*block2.DynamicBlock, error) {
	blockObj := new(block2.DynamicBlock)
	questions, err := d.questionFactory.BuildQuestionDtoForDynamicBlock(dynamicDto.Questions, blockObj)
	if err != nil {
		return nil, err
	}

	blockObj.Title = dynamicDto.Title
	blockObj.Description = dynamicDto.Description
	blockObj.Type = block2.DYNAMIC
	blockObj.Questions = questions
	return blockObj, nil
}

func (d *DynamicBlockFactory) buildFromObj(dynamicBlock *block2.DynamicBlock) (*block2.DynamicBlock, error) {
	newBlockObj := new(block2.DynamicBlock)
	newQuestions, err := d.questionFactory.BuildQuestionObjForDynamicBlock(dynamicBlock.Questions, newBlockObj)
	if err != nil {
		return nil, err
	}

	newBlockObj.Title = dynamicBlock.Title
	newBlockObj.Description = dynamicBlock.Description
	newBlockObj.Type = block2.DYNAMIC
	newBlockObj.Questions = newQuestions
	return newBlockObj, nil
}
