package migration

import (
	"hedgehog-forms/internal/core/model"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern"
	"hedgehog-forms/internal/core/model/form/pattern/block"
	"hedgehog-forms/internal/core/model/form/pattern/question"
	"hedgehog-forms/internal/core/model/form/pattern/section"
	"hedgehog-forms/internal/core/model/form/published"
	"hedgehog-forms/pkg/database"
	"log"
)

func Sync() {
	err := database.DB.AutoMigrate(
		&published.FormPublished{},
		&published.MarkConfiguration{},
		&published.ExcludedQuestion{},

		&published.Solution{},
		&published.SolutionComment{},
		&generated.Submission{},

		&generated.FormGenerated{},

		&pattern.FormPattern{},

		&section.Section{},

		&block.Block{},
		&block.StaticBlock{},
		&block.DynamicBlock{},
		&block.Variant{},

		&question.Attachment{},
		&model.File{},

		&question.Question{},

		&question.MultipleChoice{},
		&question.MultipleChoicePoints{},
		&question.MultipleChoiceOption{},

		&question.SingleChoice{},
		&question.SingleChoiceOption{},

		&question.TextInput{},
		&question.TextInputAnswer{},

		&question.Matching{},
		&question.MatchingPoints{},
		&question.MatchingTerm{},
		&question.MatchingDefinition{},
	)

	if err != nil {
		log.Fatal(err)
	}
}
