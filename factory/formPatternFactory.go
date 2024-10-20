package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/model/form/pattern/section"
)

type FormPatternFactory struct {
	sectionFactory *SectionFactory
}

func NewFormPatternFactory() *FormPatternFactory {
	return &FormPatternFactory{
		sectionFactory: NewSectionFactory(),
	}
}

func (f *FormPatternFactory) BuildPattern(dto *create.FormPatternDto) (*pattern.FormPattern, error) {
	formPattern := new(pattern.FormPattern)
	formPattern.Id = uuid.New()
	formPattern.Title = dto.Title
	formPattern.Description = dto.Description
	formPattern.SubjectId = dto.SubjectId
	sections, err := f.buildAndAddSections(dto.Sections, formPattern.Id)
	if err != nil {
		return nil, err
	}

	formPattern.Sections = sections
	return formPattern, nil
}

func (f *FormPatternFactory) buildAndAddSections(
	sectionDtos []any,
	formPatternId uuid.UUID,
) ([]section.Section, error) {
	sections := make([]section.Section, 0)
	for order, sectionDto := range sectionDtos {
		sectionObj, err := f.sectionFactory.buildSection(sectionDto)
		if err != nil {
			return nil, err
		}
		sectionObj.FormPatternId = formPatternId
		sectionObj.Order = order
		sections = append(sections, sectionObj)
	}

	return sections, nil
}
