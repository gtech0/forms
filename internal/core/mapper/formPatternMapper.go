package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/pattern"
	"hedgehog-forms/internal/core/model/form/pattern/section"
)

type FormPatternMapper struct {
	subjectMapper *SubjectMapper
	sectionMapper *SectionMapper
}

func NewFormPatternMapper() *FormPatternMapper {
	return &FormPatternMapper{
		subjectMapper: NewSubjectMapper(),
		sectionMapper: NewSectionMapper(),
	}
}

func (f *FormPatternMapper) ToDto(formPattern *pattern.FormPattern) (*get.FormPatternDto, error) {
	formPatternDto := new(get.FormPatternDto)
	formPatternDto.Id = formPattern.Id
	formPatternDto.Title = formPattern.Title
	formPatternDto.Description = formPattern.Description
	formPatternDto.OwnerId = formPattern.OwnerId
	formPatternDto.Subject = *f.subjectMapper.ToDto(formPattern.Subject)
	sections, err := f.sectionsToDto(formPattern.Sections)
	if err != nil {
		return nil, err
	}
	formPatternDto.Sections = sections
	return formPatternDto, nil
}

func (f *FormPatternMapper) ToBaseDto(formPattern pattern.FormPattern) *get.FormPatternBaseDto {
	formPatternDto := new(get.FormPatternBaseDto)
	formPatternDto.Id = formPattern.Id
	formPatternDto.Title = formPattern.Title
	formPatternDto.Description = formPattern.Description
	formPatternDto.OwnerId = formPattern.OwnerId
	formPatternDto.Subject = *f.subjectMapper.ToDto(formPattern.Subject)
	return formPatternDto
}

func (f *FormPatternMapper) sectionsToDto(sections []section.Section) ([]get.SectionDto, error) {
	mappedSections := make([]get.SectionDto, 0)
	for _, currentSection := range sections {
		mappedSection, err := f.sectionMapper.ToDto(&currentSection)
		if err != nil {
			return nil, err
		}
		mappedSections = append(mappedSections, *mappedSection)
	}
	return mappedSections, nil
}
