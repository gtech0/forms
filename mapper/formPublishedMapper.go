package mapper

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/published"
)

type FormPublishedMapper struct {
	formPatternMapper *FormPatternMapper
}

func NewFormPublishedMapper() *FormPublishedMapper {
	return &FormPublishedMapper{
		formPatternMapper: NewFormPatternMapper(),
	}
}

func (f *FormPublishedMapper) toBaseDto(publishedForm published.FormPublished) create.FormPublishedBaseDto {
	var publishedBaseDto create.FormPublishedBaseDto
	//TODO
	return publishedBaseDto
}
