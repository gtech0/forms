package common

import (
	"encoding/json"
	"strconv"
)

type MarkConfiguration map[int]int

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
