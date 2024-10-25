package processor

import (
	"hedgehog-forms/mapper"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section/block/question"
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

func (t *TextInputProcessor) markAnswer(textInput *generated.TextInput, answer string) error {
	textInput.EnteredAnswer = answer
	return nil
}

func (t *TextInputProcessor) markAnswerAndCalculatePoints(
	textInput *generated.TextInput,
	textInputObj *question.TextInput,
	enteredAnswer string,
) (int, error) {
	textInput.EnteredAnswer = enteredAnswer

	answers := t.textInputMapper.AnswersToDto(textInputObj.Answers)
	isEnteredAnswerCorrect := func() bool {
		if textInputObj.IsCaseSensitive {
			return slices.Contains(answers, enteredAnswer)
		}

		for _, answer := range answers {
			return strings.EqualFold(answer, enteredAnswer)
		}
		return false
	}

	if isEnteredAnswerCorrect() {
		textInput.Points = textInputObj.Points
	}

	return textInput.Points, nil
}
