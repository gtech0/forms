package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type Matching struct {
	Question
	Terms       []MatchingTerm
	Definitions []MatchingDefinition
	Points      []MatchingPoint
}

type MatchingPoint struct {
	model.BaseModel
	CorrectAnswer int
	Points        int
	MatchingId    uuid.UUID `gorm:"type:uuid"`
}

type MatchingTerm struct {
	model.BaseModel
	Text                 string
	MatchingId           uuid.UUID `gorm:"type:uuid"`
	MatchingDefinitionId uuid.UUID `gorm:"type:uuid"`
}

type MatchingDefinition struct {
	model.BaseModel
	Text       string
	MatchingId uuid.UUID `gorm:"type:uuid"`
	Term       MatchingTerm
}
