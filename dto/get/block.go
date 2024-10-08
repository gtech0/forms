package get

import "github.com/google/uuid"

type IBlockDto interface {
}

type BlockDto struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}

type DynamicBlockDto struct {
	BlockDto
	Questions []IQuestionDto `json:"questions"`
}

type StaticBlockDto struct {
	BlockDto
	Variants []VariantDto `json:"variants"`
}
