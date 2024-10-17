package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/factory"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
)

type FormGeneratedService struct {
	formGeneratedFactory *factory.FormGeneratedFactory
}

func NewFormGeneratedService() *FormGeneratedService {
	return &FormGeneratedService{
		formGeneratedFactory: factory.NewFormGeneratedFactory(),
	}
}

func (f *FormGeneratedService) getGeneratedForm(id uuid.UUID) (generated.FormGenerated, error) {
	var formGenerated generated.FormGenerated
	if err := database.DB.Model(&generated.FormGenerated{}).
		Where("id = ?", id).
		First(&formGenerated).
		Error; err != nil {
		return generated.FormGenerated{}, err
	}
	return formGenerated, nil
}

func (f *FormGeneratedService) buildAndCreate(
	formPublished published.FormPublished,
	userId uuid.UUID,
) (generated.FormGenerated, error) {
	generatedForm, err := f.formGeneratedFactory.BuildForm(formPublished, userId)
	if err != nil {
		return generated.FormGenerated{}, err
	}

	return generatedForm, nil
}
