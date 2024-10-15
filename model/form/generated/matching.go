package generated

import (
	"github.com/google/uuid"
)

type Matching struct {
	Question
	Terms          []Term
	Definitions    []Definition
	EnteredAnswers []EnteredMatchingPair
}

type EnteredMatchingPair struct {
	TermId       uuid.UUID
	DefinitionId uuid.UUID
}

type Term struct {
	Id   uuid.UUID
	Text string
}

type Definition struct {
	Id   uuid.UUID
	Text string
}
