package service

import (
	get2 "hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/mapper"
	"hedgehog-forms/internal/core/repository"
	"hedgehog-forms/internal/core/util"
	"net/url"
	"strconv"
)

type SectionService struct {
	sectionRepository *repository.SectionRepository
	sectionMapper     *mapper.SectionMapper
}

func NewSectionService() *SectionService {
	return &SectionService{
		sectionRepository: repository.NewSectionRepository(),
		sectionMapper:     mapper.NewSectionMapper(),
	}
}

func (s *SectionService) GetSection(sectionId string) (*get2.SectionDto, error) {
	parsedSectionId, err := util.IdCheckAndParse(sectionId)
	if err != nil {
		return nil, err
	}

	sectionEntity, err := s.sectionRepository.FindById(parsedSectionId)
	if err != nil {
		return nil, err
	}

	sectionDto, err := s.sectionMapper.ToDto(sectionEntity)
	if err != nil {
		return nil, err
	}

	return sectionDto, nil
}

func (s *SectionService) GetSections(query url.Values) (*get2.PaginationResponse[get2.SectionDto], error) {
	name := query.Get("name")
	page, _ := strconv.Atoi(query.Get("page"))
	if page <= 0 {
		page = 1
	}

	size, _ := strconv.Atoi(query.Get("size"))
	switch {
	case size > 20:
		size = 20
	case size <= 0:
		size = 5
	}

	sections, err := s.sectionRepository.FindByNameAndPaginate(name, page, size)
	if err != nil {
		return nil, err
	}

	sectionDtos := make([]get2.SectionDto, 0)
	for _, section := range sections {
		sectionDto, err := s.sectionMapper.ToDto(&section)
		if err != nil {
			return nil, err
		}

		sectionDtos = append(sectionDtos, *sectionDto)
	}

	return &get2.PaginationResponse[get2.SectionDto]{
		Page:     page,
		Size:     size,
		Elements: sectionDtos,
	}, nil
}
