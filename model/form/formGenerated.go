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
	FormPublishedID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Sections        []section.Section
	Points          int
	Mark            string
	SubmitTime      time.Time
}
