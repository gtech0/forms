package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
)

type AnswerDto struct {
	SingleChoice   map[uuid.UUID]uuid.UUID
	MultipleChoice map[uuid.UUID][]uuid.UUID
	TextInput      map[uuid.UUID]string
	Matching       map[uuid.UUID][]generated.EnteredMatchingPair
}
