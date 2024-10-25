package generated

import (
	"github.com/google/uuid"
)

type Section struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Blocks      []*Block  `json:"blocks"`
}
