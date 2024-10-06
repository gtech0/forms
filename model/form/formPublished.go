package form

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"time"
)

type FormPublished struct {
	model.BaseModel
	Deadline       time.Time
	Duration       time.Duration
	HideScore      bool
	PostModeration bool
	Groups         []FormPublishedGroup
	Users          []FormPublishedUser
	FormPatternId  uuid.UUID `gorm:"type:uuid"`
	FormsGenerated []FormGenerated
}

type MarkConfiguration struct {
	MinPoints       int
	Mark            string
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
}

type FormPublishedGroup struct {
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
	GroupId         uuid.UUID `gorm:"type:uuid"`
}

type FormPublishedUser struct {
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
	UserId          uuid.UUID `gorm:"type:uuid"`
}
