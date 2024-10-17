package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
)

type FormGeneratedFactory struct {
	sectionGeneratedFactory *SectionGeneratedFactory
	formGeneratedFactory    *FormGeneratedFactory
}

func NewFormGeneratedFactory() *FormGeneratedFactory {
	return &FormGeneratedFactory{
		sectionGeneratedFactory: NewSectionGeneratedFactory(),
		formGeneratedFactory:    NewFormGeneratedFactory(),
	}
}

func (f *FormGeneratedFactory) BuildForm(published published.FormPublished, userId uuid.UUID) (generated.FormGenerated, error) {
	var generatedForm generated.FormGenerated
	generatedForm.Id = uuid.New()
	generatedForm.Status = form.NEW
	generatedForm.FormPublishedID = published.Id
	generatedForm.UserId = userId
	sections, err := f.sectionGeneratedFactory.buildSections(published.FormPattern.Sections)
	if err != nil {
		return generatedForm, err
	}

	generatedForm.Sections = sections
	return generatedForm, nil
}

func (f *FormGeneratedFactory) buildAndCreate(
	formPublished published.FormPublished,
	userId uuid.UUID,
) (generated.FormGenerated, error) {
	generatedForm, err := f.formGeneratedFactory.BuildForm(formPublished, userId)
	if err != nil {
		return generated.FormGenerated{}, err
	}

	return generatedForm, nil
}
