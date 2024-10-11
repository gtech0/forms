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

type MatchingPoint struct {
	model.BaseModel
	CorrectAnswer int
	Points        int
	MatchingId    uuid.UUID `gorm:"type:uuid"`
}

type MatchingSlice []*Matching

func (s *MatchingSlice) ToInterface() []IQuestion {
	questions := make([]IQuestion, 0)
	for _, matching := range *s {
		questions = append(questions, matching)
	}
	return questions
}
