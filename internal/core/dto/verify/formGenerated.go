package verify

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/generated"
)

type FormGenerated struct {
	Id       uuid.UUID            `json:"id"`
	Status   generated.FormStatus `json:"status"`
	UserId   *uuid.UUID           `json:"userId"`
	Sections []Section            `json:"sections"`
}
