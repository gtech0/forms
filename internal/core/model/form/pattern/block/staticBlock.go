package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model"
)

type StaticBlock struct {
	model.Base
	Variants []Variant
	BlockId  uuid.UUID `gorm:"type:uuid"`
}
