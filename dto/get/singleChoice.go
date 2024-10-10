package get

import "github.com/google/uuid"

type SingleChoiceDto struct {
	QuestionDto
	Points  int               `json:"points"`
	Choices []SingleOptionDto `json:"choices"`
}

type SingleOptionDto struct {
	Id       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	IsAnswer bool      `json:"isAnswer"`
}
