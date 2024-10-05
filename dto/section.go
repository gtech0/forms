package dto

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/model/form/section"
	"hedgehog-forms/model/form/section/block"
)

type CreateSectionDto struct {
	Type section.SectionType `json:"type"`
}

type CreateSectionOnExistingDto struct {
	CreateSectionDto
	SectionId uuid.UUID `json:"sectionId"`
}

type CreateNewSectionDto struct {
	CreateSectionDto
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Blocks      []any             `json:"-"`
	RawBlocks   []json.RawMessage `json:"blocks"`
}

func (c *CreateNewSectionDto) UnmarshalJSON(b []byte) error {
	type sectionDto CreateNewSectionDto

	err := json.Unmarshal(b, (*sectionDto)(c))
	if err != nil {
		return err
	}

	for _, rawBlock := range c.RawBlocks {
		var blockDto CreateBlockDto
		err = json.Unmarshal(rawBlock, &blockDto)
		if err != nil {
			return err
		}

		var blockI any
		switch blockDto.Type {
		case block.STATIC:
			blockI = &CreateStaticBlockDto{}
		case block.DYNAMIC:
			blockI = &CreateDynamicBlockDto{}
		case block.EXISTING:
			blockI = &CreateBlockOnExistingDto{}
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

func (c *CreateNewSectionDto) MarshalJSON() ([]byte, error) {
	type sectionDto CreateNewSectionDto

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
