package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/model/form/section/block/question"
)

type IQuestionDto interface {
}

type QuestionDto struct {
	Id          uuid.UUID             `json:"id"`
	Description string                `json:"description"`
	Type        question.QuestionType `json:"type"`
	//Attachments
}
