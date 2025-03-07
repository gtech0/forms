package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type CommonFieldQuestionDtoMapper struct {
	subjectMapper *SubjectMapper
}

func NewCommonFieldQuestionDtoMapper() *CommonFieldQuestionDtoMapper {
	return &CommonFieldQuestionDtoMapper{
		subjectMapper: NewSubjectMapper(),
	}
}

func (q *CommonFieldQuestionDtoMapper) CommonFieldsToDto(source *question.Question, target get.IQuestionDto) {
	target.SetId(source.Id)
	target.SetDescription(source.Description)
	target.SetOwnerId(source.OwnerId)
	target.SetType(source.Type)
	attachments := make([]uuid.UUID, 0)
	for _, attachment := range source.Attachments {
		attachments = append(attachments, attachment.Id)
	}
	target.SetAttachments(attachments)
	if source.Subject.Name != "" {
		target.SetSubject(*q.subjectMapper.ToDto(source.Subject))
	}
}
