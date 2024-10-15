package generated

import (
	"github.com/google/uuid"
)

type MultipleChoice struct {
	Question
	Options        []MultipleChoiceOption
	EnteredAnswers []uuid.UUID
}

type MultipleChoiceOption struct {
	Id   uuid.UUID
	Text string
}
