package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/model/form/pattern/section"
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

func (f *FormPatternMapper) ToDto(formPattern pattern.FormPattern) (get.FormPatternDto, error) {
	var formPatternDto get.FormPatternDto
	formPatternDto.Id = formPattern.Id
	formPatternDto.Title = formPattern.Title
	formPatternDto.Description = formPattern.Description
	formPatternDto.OwnerId = formPattern.OwnerId
	formPatternDto.Subject = f.subjectMapper.toDto(formPattern.Subject)
	sections, err := f.sectionsToDto(formPattern.Sections)
	if err != nil {
		return get.FormPatternDto{}, err
	}
	formPatternDto.Sections = sections
	return formPatternDto, nil
}

func (f *FormPatternMapper) sectionsToDto(sections []section.Section) ([]get.SectionDto, error) {
	mappedSections := make([]get.SectionDto, 0)
	for _, currentSection := range sections {
		mappedSection, err := f.sectionMapper.toDto(currentSection)
		if err != nil {
			return nil, err
		}
		mappedSections = append(mappedSections, mappedSection)
	}
	return mappedSections, nil
}
