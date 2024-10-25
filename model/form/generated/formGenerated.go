package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"time"
)

type FormGenerated struct {
	model.Base
	Status          FormStatus
	FormPublishedID uuid.UUID `gorm:"type:uuid"`
	UserId          uuid.UUID `gorm:"type:uuid"`
	Sections        []Section `gorm:"type:jsonb;serializer:json"`
	Points          int
	Mark            string
	SubmitTime      time.Time
}
