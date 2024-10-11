package create

import (
	"encoding/json"
	"fmt"
	"hedgehog-forms/model/form/section/block/question"
)

type UpdateVariantDto struct {
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Questions    []any             `json:"-"`
	RawQuestions []json.RawMessage `json:"questions"`
}

func (c *UpdateVariantDto) UnmarshalJSON(b []byte) error {
	type variantDto UpdateVariantDto

	err := json.Unmarshal(b, (*variantDto)(c))
	if err != nil {
		return err
	}

	for _, rawQuestion := range c.RawQuestions {
		var questionDto QuestionDto
		err = json.Unmarshal(rawQuestion, &questionDto)
		if err != nil {
			return err
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
			return fmt.Errorf("unknown question type: %s", questionDto.Type)
		}

		err = json.Unmarshal(rawQuestion, questionI)
		if err != nil {
			return err
		}

		c.Questions = append(c.Questions, questionI)
	}

	return nil
}

func (c *UpdateVariantDto) MarshalJSON() ([]byte, error) {
	type variantDto UpdateVariantDto

	if c.Questions != nil {
		for _, questionDto := range c.Questions {
			rawQuestion, err := json.Marshal(questionDto)
			if err != nil {
				return nil, err
			}
			c.RawQuestions = append(c.RawQuestions, rawQuestion)
		}
	}

	return json.Marshal((*variantDto)(c))
}