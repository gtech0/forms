package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type Attachment struct {
	model.Base
	Question   Question
	QuestionId uuid.UUID `gorm:"type:uuid"`
	File       model.File
	FileId     uuid.UUID `gorm:"type:uuid"`
}
