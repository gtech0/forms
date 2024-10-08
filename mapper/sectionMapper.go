package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/section"
	"hedgehog-forms/model/form/section/block"
)

type SectionMapper struct {
	blockMapper *BlockMapper
}

func NewSectionMapper() *SectionMapper {
	return &SectionMapper{
		blockMapper: NewBlockMapper(),
	}
}

func (s *SectionMapper) toDto(sectionObj section.Section) get.SectionDto {
	var sectionDto get.SectionDto
	sectionDto.Id = sectionObj.Id
	sectionDto.Title = sectionObj.Title
	sectionDto.Description = sectionObj.Description

	return sectionDto
}

func (s *SectionMapper) blocksToDto(blocks []block.IBlock) []get.IBlockDto {
	mappedBlocks := make([]get.IBlockDto, 0)
	for _, currentBlock := range blocks {
		mappedBlock := s.blockMapper.toDto(currentBlock)
		mappedBlocks = append(mappedBlocks, mappedBlock)
	}
	return mappedBlocks
}
