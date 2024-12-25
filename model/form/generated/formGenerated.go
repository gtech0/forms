package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"slices"
	"time"
)

type FormGenerated struct {
	model.Base
	Status          FormStatus
	UserId          uuid.UUID `gorm:"type:uuid"`
	IsCompleted     bool
	Sections        []Section `gorm:"type:jsonb;serializer:json"`
	Points          int
	Mark            string
	StartTime       time.Time
	SubmitTime      time.Time
	FormPublishedID uuid.UUID `gorm:"type:uuid"`
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
