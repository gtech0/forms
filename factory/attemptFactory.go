package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section"
	"time"
)

type AttemptFactory struct {
	sectionGeneratedFactory *SectionGeneratedFactory
}

func NewAttemptFactory() *AttemptFactory {
	return &AttemptFactory{
		sectionGeneratedFactory: NewSectionGeneratedFactory(),
	}
}

func (f *AttemptFactory) BuildAttempt(sections []section.Section, excludedQuestions []uuid.UUID) (*generated.Attempt, error) {
	attempt := new(generated.Attempt)
	generatedSections, err := f.sectionGeneratedFactory.BuildSections(sections, excludedQuestions)
	if err != nil {
		return nil, err
	}

	attempt.Sections = generatedSections
	attempt.StartTime = time.Now()
	return attempt, nil
}
