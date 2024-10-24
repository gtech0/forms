package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form"
	"time"
)

type FormGenerated struct {
	model.Base
	Status          form.FormStatus
	FormPublishedID uuid.UUID `gorm:"type:uuid"`
	UserId          uuid.UUID `gorm:"type:uuid"`
	Sections        []Section `gorm:"type:jsonb;serializer:json"`
	Points          int
	Mark            string
	SubmitTime      time.Time
}
