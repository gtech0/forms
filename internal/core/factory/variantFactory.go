package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/model/form/pattern/section/block"
)

type VariantFactory struct {
	questionFactory *QuestionFactory
}

func NewVariantFactory() *VariantFactory {
	return &VariantFactory{
		questionFactory: NewQuestionFactory(),
	}
}

func (v *VariantFactory) buildFromDto(variantDto create.UpdateVariantDto, blockId uuid.UUID) (*block.Variant, error) {
	variant := new(block.Variant)
	variant.Title = variantDto.Title
	variant.Description = variantDto.Description
	questions, err := v.questionFactory.BuildQuestionDtoForVariant(variantDto.Questions, variant)
	if err != nil {
		return nil, err
	}
	variant.Questions = questions
	variant.StaticBlockID = blockId
	return variant, nil
}

func (v *VariantFactory) buildFromEntity(variant *block.Variant, blockId uuid.UUID) (*block.Variant, error) {
	newVariant := new(block.Variant)
	newVariant.Title = variant.Title
	newVariant.Description = variant.Description
	questions, err := v.questionFactory.BuildQuestionForVariantEntities(variant.Questions, newVariant)
	if err != nil {
		return nil, err
	}
	newVariant.Questions = questions
	newVariant.StaticBlockID = blockId
	return newVariant, nil
}

func (v *VariantFactory) buildFromDtos(variantDtos []create.UpdateVariantDto, blockId uuid.UUID) ([]block.Variant, error) {
	variants := make([]block.Variant, 0)
	for _, variantDto := range variantDtos {
		variant, err := v.buildFromDto(variantDto, blockId)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}
	return variants, nil
}

func (v *VariantFactory) buildFromEntities(variantEntities []block.Variant, blockId uuid.UUID) ([]block.Variant, error) {
	variants := make([]block.Variant, 0)
	for _, variantEntity := range variantEntities {
		variant, err := v.buildFromEntity(&variantEntity, blockId)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}
	return variants, nil
}
