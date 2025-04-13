package get

import "github.com/google/uuid"

type IntegratedMultipleChoiceDto struct {
	IntegratedQuestionDto
	Options        []IntegratedMultipleOptionDto `json:"options"`
	CorrectOptions []IntegratedMultipleOptionDto `json:"correctOptions,omitempty"`
	EnteredAnswers []uuid.UUID                   `json:"enteredAnswers"`
}

type IntegratedMultipleOptionDto struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
