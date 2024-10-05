package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/section/block/question"
)

type IBlock interface {
	GetId() uuid.UUID
	SetOrder(int)
}

type Block struct {
	model.BaseModel
	Title       string
	Description string
	Order       int
	Type        BlockType
	SectionId   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

func (b *Block) GetId() uuid.UUID {
	return b.Id
}

func (b *Block) SetOrder(order int) {
	b.Order = order
}

type DynamicBlock struct {
	Block
	Questions []question.IQuestion
}

type StaticBlock struct {
	Block
	Variants []Variant
}
