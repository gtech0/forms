package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type BlockMapper struct {
	questionMapper *QuestionMapper
}

func NewBlockMapper() *BlockMapper {
	return &BlockMapper{
		questionMapper: NewQuestionMapper(),
	}
}

func (b *BlockMapper) toDto(iBlock *block.Block) (get.IBlockDto, error) {
	switch iBlock.Type {
	case block.DYNAMIC:
		return b.dynamicToDto(iBlock)
	case block.STATIC:
		return b.staticToDto(iBlock)
	default:
		return nil, errs.New("invalid block type", 400)
	}
}

func (b *BlockMapper) dynamicToDto(dynamicBlock *block.Block) (*get.DynamicBlockDto, error) {
	dynamicBlockDto := new(get.DynamicBlockDto)
	dynamicBlockDto.Id = dynamicBlock.Id
	dynamicBlockDto.Title = dynamicBlock.Title
	dynamicBlockDto.Description = dynamicBlock.Description
	dynamicBlockDto.Type = dynamicBlock.Type
	questions, err := b.questionsToDto(dynamicBlock.DynamicBlock.Questions)
	if err != nil {
		return nil, err
	}

	dynamicBlockDto.Questions = questions
	return dynamicBlockDto, nil
}

func (b *BlockMapper) staticToDto(staticBlock *block.Block) (*get.StaticBlockDto, error) {
	staticBlockDto := new(get.StaticBlockDto)
	staticBlockDto.Id = staticBlock.Id
	staticBlockDto.Title = staticBlock.Title
	staticBlockDto.Description = staticBlock.Description
	staticBlockDto.Type = staticBlock.Type
	variants, err := b.variantsToDto(staticBlock.StaticBlock.Variants)
	if err != nil {
		return nil, err
	}
	staticBlockDto.Variants = variants
	return staticBlockDto, nil
}

func (b *BlockMapper) variantsToDto(variants []block.Variant) ([]get.VariantDto, error) {
	mappedVariants := make([]get.VariantDto, 0)
	for _, variant := range variants {
		mappedVariant, err := b.variantToDto(variant)
		if err != nil {
			return nil, err
		}
		mappedVariants = append(mappedVariants, mappedVariant)
	}
	return mappedVariants, nil
}

func (b *BlockMapper) variantToDto(variantObj block.Variant) (get.VariantDto, error) {
	var variantDto get.VariantDto
	variantDto.Id = variantObj.Id
	variantDto.Title = variantObj.Title
	variantDto.Description = variantObj.Description
	questions, err := b.questionsToDto(variantObj.Questions)
	if err != nil {
		return get.VariantDto{}, err
	}
	variantDto.Questions = questions
	return variantDto, nil
}

func (b *BlockMapper) questionsToDto(questions []question.IQuestion) ([]get.IQuestionDto, error) {
	mappedQuestions := make([]get.IQuestionDto, 0)
	for _, questionObj := range questions {
		mappedQuestion, err := b.questionMapper.toDto(questionObj)
		if err != nil {
			return nil, err
		}
		mappedQuestions = append(mappedQuestions, mappedQuestion)
	}
	return mappedQuestions, nil
}
