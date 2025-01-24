package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/dto/verify"
	"hedgehog-forms/model/form/generated"
)

type VariantGeneratedVerificationFactory struct {
	questionGeneratedVerificationFactory *QuestionGeneratedVerificationFactory
}

func NewVariantGeneratedVerificationFactory() *VariantGeneratedVerificationFactory {
	return &VariantGeneratedVerificationFactory{
		questionGeneratedVerificationFactory: NewQuestionGeneratedVerificationFactory(),
	}
}

func (v *VariantGeneratedVerificationFactory) build(
	variant *generated.Variant,
	questionsWithCorrectAnswers map[uuid.UUID]get.IQuestionDto,
) (*verify.Variant, error) {
	questions, err := v.questionGeneratedVerificationFactory.build(variant.Questions, questionsWithCorrectAnswers)
	if err != nil {
		return nil, err
	}

	return &verify.Variant{
		Id:          variant.Id,
		Name:        variant.Title,
		Description: variant.Description,
		Questions:   questions,
	}, nil
}
