package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	generated2 "hedgehog-forms/internal/core/model/form/generated"
	question2 "hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type QuestionGeneratedFactory struct{}

func NewQuestionGeneratedFactory() *QuestionGeneratedFactory {
	return &QuestionGeneratedFactory{}
}

func (q *QuestionGeneratedFactory) buildQuestions(questions []*question2.Question) ([]generated2.IQuestion, error) {
	generatedQuestions := make([]generated2.IQuestion, 0)
	for _, iQuestion := range questions {
		generatedQuestion, err := q.buildQuestion(iQuestion)
		if err != nil {
			return nil, err
		}

		generatedQuestions = append(generatedQuestions, generatedQuestion)
	}
	return generatedQuestions, nil
}

func (q *QuestionGeneratedFactory) buildQuestion(questionEntity *question2.Question) (generated2.IQuestion, error) {
	switch questionEntity.Type {
	case question2.TEXT_INPUT:
		return q.buildTextInput(questionEntity), nil
	case question2.SINGLE_CHOICE:
		return q.buildSingleChoice(questionEntity), nil
	case question2.MULTIPLE_CHOICE:
		return q.buildMultipleChoice(questionEntity), nil
	case question2.MATCHING:
		return q.buildMatching(questionEntity), nil
	default:
		return nil, errs.New("unsupported question type", 400)
	}
}

func (q *QuestionGeneratedFactory) buildTextInput(textInput *question2.Question) *generated2.TextInput {
	generatedTextInput := new(generated2.TextInput)
	q.buildCommonFields(textInput, &generatedTextInput.Question)
	return generatedTextInput
}

func (q *QuestionGeneratedFactory) buildSingleChoice(questionEntity *question2.Question) *generated2.SingleChoice {
	generatedSingleChoice := new(generated2.SingleChoice)
	q.buildCommonFields(questionEntity, &generatedSingleChoice.Question)
	generatedSingleChoice.Points = questionEntity.SingleChoice.Points
	options := make([]generated2.SingleChoiceOption, 0)
	for _, option := range questionEntity.SingleChoice.Options {
		var generatedOption generated2.SingleChoiceOption
		generatedOption.Id = option.Id
		generatedOption.Text = option.Text
		options = append(options, generatedOption)
	}
	generatedSingleChoice.Options = options
	return generatedSingleChoice
}

func (q *QuestionGeneratedFactory) buildMultipleChoice(questionEntity *question2.Question) *generated2.MultipleChoice {
	generatedMultipleChoice := new(generated2.MultipleChoice)
	q.buildCommonFields(questionEntity, &generatedMultipleChoice.Question)
	options := make([]generated2.MultipleChoiceOption, 0)
	for _, option := range questionEntity.MultipleChoice.Options {
		var generatedOption generated2.MultipleChoiceOption
		generatedOption.Id = option.Id
		generatedOption.Text = option.Text
		options = append(options, generatedOption)
	}
	generatedMultipleChoice.Options = options
	return generatedMultipleChoice
}

func (q *QuestionGeneratedFactory) buildMatching(questionEntity *question2.Question) *generated2.Matching {
	generatedMatching := new(generated2.Matching)
	q.buildCommonFields(questionEntity, &generatedMatching.Question)
	terms := make([]generated2.Term, 0)
	for _, term := range questionEntity.Matching.Terms {
		var generatedTerm generated2.Term
		generatedTerm.Id = term.Id
		generatedTerm.Text = term.Text
		terms = append(terms, generatedTerm)
	}
	generatedMatching.Terms = terms

	definitions := make([]generated2.Definition, 0)
	for _, definition := range questionEntity.Matching.Definitions {
		var generatedDefinition generated2.Definition
		generatedDefinition.Id = definition.Id
		generatedDefinition.Text = definition.Text
		definitions = append(definitions, generatedDefinition)
	}
	generatedMatching.Definitions = definitions

	return generatedMatching
}

func (q *QuestionGeneratedFactory) buildCommonFields(source *question2.Question, target *generated2.Question) {
	target.Id = source.Id
	target.Description = source.Description
	target.Type = source.Type
	target.Attachments = q.buildAttachments(source.Attachments)
}

func (q *QuestionGeneratedFactory) buildAttachments(attachments []question2.Attachment) []uuid.UUID {
	attachmentIds := make([]uuid.UUID, 0)
	for _, attachment := range attachments {
		attachmentIds = append(attachmentIds, attachment.Id)
	}
	return attachmentIds
}
