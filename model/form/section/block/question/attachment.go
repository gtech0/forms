package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type Attachment struct {
	model.BaseModel
	QuestionId uuid.UUID `gorm:"type:uuid"`
}
