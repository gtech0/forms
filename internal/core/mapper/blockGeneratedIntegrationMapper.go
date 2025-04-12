package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/generated"
)

type BlockGeneratedIntegrationMapper struct {
	questionGeneratedIntegrationMapper *QuestionGeneratedIntegrationMapper
}

func NewBlockGeneratedIntegrationMapper() *BlockGeneratedIntegrationMapper {
	return &BlockGeneratedIntegrationMapper{
		questionGeneratedIntegrationMapper: NewQuestionGeneratedIntegrationMapper(),
	}
}

func (b *BlockGeneratedIntegrationMapper) toDto(block *generated.Block) (*get.IntegrationBlockDto, error) {
	blockDto := new(get.IntegrationBlockDto)
	blockDto.Id = block.Id
	blockDto.Title = block.Title
	blockDto.Description = block.Description
	blockDto.Type = block.Type
	if block.Variant != nil {
		variant, err := b.variantToDto(block.Variant)
		if err != nil {
			return nil, err
		}
		blockDto.Variant = variant
	}

	questions, err := b.questionsToDto(block.Questions)
	if err != nil {
		return nil, err
	}

	blockDto.Questions = questions
	return blockDto, nil
}

func (b *BlockGeneratedIntegrationMapper) variantToDto(variantEntity *generated.Variant) (*get.IntegratedVariantDto, error) {
	variant := new(get.IntegratedVariantDto)
	variant.Id = variantEntity.Id
	variant.Title = variantEntity.Title
	variant.Description = variantEntity.Description
	questions, err := b.questionsToDto(variantEntity.Questions)
	if err != nil {
		return nil, err
	}
	variant.Questions = questions
	return variant, nil
}

func (b *BlockGeneratedIntegrationMapper) questionsToDto(questions []generated.IQuestion) ([]get.IntegratedIQuestionDto, error) {
	mappedQuestions := make([]get.IntegratedIQuestionDto, 0)
	for _, questionEntity := range questions {
		mappedQuestion, err := b.questionGeneratedIntegrationMapper.ToDto(questionEntity)
		if err != nil {
			return nil, err
		}
		mappedQuestions = append(mappedQuestions, mappedQuestion)
	}
	return mappedQuestions, nil
}
