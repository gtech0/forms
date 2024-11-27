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
	sectionEntity := new(section.Section)
	sectionEntity.Title = sectionDto.Title
	sectionEntity.Description = sectionDto.Description
	err := s.buildAndAddBlocksFromDto(sectionDto.Blocks, sectionEntity)
	if err != nil {
		return nil, err
	}
	return sectionEntity, nil
}

func (s *SectionFactory) buildAndAddBlocksFromDto(blockDtos []any, sectionEntity *section.Section) error {
	blocks := make([]*block.Block, 0)
	for order, blockDto := range blockDtos {
		blockEntity, err := s.blockFactory.BuildFromDto(blockDto)
		if err != nil {
			return err
		}

		blockEntity.Order = order
		blocks = append(blocks, blockEntity)
	}

	sectionEntity.Blocks = blocks
	return nil
}

func (s *SectionFactory) buildSectionExisting(sectionDto *create.SectionOnExistingDto) (*section.Section, error) {
	sectEntity, err := s.sectionRepository.FindById(sectionDto.SectionId)
	if err != nil {
		return nil, err
	}

	sectionEntity := new(section.Section)
	sectionEntity.Title = sectEntity.Title
	sectionEntity.Description = sectEntity.Description
	s.buildAndAddBlocksFromEntity(sectEntity.Blocks, sectionEntity)
	return sectionEntity, nil
}

func (s *SectionFactory) buildAndAddBlocksFromEntity(blockEntities []*block.Block, sectionEntity *section.Section) {
	blocks := make([]*block.Block, 0)
	for order, blockEntity := range blockEntities {
		blockEntity.Order = order
		blocks = append(blocks, blockEntity)
	}

	sectionEntity.Blocks = blocks
}
