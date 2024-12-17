package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type SolutionObject struct {
	model.Base
	NumberAttempts int        `json:"numberAttempts"`
	IsIndividual   bool       `json:"isIndividual"`
	TeamOwnerID    uuid.UUID  `json:"teamOwnerId"`
	UserOwnerID    uuid.UUID  `json:"userOwnerId"`
	Score          int        `json:"score"`
	Mark           string     `json:"mark"`
	Verdict        FormStatus `json:"verdict"`
}
