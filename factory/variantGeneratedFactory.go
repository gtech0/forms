package factory

import (
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section/block"
	"math/rand"
)

type VariantGeneratedFactory struct {
	questionGeneratedFactory *QuestionGeneratedFactory
}

func NewVariantGeneratedFactory() *VariantGeneratedFactory {
	return &VariantGeneratedFactory{
		questionGeneratedFactory: NewQuestionGeneratedFactory(),
	}
}

func (v *VariantGeneratedFactory) buildVariant(variants []block.Variant) (generated.Variant, error) {
	randomIndex := rand.Intn(len(variants))
	variant := variants[randomIndex]
	questions, err := v.questionGeneratedFactory.buildQuestions(variant.Questions)
	if err != nil {
		return generated.Variant{}, err
	}

	return generated.Variant{
		Id:          variant.Id,
		Title:       variant.Title,
		Description: variant.Description,
		Questions:   questions,
	}, nil
}
