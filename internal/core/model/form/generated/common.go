package generated

import (
	"encoding/json"
	"fmt"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

func CommonQuestionUnmarshal(rawQuestion json.RawMessage) (IQuestion, error) {
	var generatedQuestion Question
	err := json.Unmarshal(rawQuestion, &generatedQuestion)
	if err != nil {
		return nil, err
	}

	var questionI IQuestion
	switch generatedQuestion.Type {
	case question.MATCHING:
		questionI = &Matching{}
	case question.MULTIPLE_CHOICE:
		questionI = &MultipleChoice{}
	case question.SINGLE_CHOICE:
		questionI = &SingleChoice{}
	case question.TEXT_INPUT:
		questionI = &TextInput{}
	default:
		return nil, fmt.Errorf("unknown question type: %s", generatedQuestion.Type)
	}

	err = json.Unmarshal(rawQuestion, questionI)
	if err != nil {
		return nil, err
	}

	return questionI, nil
}
