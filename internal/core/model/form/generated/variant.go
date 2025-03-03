package generated

import (
	"encoding/json"
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/util"
)

type Variant struct {
	Id           uuid.UUID         `json:"id"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Questions    []IQuestion       `json:"-"`
	RawQuestions []json.RawMessage `json:"questions"`
}

func (c *Variant) UnmarshalJSON(b []byte) error {
	type questions Variant

	err := json.Unmarshal(b, (*questions)(c))
	if err != nil {
		return err
	}

	for _, rawQuestion := range c.RawQuestions {
		questionI, err := CommonQuestionUnmarshal(rawQuestion)
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
