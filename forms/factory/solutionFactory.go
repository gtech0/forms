package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
	"hedgehog-forms/util"
)

type SolutionFactory struct{}

func NewSolutionFactory() *SolutionFactory {
	return &SolutionFactory{}
}

func (s *SolutionFactory) BuildFromPublished(
	formPublished *published.FormPublished,
	userId *uuid.UUID,
	submission *generated.Submission,
) *published.Solution {
	solution := new(published.Solution)
	solution.IsIndividual = util.Pointer(false)
	if formPublished.Teams == nil || len(formPublished.Teams) == 0 {
		solution.IsIndividual = util.Pointer(true)
	}

	solution.UserOwnerId = userId
	solution.ClassTaskId = formPublished.Id
	solution.Submissions = []generated.Submission{*submission}
	return solution
}
