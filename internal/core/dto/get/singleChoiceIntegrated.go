package get

import "github.com/google/uuid"

type IntegratedSingleChoiceDto struct {
	IntegratedQuestionDto
	Options       []IntegratedSingleOptionDto `json:"choices"`
	Answer        IntegratedSingleOptionDto   `json:"answer"`
	EnteredAnswer uuid.UUID                   `json:"enteredAnswer"`
}

type IntegratedSingleOptionDto struct {
	Id   uuid.UUID `json:"id"`
	Text string    `json:"text"`
}
