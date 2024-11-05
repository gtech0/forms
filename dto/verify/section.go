package verify

import "github.com/google/uuid"

type Section struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Blocks      []Block   `json:"blocks"`
}
