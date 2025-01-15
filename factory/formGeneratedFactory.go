package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
	"time"
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
	userId uuid.UUID,
) (*generated.FormGenerated, error) {
	generatedForm := new(generated.FormGenerated)
	generatedForm.Id = uuid.New()
	generatedForm.Status = generated.NEW
	generatedForm.UserId = userId
	generatedForm.FormPublishedId = published.Id
	generatedForm.StartTime = time.Now()
	sections, err := f.sectionFactory.BuildSections(published.FormPattern.Sections, published.ExcludedQuestionsToSlice())
	if err != nil {
		return nil, err
	}

	generatedForm.Sections = sections
	return generatedForm, nil
}
