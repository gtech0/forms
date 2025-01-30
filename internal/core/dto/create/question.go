package create

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type QuestionDto struct {
	Type question.QuestionType `json:"type"`
}

type QuestionOnExistingDto struct {
	QuestionDto
	QuestionId uuid.UUID `json:"questionId"`
}

type NewQuestionDto struct {
	QuestionDto
	Description string      `json:"description"`
	Attachments []uuid.UUID `json:"attachments"`
}

func CommonQuestionDtoUnmarshal(rawQuestion json.RawMessage) (any, error) {
	var questionDto QuestionDto
	if err := json.Unmarshal(rawQuestion, &questionDto); err != nil {
		return nil, err
	}

	var questionI any
	switch questionDto.Type {
	case question.EXISTING:
		questionI = &QuestionOnExistingDto{}
	case question.MATCHING:
		questionI = &MatchingQuestionDto{}
	case question.MULTIPLE_CHOICE:
		questionI = &MultipleChoiceQuestionDto{}
	case question.SINGLE_CHOICE:
		questionI = &SingleChoiceQuestionDto{}
	case question.TEXT_INPUT:
		questionI = &TextQuestionDto{}
	default:
		return nil, errs.New(fmt.Sprintf("unknown question type: %s", questionDto.Type), 400)
	}

	if err := json.Unmarshal(rawQuestion, questionI); err != nil {
		return nil, err
	}

	return questionI, nil
}
