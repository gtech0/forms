package create

import (
	"encoding/json"
	"github.com/google/uuid"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/util"
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
	QuestionCount int               `json:"questionCount"`
	Questions     []any             `json:"-"`
	RawQuestions  []json.RawMessage `json:"questions"`
}

func (d *DynamicBlockDto) UnmarshalJSON(b []byte) error {
	type blockDto DynamicBlockDto

	err := json.Unmarshal(b, (*blockDto)(d))
	if err != nil {
		return err
	}

	for _, rawQuestion := range d.RawQuestions {
		questionI, err := CommonQuestionDtoUnmarshal(rawQuestion)
		if err != nil {
			return err
		}

		d.Questions = append(d.Questions, questionI)
	}

	return nil
}

func (d *DynamicBlockDto) MarshalJSON() ([]byte, error) {
	type blockDto DynamicBlockDto

	rawMessage, err := util.CommonMarshal(d.Questions)
	if err != nil {
		return nil, err
	}
	d.RawQuestions = rawMessage

	return json.Marshal((*blockDto)(d))
}
