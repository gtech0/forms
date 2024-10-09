package factory

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/section"
)

type FormPatternFactory struct {
	sectionFactory *SectionFactory
}

func NewFormPatternFactory() *FormPatternFactory {
	return &FormPatternFactory{
		sectionFactory: NewSectionFactory(),
	}
}

func (f *FormPatternFactory) BuildFormPattern(dto *create.FormPatternDto) (form.FormPattern, error) {
	var formPattern form.FormPattern
	formPattern.Title = dto.Title
	formPattern.Description = dto.Description
	formPattern.SubjectId = dto.SubjectId
	sections, err := f.buildAndAddSections(dto.Sections, formPattern)
	if err != nil {
		return form.FormPattern{}, err
	}

	formPattern.Sections = sections
	return formPattern, nil
}

func (f *FormPatternFactory) buildAndAddSections(
	sectionDtos []any,
	formPattern form.FormPattern,
) ([]section.Section, error) {
	sections := make([]section.Section, 0)
	for order, sectionDto := range sectionDtos {
		sectionObj, err := f.sectionFactory.buildSection(sectionDto)
		if err != nil {
			return nil, err
		}
		sectionObj.FormPatternId = formPattern.Id
		sectionObj.Order = order
		sections = append(sections, sectionObj)
	}

	return sections, nil
}
