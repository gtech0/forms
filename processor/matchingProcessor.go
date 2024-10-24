package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"slices"
)

type MatchingProcessor struct{}

func NewMatchingProcessor() *MatchingProcessor {
	return &MatchingProcessor{}
}

func (m *MatchingProcessor) markAnswers(matchingQuestions []*generated.Matching, answersDto get.AnswerDto) error {
	for questionId, termsAndDefinitions := range answersDto.Matching {
		matchingQuestion, err := findQuestion[*generated.Matching](matchingQuestions, questionId)
		if err != nil {
			return err
		}

		if err = m.markAnswer(questionId, termsAndDefinitions, matchingQuestion); err != nil {
			return err
		}
	}
	return nil
}

func (m *MatchingProcessor) markAndCalculatePoints(
	matchingQuestions []*generated.Matching,
	matchingObjs []*question.Matching,
	answersDto get.AnswerDto,
) (int, error) {
	var points int
	for questionId, termsAndDefinitions := range answersDto.Matching {
		matchingQuestion, err := findQuestion[*generated.Matching](matchingQuestions, questionId)
		if err != nil {
			return 0, err
		}

		if err = m.markAnswer(questionId, termsAndDefinitions, matchingQuestion); err != nil {
			return 0, err
		}

		matchingQuestionObj, err := findQuestionObj[*question.Matching](matchingObjs, questionId)
		if err != nil {
			return 0, err
		}

		points += m.calculateAndSetPoints(matchingQuestion, matchingQuestionObj)
	}

	return points, nil
}

func (m *MatchingProcessor) calculateAndSetPoints(
	matching *generated.Matching,
	matchingObj *question.Matching,
) int {
	termsAndDefinitions := m.extractTermDefinitionPairs(matchingObj)

	var correctAnswers int
	for _, enteredAnswer := range matching.EnteredAnswers {
		if slices.Contains(termsAndDefinitions, enteredAnswer) {
			correctAnswers++
		}
	}

	return m.calculatePoints(matchingObj.Points, correctAnswers)
}

func (m *MatchingProcessor) calculatePoints(matchingPoints []question.MatchingPoint, correctAnswers int) int {
	var points int
	for _, matchingPoint := range matchingPoints {
		if matchingPoint.CorrectAnswers > points && matchingPoint.CorrectAnswers <= correctAnswers {
			points = matchingPoint.CorrectAnswers
		}
	}

	return points
}

func (m *MatchingProcessor) extractTermDefinitionPairs(matchingObj *question.Matching) []generated.EnteredMatchingPair {
	pairs := make([]generated.EnteredMatchingPair, 0)
	for _, matchingTerm := range matchingObj.Terms {
		var pair generated.EnteredMatchingPair
		pair.TermId = matchingTerm.Id
		pair.DefinitionId = matchingTerm.MatchingDefinitionId
		pairs = append(pairs, pair)
	}
	return pairs
}

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
