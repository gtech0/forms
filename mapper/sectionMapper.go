package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/section"
)

type SectionMapper struct {
	blockMapper *BlockMapper
}

func NewSectionMapper() *SectionMapper {
	return &SectionMapper{
		blockMapper: NewBlockMapper(),
	}
}

func (s *SectionMapper) toDto(sectionObj section.Section) (get.SectionDto, error) {
	var sectionDto get.SectionDto
	sectionDto.Id = sectionObj.Id
	sectionDto.Title = sectionObj.Title
	sectionDto.Description = sectionObj.Description
	blocks, err := s.blocksToDto(sectionObj)
	if err != nil {
		return get.SectionDto{}, err
	}

	sectionDto.Blocks = blocks
	return sectionDto, nil
}

func (s *SectionMapper) blocksToDto(sectionObj section.Section) ([]get.IBlockDto, error) {
	mappedBlocks := make([]get.IBlockDto, 0)
	for _, currentBlock := range sectionObj.Blocks {
		mappedBlock, err := s.blockMapper.toDto(currentBlock)
		if err != nil {
			return nil, err
		}

		mappedBlocks = append(mappedBlocks, mappedBlock)
	}
	return mappedBlocks, nil
}
