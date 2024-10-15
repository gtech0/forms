package generated

import (
	"github.com/google/uuid"
)

type SingleChoice struct {
	Question
	Options       []SingleChoiceOption
	EnteredAnswer uuid.UUID
}

type SingleChoiceOption struct {
	Id   uuid.UUID
	Text string
}
