package factory

import (
	"errors"
	"hedgehog-forms/dto"
	blockF "hedgehog-forms/factory/block"
	"hedgehog-forms/model/form/section"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/service"
)

type SectionFactory struct {
	blockFactory *blockF.BlockFactory
}

func NewSectionFactory() *SectionFactory {
	return &SectionFactory{
		blockFactory: blockF.NewBlockFactory(),
	}
}

func (s *SectionFactory) buildSection(sectionDto any) (section.Section, error) {
	switch sect := sectionDto.(type) {
	case *dto.CreateNewSectionDto:
		return s.buildSectionNew(sect)
	case *dto.CreateSectionOnExistingDto:
		return s.buildSectionExisting(sect)
	default:
		return section.Section{}, errors.New("invalid section type")
	}
}

func (s *SectionFactory) buildSectionNew(sectionDto *dto.CreateNewSectionDto) (section.Section, error) {
	var sectionObj section.Section
	sectionObj.Title = sectionDto.Title
	sectionObj.Description = sectionDto.Description
	err := s.buildAndAddBlocksFromDto(sectionDto.Blocks, &sectionObj)
	if err != nil {
		return section.Section{}, err
	}
	return sectionObj, nil
}

func (s *SectionFactory) buildAndAddBlocksFromDto(blockDtos []any, sectionObj *section.Section) error {
	blocks := make([]block.IBlock, len(blockDtos))
	for order, blockDto := range blockDtos {
		blockObj, err := s.blockFactory.BuildFromDto(blockDto)
		if err != nil {
			return err
		}

		switch blockTyped := blockObj.(type) {
		case block.IBlock:
			blockTyped.SetOrder(order)
			blocks = append(blocks, blockTyped)
		default:
			return errors.New("invalid block type")
		}
	}

	sectionObj.Blocks = blocks
	return nil
}

func (s *SectionFactory) buildSectionExisting(sectionDto *dto.CreateSectionOnExistingDto) (section.Section, error) {
	sectObj, err := service.NewSectionService().GetSectionObjectById(sectionDto.SectionId)
	if err != nil {
		return section.Section{}, err
	}

	var sectionObj section.Section
	sectionObj.Title = sectObj.Title
	sectionObj.Description = sectObj.Description
	s.buildAndAddBlocksFromObj(sectObj.Blocks, &sectionObj)
	return sectionObj, nil
}

func (s *SectionFactory) buildAndAddBlocksFromObj(blockObjs []block.IBlock, sectionObj *section.Section) {
	blocks := make([]block.IBlock, len(blockObjs))
	for order, blockObj := range blockObjs {
		blockObj.SetOrder(order)
		blocks = append(blocks, blockObj)
	}

	sectionObj.Blocks = blocks
}
