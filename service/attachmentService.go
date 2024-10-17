package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/pattern"
)

type AttachmentService struct {
	patternService *FormPatternService
}

func NewAttachmentService() *AttachmentService {
	return &AttachmentService{
		patternService: NewFormPatternService(),
	}
}

func (f *AttachmentService) ValidatePatternAttachments(pattern pattern.FormPattern) ([]uuid.UUID, error) {
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
		var attachment model.File
		if err := database.DB.Model(&model.File{}).
			Where("id = ?", attachmentId).
			Find(&attachment).Error; err != nil {
			return nil, err
		}

		if attachment.Id == uuid.Nil {
			nonexistentAttachmentIds = append(nonexistentAttachmentIds, attachmentId)
		}
	}

	return nonexistentAttachmentIds, nil
}
