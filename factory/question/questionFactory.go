package question

import (
	"errors"
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/model/form/section/block/question"
)

type QuestionFactory struct {
}

func (q *QuestionFactory) buildQuestion(questionDto dto.CreateQuestion) (any, error) {
	switch questionDto.GetDtoType() {
	case question.EXISTING:
		return NewExistingQuestionFactory().BuildFromDto(questionDto.(dto.CreateQuestionOnExistingDto))
	case question.MATCHING:
		return NewMatchingFactory().BuildFromDto(questionDto.(dto.CreateMatchingQuestionDto))
	case question.TEXT_INPUT:
		return NewTextInputFactory().BuildFromDto(questionDto.(dto.CreateTextQuestionDto))
	case question.SINGLE_CHOICE:
		return NewSingleChoiceFactory().BuildFromDto(questionDto.(dto.CreateSingleChoiceQuestionDto))
	case question.MULTIPLE_CHOICE:
		return NewMultipleChoiceFactory().BuildFromDto(questionDto.(dto.CreateMultipleChoiceQuestionDto))
	default:
		return nil, errors.New("unknown question type")
	}
}

func (q *QuestionFactory) buildQuestionDtoForDynamicBlock(questionDtos []dto.CreateQuestionDto, dynamicBlock block.DynamicBlock) ([]question.Question, error) {
	questionObjs := make([]question.Question, len(questionDtos))
	for _, questionDto := range questionDtos {
		questionObj, err := q.buildQuestion(questionDto)
		if err != nil {
			return nil, err
		}

		questionObj.(*question.Question).DynamicBlockId = dynamicBlock.Id
		questionObjs = append(questionObjs, questionObj.(question.Question))
	}
	return questionObjs, nil
}

//    public List<QuestionEntity> buildQuestionsForDynamicBlockFromEntities(
//            List<QuestionEntity> questionEntities,
//            DynamicBlockEntity blockEntity
//    ) {
//        var newQuestionEntities = new ArrayList<QuestionEntity>(questionEntities.size());
//
//        for (QuestionEntity questionEntity : questionEntities) {
//            var newQuestionEntity = buildQuestionFromEntity(questionEntity);
//            newQuestionEntity.setDynamicBlock(blockEntity);
//
//            newQuestionEntities.add(newQuestionEntity);
//        }
//
//        return newQuestionEntities;
//    }
//
//    public List<QuestionEntity> buildQuestionsForVariantFromDtos(List<CreateQuestionDto> questionDtos,
//                                                                 VariantEntity variantEntity
//    ) {
//        var questionEntities = new ArrayList<QuestionEntity>(questionDtos.size());
//        for (int i = 0; i < questionDtos.size(); i++) {
//            var questionDto = questionDtos.get(i);
//
//            var questionEntity = buildQuestionFromDto(questionDto);
//            questionEntity.setVariant(variantEntity);
//            questionEntity.setOrder(i);
//
//            questionEntities.add(questionEntity);
//        }
//
//        return questionEntities;
//    }
//
//    public List<QuestionEntity> buildQuestionsForVariantFromEntities(
//            List<QuestionEntity> questionEntities,
//            VariantEntity variantEntity
//    ) {
//        var newQuestionEntities = new ArrayList<QuestionEntity>(questionEntities.size());
//
//        for (int i = 0; i < questionEntities.size(); i++) {
//            var questionEntity = questionEntities.get(i);
//            var newQuestionEntity = buildQuestionFromEntity(questionEntity);
//            newQuestionEntity.setVariant(variantEntity);
//            newQuestionEntity.setOrder(i);
//
//            newQuestionEntities.add(newQuestionEntity);
//        }
//
//        return newQuestionEntities;
//    }
//
//    private QuestionEntity buildQuestionFromEntity(QuestionEntity questionEntity) {
//        var questionDto = new CreateQuestionBasedOnExistingDto();
//        questionDto.setQuestionId(questionEntity.getId());
//
//        return buildQuestionFromDto(questionDto);
//    }
