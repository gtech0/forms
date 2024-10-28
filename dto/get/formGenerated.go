package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
	"time"
)

type FormGeneratedDto struct {
	Id            uuid.UUID            `json:"id"`
	Status        generated.FormStatus `json:"status"`
	FormPublished FormPublishedBaseDto `json:"formPublished"`
	UserId        uuid.UUID            `json:"userId"`
	Sections      []generated.Section  `json:"sections"`
}

type MyGeneratedDto struct {
	Id            uuid.UUID            `json:"id"`
	Status        generated.FormStatus `json:"status"`
	FormPublished FormPublishedBaseDto `json:"formPublished"`
	SubmitTime    time.Time            `json:"submitTime"`
	Points        int                  `json:"points"`
	Mark          string               `json:"mark"`
}

type SubmittedDto struct {
	Status     generated.FormStatus `json:"status"`
	Points     int                  `json:"points"`
	Mark       string               `json:"mark"`
	SubmitTime time.Time            `json:"submitTime"`
}
