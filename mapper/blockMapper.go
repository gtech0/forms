package mapper

import (
	"errors"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/section/block"
)

type BlockMapper struct {
	questionMapper *QuestionMapper
}

func NewBlockMapper() *BlockMapper {
	return &BlockMapper{
		questionMapper: NewQuestionMapper(),
	}
}

func (b *BlockMapper) toDto(blockObj block.IBlock) (get.IBlockDto, error) {
	switch assertedBlock := blockObj.(type) {
	case *block.DynamicBlock:
		return b.dynamicToDto(assertedBlock), nil
	case *block.StaticBlock:
		return b.staticToDto(assertedBlock), nil
	default:
		return nil, errors.New("block type not found")
	}
}

func (b *BlockMapper) dynamicToDto(dynamicBlock *block.DynamicBlock) *get.DynamicBlockDto {
	dynamicBlockDto := new(get.DynamicBlockDto)
	dynamicBlockDto.Id = dynamicBlock.Id
	dynamicBlockDto.Title = dynamicBlock.Title
	dynamicBlockDto.Description = dynamicBlock.Description
	dynamicBlockDto.Type = dynamicBlock.Type
	//dynamicBlockDto.Questions = b.questionMapper.toDto()
	return dynamicBlockDto
}

func (b *BlockMapper) staticToDto(staticBlock *block.StaticBlock) *get.StaticBlockDto {
	staticBlockDto := new(get.StaticBlockDto)
	staticBlockDto.Id = staticBlock.Id
	staticBlockDto.Title = staticBlock.Title
	staticBlockDto.Description = staticBlock.Description
	staticBlockDto.Type = staticBlock.Type
	//staticBlockDto.Variants = b.questionMapper.toDto()
	return staticBlockDto
}
