package dto

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/model/form/section"
)

type CreateFormPatternDto struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	SubjectId   uuid.UUID         `json:"subjectId"`
	Sections    []any             `json:"-"`
	RawSections []json.RawMessage `json:"sections"`
}

func (c *CreateFormPatternDto) UnmarshalJSON(b []byte) error {
	type patternDto CreateFormPatternDto

	err := json.Unmarshal(b, (*patternDto)(c))
	if err != nil {
		return err
	}

	for _, rawSection := range c.RawSections {
		var sectionDto CreateSectionDto
		err = json.Unmarshal(rawSection, &sectionDto)
		if err != nil {
			return err
		}

		var sectionI any
		switch sectionDto.Type {
		case section.NEW:
			sectionI = &CreateNewSectionDto{}
		case section.EXISTING:
			sectionI = &CreateSectionOnExistingDto{}
		default:
			return fmt.Errorf("unknown section type: %s", sectionDto.Type)
		}

		err = json.Unmarshal(rawSection, sectionI)
		if err != nil {
			return err
		}

		c.Sections = append(c.Sections, sectionI)
	}

	return nil
}

func (c *CreateFormPatternDto) MarshalJSON() ([]byte, error) {
	type patternDto CreateFormPatternDto

	if c.Sections != nil {
		for _, sectionDto := range c.Sections {
			rawSection, err := json.Marshal(sectionDto)
			if err != nil {
				return nil, err
			}
			c.RawSections = append(c.RawSections, rawSection)
		}
	}

	return json.Marshal((*patternDto)(c))
}
