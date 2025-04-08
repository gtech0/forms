package published

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern"
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
	FormPattern       pattern.FormPattern
	FormPatternId     uuid.UUID `gorm:"type:uuid"`
	FormsGenerated    []*generated.FormGenerated
	ExcludedQuestions []ExcludedQuestion
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

func (f *FormPublished) GetMarkConfigMap() map[int]int {
	config := make(map[int]int)
	for _, markConfiguration := range f.MarkConfiguration {
		config[markConfiguration.Mark] = markConfiguration.MinPoints
	}
	return config
}

type MarkConfiguration struct {
	model.Base
	Mark            int
	MinPoints       int
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
}
