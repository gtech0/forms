package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type CommonFieldQuestionDtoMapper struct {
	subjectMapper *SubjectMapper
}

func NewCommonFieldQuestionDtoMapper() *CommonFieldQuestionDtoMapper {
	return &CommonFieldQuestionDtoMapper{
		subjectMapper: NewSubjectMapper(),
	}
}

func (q *CommonFieldQuestionDtoMapper) commonFieldsToDto(source question.Question, target get.IQuestionDto) {
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
		target.SetSubject(q.subjectMapper.toDto(source.Subject))
	}
}
