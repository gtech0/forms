package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model"
	"slices"
)

type FormGenerated struct {
	model.Base
	Status          FormStatus
	Sections        []Section `gorm:"type:jsonb;serializer:json"`
	Points          int
	Mark            int
	FormPublishedId uuid.UUID `gorm:"type:uuid"`
	//SubmissionId    uuid.UUID `gorm:"type:uuid"`
}

func (f *FormGenerated) ExtractQuestions() []IQuestion {
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
