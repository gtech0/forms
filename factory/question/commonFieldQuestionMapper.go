package question

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/section/block/question"
)

type CommonFieldQuestionMapper struct{}

func NewCommonFieldQuestionMapper() *CommonFieldQuestionMapper {
	return &CommonFieldQuestionMapper{}
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsDto(source create.NewQuestionDto, target question.IQuestion) {
	target.SetDescription(source.Description)
	//target.Attachments = source.Attachments
	target.SetType(source.Type)
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsObj(source question.Question, target question.IQuestion) {
	target.SetDescription(source.Description)
	//target.Attachments = source.Attachments
	target.SetType(source.Type)
	target.SetIsQuestionFromBank(source.IsQuestionFromBank)
}
