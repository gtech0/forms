package generated

import (
	"encoding/json"
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model/form/pattern/block"
	"hedgehog-forms/internal/core/util"
)

type Block struct {
	Id           uuid.UUID         `json:"id"`
	Type         block.BlockType   `json:"type"`
	Title        string            `json:"title"`
	Description  string            `json:"description"`
	Variant      *Variant          `json:"variant"`
	Questions    []IQuestion       `json:"-"`
	RawQuestions []json.RawMessage `json:"questions"`
}

func (c *Block) UnmarshalJSON(b []byte) error {
	type questions Block

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

func (c *Block) MarshalJSON() ([]byte, error) {
	type questions Block

	rawMessage, err := util.CommonMarshal(c.Questions)
	if err != nil {
		return nil, err
	}
	c.RawQuestions = rawMessage

	return json.Marshal((*questions)(c))
}
