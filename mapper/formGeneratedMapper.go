package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
)

type GeneratedMapper struct {
	sectionMapper *SectionMapper
}

func NewGeneratedMapper() *GeneratedMapper {
	return &GeneratedMapper{
		sectionMapper: NewSectionMapper(),
	}
}

func (g *GeneratedMapper) Build(published published.FormPublished, userId uuid.UUID) generated.FormGenerated {
	var generatedForm generated.FormGenerated
	generatedForm.Id = uuid.New()
	generatedForm.Status = form.NEW
	generatedForm.FormPublishedID = published.Id
	generatedForm.UserId = userId
	generatedForm.Sections = g.sectionMapper.build(published.FormPattern.Sections)
	return generatedForm
}
