package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
)

type FormGeneratedFactory struct {
	attemptFactory *AttemptFactory
}

func NewFormGeneratedFactory() *FormGeneratedFactory {
	return &FormGeneratedFactory{
		attemptFactory: NewAttemptFactory(),
	}
}

func (f *FormGeneratedFactory) BuildForm(
	published *published.FormPublished,
	userId uuid.UUID,
) (*generated.FormGenerated, error) {
	generatedForm := new(generated.FormGenerated)
	generatedForm.Id = uuid.New()
	generatedForm.Status = generated.NEW
	generatedForm.FormPublishedID = published.Id
	generatedForm.UserId = userId
	attempts := make([]*generated.Attempt, 0)
	attempt, err := f.attemptFactory.BuildAttempt(published.FormPattern.Sections, nil)
	if err != nil {
		return nil, err
	}
	attempts = append(attempts, attempt)
	generatedForm.Attempts = attempts
	return generatedForm, nil
}
