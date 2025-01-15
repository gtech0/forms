package create

import (
	"encoding/json"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type MarkConfiguration map[int]int

type FormPublishDto struct {
	FormPatternId     uuid.UUID         `json:"formPatternId"`
	TeamIds           []uuid.UUID       `json:"teamIds"`
	UserIds           []uuid.UUID       `json:"userIds"`
	Deadline          time.Time         `json:"deadline"`
	Duration          time.Duration     `json:"duration"`
	HideScore         bool              `json:"hideScore"`
	PostModeration    bool              `json:"postModeration"`
	MaxAttempts       int               `json:"maxAttempts"`
	MarkConfiguration MarkConfiguration `json:"markConfiguration"`
}

type UpdateFormPublishedDto struct {
	Deadline          time.Time         `json:"deadline"`
	Duration          time.Duration     `json:"duration"`
	TeamIds           []uuid.UUID       `json:"teamIds"`
	UserIds           []uuid.UUID       `json:"userIds"`
	HideScore         bool              `json:"hideScore"`
	MarkConfiguration MarkConfiguration `json:"markConfiguration"`
}

func (m *MarkConfiguration) UnmarshalJSON(bytes []byte) error {
	raw := make(map[string]int)
	if err := json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	newMap := make(map[int]int)
	for key, value := range raw {
		newKey, err := strconv.Atoi(key)
		if err != nil {
			return err
		}
		newMap[newKey] = value
	}
	*m = newMap
	return nil
}
