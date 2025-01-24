package util

import (
	"encoding/json"
	"hedgehog-forms/errs"
)

func CommonMarshal[T any](data []T) ([]json.RawMessage, error) {
	rawMessage := make([]json.RawMessage, 0)
	if data != nil {
		for _, dto := range data {
			raw, err := json.Marshal(dto)
			if err != nil {
				return nil, errs.New(err.Error(), 500)
			}
			rawMessage = append(rawMessage, raw)
		}
	}
	return rawMessage, nil
}
