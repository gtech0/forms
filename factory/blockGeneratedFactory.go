package factory

import (
	"hedgehog-forms/errs"
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

func (b *BlockGeneratedFactory) buildBlocks(blocks []*block.Block) ([]*generated.Block, error) {
	generatedBlocks := make([]*generated.Block, 0)
	for _, iBlock := range blocks {
		newBlock, err := b.buildBlock(iBlock)
		if err != nil {
			return nil, err
		}

		generatedBlocks = append(generatedBlocks, newBlock)
	}
	return generatedBlocks, nil
}

func (b *BlockGeneratedFactory) buildBlock(iBlock *block.Block) (*generated.Block, error) {
	switch iBlock.Type {
	case block.DYNAMIC:
		return b.buildDynamicBlock(iBlock)
	case block.STATIC:
		return b.buildStaticBlock(iBlock)
	default:
		return nil, errs.New("unsupported block type", 400)
	}
}

func (b *BlockGeneratedFactory) buildDynamicBlock(dynamicBlock *block.Block) (*generated.Block, error) {
	questionCount := dynamicBlock.DynamicBlock.QuestionCount
	questions := dynamicBlock.DynamicBlock.Questions
	questionsForBlock := make([]generated.IQuestion, 0, questionCount)
	for i := 0; i < questionCount; i++ {
		randomIndex := rand.Intn(len(questions))
		randomQuestion := questions[randomIndex]
		generatedQuestion, err := b.questionGeneratedFactory.buildQuestion(randomQuestion)
		if err != nil {
			return nil, err
		}

		questions = slices.Delete(questions, randomIndex, randomIndex+1)
		questionsForBlock = append(questionsForBlock, generatedQuestion)
	}

	return &generated.Block{
		Id:          dynamicBlock.Id,
		Type:        dynamicBlock.Type,
		Title:       dynamicBlock.Title,
		Description: dynamicBlock.Description,
		Questions:   questionsForBlock,
	}, nil
}

func (b *BlockGeneratedFactory) buildStaticBlock(staticBlock *block.Block) (*generated.Block, error) {
	variant, err := b.variantGeneratedFactory.buildVariant(staticBlock.StaticBlock.Variants)
	if err != nil {
		return nil, err
	}

	return &generated.Block{
		Id:          staticBlock.Id,
		Type:        staticBlock.Type,
		Title:       staticBlock.Title,
		Description: staticBlock.Description,
		Variant:     variant,
	}, nil
}
