package get

import "github.com/google/uuid"

type SubjectDto struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
