package service

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"hedgehog-forms/repository"
)

type AttachmentService struct {
	attachmentRepository *repository.AttachmentRepository
}

func NewAttachmentService() *AttachmentService {
	return &AttachmentService{
		attachmentRepository: repository.NewAttachmentRepository(),
	}
}

func (f *AttachmentService) ValidateAttachments(questions ...*question.Question) error {
	attachmentIds := make([]uuid.UUID, 0)
	for _, questionEntity := range questions {
		for _, attachment := range questionEntity.Attachments {
			attachmentIds = append(attachmentIds, attachment.Id)
		}
	}

	nonexistentAttachmentIds := make([]uuid.UUID, 0)
	for _, attachmentId := range attachmentIds {
		attachment, err := f.attachmentRepository.FindById(attachmentId)
		if err != nil {
			return err
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
