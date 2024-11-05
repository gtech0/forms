package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type CommonFieldQuestionMapper struct{}

func NewCommonFieldQuestionMapper() *CommonFieldQuestionMapper {
	return &CommonFieldQuestionMapper{}
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsDto(source create.NewQuestionDto, target *question.Question) error {
	if target == nil {
		return errs.New("nil question pointer", 500)
	}

	target.Id = uuid.New()
	target.Description = source.Description
	attachments := make([]question.Attachment, 0)
	for _, attachmentId := range source.Attachments {
		attachmentObj := new(question.Attachment)
		attachmentObj.Id = attachmentId
		attachmentObj.QuestionId = target.Id
		attachments = append(attachments, *attachmentObj)
	}
	target.Attachments = attachments
	target.QuestionType = source.Type
	return nil
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsObj(source question.Question, target *question.Question) error {
	if target == nil {
		return errs.New("nil question pointer", 500)
	}

	target.Id = uuid.New()
	target.Description = source.Description
	target.Attachments = source.Attachments
	target.QuestionType = source.QuestionType
	target.IsQuestionFromBank = source.IsQuestionFromBank
	return nil
}
