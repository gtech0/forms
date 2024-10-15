package create

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/model/form"
	"hedgehog-forms/model/form/pattern/section/block"
)

type BlockDto struct {
	Type block.BlockType `json:"type"`
}

type BlockOnExistingDto struct {
	BlockDto
	BlockId uuid.UUID `json:"block_id"`
}

type NewBlockDto struct {
	BlockDto
	Title       string `json:"title"`
	Description string `json:"description"`
}

type StaticBlockDto struct {
	NewBlockDto
	Variants []UpdateVariantDto `json:"variants"`
}

type DynamicBlockDto struct {
	NewBlockDto
	Questions    []any             `json:"-"`
	RawQuestions []json.RawMessage `json:"questions"`
}

func (c *DynamicBlockDto) UnmarshalJSON(b []byte) error {
	type blockDto DynamicBlockDto

	err := json.Unmarshal(b, (*blockDto)(c))
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
		case form.EXISTING:
			questionI = &QuestionOnExistingDto{}
		case form.MATCHING:
			questionI = &MatchingQuestionDto{}
		case form.MULTIPLE_CHOICE:
			questionI = &MultipleChoiceQuestionDto{}
		case form.SINGLE_CHOICE:
			questionI = &SingleChoiceQuestionDto{}
		case form.TEXT_INPUT:
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

func (c *DynamicBlockDto) MarshalJSON() ([]byte, error) {
	type blockDto DynamicBlockDto

	if c.Questions != nil {
		for _, questionDto := range c.Questions {
			rawQuestion, err := json.Marshal(questionDto)
			if err != nil {
				return nil, err
			}
			c.RawQuestions = append(c.RawQuestions, rawQuestion)
		}
	}

	return json.Marshal((*blockDto)(c))
}
