package dto

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
		var questionDto CreateQuestionDto
		err = json.Unmarshal(rawQuestion, &questionDto)
		if err != nil {
			return err
		}

		var questionI any
		switch questionDto.Type {
		case question.EXISTING:
			questionI = &CreateQuestionOnExistingDto{}
		case question.MATCHING:
			questionI = &CreateMatchingQuestionDto{}
		case question.MULTIPLE_CHOICE:
			questionI = &CreateMultipleChoiceQuestionDto{}
		case question.SINGLE_CHOICE:
			questionI = &CreateSingleChoiceQuestionDto{}
		case question.TEXT_INPUT:
			questionI = &CreateTextQuestionDto{}
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
