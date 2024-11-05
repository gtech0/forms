package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/dto/verify"
	"hedgehog-forms/model/form/generated"
)

type SectionGeneratedVerificationFactory struct{}

func NewSectionGeneratedVerificationFactory() *SectionGeneratedVerificationFactory {
	return &SectionGeneratedVerificationFactory{}
}

func (s *SectionGeneratedVerificationFactory) build(
	sections []generated.Section,
	questionsWithCorrectAnswers map[uuid.UUID]get.IQuestionDto,
) ([]verify.Section, error) {
	//TODO
	return nil, nil
}
