package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type StaticBlock struct {
	model.Base
	Variants []Variant
	BlockId  uuid.UUID `gorm:"type:uuid"`
}
