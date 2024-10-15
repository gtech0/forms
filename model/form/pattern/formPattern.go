package pattern

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/pattern/section"
)

type FormPattern struct {
	model.Base
	Title       string
	Description string
	OwnerId     uuid.NullUUID `gorm:"type:uuid"`
	Subject     model.Subject
	SubjectId   uuid.UUID `gorm:"type:uuid"`
	Sections    []section.Section
}
