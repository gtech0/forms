package create

import (
	"encoding/json"
	"hedgehog-forms/util"
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
		questionI, err := CommonQuestionDtoUnmarshal(rawQuestion)
		if err != nil {
			return err
		}

		c.Questions = append(c.Questions, questionI)
	}

	return nil
}

func (c *UpdateVariantDto) MarshalJSON() ([]byte, error) {
	type variantDto UpdateVariantDto

	rawMessage, err := util.CommonMarshal(c.Questions)
	if err != nil {
		return nil, err
	}
	c.RawQuestions = rawMessage

	return json.Marshal((*variantDto)(c))
}
