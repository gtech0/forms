package get

import (
	"github.com/google/uuid"
)

type FormPatternBaseDto struct {
	Id          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	OwnerId     uuid.NullUUID `json:"ownerId"`
	Subject     SubjectDto    `json:"subject"`
}

type FormPatternDto struct {
	Id          uuid.UUID     `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	OwnerId     uuid.NullUUID `json:"ownerId"`
	Subject     SubjectDto    `json:"subject"`
	Sections    []SectionDto  `json:"sections"`
}
