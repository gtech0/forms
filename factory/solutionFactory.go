package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
)

type SolutionFactory struct {
}

func NewSolutionFactory() *SolutionFactory {
	return &SolutionFactory{}
}

func (s *SolutionFactory) BuildFromPublished(
	formPublished *published.FormPublished,
	userId *uuid.UUID,
	submission *generated.FormGenerated,
) *published.Solution {
	solution := new(published.Solution)
	solution.UserOwnerId = userId
	solution.ClassTaskId = formPublished.Id
	solution.Submissions = []generated.FormGenerated{*submission}
	return solution
}
