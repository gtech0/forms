package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/published"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) FindById(id uuid.UUID) (*published.User, error) {
	user := new(published.User)
	if err := database.DB.Model(&published.User{}).
		Where("id = ?", id).
		First(user).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return user, nil
}
