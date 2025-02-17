package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/question"
	"hedgehog-forms/pkg/database"
)

type AttachmentRepository struct{}

func NewAttachmentRepository() *AttachmentRepository {
	return &AttachmentRepository{}
}

func (f *AttachmentRepository) FindById(id uuid.UUID) (*question.Attachment, error) {
	attachment := new(question.Attachment)
	if err := database.DB.Model(&question.Attachment{}).
		Where("id = ?", id).
		Find(&attachment).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return attachment, nil
}
