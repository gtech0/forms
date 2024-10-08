package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/section/block"
)

type BlockMapper struct{}

func NewBlockMapper() *BlockMapper {
	return &BlockMapper{}
}

func (b *BlockMapper) toDto(blockObj block.IBlock) get.IBlockDto {
	var blockDto get.IBlockDto
	//newSlice := []interface{}{v.MultipleChoice, v.SingleChoice, v.Matching, v.TextInput}
	//for _, slice := range newSlice {
	//	v.Questions = append(v.Questions, slice.(question.IQuestion))
	//}
	return blockDto
}
