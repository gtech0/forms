package block

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type DynamicBlock struct {
	model.Base
	QuestionCount int
	Questions     []*question.Question
	BlockId       uuid.UUID `gorm:"type:uuid"`
}
