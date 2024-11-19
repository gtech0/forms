package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type Variant struct {
	model.Base
	Title         string
	Description   string
	StaticBlockID uuid.UUID `gorm:"type:uuid"`
	Questions     []*question.Question
}
