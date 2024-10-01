package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type Matching struct {
	Question
	Terms       []MatchingTerm
	Definitions []MatchingDefinition
}

type MatchingPoints struct {
	CorrectAnswers int
	Points         int
	MatchingId     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

type MatchingTerm struct {
	model.BaseModel
	Text                 string
	MatchingId           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	MatchingDefinitionId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

type MatchingDefinition struct {
	model.BaseModel
	Text       string
	MatchingId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
