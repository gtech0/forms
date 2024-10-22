package generated

import (
	"github.com/google/uuid"
)

type MultipleChoice struct {
	Question
	Options        []MultipleChoiceOption `json:"options"`
	EnteredAnswers []uuid.UUID            `json:"enteredAnswers"`
}

type MultipleChoiceOption struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
