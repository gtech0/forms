package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/question"
	"hedgehog-forms/internal/core/repository"
)

type QuestionGeneratedIntegrationMapper struct {
	questionRepository *repository.QuestionRepository

	singleChoiceGeneratedIntegrationMapper   *SingleChoiceGeneratedIntegrationMapper
	textInputGeneratedIntegrationMapper      *TextInputGeneratedIntegrationMapper
	multipleChoiceGeneratedIntegrationMapper *MultipleChoiceGeneratedIntegrationMapper
	matchingGeneratedIntegrationMapper       *MatchingGeneratedIntegrationMapper
}

func NewQuestionGeneratedIntegrationMapper() *QuestionGeneratedIntegrationMapper {
	return &QuestionGeneratedIntegrationMapper{
		questionRepository: repository.NewQuestionRepository(),

		singleChoiceGeneratedIntegrationMapper:   NewSingleChoiceGeneratedIntegrationMapper(),
		textInputGeneratedIntegrationMapper:      NewTextInputGeneratedIntegrationMapper(),
		multipleChoiceGeneratedIntegrationMapper: NewMultipleChoiceGeneratedIntegrationMapper(),
		matchingGeneratedIntegrationMapper:       NewMatchingGeneratedIntegrationMapper(),
	}
}

func (q *QuestionGeneratedIntegrationMapper) ToDto(questionGenerated generated.IQuestion, isAnswerRequired bool) (get.IntegratedIQuestionDto, error) {
	questionEntity, err := q.questionRepository.FindById(questionGenerated.GetId())
	if err != nil {
		return nil, err
	}

	switch questionGenerated.GetType() {
	case question.SINGLE_CHOICE:
		return q.singleChoiceGeneratedIntegrationMapper.toDto(questionGenerated.(*generated.SingleChoice), questionEntity, isAnswerRequired)
	case question.TEXT_INPUT:
		return q.textInputGeneratedIntegrationMapper.toDto(questionGenerated.(*generated.TextInput), questionEntity, isAnswerRequired)
	case question.MULTIPLE_CHOICE:
		return q.multipleChoiceGeneratedIntegrationMapper.toDto(questionGenerated.(*generated.MultipleChoice), questionEntity, isAnswerRequired)
	case question.MATCHING:
		return q.matchingGeneratedIntegrationMapper.toDto(questionGenerated.(*generated.Matching), questionEntity, isAnswerRequired)
	default:
		return nil, errs.New("invalid question type", 400)
	}
}
