package factory

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/section"
)

type FormPatternFactory struct {
}

func (f *FormPatternFactory) BuildFormPattern(dto *dto.CreateFormPatternDto) (form.FormPattern, error) {
	var formPattern form.FormPattern
	formPattern.Title = dto.Title
	formPattern.Description = dto.Description
	formPattern.SubjectId = dto.SubjectId
	sections, err := f.BuildAndAddSections(dto.Sections, formPattern)
	if err != nil {
		return form.FormPattern{}, err
	}

	formPattern.Sections = sections
	return formPattern, nil
}

func (f *FormPatternFactory) BuildAndAddSections(
	sectionDtos []any,
	formPattern form.FormPattern,
) ([]section.Section, error) {
	sections := make([]section.Section, len(sectionDtos))
	for order, sectionDto := range sectionDtos {
		sectionObj, err := NewSectionFactory().buildSection(sectionDto)
		if err != nil {
			return nil, err
		}
		sectionObj.FormPatternId = formPattern.Id
		sectionObj.Order = order
		sections = append(sections, sectionObj)
	}

	return sections, nil
}
