package util

import (
	"github.com/google/uuid"
	"hedgehog-forms/errs"
)

func IdCheckAndParse(id string) (uuid.UUID, error) {
	if id == "" {
		return uuid.Nil, errs.New("id is required", 400)
	}

	parsedId, err := uuid.Parse(id)
	if err != nil {
		return uuid.Nil, errs.New(err.Error(), 400)
	}

	return parsedId, nil
}
