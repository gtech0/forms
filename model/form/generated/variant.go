package generated

import (
	"github.com/google/uuid"
)

type Variant struct {
	Id          uuid.UUID
	Title       string
	Description string
	Questions   []IQuestion
}
