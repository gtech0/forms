package generated

import (
	"github.com/google/uuid"
)

type SingleChoice struct {
	Question
	Options       []SingleChoiceOption `json:"options"`
	EnteredAnswer uuid.UUID            `json:"enteredAnswer"`
}

type SingleChoiceOption struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
