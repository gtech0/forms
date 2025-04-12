package get

import (
	"hedgehog-forms/internal/core/model/form/generated"
)

type IntegratedMatchingDto struct {
	IntegratedQuestionDto
	Terms               []generated.Term                `json:"terms"`
	Definitions         []generated.Definition          `json:"definitions"`
	TermsAndDefinitions []MatchingTermDefinitionDto     `json:"termsAndDefinitions"`
	EnteredAnswers      []generated.EnteredMatchingPair `json:"enteredAnswers"`
}
