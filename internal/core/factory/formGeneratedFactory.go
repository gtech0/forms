package factory

import (
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/published"
)

type FormGeneratedFactory struct {
	sectionFactory *SectionGeneratedFactory
}

func NewFormGeneratedFactory() *FormGeneratedFactory {
	return &FormGeneratedFactory{
		sectionFactory: NewSectionGeneratedFactory(),
	}
}

func (f *FormGeneratedFactory) BuildForm(
	published *published.FormPublished,
) (*generated.FormGenerated, error) {
	generatedForm := new(generated.FormGenerated)
	generatedForm.Status = generated.NEW
	generatedForm.FormPublishedId = published.Id
	sections, err := f.sectionFactory.BuildSections(published.FormPattern.Sections, published.ExcludedQuestionsToSlice())
	if err != nil {
		return nil, err
	}

	generatedForm.Sections = sections
	return generatedForm, nil
}
