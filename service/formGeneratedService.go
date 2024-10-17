package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/model/form/generated"
)

type FormGeneratedService struct{}

func NewFormGeneratedService() *FormGeneratedService {
	return &FormGeneratedService{}
}

func (f *FormGeneratedService) getForm(id uuid.UUID) (generated.FormGenerated, error) {
	var formGenerated generated.FormGenerated
	if err := database.DB.Model(&generated.FormGenerated{}).
		Where("id = ?", id).
		First(&formGenerated).
		Error; err != nil {
		return generated.FormGenerated{}, err
	}
	return formGenerated, nil
}
