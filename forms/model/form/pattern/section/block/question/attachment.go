package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type Attachment struct {
	model.Base
	File       model.File
	FileId     uuid.UUID `gorm:"type:uuid"`
	QuestionId uuid.UUID `gorm:"type:uuid"`
}
