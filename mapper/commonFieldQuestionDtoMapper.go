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

func (q *CommonFieldQuestionDtoMapper) commonFieldsToDto(source question.IQuestion, target get.IQuestionDto) {
	target.SetId(source.GetId())
	target.SetDescription(source.GetDescription())
	target.SetOwnerId(source.GetOwnerId())
	target.SetType(source.GetType())
	attachments := make([]uuid.UUID, 0)
	for _, attachment := range source.GetAttachments() {
		attachments = append(attachments, attachment.Id)
	}
	target.SetAttachments(attachments)
	if source.GetSubject().Name != "" {
		target.SetSubject(*q.subjectMapper.ToDto(source.GetSubject()))
	}
}
