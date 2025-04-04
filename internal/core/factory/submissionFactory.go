package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/published"
	"time"
)

type SubmissionFactory struct {
	formGeneratedFactory *FormGeneratedFactory
}

func NewSubmissionFactory() *SubmissionFactory {
	return &SubmissionFactory{
		formGeneratedFactory: NewFormGeneratedFactory(),
	}
}

func (s *SubmissionFactory) Build(
	userId *uuid.UUID,
	formPublished *published.FormPublished,
) (*generated.Submission, error) {
	submission := new(generated.Submission)
	submission.UserId = userId
	submission.StartTime = time.Now()
	formGenerated, err := s.formGeneratedFactory.BuildForm(formPublished)
	if err != nil {
		return nil, err
	}

	submission.FormGenerated = formGenerated
	return submission, nil
}
