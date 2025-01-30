package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/generated"
)

type AnswerDto struct {
	SingleChoice   map[uuid.UUID]uuid.UUID                       `json:"singleChoice"`
	MultipleChoice map[uuid.UUID][]uuid.UUID                     `json:"multipleChoice"`
	TextInput      map[uuid.UUID]string                          `json:"textInput"`
	Matching       map[uuid.UUID][]generated.EnteredMatchingPair `json:"matching"`
}
