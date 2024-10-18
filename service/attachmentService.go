package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model"
)

type AttachmentService struct{}

func NewAttachmentService() *AttachmentService {
	return &AttachmentService{}
}

func (f *AttachmentService) validateAttachments(attachmentIds []uuid.UUID) ([]uuid.UUID, error) {
	nonexistentAttachmentIds := make([]uuid.UUID, 0)
	for _, attachmentId := range attachmentIds {
		var attachment model.File
		if err := database.DB.Model(&model.File{}).
			Where("id = ?", attachmentId).
			Find(&attachment).Error; err != nil {
			return nil, errs.New(err.Error(), 500)
		}

		if attachment.Id == uuid.Nil {
			nonexistentAttachmentIds = append(nonexistentAttachmentIds, attachmentId)
		}
	}

	return nonexistentAttachmentIds, nil
}
