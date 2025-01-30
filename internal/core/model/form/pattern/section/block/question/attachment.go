package question

import (
	"github.com/google/uuid"
	model2 "hedgehog-forms/internal/core/model"
)

type Attachment struct {
	model2.Base
	File       model2.File
	FileId     uuid.UUID `gorm:"type:uuid"`
	QuestionId uuid.UUID `gorm:"type:uuid"`
}
