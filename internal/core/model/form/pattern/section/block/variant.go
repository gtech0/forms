package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type Variant struct {
	model.Base
	Title         string
	Description   string
	StaticBlockID uuid.UUID `gorm:"type:uuid"`
	Questions     []*question.Question
}
