package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/errs"
	"hedgehog-forms/model"
	"slices"
	"time"
)

type FormGenerated struct {
	model.Base
	Status            FormStatus
	FormPublishedID   uuid.UUID `gorm:"type:uuid"`
	UserId            uuid.UUID `gorm:"type:uuid"`
	Attempts          []*Attempt
	FinalPoints       int
	FinalMark         string
	SubmitTime        time.Time
	ExcludedQuestions []uuid.UUID `gorm:"type:uuid[]"`
}

func (f *FormGenerated) ExtractQuestionsFromGeneratedForm() []IQuestion {
	questions := make([]IQuestion, 0)
	currentAttempt := new(Attempt)
	for _, attempt := range f.Attempts {
		if attempt.IsComplete == false {
			currentAttempt = attempt
			break
		}
	}

	for _, generatedSection := range currentAttempt.Sections {
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

func (f *FormGenerated) ExtractCurrentAttempt() (*Attempt, error) {
	for _, attempt := range f.Attempts {
		if !attempt.IsComplete {
			return attempt, nil
		}
	}

	return nil, errs.New("Attempt limit reached", 400)
}
