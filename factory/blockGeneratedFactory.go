package factory

import (
	"errors"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section/block"
	"math/rand"
	"slices"
)

type BlockGeneratedFactory struct {
	questionGeneratedFactory *QuestionGeneratedFactory
	variantGeneratedFactory  *VariantGeneratedFactory
}

func NewBlockGeneratedFactory() *BlockGeneratedFactory {
	return &BlockGeneratedFactory{
		questionGeneratedFactory: NewQuestionGeneratedFactory(),
		variantGeneratedFactory:  NewVariantGeneratedFactory(),
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
		return b.buildDynamicBlock(assertedBlock)
	case *block.StaticBlock:
		return b.buildStaticBlock(assertedBlock)
	default:
		return generated.Block{}, errors.New("unsupported block type")
	}
}

func (b *BlockGeneratedFactory) buildDynamicBlock(dynamicBlock *block.DynamicBlock) (generated.Block, error) {
	questionCount := dynamicBlock.QuestionCount
	questions := dynamicBlock.Questions
	questionsForBlock := make([]generated.IQuestion, 0, questionCount)
	for i := 0; i < questionCount; i++ {
		randomIndex := rand.Intn(len(questions))
		randomQuestion := questions[randomIndex]
		generatedQuestion, err := b.questionGeneratedFactory.buildQuestion(randomQuestion)
		if err != nil {
			return generated.Block{}, err
		}

		questions = slices.Delete(questions, randomIndex, randomIndex+1)
		questionsForBlock = append(questionsForBlock, generatedQuestion)
	}

	var generatedBlock generated.Block
	generatedBlock.Id = dynamicBlock.Id
	generatedBlock.Type = dynamicBlock.Type
	generatedBlock.Title = dynamicBlock.Title
	generatedBlock.Description = dynamicBlock.Description
	generatedBlock.Questions = questionsForBlock
	return generatedBlock, nil
}

func (b *BlockGeneratedFactory) buildStaticBlock(staticBlock *block.StaticBlock) (generated.Block, error) {
	variant, err := b.variantGeneratedFactory.buildVariant(staticBlock.Variants)
	if err != nil {
		return generated.Block{}, err
	}

	return generated.Block{
		Id:          staticBlock.Id,
		Type:        staticBlock.Type,
		Title:       staticBlock.Title,
		Description: staticBlock.Description,
		Variant:     variant,
	}, nil
}
