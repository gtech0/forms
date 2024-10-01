package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/section/block/question"
)

type Block struct {
	model.BaseModel
	Title       string
	Description string
	Order       int
	Type        BlockType
	SectionId   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}

type DynamicBlock struct {
	Block
	BlockObjects []question.Question
}

type StaticBlock struct {
	Block
	Variants []Variant
}
