package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block"
)

type StaticBlockRepository struct{}

func NewStaticBlockRepository() *StaticBlockRepository {
	return &StaticBlockRepository{}
}

func (b *StaticBlockRepository) GetById(id uuid.UUID) (*block.StaticBlock, error) {
	staticBlock := new(block.StaticBlock)
	if err := database.DB.Model(&block.StaticBlock{}).
		Where("id = ?", id).
		First(staticBlock).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return staticBlock, nil
}
