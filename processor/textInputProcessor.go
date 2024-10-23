package processor

import (
	"hedgehog-forms/dto/get"
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

func (t *TextInputProcessor) markAnswers(textInputQuestions []*generated.TextInput, answersDto get.AnswerDto) error {
	for questionId, enteredAnswer := range answersDto.TextInput {
		textInput, err := findQuestion[*generated.TextInput](textInputQuestions, questionId)
		if err != nil {
			return err
		}

		textInput.EnteredAnswer = enteredAnswer
	}
	return nil
}

func (t *TextInputProcessor) markAnswersAndCalculatePoints(
	textInputQuestions []*generated.TextInput,
	textInputObjs []*question.TextInput,
	answersDto get.AnswerDto,
) (int, error) {
	var points int
	for questionId, enteredAnswer := range answersDto.TextInput {
		textInputQuestion, err := findQuestion[*generated.TextInput](textInputQuestions, questionId)
		if err != nil {
			return 0, err
		}

		textInputObj, err := findQuestionObj[*question.TextInput](textInputObjs, questionId)
		if err != nil {
			return 0, err
		}

		points += t.calculateAndSetPoints(textInputQuestion, textInputObj, enteredAnswer)
	}
	return points, nil
}

func (t *TextInputProcessor) calculateAndSetPoints(
	textInput *generated.TextInput,
	textInputObj *question.TextInput,
	enteredAnswer string,
) int {
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

	return textInput.Points
}
