package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type CommonFieldQuestionMapper struct{}

func NewCommonFieldQuestionMapper() *CommonFieldQuestionMapper {
	return &CommonFieldQuestionMapper{}
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsDto(source create.NewQuestionDto, target question.IQuestion) {
	target.SetId(uuid.New())
	target.SetDescription(source.Description)
	attachments := make([]question.Attachment, 0)
	for _, attachmentId := range source.Attachments {
		attachmentObj := new(question.Attachment)
		attachmentObj.Id = attachmentId
		attachmentObj.QuestionId = target.GetId()
		attachments = append(attachments, *attachmentObj)
	}
	target.SetAttachments(attachments)
	target.SetType(source.Type)
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsObj(source question.Question, target question.IQuestion) {
	target.SetId(uuid.New())
	target.SetDescription(source.Description)
	target.SetAttachments(source.Attachments)
	target.SetType(source.Type)
	target.SetIsQuestionFromBank(source.IsQuestionFromBank)
}
