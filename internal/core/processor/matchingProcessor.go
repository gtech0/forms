package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/question"
	"slices"
)

type MatchingProcessor struct{}

func NewMatchingProcessor() *MatchingProcessor {
	return &MatchingProcessor{}
}

func (m *MatchingProcessor) calculateAndSetPoints(
	matching *generated.Matching,
	matchingEntity *question.Matching,
) int {
	termsAndDefinitions := m.extractTermDefinitionPairs(matchingEntity)

	var correctAnswers int
	for _, enteredAnswer := range matching.EnteredAnswers {
		if slices.Contains(termsAndDefinitions, enteredAnswer) {
			correctAnswers++
		}
	}

	matching.Points = m.calculatePoints(matchingEntity.Points, correctAnswers)
	return matching.Points
}

func (m *MatchingProcessor) calculatePoints(matchingPoints []question.MatchingPoints, correctAnswers int) int {
	var points int
	for _, matchingPoint := range matchingPoints {
		if matchingPoint.Points > points && matchingPoint.CorrectAnswer <= correctAnswers {
			points = matchingPoint.Points
		}
	}

	return points
}

func (m *MatchingProcessor) extractTermDefinitionPairs(matchingEntity *question.Matching) []generated.EnteredMatchingPair {
	pairs := make([]generated.EnteredMatchingPair, 0)
	for _, matchingTerm := range matchingEntity.Terms {
		var pair generated.EnteredMatchingPair
		pair.TermId = matchingTerm.Id
		pair.DefinitionId = matchingTerm.MatchingDefinitionId
		pairs = append(pairs, pair)
	}
	return pairs
}

func (m *MatchingProcessor) saveAnswer(
	matching *generated.Matching,
	termsAndDefinitions []generated.EnteredMatchingPair,
	questionId uuid.UUID,
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
