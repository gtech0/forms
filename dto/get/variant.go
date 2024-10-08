package get

import "github.com/google/uuid"

type VariantDto struct {
	Id          uuid.UUID      `json:"id"`
	Title       string         `json:"title"`
	Description string         `json:"description"`
	Questions   []IQuestionDto `json:"questions"`
}
