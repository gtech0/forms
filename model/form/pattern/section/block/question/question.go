package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type IQuestion interface {
	GetId() uuid.UUID
	SetId(uuid.UUID)

	GetDescription() string
	SetDescription(string)

	SetOrder(int)

	GetOwnerId() uuid.NullUUID

	GetType() QuestionType
	SetType(QuestionType)

	GetAttachments() []Attachment
	SetAttachments([]Attachment)

	SetVariantId(nullUUID uuid.NullUUID)

	SetDynamicBlockId(nullUUID uuid.NullUUID)

	GetSubject() model.Subject

	SetIsQuestionFromBank(bool)
}

type Question struct {
	model.Base
	Description        string
	Order              int
	OwnerId            uuid.NullUUID `gorm:"type:uuid"`
	Type               QuestionType
	Attachments        []Attachment
	VariantId          uuid.NullUUID `gorm:"type:uuid"`
	DynamicBlockId     uuid.NullUUID `gorm:"type:uuid"`
	Subject            model.Subject
	SubjectId          uuid.NullUUID `gorm:"type:uuid"`
	IsQuestionFromBank bool
}

func (q *Question) SetId(id uuid.UUID) {
	q.Id = id
}

func (q *Question) GetId() uuid.UUID {
	return q.Id
}

func (q *Question) GetDescription() string {
	return q.Description
}

func (q *Question) SetDescription(description string) {
	q.Description = description
}

func (q *Question) SetOrder(order int) {
	q.Order = order
}

func (q *Question) GetOwnerId() uuid.NullUUID {
	return q.OwnerId
}

func (q *Question) GetType() QuestionType {
	return q.Type
}

func (q *Question) SetType(t QuestionType) {
	q.Type = t
}

func (q *Question) GetAttachments() []Attachment {
	return q.Attachments
}

func (q *Question) SetAttachments(attachments []Attachment) {
	q.Attachments = attachments
}

func (q *Question) SetVariantId(id uuid.NullUUID) {
	q.VariantId = id
}

func (q *Question) SetDynamicBlockId(id uuid.NullUUID) {
	q.DynamicBlockId = id
}

func (q *Question) GetSubject() model.Subject {
	return q.Subject
}

func (q *Question) SetIsQuestionFromBank(isQuestionFromBank bool) {
	q.IsQuestionFromBank = isQuestionFromBank
}
