package dto

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/section/block/question"
)

type CreateQuestionDto struct {
	Type question.QuestionType `json:"type"`
}

type CreateQuestionOnExistingDto struct {
	CreateQuestionDto
	QuestionId uuid.UUID `json:"questionId"`
}

type NewQuestionDto struct {
	CreateQuestionDto
	Description string      `json:"description"`
	Attachments []uuid.UUID `json:"attachments"`
}

type CreateMultipleChoiceQuestionDto struct {
	NewQuestionDto
	Options        []string    `json:"options"`
	CorrectOptions []int       `json:"correctOptions"`
	Points         map[int]int `json:"points"`
}

type CreateTextQuestionDto struct {
	NewQuestionDto
	IsCaseSensitive bool     `json:"isCaseSensitive"`
	Answers         []string `json:"answers"`
	Points          int      `json:"points"`
}

type CreateMatchingQuestionDto struct {
	NewQuestionDto
	TermsAndDefinitions map[string]string `json:"termsAndDefinitions"`
	Points              map[int]int       `json:"points"`
}

type CreateSingleChoiceQuestionDto struct {
	NewQuestionDto
	Options       []string `json:"options"`
	CorrectOption int      `json:"correctOption"`
	Points        int      `json:"points"`
}
