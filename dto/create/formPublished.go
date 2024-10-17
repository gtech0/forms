package create

import (
	"github.com/google/uuid"
	"time"
)

type FormPublishedDto struct {
	PatternId         uuid.UUID      `json:"patternId"`
	GroupIds          []uuid.UUID    `json:"groupIds"`
	UserIds           []uuid.UUID    `json:"userIds"`
	Deadline          time.Time      `json:"deadline"`
	Duration          time.Duration  `json:"duration"`
	HideScore         bool           `json:"hideScore"`
	PostModeration    bool           `json:"postModeration"`
	MarkConfiguration map[string]int `json:"markConfiguration"`
}

type FormPublishedBaseDto struct {
	Id                uuid.UUID      `json:"id"`
	Deadline          time.Time      `json:"deadline"`
	Duration          time.Duration  `json:"duration"`
	GroupIds          []uuid.UUID    `json:"groupIds"`
	UserIds           []uuid.UUID    `json:"userIds"`
	PatternId         uuid.UUID      `json:"patternId"`
	HideScore         bool           `json:"hideScore"`
	MarkConfiguration map[string]int `json:"markConfiguration"`
}
