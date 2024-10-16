package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
)

type FormGeneratedFactory struct {
	sectionGeneratedFactory *SectionGeneratedFactory
}

func NewFormGeneratedFactory() *FormGeneratedFactory {
	return &FormGeneratedFactory{
		sectionGeneratedFactory: NewSectionGeneratedFactory(),
	}
}

func (s *FormGeneratedFactory) BuildForm(published published.FormPublished, userId uuid.UUID) (generated.FormGenerated, error) {
	var generatedForm generated.FormGenerated
	generatedForm.Id = uuid.New()
	generatedForm.Status = form.NEW
	generatedForm.FormPublishedID = published.Id
	generatedForm.UserId = userId
	sections, err := s.sectionGeneratedFactory.buildSections(published.FormPattern.Sections)
	if err != nil {
		return generatedForm, err
	}

	generatedForm.Sections = sections
	return generatedForm, nil
}
