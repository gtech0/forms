package section

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/section/block"
)

type Section struct {
	model.BaseModel
	Title           string
	Description     string
	Order           int
	Blocks          []block.Block
	FormPatternId   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	FormGeneratedId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
}
