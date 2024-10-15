package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type Attachment struct {
	model.Base
	QuestionId uuid.UUID `gorm:"type:uuid"`
}
