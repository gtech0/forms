package factory

import (
	"errors"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section/block"
)

type BlockGeneratedFactory struct {
	questionGeneratedFactory *QuestionGeneratedFactory
}

func NewBlockGeneratedFactory() *BlockGeneratedFactory {
	return &BlockGeneratedFactory{
		questionGeneratedFactory: NewQuestionGeneratedFactory(),
	}
}

func (b *BlockGeneratedFactory) buildBlocks(iBlocks []block.IBlock) ([]generated.Block, error) {
	var generatedBlocks []generated.Block
	for _, iBlock := range iBlocks {
		newBlock, err := b.buildBlock(iBlock)
		if err != nil {
			return nil, err
		}

		generatedBlocks = append(generatedBlocks, newBlock)
	}
	return generatedBlocks, nil
}

func (b *BlockGeneratedFactory) buildBlock(iBlock block.IBlock) (generated.Block, error) {
	switch assertedBlock := iBlock.(type) {
	case *block.DynamicBlock:
		return b.buildDynamicBlock(assertedBlock), nil
	case *block.StaticBlock:
		return b.buildStaticBlock(assertedBlock), nil
	default:
		return generated.Block{}, errors.New("unsupported block type")
	}
}

func (b *BlockGeneratedFactory) buildDynamicBlock(dynamicBlock *block.DynamicBlock) generated.Block {
	//TODO
	//questionCount := dynamicBlock.QuestionCount
	//questions := dynamicBlock.Questions
	//questionsForBlock := make([]generated.IQuestion, 0, questionCount)
	return generated.Block{}
}

func (b *BlockGeneratedFactory) buildStaticBlock(staticBlock *block.StaticBlock) generated.Block {
	//TODO
	return generated.Block{}
}
