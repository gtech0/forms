package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type Attachment struct {
	model.BaseModel
	Description string
	QuestionId  uuid.UUID `gorm:"type:uuid"`
}
