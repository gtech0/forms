package factory

import (
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section"
)

type SectionGeneratedFactory struct {
	blockGeneratedFactory *BlockGeneratedFactory
}

func NewSectionGeneratedFactory() *SectionGeneratedFactory {
	return &SectionGeneratedFactory{
		blockGeneratedFactory: NewBlockGeneratedFactory(),
	}
}

func (s *SectionGeneratedFactory) buildSections(blocks []section.Section) ([]generated.Section, error) {
	var generatedSections []generated.Section
	for _, currBlock := range blocks {
		newBlock, err := s.buildSection(currBlock)
		if err != nil {
			return nil, err
		}
		generatedSections = append(generatedSections, newBlock)
	}
	return generatedSections, nil
}

func (s *SectionGeneratedFactory) buildSection(newSection section.Section) (generated.Section, error) {
	var generatedSection generated.Section
	generatedSection.Id = newSection.Id
	generatedSection.Title = newSection.Title
	generatedSection.Description = newSection.Description
	blocks, err := s.blockGeneratedFactory.buildBlocks(newSection.Blocks)
	if err != nil {
		return generatedSection, err
	}

	generatedSection.Blocks = blocks
	return generatedSection, nil
}
