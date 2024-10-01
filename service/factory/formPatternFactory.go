package factory

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form"
)

type FormPatternFactory struct {
}

func (f *FormPatternFactory) BuildFormPattern(dto dto.CreateFormPatternDto) form.FormPattern {
	var formPattern form.FormPattern
	formPattern.Title = dto.Title
	formPattern.Description = dto.Description
	formPattern.SubjectId = dto.SubjectId
}
