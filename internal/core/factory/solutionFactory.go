package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/published"
	"hedgehog-forms/internal/core/util"
)

type SolutionFactory struct {
	submissionFactory *SubmissionFactory
}

func NewSolutionFactory() *SolutionFactory {
	return &SolutionFactory{
		submissionFactory: NewSubmissionFactory(),
	}
}

func (s *SolutionFactory) BuildFromPublished(
	formPublished *published.FormPublished,
	userId *uuid.UUID,
) (*published.Solution, error) {
	solution := new(published.Solution)
	solution.IsIndividual = util.Pointer(true)
	solution.UserOwnerId = userId
	solution.ClassTaskId = formPublished.Id

	submission, err := s.submissionFactory.Build(userId, formPublished)
	if err != nil {
		return nil, err
	}

	solution.Submissions = append(solution.Submissions, *submission)
	return solution, nil
}
