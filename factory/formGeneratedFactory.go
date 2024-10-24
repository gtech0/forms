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

func (f *FormGeneratedFactory) BuildForm(published published.FormPublished) (*generated.FormGenerated, error) {
	generatedForm := new(generated.FormGenerated)
	generatedForm.Id = uuid.New()
	generatedForm.Status = form.NEW
	generatedForm.FormPublishedID = published.Id
	//generatedForm.UserId = userId
	sections, err := f.sectionGeneratedFactory.buildSections(published.FormPattern.Sections)
	if err != nil {
		return nil, err
	}

	generatedForm.Sections = sections
	return generatedForm, nil
}
