package get

import "github.com/google/uuid"

type SectionDto struct {
	Id          uuid.UUID   `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Blocks      []IBlockDto `json:"blocks"`
}

type IntegrationSectionDto struct {
	Id          uuid.UUID             `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Blocks      []IntegrationBlockDto `json:"blocks"`
}
