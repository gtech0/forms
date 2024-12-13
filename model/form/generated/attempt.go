package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"time"
)

type Attempt struct {
	model.Base
	IsComplete      bool
	Sections        []Section `gorm:"type:jsonb;serializer:json"`
	Points          int
	Mark            string
	StartTime       time.Time
	SubmitTime      time.Time
	FormGeneratedId uuid.UUID `gorm:"type:uuid"`
}
