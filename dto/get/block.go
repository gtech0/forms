package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/pattern/section/block"
)

type IBlockDto interface {
	GetType() block.BlockType
}

type BlockDto struct {
	Id          uuid.UUID       `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Type        block.BlockType `json:"type"`
}

func (b *BlockDto) GetType() block.BlockType {
	return b.Type
}

type DynamicBlockDto struct {
	BlockDto
	Questions []IQuestionDto `json:"questions"`
}

type StaticBlockDto struct {
	BlockDto
	Variants []VariantDto `json:"variants"`
}
