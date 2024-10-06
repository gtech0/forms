package form

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/section"
)

type FormPattern struct {
	model.BaseModel
	Title       string
	Description string
	OwnerId     uuid.UUID `gorm:"type:uuid"`
	SubjectId   uuid.UUID `gorm:"type:uuid"`
	Sections    []section.Section
}
