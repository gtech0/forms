package create

import (
	"encoding/json"
	"strconv"
)

type MatchingPoints map[int]int

type MatchingQuestionDto struct {
	NewQuestionDto
	TermsAndDefinitions map[string]string `json:"termsAndDefinitions"`
	Points              MatchingPoints    `json:"points"`
}

func (m *MatchingPoints) UnmarshalJSON(bytes []byte) error {
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
