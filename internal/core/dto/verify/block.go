package verify

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/pattern/section/block"
)

type Block struct {
	Id          uuid.UUID       `json:"id"`
	Type        block.BlockType `json:"type"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Variant     *Variant        `json:"variant"`
	Questions   []Question      `json:"questions"`
}
