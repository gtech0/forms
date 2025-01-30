package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type IQuestionDto interface {
	SetId(uuid.UUID)

	SetDescription(string)

	SetOwnerId(nullUUID uuid.NullUUID)

	GetType() question.QuestionType
	SetType(question.QuestionType)

	SetAttachments([]uuid.UUID)

	SetSubject(SubjectDto)
}

type QuestionDto struct {
	Id          uuid.UUID             `json:"id"`
	Description string                `json:"description"`
	OwnerId     uuid.NullUUID         `json:"ownerId"`
	Type        question.QuestionType `json:"type"`
	Attachments []uuid.UUID           `json:"attachments"`
	Subject     SubjectDto            `json:"subject"`
}

func (q *QuestionDto) SetId(id uuid.UUID) {
	q.Id = id
}

func (q *QuestionDto) SetDescription(description string) {
	q.Description = description
}

func (q *QuestionDto) SetOwnerId(ownerId uuid.NullUUID) {
	q.OwnerId = ownerId
}

func (q *QuestionDto) GetType() question.QuestionType {
	return q.Type
}

func (q *QuestionDto) SetType(t question.QuestionType) {
	q.Type = t
}

func (q *QuestionDto) SetAttachments(attachments []uuid.UUID) {
	q.Attachments = attachments
}

func (q *QuestionDto) SetSubject(subject SubjectDto) {
	q.Subject = subject
}
