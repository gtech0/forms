package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type QuestionGeneratedFactory struct{}

func NewQuestionGeneratedFactory() *QuestionGeneratedFactory {
	return &QuestionGeneratedFactory{}
}

func (q *QuestionGeneratedFactory) buildQuestions(questions []question.IQuestion) ([]generated.IQuestion, error) {
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

func (q *QuestionGeneratedFactory) buildQuestion(iQuestion question.IQuestion) (generated.IQuestion, error) {
	switch assertedQuestion := iQuestion.(type) {
	case *question.TextInput:
		return q.buildTextInput(assertedQuestion), nil
	case *question.SingleChoice:
		return q.buildSingleChoice(assertedQuestion), nil
	case *question.MultipleChoice:
		return q.buildMultipleChoice(assertedQuestion), nil
	case *question.Matching:
		return q.buildMatching(assertedQuestion), nil
	default:
		return nil, errs.New("unsupported question type", 400)
	}
}

func (q *QuestionGeneratedFactory) buildTextInput(textInput *question.TextInput) *generated.TextInput {
	generatedTextInput := new(generated.TextInput)
	q.buildCommonFields(textInput.Question, &generatedTextInput.Question)
	return generatedTextInput
}

func (q *QuestionGeneratedFactory) buildSingleChoice(singleChoice *question.SingleChoice) *generated.SingleChoice {
	generatedSingleChoice := new(generated.SingleChoice)
	q.buildCommonFields(singleChoice.Question, &generatedSingleChoice.Question)
	generatedSingleChoice.Points = singleChoice.Points
	options := make([]generated.SingleChoiceOption, 0)
	for _, option := range singleChoice.Options {
		var generatedOption generated.SingleChoiceOption
		generatedOption.Id = option.Id
		generatedOption.Text = option.Text
		options = append(options, generatedOption)
	}
	generatedSingleChoice.Options = options
	return generatedSingleChoice
}

func (q *QuestionGeneratedFactory) buildMultipleChoice(multipleChoice *question.MultipleChoice) *generated.MultipleChoice {
	generatedMultipleChoice := new(generated.MultipleChoice)
	q.buildCommonFields(multipleChoice.Question, &generatedMultipleChoice.Question)
	options := make([]generated.MultipleChoiceOption, 0)
	for _, option := range multipleChoice.Options {
		var generatedOption generated.MultipleChoiceOption
		generatedOption.Id = option.Id
		generatedOption.Text = option.Text
		options = append(options, generatedOption)
	}
	generatedMultipleChoice.Options = options
	return generatedMultipleChoice
}

func (q *QuestionGeneratedFactory) buildMatching(matching *question.Matching) *generated.Matching {
	generatedMatching := new(generated.Matching)
	q.buildCommonFields(matching.Question, &generatedMatching.Question)
	terms := make([]generated.Term, 0)
	for _, term := range matching.Terms {
		var generatedTerm generated.Term
		generatedTerm.Id = term.Id
		generatedTerm.Text = term.Text
		terms = append(terms, generatedTerm)
	}
	generatedMatching.Terms = terms

	definitions := make([]generated.Definition, 0)
	for _, definition := range matching.Definitions {
		var generatedDefinition generated.Definition
		generatedDefinition.Id = definition.Id
		generatedDefinition.Text = definition.Text
		definitions = append(definitions, generatedDefinition)
	}
	generatedMatching.Definitions = definitions

	return generatedMatching
}

func (q *QuestionGeneratedFactory) buildCommonFields(source question.Question, target *generated.Question) {
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
