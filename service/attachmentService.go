package service

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type AttachmentService struct{}

func NewAttachmentService() *AttachmentService {
	return &AttachmentService{}
}

func (f *AttachmentService) ValidateAttachments(questions ...question.IQuestion) error {
	attachmentIds := make([]uuid.UUID, 0)
	for _, iQuestion := range questions {
		for _, attachment := range iQuestion.GetAttachments() {
			attachmentIds = append(attachmentIds, attachment.Id)
		}
	}

	nonexistentAttachmentIds := make([]uuid.UUID, 0)
	for _, attachmentId := range attachmentIds {
		var attachment model.File
		if err := database.DB.Model(&model.File{}).
			Where("id = ?", attachmentId).
			Find(&attachment).Error; err != nil {
			return errs.New(err.Error(), 500)
		}

		if attachment.Id == uuid.Nil {
			nonexistentAttachmentIds = append(nonexistentAttachmentIds, attachmentId)
		}
	}

	if len(nonexistentAttachmentIds) > 0 {
		return errs.New(fmt.Sprintf("incorrect attachment ids: %v", nonexistentAttachmentIds), 400)
	}

	return nil
}
