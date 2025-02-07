package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type CommonFieldQuestionMapper struct{}

func NewCommonFieldQuestionMapper() *CommonFieldQuestionMapper {
	return &CommonFieldQuestionMapper{}
}

func (c *CommonFieldQuestionMapper) MapCommonDtoFields(source create.NewQuestionDto, target *question.Question) error {
	if target == nil {
		return errs.New("nil question pointer", 500)
	}

	target.Id = uuid.New()
	target.Description = source.Description
	attachments := make([]question.Attachment, 0)
	for _, fileId := range source.Attachments {
		attachmentEntity := new(question.Attachment)
		attachmentEntity.FileId = fileId
		attachmentEntity.QuestionId = target.Id
		attachments = append(attachments, *attachmentEntity)
	}
	target.Attachments = attachments
	target.Type = source.Type
	return nil
}

func (c *CommonFieldQuestionMapper) MapCommonEntityFields(source question.Question, target *question.Question) error {
	if target == nil {
		return errs.New("nil question pointer", 500)
	}

	target.Id = uuid.New()
	target.Description = source.Description
	target.Attachments = source.Attachments
	target.Type = source.Type
	target.IsQuestionFromBank = source.IsQuestionFromBank
	return nil
}
