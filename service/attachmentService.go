package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form"
)

type AttachmentService struct {
	patternService *PatternService
}

func NewAttachmentService() *AttachmentService {
	return &AttachmentService{
		patternService: NewPatternService(),
	}
}

func (f *AttachmentService) ValidatePatternAttachments(pattern form.FormPattern) ([]uuid.UUID, error) {
	questions := f.patternService.ExtractQuestionsFromPattern(pattern)
	attachmentIds := make([]uuid.UUID, 0)
	for _, iQuestion := range questions {
		for _, attachment := range iQuestion.GetAttachments() {
			attachmentIds = append(attachmentIds, attachment.Id)
		}
	}

	return f.validateAttachments(attachmentIds)
}

func (f *AttachmentService) validateAttachments(attachmentIds []uuid.UUID) ([]uuid.UUID, error) {
	nonexistentAttachmentIds := make([]uuid.UUID, 0)
	for _, attachmentId := range attachmentIds {
		var exists bool
		if err := database.DB.Model(&model.File{}).
			Where("id = ?", attachmentId).
			Find(&exists).Error; err != nil {
			return nil, err
		}

		if !exists {
			nonexistentAttachmentIds = append(nonexistentAttachmentIds, attachmentId)
		}
	}

	return nonexistentAttachmentIds, nil
}
