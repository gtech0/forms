package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/section/block/question"
)

type Variant struct {
	model.BaseModel
	Title         string
	Description   string
	StaticBlockID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	Questions     []question.IQuestion
}
