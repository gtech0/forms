package mapper

import (
	get2 "hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	block2 "hedgehog-forms/internal/core/model/form/pattern/section/block"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type BlockMapper struct {
	questionMapper *QuestionMapper
}

func NewBlockMapper() *BlockMapper {
	return &BlockMapper{
		questionMapper: NewQuestionMapper(),
	}
}

func (b *BlockMapper) toDto(iBlock *block2.Block) (get2.IBlockDto, error) {
	switch iBlock.Type {
	case block2.DYNAMIC:
		return b.dynamicToDto(iBlock)
	case block2.STATIC:
		return b.staticToDto(iBlock)
	default:
		return nil, errs.New("invalid block type", 400)
	}
}

func (b *BlockMapper) dynamicToDto(dynamicBlock *block2.Block) (*get2.DynamicBlockDto, error) {
	dynamicBlockDto := new(get2.DynamicBlockDto)
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

func (b *BlockMapper) staticToDto(staticBlock *block2.Block) (*get2.StaticBlockDto, error) {
	staticBlockDto := new(get2.StaticBlockDto)
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

func (b *BlockMapper) variantsToDto(variants []block2.Variant) ([]get2.VariantDto, error) {
	mappedVariants := make([]get2.VariantDto, 0)
	for _, variant := range variants {
		mappedVariant, err := b.variantToDto(variant)
		if err != nil {
			return nil, err
		}
		mappedVariants = append(mappedVariants, mappedVariant)
	}
	return mappedVariants, nil
}

func (b *BlockMapper) variantToDto(variantEntity block2.Variant) (get2.VariantDto, error) {
	var variantDto get2.VariantDto
	variantDto.Id = variantEntity.Id
	variantDto.Title = variantEntity.Title
	variantDto.Description = variantEntity.Description
	questions, err := b.questionsToDto(variantEntity.Questions)
	if err != nil {
		return get2.VariantDto{}, err
	}
	variantDto.Questions = questions
	return variantDto, nil
}

func (b *BlockMapper) questionsToDto(questions []*question.Question) ([]get2.IQuestionDto, error) {
	mappedQuestions := make([]get2.IQuestionDto, 0)
	for _, questionEntity := range questions {
		mappedQuestion, err := b.questionMapper.ToDto(questionEntity)
		if err != nil {
			return nil, err
		}
		mappedQuestions = append(mappedQuestions, mappedQuestion)
	}
	return mappedQuestions, nil
}
