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
	FormPublishedID   uuid.UUID `gorm:"type:uuid"`
	UserId            uuid.UUID `gorm:"type:uuid"`
	Sections          []Section `gorm:"type:jsonb;serializer:json"`
	Points            int
	Mark              string
	SubmitTime        time.Time
	CurrentAttempts   int
	IsGenerated       bool
	ExcludedQuestions []uuid.UUID `gorm:"type:uuid[]"`
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
