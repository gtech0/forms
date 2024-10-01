package factory

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/section"
)

type FormPatternFactory struct {
}

func (f *FormPatternFactory) BuildFormPattern(dto dto.CreateFormPatternDto) form.FormPattern {
	var formPattern form.FormPattern
	formPattern.Title = dto.Title
	formPattern.Description = dto.Description
	formPattern.SubjectId = dto.SubjectId
}

func (f *FormPatternFactory) BuildAndAddSections(dtos []dto.CreateSectionDto, formPattern form.FormPattern) {
	sections := make([]section.Section, len(dtos))
	for i := 0; i < len(dtos); i++ {

	}
}
