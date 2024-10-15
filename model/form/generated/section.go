package generated

import (
	"github.com/google/uuid"
)

type Section struct {
	Id          uuid.UUID
	Title       string
	Description string
	Blocks      []Block
}
