package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"slices"
	"time"
)

type FormGenerated struct {
	model.Base
	Status            FormStatus
	UserId            uuid.UUID `gorm:"type:uuid"`
	IsCompleted       bool
	Sections          []Section `gorm:"type:jsonb;serializer:json"`
	Points            int
	Mark              string
	StartTime         time.Time
	SubmitTime        time.Time
	ExcludedQuestions []ExcludedQuestion
	FormPublishedID   uuid.UUID `gorm:"type:uuid"`
}

func (f *FormGenerated) ExcludedQuestionsToSlice() []uuid.UUID {
	questions := make([]uuid.UUID, 0)
	for _, excludedQuestion := range f.ExcludedQuestions {
		questions = append(questions, excludedQuestion.QuestionId)
	}
	return questions
}

type ExcludedQuestion struct {
	model.Base
	QuestionId      uuid.UUID `gorm:"type:uuid"`
	FormGeneratedId uuid.UUID `gorm:"type:uuid"`
}

func (f *FormGenerated) ExtractQuestionsFromGeneratedForm() []IQuestion {
	questions := make([]IQuestion, 0)
	for _, generatedSection := range f.Sections {
		for _, generatedBlock := range generatedSection.Blocks {
			if generatedBlock != nil {
				questions = slices.Concat(questions, generatedBlock.Questions)

				if generatedBlock.Variant != nil {
					questions = slices.Concat(questions, generatedBlock.Variant.Questions)
				}
			}
		}
	}
	return questions
}
