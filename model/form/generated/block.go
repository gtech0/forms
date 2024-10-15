package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/pattern/section/block"
)

type Block struct {
	Id          uuid.UUID
	Type        block.BlockType
	Title       string
	Description string
	Variant     Variant
	Questions   []IQuestion
}
