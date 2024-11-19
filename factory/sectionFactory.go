package factory

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/repository"
)

type SectionFactory struct {
	blockFactory      *BlockFactory
	sectionRepository *repository.SectionRepository
}

func NewSectionFactory() *SectionFactory {
	return &SectionFactory{
		blockFactory:      NewBlockFactory(),
		sectionRepository: repository.NewSectionRepository(),
	}
}

func (s *SectionFactory) buildSection(sectionDto any) (*section.Section, error) {
	switch sect := sectionDto.(type) {
	case *create.NewSectionDto:
		return s.buildSectionNew(sect)
	case *create.SectionOnExistingDto:
		return s.buildSectionExisting(sect)
	default:
		return nil, errs.New("invalid section type", 400)
	}
}

func (s *SectionFactory) buildSectionNew(sectionDto *create.NewSectionDto) (*section.Section, error) {
	sectionObj := new(section.Section)
	sectionObj.Title = sectionDto.Title
	sectionObj.Description = sectionDto.Description
	err := s.buildAndAddBlocksFromDto(sectionDto.Blocks, sectionObj)
	if err != nil {
		return nil, err
	}
	return sectionObj, nil
}

func (s *SectionFactory) buildAndAddBlocksFromDto(blockDtos []any, sectionObj *section.Section) error {
	blocks := make([]*block.Block, 0)
	for order, blockDto := range blockDtos {
		blockObj, err := s.blockFactory.BuildFromDto(blockDto)
		if err != nil {
			return err
		}

		blockObj.Order = order
		blocks = append(blocks, blockObj)
	}

	sectionObj.Blocks = blocks
	return nil
}

func (s *SectionFactory) buildSectionExisting(sectionDto *create.SectionOnExistingDto) (*section.Section, error) {
	sectObj, err := s.sectionRepository.FindById(sectionDto.SectionId)
	if err != nil {
		return nil, err
	}

	sectionObj := new(section.Section)
	sectionObj.Title = sectObj.Title
	sectionObj.Description = sectObj.Description
	s.buildAndAddBlocksFromObj(sectObj.Blocks, sectionObj)
	return sectionObj, nil
}

func (s *SectionFactory) buildAndAddBlocksFromObj(blockObjs []*block.Block, sectionObj *section.Section) {
	blocks := make([]*block.Block, 0)
	for order, blockObj := range blockObjs {
		blockObj.Order = order
		blocks = append(blocks, blockObj)
	}

	sectionObj.Blocks = blocks
}
