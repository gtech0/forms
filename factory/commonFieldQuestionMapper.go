package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type CommonFieldQuestionMapper struct{}

func NewCommonFieldQuestionMapper() *CommonFieldQuestionMapper {
	return &CommonFieldQuestionMapper{}
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsDto(source create.NewQuestionDto, target *question.Question) {
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
	target.Type = source.Type
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsObj(source question.Question, target *question.Question) {
	target.Id = uuid.New()
	target.Description = source.Description
	target.Attachments = source.Attachments
	target.Type = source.Type
	target.IsQuestionFromBank = source.IsQuestionFromBank
}
