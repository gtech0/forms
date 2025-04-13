package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/generated"
)

type SectionGeneratedIntegrationMapper struct {
	blockGeneratedIntegrationMapper *BlockGeneratedIntegrationMapper
}

func NewSectionGeneratedIntegrationMapper() *SectionGeneratedIntegrationMapper {
	return &SectionGeneratedIntegrationMapper{
		blockGeneratedIntegrationMapper: NewBlockGeneratedIntegrationMapper(),
	}
}

func (m *SectionGeneratedIntegrationMapper) ToIntegrationDto(section generated.Section, isAnswerRequired bool) (*get.IntegrationSectionDto, error) {
	sectionDto := new(get.IntegrationSectionDto)
	sectionDto.Id = section.Id
	sectionDto.Title = section.Title
	sectionDto.Description = section.Description
	blocks, err := m.blocksToDto(section.Blocks, isAnswerRequired)
	if err != nil {
		return nil, err
	}

	sectionDto.Blocks = blocks
	return sectionDto, nil
}

func (m *SectionGeneratedIntegrationMapper) blocksToDto(blocks []*generated.Block, isAnswerRequired bool) ([]get.IntegrationBlockDto, error) {
	mappedBlocks := make([]get.IntegrationBlockDto, 0)
	for _, currentBlock := range blocks {
		mappedBlock, err := m.blockGeneratedIntegrationMapper.toDto(currentBlock, isAnswerRequired)
		if err != nil {
			return nil, err
		}

		mappedBlocks = append(mappedBlocks, *mappedBlock)
	}
	return mappedBlocks, nil
}
