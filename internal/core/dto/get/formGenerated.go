package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/generated"
)

type FormGeneratedDto struct {
	Id            uuid.UUID            `json:"id"`
	Status        generated.FormStatus `json:"status"`
	FormPublished FormPublishedBaseDto `json:"formPublished"`
	UserId        *uuid.UUID           `json:"userId"`
	Sections      []generated.Section  `json:"sections"`
}

type MyGeneratedDto struct {
	Id            uuid.UUID            `json:"id"`
	Status        generated.FormStatus `json:"status"`
	FormPublished FormPublishedBaseDto `json:"formPublished"`
	Points        int                  `json:"points"`
	Mark          int                  `json:"mark"`
}

type SubmittedDto struct {
	Status generated.FormStatus `json:"status"`
	Points int                  `json:"points"`
	Mark   int                  `json:"mark"`
}
