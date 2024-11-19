package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type DynamicBlock struct {
	model.Base
	QuestionCount int
	Questions     []*question.Question
	BlockId       uuid.UUID `gorm:"type:uuid"`
}
