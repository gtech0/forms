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
	Groups            []Group `gorm:"many2many:form_groups"`
	Users             []User  `gorm:"many2many:form_users"`
	FormPattern       pattern.FormPattern
	FormPatternId     uuid.UUID `gorm:"type:uuid"`
	FormsGenerated    []generated.FormGenerated
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
