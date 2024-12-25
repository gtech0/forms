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
	Teams             []FormPublishedTeam
	Users             []FormPublishedUser
	FormPattern       pattern.FormPattern
	FormPatternId     uuid.UUID `gorm:"type:uuid"`
	FormsGenerated    []*generated.FormGenerated
	ExcludedQuestions []ExcludedQuestion
}

type FormPublishedTeam struct {
	model.Base
	GroupId         uuid.UUID `gorm:"type:uuid"`
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
}

type FormPublishedUser struct {
	model.Base
	UserId          uuid.UUID `gorm:"type:uuid"`
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
}

type ExcludedQuestion struct {
	model.Base
	UserId          uuid.UUID `gorm:"type:uuid"`
	QuestionId      uuid.UUID `gorm:"type:uuid"`
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
}

func (f *FormPublished) ExcludedQuestionsToSlice() []uuid.UUID {
	questions := make([]uuid.UUID, 0)
	for _, excludedQuestion := range f.ExcludedQuestions {
		questions = append(questions, excludedQuestion.QuestionId)
	}
	return questions
}

func (f *FormPublished) GetMarkConfigMap() map[string]int {
	config := make(map[string]int)
	for _, markConfiguration := range f.MarkConfiguration {
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
