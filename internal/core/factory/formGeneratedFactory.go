package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/published"
)

type FormGeneratedFactory struct {
	sectionFactory    *SectionGeneratedFactory
	submissionFactory *SubmissionFactory
}

func NewFormGeneratedFactory() *FormGeneratedFactory {
	return &FormGeneratedFactory{
		sectionFactory:    NewSectionGeneratedFactory(),
		submissionFactory: NewSubmissionFactory(),
	}
}

func (f *FormGeneratedFactory) BuildForm(
	published *published.FormPublished,
	userId uuid.UUID,
) (*generated.FormGenerated, error) {
	generatedForm := new(generated.FormGenerated)
	generatedForm.Id = uuid.New()
	generatedForm.Status = generated.NEW
	generatedForm.FormPublishedId = published.Id
	sections, err := f.sectionFactory.BuildSections(published.FormPattern.Sections, published.ExcludedQuestionsToSlice())
	if err != nil {
		return nil, err
	}

	generatedForm.Sections = sections
	generatedForm.Submission = f.submissionFactory.Build(userId, generatedForm.Id)
	return generatedForm, nil
}
