package dto

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/model/form/section/block"
	"hedgehog-forms/model/form/section/block/question"
)

type CreateBlockDto struct {
	Type block.BlockType `json:"type"`
}

type CreateBlockOnExistingDto struct {
	CreateBlockDto
	BlockId uuid.UUID `json:"block_id"`
}

type NewBlockDto struct {
	CreateBlockDto
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CreateStaticBlockDto struct {
	NewBlockDto
	Variants []UpdateVariantDto `json:"variants"`
}

type CreateDynamicBlockDto struct {
	NewBlockDto
	Questions    []any             `json:"-"`
	RawQuestions []json.RawMessage `json:"questions"`
}

func (c *CreateDynamicBlockDto) UnmarshalJSON(b []byte) error {
	type blockDto CreateDynamicBlockDto

	err := json.Unmarshal(b, (*blockDto)(c))
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

func (c *CreateDynamicBlockDto) MarshalJSON() ([]byte, error) {
	type blockDto CreateDynamicBlockDto

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
