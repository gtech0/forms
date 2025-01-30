package section

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model"
	"hedgehog-forms/internal/core/model/form/pattern/section/block"
)

type Section struct {
	model.Base
	Title           string
	Description     string
	Order           int
	Blocks          []*block.Block
	FormPatternId   uuid.UUID     `gorm:"type:uuid"`
	FormGeneratedId uuid.NullUUID `gorm:"type:uuid"`
}
