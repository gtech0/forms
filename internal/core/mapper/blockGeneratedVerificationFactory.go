package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/get"
	verify2 "hedgehog-forms/internal/core/dto/verify"
	"hedgehog-forms/internal/core/model/form/generated"
)

type BlockGeneratedVerificationFactory struct {
	questionGeneratedVerificationFactory *QuestionGeneratedVerificationFactory
	variantGeneratedVerificationFactory  *VariantGeneratedVerificationFactory
}

func NewBlockGeneratedVerificationFactory() *BlockGeneratedVerificationFactory {
	return &BlockGeneratedVerificationFactory{
		questionGeneratedVerificationFactory: NewQuestionGeneratedVerificationFactory(),
		variantGeneratedVerificationFactory:  NewVariantGeneratedVerificationFactory(),
	}
}

func (b *BlockGeneratedVerificationFactory) build(
	blocks []*generated.Block,
	questionsWithCorrectAnswers map[uuid.UUID]get.IQuestionDto,
) ([]verify2.Block, error) {
	verifiedBlocks := make([]verify2.Block, 0)
	for _, currBlock := range blocks {
		newBlock, err := b.buildBlock(currBlock, questionsWithCorrectAnswers)
		if err != nil {
			return nil, err
		}
		verifiedBlocks = append(verifiedBlocks, *newBlock)
	}
	return verifiedBlocks, nil
}

func (b *BlockGeneratedVerificationFactory) buildBlock(
	generatedBlock *generated.Block,
	questionsWithCorrectAnswers map[uuid.UUID]get.IQuestionDto,
) (*verify2.Block, error) {
	questions, err := b.questionGeneratedVerificationFactory.build(generatedBlock.Questions, questionsWithCorrectAnswers)
	if err != nil {
		return nil, err
	}

	variant := new(verify2.Variant)
	if generatedBlock.Variant != nil {
		variant, err = b.variantGeneratedVerificationFactory.build(generatedBlock.Variant, questionsWithCorrectAnswers)
		if err != nil {
			return nil, err
		}
	}

	return &verify2.Block{
		Id:          generatedBlock.Id,
		Type:        generatedBlock.Type,
		Name:        generatedBlock.Title,
		Description: generatedBlock.Description,
		Variant:     variant,
		Questions:   questions,
	}, nil
}
