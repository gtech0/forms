package create

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/section/block/question"
)

type QuestionDto struct {
	Type question.QuestionType `json:"type"`
}

type QuestionOnExistingDto struct {
	QuestionDto
	QuestionId uuid.UUID `json:"questionId"`
}

type NewQuestionDto struct {
	QuestionDto
	Description string      `json:"description"`
	Attachments []uuid.UUID `json:"attachments"`
}

type MultipleChoiceQuestionDto struct {
	NewQuestionDto
	Options        []string    `json:"options"`
	CorrectOptions []int       `json:"correctOptions"`
	Points         map[int]int `json:"points"`
}

type TextQuestionDto struct {
	NewQuestionDto
	IsCaseSensitive bool     `json:"isCaseSensitive"`
	Answers         []string `json:"answers"`
	Points          int      `json:"points"`
}

type MatchingQuestionDto struct {
	NewQuestionDto
	TermsAndDefinitions map[string]string `json:"termsAndDefinitions"`
	Points              map[int]int       `json:"points"`
}

type SingleChoiceQuestionDto struct {
	NewQuestionDto
	Options       []string `json:"options"`
	CorrectOption int      `json:"correctOption"`
	Points        int      `json:"points"`
}
