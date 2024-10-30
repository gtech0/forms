package create

import (
	"github.com/google/uuid"
	"time"
)

type FormPublishDto struct {
	FormPatternId     uuid.UUID      `json:"formPatternId"`
	GroupIds          []uuid.UUID    `json:"groupIds"`
	UserIds           []uuid.UUID    `json:"userIds"`
	Deadline          time.Time      `json:"deadline"`
	Duration          time.Duration  `json:"duration"`
	HideScore         bool           `json:"hideScore"`
	PostModeration    bool           `json:"postModeration"`
	MarkConfiguration map[string]int `json:"markConfiguration"`
}

type UpdateFormPublishedDto struct {
	Deadline          time.Time      `json:"deadline"`
	Duration          time.Duration  `json:"duration"`
	GroupIds          []uuid.UUID    `json:"groupIds"`
	UserIds           []uuid.UUID    `json:"userIds"`
	HideScore         bool           `json:"hideScore"`
	MarkConfiguration map[string]int `json:"markConfiguration"`
}
