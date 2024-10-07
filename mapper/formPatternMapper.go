package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/section"
)

type FormPatternMapper struct {
	subjectMapper *SubjectMapper
	sectionMapper *SectionMapper
}

func NewFormPatternMapper() *FormPatternMapper {
	return &FormPatternMapper{
		subjectMapper: NewSubjectMapper(),
	}
}

func (f *FormPatternMapper) toDto(formPattern form.FormPattern) get.FormPatternDto {
	var formPatternDto get.FormPatternDto
	formPatternDto.Id = formPattern.Id
	formPatternDto.Title = formPattern.Title
	formPatternDto.Description = formPattern.Description
	formPatternDto.Subject = f.subjectMapper.toDto(formPattern.Subject)
	formPatternDto.Sections = f.sectionsToDto(formPattern.Sections)
	return formPatternDto
}

func (f *FormPatternMapper) sectionsToDto(sections []section.Section) []get.SectionDto {
	mappedSections := make([]get.SectionDto, 0)
	for _, currentSection := range sections {
		mappedSection := f.sectionMapper.toDto(currentSection)
		mappedSections = append(mappedSections, mappedSection)
	}
	return mappedSections
}
