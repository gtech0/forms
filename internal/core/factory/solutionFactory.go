package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/generated"
	published2 "hedgehog-forms/internal/core/model/form/published"
	"hedgehog-forms/internal/core/util"
)

type SolutionFactory struct{}

func NewSolutionFactory() *SolutionFactory {
	return &SolutionFactory{}
}

func (s *SolutionFactory) BuildFromPublished(
	formPublished *published2.FormPublished,
	userId *uuid.UUID,
	submission *generated.Submission,
) *published2.Solution {
	solution := new(published2.Solution)
	solution.IsIndividual = util.Pointer(false)
	if formPublished.Teams == nil || len(formPublished.Teams) == 0 {
		solution.IsIndividual = util.Pointer(true)
	}

	solution.UserOwnerId = userId
	solution.ClassTaskId = formPublished.Id
	solution.Submissions = []generated.Submission{*submission}
	return solution
}
