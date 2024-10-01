package question

import (
	"github.com/google/uuid"
)

type TextInput struct {
	Question
	Points          int
	IsCaseSensitive bool
}

type TextInputAnswer struct {
	Answer      string
	TextInputId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
