package service

import (
	"hedgehog-forms/internal/core/repository"
)

type AttachmentService struct {
	attachmentRepository *repository.AttachmentRepository
}

func NewAttachmentService() *AttachmentService {
	return &AttachmentService{
		attachmentRepository: repository.NewAttachmentRepository(),
	}
}
