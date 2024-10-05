package block

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/factory/question"
	"hedgehog-forms/model/form/section/block"
)

type VariantFactory struct {
	questionFactory *question.QuestionFactory
}

func NewVariantFactory() *VariantFactory {
	return &VariantFactory{
		questionFactory: question.NewQuestionFactory(),
	}
}

func (v *VariantFactory) buildFromDto(variantDto dto.UpdateVariantDto, blockObj *block.StaticBlock) (*block.Variant, error) {
	var variant *block.Variant
	variant.Title = variantDto.Title
	variant.Description = variantDto.Description
	questions, err := v.questionFactory.BuildQuestionDtoForVariant(variantDto.Questions, variant)
	if err != nil {
		return nil, err
	}
	variant.Questions = questions
	variant.StaticBlockID = blockObj.Id
	return variant, nil
}

func (v *VariantFactory) buildFromObj(variant *block.Variant, blockObj *block.StaticBlock) (*block.Variant, error) {
	var newVariant *block.Variant
	newVariant.Title = variant.Title
	newVariant.Description = variant.Description
	questions, err := v.questionFactory.BuildQuestionForVariantObj(variant.Questions, newVariant)
	if err != nil {
		return nil, err
	}
	newVariant.Questions = questions
	newVariant.StaticBlockID = blockObj.Id
	return newVariant, nil
}

func (v *VariantFactory) buildFromDtos(variantDtos []dto.UpdateVariantDto, blockObj *block.StaticBlock) ([]block.Variant, error) {
	variants := make([]block.Variant, len(variantDtos))
	for _, variantDto := range variantDtos {
		variant, err := v.buildFromDto(variantDto, blockObj)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}
	return variants, nil
}

func (v *VariantFactory) buildFromObjs(variantObjs []block.Variant, blockObj *block.StaticBlock) ([]block.Variant, error) {
	variants := make([]block.Variant, len(variantObjs))
	for _, variantObj := range variantObjs {
		variant, err := v.buildFromObj(&variantObj, blockObj)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}
	return variants, nil
}
