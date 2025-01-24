package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type Block struct {
	model.Base
	Title       string
	Description string
	Order       int
	Type        BlockType
	SectionId   uuid.UUID `gorm:"type:uuid"`

	DynamicBlock *DynamicBlock
	StaticBlock  *StaticBlock
}
