package form

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/section"
	"time"
)

type FormGenerated struct {
	model.BaseModel
	Status          FormStatus
	FormPublishedID uuid.UUID `gorm:"type:uuid"`
	UserId          uuid.UUID `gorm:"type:uuid"`
	Sections        []section.Section
	Points          int
	Mark            string
	SubmitTime      time.Time
}
