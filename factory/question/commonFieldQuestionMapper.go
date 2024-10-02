package question

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block/question"
)

type CommonFieldQuestionMapper struct{}

func NewCommonFieldQuestionMapper() *CommonFieldQuestionMapper {
	return &CommonFieldQuestionMapper{}
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsDto(source dto.NewQuestionDto, target question.Question) {
	target.Description = source.Description
	//target.Attachments = source.Attachments
	target.Type = source.Type
}

func (c *CommonFieldQuestionMapper) MapCommonFieldsObj(source question.Question, target question.Question) {
	target.Description = source.Description
	target.Attachments = source.Attachments
	target.Type = source.Type
	target.IsQuestionFromBank = source.IsQuestionFromBank
}
