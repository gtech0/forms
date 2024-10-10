package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/section/block/question"
)

type IQuestionDto interface {
	GetType() question.QuestionType
}

type QuestionDto struct {
	Id          uuid.UUID             `json:"id"`
	Description string                `json:"description"`
	OwnerId     uuid.UUID             `json:"ownerId"`
	Type        question.QuestionType `json:"type"`
	Attachments []AttachmentDto       `json:"attachments"`
	Subject     SubjectDto            `json:"subject"`
}

func (q *QuestionDto) GetType() question.QuestionType {
	return q.Type
}

type TextInputDto struct {
	QuestionDto
	Points          int      `json:"points"`
	IsCaseSensitive bool     `json:"isCaseSensitive"`
	Answers         []string `json:"answers"`
}

type MatchingDto struct {
	QuestionDto
	Points              map[int]int         `json:"points"`
	TermsAndDefinitions []TermDefinitionDto `json:"termsAndDefinitions"`
}

type TermDefinitionDto struct {
	Term       MatchingOptionDto `json:"term"`
	Definition MatchingOptionDto `json:"definition"`
}

type MatchingOptionDto struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}

type MultipleChoiceDto struct {
	Points  map[int]int         `json:"points"`
	Choices []MultipleOptionDto `json:"choices"`
}

type MultipleOptionDto struct {
	Id       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	IsAnswer bool      `json:"isAnswer"`
}

type SingleChoice struct {
	QuestionDto
	Points  int               `json:"points"`
	Choices []SingleOptionDto `json:"choices"`
}

type SingleOptionDto struct {
	Id       uuid.UUID `json:"id"`
	Text     string    `json:"text"`
	IsAnswer bool      `json:"isAnswer"`
}
