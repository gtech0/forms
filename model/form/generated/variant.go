package generated

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"hedgehog-forms/util"
)

type Variant struct {
	Id           uuid.UUID         `json:"id"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Questions    []any             `json:"-"`
	RawQuestions []json.RawMessage `json:"questions"`
}

func (c *Variant) UnmarshalJSON(b []byte) error {
	type questions Variant

	err := json.Unmarshal(b, (*questions)(c))
	if err != nil {
		return err
	}

	for _, rawQuestion := range c.RawQuestions {
		var generatedQuestion Question
		err = json.Unmarshal(rawQuestion, &generatedQuestion)
		if err != nil {
			return err
		}

		var questionI any
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
			return fmt.Errorf("unknown question type: %s", generatedQuestion.Type)
		}

		err = json.Unmarshal(rawQuestion, questionI)
		if err != nil {
			return err
		}

		c.Questions = append(c.Questions, questionI)
	}

	return nil
}

func (c *Variant) MarshalJSON() ([]byte, error) {
	type questions Variant

	rawMessage, err := util.CommonMarshal(c.Questions)
	if err != nil {
		return nil, err
	}
	c.RawQuestions = rawMessage

	return json.Marshal((*questions)(c))
}
