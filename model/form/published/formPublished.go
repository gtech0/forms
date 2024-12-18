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
	MaxAttempts       int
	MarkConfiguration []MarkConfiguration
	Groups            []FormPublishedGroup
	Users             []FormPublishedUser
	FormPattern       pattern.FormPattern
	FormPatternId     uuid.UUID `gorm:"type:uuid"`
	FormsGenerated    []*generated.FormGenerated
}

type FormPublishedGroup struct {
	model.Base
	GroupId         uuid.UUID `gorm:"type:uuid"`
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
}

type FormPublishedUser struct {
	model.Base
	UserId          uuid.UUID `gorm:"type:uuid"`
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
}

func (m *FormPublished) GetMarkConfigMap() map[string]int {
	config := make(map[string]int)
	for _, markConfiguration := range m.MarkConfiguration {
		config[markConfiguration.Mark] = markConfiguration.MinPoints
	}
	return config
}

type MarkConfiguration struct {
	model.Base
	Mark            string
	MinPoints       int
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
}
