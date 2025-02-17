package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/generated"
	question "hedgehog-forms/internal/core/model/form/pattern/question"
)

type QuestionGeneratedFactory struct{}

func NewQuestionGeneratedFactory() *QuestionGeneratedFactory {
	return &QuestionGeneratedFactory{}
}

func (q *QuestionGeneratedFactory) buildQuestions(questions []*question.Question) ([]generated.IQuestion, error) {
	generatedQuestions := make([]generated.IQuestion, 0)
	for _, iQuestion := range questions {
		generatedQuestion, err := q.buildQuestion(iQuestion)
		if err != nil {
			return nil, err
		}

		generatedQuestions = append(generatedQuestions, generatedQuestion)
	}
	return generatedQuestions, nil
}

func (q *QuestionGeneratedFactory) buildQuestion(questionEntity *question.Question) (generated.IQuestion, error) {
	switch questionEntity.Type {
	case question.TEXT_INPUT:
		return q.buildTextInput(questionEntity), nil
	case question.SINGLE_CHOICE:
		return q.buildSingleChoice(questionEntity), nil
	case question.MULTIPLE_CHOICE:
		return q.buildMultipleChoice(questionEntity), nil
	case question.MATCHING:
		return q.buildMatching(questionEntity), nil
	default:
		return nil, errs.New("unsupported question type", 400)
	}
}

func (q *QuestionGeneratedFactory) buildTextInput(textInput *question.Question) *generated.TextInput {
	generatedTextInput := new(generated.TextInput)
	q.buildCommonFields(textInput, &generatedTextInput.Question)
	return generatedTextInput
}

func (q *QuestionGeneratedFactory) buildSingleChoice(questionEntity *question.Question) *generated.SingleChoice {
	generatedSingleChoice := new(generated.SingleChoice)
	q.buildCommonFields(questionEntity, &generatedSingleChoice.Question)
	generatedSingleChoice.Points = questionEntity.SingleChoice.Points
	options := make([]generated.SingleChoiceOption, 0)
	for _, option := range questionEntity.SingleChoice.Options {
		var generatedOption generated.SingleChoiceOption
		generatedOption.Id = option.Id
		generatedOption.Text = option.Text
		options = append(options, generatedOption)
	}
	generatedSingleChoice.Options = options
	return generatedSingleChoice
}

func (q *QuestionGeneratedFactory) buildMultipleChoice(questionEntity *question.Question) *generated.MultipleChoice {
	generatedMultipleChoice := new(generated.MultipleChoice)
	q.buildCommonFields(questionEntity, &generatedMultipleChoice.Question)
	options := make([]generated.MultipleChoiceOption, 0)
	for _, option := range questionEntity.MultipleChoice.Options {
		var generatedOption generated.MultipleChoiceOption
		generatedOption.Id = option.Id
		generatedOption.Text = option.Text
		options = append(options, generatedOption)
	}
	generatedMultipleChoice.Options = options
	return generatedMultipleChoice
}

func (q *QuestionGeneratedFactory) buildMatching(questionEntity *question.Question) *generated.Matching {
	generatedMatching := new(generated.Matching)
	q.buildCommonFields(questionEntity, &generatedMatching.Question)
	terms := make([]generated.Term, 0)
	for _, term := range questionEntity.Matching.Terms {
		var generatedTerm generated.Term
		generatedTerm.Id = term.Id
		generatedTerm.Text = term.Text
		terms = append(terms, generatedTerm)
	}
	generatedMatching.Terms = terms

	definitions := make([]generated.Definition, 0)
	for _, definition := range questionEntity.Matching.Definitions {
		var generatedDefinition generated.Definition
		generatedDefinition.Id = definition.Id
		generatedDefinition.Text = definition.Text
		definitions = append(definitions, generatedDefinition)
	}
	generatedMatching.Definitions = definitions

	return generatedMatching
}

func (q *QuestionGeneratedFactory) buildCommonFields(source *question.Question, target *generated.Question) {
	target.Id = source.Id
	target.Description = source.Description
	target.Type = source.Type
	target.Attachments = q.buildAttachments(source.Attachments)
}

func (q *QuestionGeneratedFactory) buildAttachments(attachments []question.Attachment) []uuid.UUID {
	attachmentIds := make([]uuid.UUID, 0)
	for _, attachment := range attachments {
		attachmentIds = append(attachmentIds, attachment.Id)
	}
	return attachmentIds
}
