package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type IBlock interface {
	GetId() uuid.UUID
	GetType() BlockType
	SetOrder(int)
}

type Block struct {
	model.Base
	Title       string
	Description string
	Order       int
	Type        BlockType
	SectionId   uuid.UUID `gorm:"type:uuid"`
}

func (b *Block) GetId() uuid.UUID {
	return b.Id
}

func (b *Block) GetType() BlockType {
	return b.Type
}

func (b *Block) SetOrder(order int) {
	b.Order = order
}
