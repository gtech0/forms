package get

import "github.com/google/uuid"

type MultipleChoiceDto struct {
	QuestionDto
	Points  map[int]int         `json:"points"`
	Options []MultipleOptionDto `json:"options"`
}

type MultipleOptionDto struct {
	Id       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	IsAnswer bool      `json:"isAnswer"`
}
