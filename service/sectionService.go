package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/model/form/section"
)

type SectionService struct{}

func NewSectionService() *SectionService {
	return &SectionService{}
}

func (b *SectionService) GetSectionObjectById(id uuid.UUID) (section.Section, error) {
	var obj section.Section
	if err := database.DB.Model(section.Section{}).
		Where("id = ?", id).
		First(&obj).Error; err != nil {
		return section.Section{}, err
	}

	return obj, nil
}
