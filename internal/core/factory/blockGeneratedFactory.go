package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/section/block"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
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

func (b *BlockGeneratedFactory) buildBlocks(
	blocks []*block.Block,
	excluded []uuid.UUID,
) ([]*generated.Block, error) {
	generatedBlocks := make([]*generated.Block, 0)
	for _, iBlock := range blocks {
		newBlock, err := b.buildBlock(iBlock, excluded)
		if err != nil {
			return nil, err
		}

		generatedBlocks = append(generatedBlocks, newBlock)
	}
	return generatedBlocks, nil
}

func (b *BlockGeneratedFactory) buildBlock(
	iBlock *block.Block,
	excluded []uuid.UUID,
) (*generated.Block, error) {
	switch iBlock.Type {
	case block.DYNAMIC:
		return b.buildDynamicBlock(iBlock, excluded)
	case block.STATIC:
		return b.buildStaticBlock(iBlock)
	default:
		return nil, errs.New("unsupported block type", 400)
	}
}

func (b *BlockGeneratedFactory) buildDynamicBlock(
	dynamicBlock *block.Block,
	excluded []uuid.UUID,
) (*generated.Block, error) {
	questionCount := dynamicBlock.DynamicBlock.QuestionCount
	questions := make([]*question.Question, 0)
	for _, currQuestion := range dynamicBlock.DynamicBlock.Questions {
		if !slices.Contains(excluded, currQuestion.Id) {
			questions = append(questions, currQuestion)
		}
	}

	questionsForBlock := make([]generated.IQuestion, 0)

	if questionCount > len(questions) {
		return nil, errs.New("no questions in the bank", 500)
	}

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
