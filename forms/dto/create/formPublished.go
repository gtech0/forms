package create

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/common"
	"time"
)

type FormPublishDto struct {
	FormPatternId     uuid.UUID                `json:"formPatternId"`
	TeamIds           []uuid.UUID              `json:"teamIds"`
	UserIds           []uuid.UUID              `json:"userIds"`
	Deadline          time.Time                `json:"deadline"`
	Duration          time.Duration            `json:"duration"`
	HideScore         bool                     `json:"hideScore"`
	PostModeration    bool                     `json:"postModeration"`
	MaxAttempts       int                      `json:"maxAttempts"`
	MarkConfiguration common.MarkConfiguration `json:"markConfiguration"`
}

type UpdateFormPublishedDto struct {
	Deadline          time.Time                `json:"deadline"`
	Duration          time.Duration            `json:"duration"`
	TeamIds           []uuid.UUID              `json:"teamIds"`
	UserIds           []uuid.UUID              `json:"userIds"`
	HideScore         bool                     `json:"hideScore"`
	MarkConfiguration common.MarkConfiguration `json:"markConfiguration"`
}
