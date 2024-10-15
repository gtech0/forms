package published

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern"
	"time"
)

type FormPublished struct {
	model.Base
	Deadline          time.Time
	Duration          time.Duration
	HideScore         bool
	PostModeration    bool
	MarkConfiguration MarkConfiguration
	Groups            []FormPublishedGroup
	Users             []FormPublishedUser
	FormPattern       pattern.FormPattern
	FormPatternId     uuid.UUID `gorm:"type:uuid"`
	FormsGenerated    []generated.FormGenerated
}

type MarkConfiguration struct {
	Mark            string
	MinPoints       int
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
