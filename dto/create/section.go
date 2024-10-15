package create

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/model/form/pattern/section"
	"hedgehog-forms/model/form/pattern/section/block"
)

type SectionDto struct {
	Type section.SectionType `json:"type"`
}

type SectionOnExistingDto struct {
	SectionDto
	SectionId uuid.UUID `json:"sectionId"`
}

type NewSectionDto struct {
	SectionDto
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Blocks      []any             `json:"-"`
	RawBlocks   []json.RawMessage `json:"blocks"`
}

func (c *NewSectionDto) UnmarshalJSON(b []byte) error {
	type sectionDto NewSectionDto

	err := json.Unmarshal(b, (*sectionDto)(c))
	if err != nil {
		return err
	}

	for _, rawBlock := range c.RawBlocks {
		var blockDto BlockDto
		err = json.Unmarshal(rawBlock, &blockDto)
		if err != nil {
			return err
		}

		var blockI any
		switch blockDto.Type {
		case block.STATIC:
			blockI = &StaticBlockDto{}
		case block.DYNAMIC:
			blockI = &DynamicBlockDto{}
		case block.EXISTING:
			blockI = &BlockOnExistingDto{}
		default:
			return fmt.Errorf("unknown block type: %s", blockDto.Type)
		}

		err = json.Unmarshal(rawBlock, blockI)
		if err != nil {
			return err
		}

		c.Blocks = append(c.Blocks, blockI)
	}

	return nil
}

func (c *NewSectionDto) MarshalJSON() ([]byte, error) {
	type sectionDto NewSectionDto

	if c.Blocks != nil {
		for _, blockDto := range c.Blocks {
			rawBlock, err := json.Marshal(blockDto)
			if err != nil {
				return nil, err
			}
			c.RawBlocks = append(c.RawBlocks, rawBlock)
		}
	}

	return json.Marshal((*sectionDto)(c))
}
