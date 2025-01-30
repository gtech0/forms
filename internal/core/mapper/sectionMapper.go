package mapper

import (
	get2 "hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/pattern/section"
)

type SectionMapper struct {
	blockMapper *BlockMapper
}

func NewSectionMapper() *SectionMapper {
	return &SectionMapper{
		blockMapper: NewBlockMapper(),
	}
}

func (s *SectionMapper) ToDto(sectionEntity *section.Section) (*get2.SectionDto, error) {
	sectionDto := new(get2.SectionDto)
	sectionDto.Id = sectionEntity.Id
	sectionDto.Title = sectionEntity.Title
	sectionDto.Description = sectionEntity.Description
	blocks, err := s.blocksToDto(sectionEntity)
	if err != nil {
		return nil, err
	}

	sectionDto.Blocks = blocks
	return sectionDto, nil
}

func (s *SectionMapper) blocksToDto(sectionEntity *section.Section) ([]get2.IBlockDto, error) {
	mappedBlocks := make([]get2.IBlockDto, 0)
	for _, currentBlock := range sectionEntity.Blocks {
		mappedBlock, err := s.blockMapper.toDto(currentBlock)
		if err != nil {
			return nil, err
		}

		mappedBlocks = append(mappedBlocks, mappedBlock)
	}
	return mappedBlocks, nil
}
