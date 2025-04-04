package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model"
)

type Matching struct {
	Base
	Definitions []MatchingDefinition
	Terms       []MatchingTerm
	Points      []MatchingPoints
}

type MatchingTerm struct {
	model.Base
	Text                 string
	MatchingId           uuid.UUID `gorm:"type:uuid"`
	MatchingDefinitionId uuid.UUID `gorm:"type:uuid"`
}

type MatchingDefinition struct {
	model.Base
	Text         string
	MatchingId   uuid.UUID `gorm:"type:uuid"`
	MatchingTerm MatchingTerm
}

type MatchingPoints struct {
	model.Base
	CorrectAnswers int
	Points         int
	MatchingId     uuid.UUID `gorm:"type:uuid"`
}
