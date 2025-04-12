package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/generated"
)

type IntegrationGeneratedFormDto struct {
	Id       uuid.UUID               `json:"id"`
	Status   generated.FormStatus    `json:"status"`
	Sections []IntegrationSectionDto `json:"sections"`
}

type FormGeneratedDto struct {
	Id     uuid.UUID            `json:"id"`
	Status generated.FormStatus `json:"status"`
	//FormPublished FormPublishedBaseDto `json:"formPublished"`
	Sections []generated.Section `json:"sections"`
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
