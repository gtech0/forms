package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
	"time"
)

type SubmissionFactory struct{}

func NewSubmissionFactory() *SubmissionFactory {
	return &SubmissionFactory{}
}

func (s *SubmissionFactory) Build(
	userId uuid.UUID,
	generatedId uuid.UUID,
) *generated.Submission {
	submission := new(generated.Submission)
	submission.UserId = &userId
	submission.StartTime = time.Now()
	submission.FormGeneratedId = generatedId
	return submission
}
