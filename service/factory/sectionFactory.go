package factory

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section"
	"hedgehog-forms/model/form/section/block"
)

type SectionFactory struct{}

func (s *SectionFactory) buildSection(dto dto.CreateSectionDto) section.Section {
	switch dto.Type {
	case section.NEW:
	case section.EXISTING:
	default:

	}
}

func (s *SectionFactory) buildSectionFromDto(sectionDto dto.CreateNewSectionDto) section.Section {
	var section section.Section
	section.Title = sectionDto.Title
	section.Description = sectionDto.Description

}

func (s *SectionFactory) BuildAndAddBlocks(dtos []dto.CreateBlockDto, section section.Section) {
	blocks := make([]block.Block, len(dtos))
	for i := 0; i < len(dtos); i++ {

	}

	section.Blocks = blocks
}
