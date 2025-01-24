package service

import (
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
