package get

import (
	"github.com/google/uuid"
	"time"
)

type FormPublishedBaseDto struct {
	Id                uuid.UUID      `json:"id"`
	Deadline          time.Time      `json:"deadline"`
	Duration          time.Duration  `json:"duration"`
	GroupIds          []uuid.UUID    `json:"groupIds"`
	UserIds           []uuid.UUID    `json:"userIds"`
	FormPatternId     uuid.UUID      `json:"formPatternId"`
	HideScore         bool           `json:"hideScore"`
	MarkConfiguration map[string]int `json:"markConfiguration"`
}

type FormPublishedDto struct {
	Id          uuid.UUID      `json:"id"`
	Deadline    time.Time      `json:"deadline"`
	Duration    time.Duration  `json:"duration"`
	GroupIds    []uuid.UUID    `json:"groupIds"`
	UserIds     []uuid.UUID    `json:"userIds"`
	FormPattern FormPatternDto `json:"formPattern"`
	HideScore   bool           `json:"hideScore"`
}
