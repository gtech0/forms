package factory

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/model/form/pattern/section"
)

type PatternFactory struct {
	sectionFactory *SectionFactory
}

func NewPatternFactory() *PatternFactory {
	return &PatternFactory{
		sectionFactory: NewSectionFactory(),
	}
}

func (f *PatternFactory) BuildPattern(dto *create.FormPatternDto) (pattern.FormPattern, error) {
	var formPattern pattern.FormPattern
	formPattern.Title = dto.Title
	formPattern.Description = dto.Description
	formPattern.SubjectId = dto.SubjectId
	sections, err := f.buildAndAddSections(dto.Sections, formPattern)
	if err != nil {
		return pattern.FormPattern{}, err
	}

	formPattern.Sections = sections
	return formPattern, nil
}

func (f *PatternFactory) buildAndAddSections(
	sectionDtos []any,
	formPattern pattern.FormPattern,
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
