package get

import "github.com/google/uuid"

type MatchingDto struct {
	QuestionDto
	Points              map[int]int         `json:"points"`
	TermsAndDefinitions []TermDefinitionDto `json:"termsAndDefinitions"`
}

type TermDefinitionDto struct {
	Term       MatchingOptionDto `json:"term"`
	Definition MatchingOptionDto `json:"definition"`
}

type MatchingOptionDto struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
