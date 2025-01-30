package pattern

import (
	"github.com/google/uuid"
	model2 "hedgehog-forms/internal/core/model"
	"hedgehog-forms/internal/core/model/form/pattern/section"
	"hedgehog-forms/internal/core/model/form/pattern/section/block"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
	"slices"
)

type FormPattern struct {
	model2.Base
	Title       string
	Description string
	OwnerId     uuid.NullUUID `gorm:"type:uuid"`
	Subject     model2.Subject
	SubjectId   uuid.UUID `gorm:"type:uuid"`
	Sections    []section.Section
}

func (f *FormPattern) ExtractQuestionEntities() []*question.Question {
	questions := make([]*question.Question, 0)
	for _, patternSection := range f.Sections {
		for _, sectionBlock := range patternSection.Blocks {
			switch sectionBlock.Type {
			case block.DYNAMIC:
				questions = slices.Concat(questions, sectionBlock.DynamicBlock.Questions)
			case block.STATIC:
				variants := sectionBlock.StaticBlock.Variants
				for _, variant := range variants {
					questions = slices.Concat(questions, variant.Questions)
				}
			}
		}
	}
	return questions
}
