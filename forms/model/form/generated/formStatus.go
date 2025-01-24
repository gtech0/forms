package generated

import (
	"slices"
)

type FormStatus string

const (
	NEW         FormStatus = "NEW"
	IN_PROGRESS FormStatus = "IN_PROGRESS"
	RETURNED    FormStatus = "RETURNED"
	SUBMITTED   FormStatus = "SUBMITTED"
	COMPLETED   FormStatus = "COMPLETED"
)

func (s FormStatus) String() string {
	return string(s)
}

func CheckStatusAndGet(str string) FormStatus {
	statusSlice := []string{NEW.String(), IN_PROGRESS.String(), SUBMITTED.String(), COMPLETED.String()}
	if slices.Contains(statusSlice, str) {
		return FormStatus(str)
	}

	return IN_PROGRESS
}
