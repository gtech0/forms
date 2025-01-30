package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/dto/verify"
	"hedgehog-forms/internal/core/model/form/generated"
)

type SectionGeneratedVerificationFactory struct {
	blockGeneratedVerificationFactory *BlockGeneratedVerificationFactory
}

func NewSectionGeneratedVerificationFactory() *SectionGeneratedVerificationFactory {
	return &SectionGeneratedVerificationFactory{
		blockGeneratedVerificationFactory: NewBlockGeneratedVerificationFactory(),
	}
}

func (s *SectionGeneratedVerificationFactory) build(
	sections []generated.Section,
	questionsWithCorrectAnswers map[uuid.UUID]get.IQuestionDto,
) ([]verify.Section, error) {
	verifiedSections := make([]verify.Section, 0)
	for _, currSection := range sections {
		newSection, err := s.buildSection(currSection, questionsWithCorrectAnswers)
		if err != nil {
			return nil, err
		}
		verifiedSections = append(verifiedSections, *newSection)
	}
	return verifiedSections, nil
}

func (s *SectionGeneratedVerificationFactory) buildSection(
	generatedSection generated.Section,
	questionsWithCorrectAnswers map[uuid.UUID]get.IQuestionDto,
) (*verify.Section, error) {
	verifiedBlocks, err := s.blockGeneratedVerificationFactory.build(generatedSection.Blocks, questionsWithCorrectAnswers)
	if err != nil {
		return nil, err
	}

	return &verify.Section{
		Id:          generatedSection.Id,
		Name:        generatedSection.Title,
		Description: generatedSection.Description,
		Blocks:      verifiedBlocks,
	}, nil
}
