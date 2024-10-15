package section

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/pattern/section/block"
	"slices"
)

type Section struct {
	model.Base
	Title           string
	Description     string
	Order           int
	Blocks          []block.IBlock `gorm:"-"`
	DynamicBlocks   block.DynamicBlockSlice
	StaticBlocks    block.StaticBlockSlice
	FormPatternId   uuid.UUID     `gorm:"type:uuid"`
	FormGeneratedId uuid.NullUUID `gorm:"type:uuid"`
}

func (s *Section) BeforeSave(*gorm.DB) error {
	for _, iBlock := range s.Blocks {
		switch iBlock.GetType() {
		case block.DYNAMIC:
			s.DynamicBlocks = append(s.DynamicBlocks, iBlock.(*block.DynamicBlock))
		case block.STATIC:
			s.StaticBlocks = append(s.StaticBlocks, iBlock.(*block.StaticBlock))
		}
	}
	return nil
}

func (s *Section) AfterFind(*gorm.DB) error {
	s.Blocks = slices.Concat(s.DynamicBlocks.ToInterface(), s.StaticBlocks.ToInterface())
	return nil
}
