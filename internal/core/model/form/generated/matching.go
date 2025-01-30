package generated

import (
	"github.com/google/uuid"
)

type Matching struct {
	Question
	Terms          []Term                `json:"terms"`
	Definitions    []Definition          `json:"definitions"`
	EnteredAnswers []EnteredMatchingPair `json:"enteredAnswers"`
}

type EnteredMatchingPair struct {
	TermId       uuid.UUID `json:"termId"`
	DefinitionId uuid.UUID `json:"definitionId"`
}

type Term struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

type Definition struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
