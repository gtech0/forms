package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/section"
)

type SectionMapper struct{}

func NewSectionMapper() *SectionMapper {
	return &SectionMapper{}
}

func (s *SectionMapper) toDto(sections section.Section) get.SectionDto {

}
