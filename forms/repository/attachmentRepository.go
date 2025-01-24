package repository

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
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
