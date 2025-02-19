package processor

import (
	"hedgehog-forms/internal/core/mapper"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/question"
	"slices"
	"strings"
)

type TextInputProcessor struct {
	textInputMapper *mapper.TextInputMapper
}

func NewTextInputProcessor() *TextInputProcessor {
	return &TextInputProcessor{
		textInputMapper: mapper.NewTextInputMapper(),
	}
}

func (t *TextInputProcessor) saveAnswer(textInput *generated.TextInput, answer string) error {
	textInput.EnteredAnswer = answer
	return nil
}

func (t *TextInputProcessor) saveAnswerAndCalculatePoints(
	textInput *generated.TextInput,
	textInputEntity *question.TextInput,
) (int, error) {
	answers := t.textInputMapper.AnswersToString(textInputEntity.Answers)
	isEnteredAnswerCorrect := func() bool {
		if textInputEntity.IsCaseSensitive {
			return slices.Contains(answers, textInput.EnteredAnswer)
		}

		for _, answer := range answers {
			return strings.EqualFold(answer, textInput.EnteredAnswer)
		}
		return false
	}

	if isEnteredAnswerCorrect() {
		textInput.Points = textInputEntity.Points
	}

	return textInput.Points, nil
}
