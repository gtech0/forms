package block

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/factory/question"
	block2 "hedgehog-forms/model/form/pattern/section/block"
)

type VariantFactory struct {
	questionFactory *question.QuestionFactory
}

func NewVariantFactory() *VariantFactory {
	return &VariantFactory{
		questionFactory: question.NewQuestionFactory(),
	}
}

func (v *VariantFactory) buildFromDto(variantDto create.UpdateVariantDto, blockObj *block2.StaticBlock) (*block2.Variant, error) {
	variant := new(block2.Variant)
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

func (v *VariantFactory) buildFromObj(variant *block2.Variant, blockObj *block2.StaticBlock) (*block2.Variant, error) {
	newVariant := new(block2.Variant)
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

func (v *VariantFactory) buildFromDtos(variantDtos []create.UpdateVariantDto, blockObj *block2.StaticBlock) ([]block2.Variant, error) {
	variants := make([]block2.Variant, 0)
	for _, variantDto := range variantDtos {
		variant, err := v.buildFromDto(variantDto, blockObj)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}
	return variants, nil
}

func (v *VariantFactory) buildFromObjs(variantObjs []block2.Variant, blockObj *block2.StaticBlock) ([]block2.Variant, error) {
	variants := make([]block2.Variant, 0)
	for _, variantObj := range variantObjs {
		variant, err := v.buildFromObj(&variantObj, blockObj)
		if err != nil {
			return nil, err
		}
		variants = append(variants, *variant)
	}
	return variants, nil
}
