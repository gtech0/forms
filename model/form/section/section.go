package section

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/section/block"
)

type Section struct {
	model.BaseModel
	Title           string
	Description     string
	Order           int
	Blocks          []block.IBlock `gorm:"-"`
	DynamicBlocks   []*block.DynamicBlock
	StaticBlocks    []*block.StaticBlock
	FormPatternId   uuid.UUID     `gorm:"type:uuid"`
	FormGeneratedId uuid.NullUUID `gorm:"type:uuid"`
}

func (d *Section) BeforeSave(*gorm.DB) error {
	for _, iBlock := range d.Blocks {
		switch iBlock.GetType() {
		case block.DYNAMIC:
			d.DynamicBlocks = append(d.DynamicBlocks, iBlock.(*block.DynamicBlock))
		case block.STATIC:
			d.StaticBlocks = append(d.StaticBlocks, iBlock.(*block.StaticBlock))
		}
	}
	return nil
}
