package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/common"
	"time"
)

type FormPublishedBaseDto struct {
	Id                uuid.UUID                `json:"id"`
	Deadline          time.Time                `json:"deadline"`
	Duration          time.Duration            `json:"duration"`
	TeamIds           []uuid.UUID              `json:"teamIds"`
	UserIds           []uuid.UUID              `json:"userIds"`
	FormPatternId     uuid.UUID                `json:"formPatternId"`
	HideScore         bool                     `json:"hideScore"`
	MaxAttempts       int                      `json:"maxAttempts"`
	MarkConfiguration common.MarkConfiguration `json:"markConfiguration"`
}

type FormPublishedDto struct {
	Id          uuid.UUID      `json:"id"`
	Deadline    time.Time      `json:"deadline"`
	Duration    time.Duration  `json:"duration"`
	TeamIds     []uuid.UUID    `json:"teamIds"`
	UserIds     []uuid.UUID    `json:"userIds"`
	FormPattern FormPatternDto `json:"formPattern"`
	HideScore   bool           `json:"hideScore"`
}
