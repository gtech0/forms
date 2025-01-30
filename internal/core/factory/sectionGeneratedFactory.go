package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/section"
)

type SectionGeneratedFactory struct {
	blockGeneratedFactory *BlockGeneratedFactory
}

func NewSectionGeneratedFactory() *SectionGeneratedFactory {
	return &SectionGeneratedFactory{
		blockGeneratedFactory: NewBlockGeneratedFactory(),
	}
}

func (s *SectionGeneratedFactory) BuildSections(
	sections []section.Section,
	excluded []uuid.UUID,
) ([]generated.Section, error) {
	var generatedSections []generated.Section
	for _, currSection := range sections {
		newSection, err := s.buildSection(currSection, excluded)
		if err != nil {
			return nil, err
		}
		generatedSections = append(generatedSections, newSection)
	}
	return generatedSections, nil
}

func (s *SectionGeneratedFactory) buildSection(
	newSection section.Section,
	excluded []uuid.UUID,
) (generated.Section, error) {
	var generatedSection generated.Section
	generatedSection.Id = newSection.Id
	generatedSection.Title = newSection.Title
	generatedSection.Description = newSection.Description
	blocks, err := s.blockGeneratedFactory.buildBlocks(newSection.Blocks, excluded)
	if err != nil {
		return generatedSection, err
	}

	generatedSection.Blocks = blocks
	return generatedSection, nil
}
