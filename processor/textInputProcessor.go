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

func (t *TextInputProcessor) calculateAndSetPoints(
	textInputObj *question.TextInput,
	textInput *generated.TextInput,
	enteredAnswer string,
) int {
	textInput.EnteredAnswer = enteredAnswer //FIXME

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

	return textInput.Points
}
