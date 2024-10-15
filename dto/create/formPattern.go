package create

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/model/form/pattern/section"
)

type FormPatternDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	//OwnerId     uuid.UUID         `json:"ownerId"`
	SubjectId   uuid.UUID         `json:"subjectId"`
	Sections    []any             `json:"-"`
	RawSections []json.RawMessage `json:"sections"`
}

func (c *FormPatternDto) UnmarshalJSON(b []byte) error {
	type patternDto FormPatternDto

	err := json.Unmarshal(b, (*patternDto)(c))
	if err != nil {
		return err
	}

	for _, rawSection := range c.RawSections {
		var sectionDto SectionDto
		err = json.Unmarshal(rawSection, &sectionDto)
		if err != nil {
			return err
		}

		var sectionI any
		switch sectionDto.Type {
		case section.NEW:
			sectionI = new(NewSectionDto)
		case section.EXISTING:
			sectionI = new(SectionOnExistingDto)
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

func (c *FormPatternDto) MarshalJSON() ([]byte, error) {
	type patternDto FormPatternDto

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
