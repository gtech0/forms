package util

import (
	"encoding/json"
)

func commonUnmarshal() {
	//TODO
}

func CommonMarshal[T any](data []T) ([]json.RawMessage, error) {
	rawMessage := make([]json.RawMessage, 0)
	if data != nil {
		for _, dto := range data {
			raw, err := json.Marshal(dto)
			if err != nil {
				return nil, err
			}
			rawMessage = append(rawMessage, raw)
		}
	}
	return rawMessage, nil
}
