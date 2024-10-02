package factory

import (
	"errors"
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/service"
)

type SectionFactory struct{}

func NewSectionFactory() *SectionFactory {
	return &SectionFactory{}
}

func (s *SectionFactory) buildSection(sectionDto any) (section.Section, error) {
	switch sect := sectionDto.(type) {
	case *dto.CreateNewSectionDto:
		return s.buildSectionNew(sect), nil
	case *dto.CreateSectionOnExistingDto:
		return s.buildSectionExisting(sect)
	default:
		return section.Section{}, errors.New("invalid section type")
	}
}

func (s *SectionFactory) buildSectionNew(sectionDto *dto.CreateNewSectionDto) section.Section {
	var sectionObj section.Section
	sectionObj.Title = sectionDto.Title
	sectionObj.Description = sectionDto.Description
	s.buildAndAddBlocksFromDto(sectionDto.Blocks, sectionObj)
	return sectionObj
}

func (s *SectionFactory) buildSectionExisting(sectionDto *dto.CreateSectionOnExistingDto) (section.Section, error) {
	sectObj, err := service.NewSectionService().GetSectionObjectById(sectionDto.SectionId)
	if err != nil {
		return section.Section{}, err
	}

	var sectionObj section.Section
	sectionObj.Title = sectObj.Title
	sectionObj.Description = sectObj.Description
	s.buildAndAddBlocksFromObj(sectObj.Blocks, sectionObj)
	return sectionObj, nil
}

func (s *SectionFactory) buildAndAddBlocksFromDto(blockDtos []any, sectionObj section.Section) {
	blocks := make([]block.Block, len(blockDtos))
	for i := 0; i < len(blockDtos); i++ {

	}

	sectionObj.Blocks = blocks
}

func (s *SectionFactory) buildAndAddBlocksFromObj(blockObjs []block.Block, sectionObj section.Section) {

}
