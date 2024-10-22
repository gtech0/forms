package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
	"slices"
)

type MatchingProcessor struct{}

func NewMatchingProcessor() *MatchingProcessor {
	return &MatchingProcessor{}
}

// TODO
func (m *MatchingProcessor) markAnswer(
	questionId uuid.UUID,
	termsAndDefinitions []generated.EnteredMatchingPair,
	matching *generated.Matching,
) error {
	enteredAnswers := make([]generated.EnteredMatchingPair, 0)
	termIds := make([]uuid.UUID, 0)
	for _, term := range matching.Terms {
		termIds = append(termIds, term.Id)
	}

	definitionIds := make([]uuid.UUID, 0)
	for _, definition := range matching.Definitions {
		definitionIds = append(definitionIds, definition.Id)
	}

	for _, termAndDefinition := range termsAndDefinitions {
		if !slices.Contains(termIds, termAndDefinition.TermId) {
			return errs.New(
				fmt.Sprintf("term %v isn't in question %v", termAndDefinition.TermId, questionId),
				400,
			)
		}

		if !slices.Contains(definitionIds, termAndDefinition.DefinitionId) {
			return errs.New(
				fmt.Sprintf("definition %v isn't in question %v", termAndDefinition.TermId, questionId),
				400,
			)
		}

		var matchingPair generated.EnteredMatchingPair
		matchingPair.TermId = termAndDefinition.TermId
		matchingPair.DefinitionId = termAndDefinition.DefinitionId
		enteredAnswers = append(enteredAnswers, matchingPair)
	}

	matching.EnteredAnswers = enteredAnswers
	return nil
}

func (m *MatchingProcessor) checkIdExisting(ids []uuid.UUID, id uuid.UUID, questionId uuid.UUID) bool {
	if slices.Contains(ids, id) {
		return true
	}
	return false
}
