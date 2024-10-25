package create

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
)

type CheckDto struct {
	Status generated.FormStatus `json:"status"`
	Points map[uuid.UUID]int    `json:"points"`
}
