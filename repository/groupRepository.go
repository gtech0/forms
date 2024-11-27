package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/published"
)

type GroupRepository struct{}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{}
}

func (g *GroupRepository) FindById(id uuid.UUID) (*published.Group, error) {
	group := new(published.Group)
	if err := database.DB.Model(&published.Group{}).
		Where("id = ?", id).
		First(group).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return group, nil
}
