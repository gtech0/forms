package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"time"
)

type Submission struct {
	model.Base
	UserId                *uuid.UUID
	StartTime             time.Time
	SubmitTime            *time.Time
	PostmoderationVerdict string    `gorm:"default:PENDING"`
	SolutionId            uuid.UUID `gorm:"type:uuid"`
	FormGeneratedId       uuid.UUID `gorm:"type:uuid"`
}
