package published

import (
	"github.com/google/uuid"
	model2 "hedgehog-forms/internal/core/model"
	"hedgehog-forms/internal/core/model/form/generated"
)

type Solution struct {
	model2.Base

	IsIndividual   *bool `json:"isIndividual"`
	NumberAttempts int   `json:"numberAttempts" gorm:"default:0"`
	Score          int   `json:"score"`

	UserOwnerId *uuid.UUID `json:"userOwnerId,omitempty"`
	TeamOwnerId *uuid.UUID `json:"teamOwnerId,omitempty"`

	ClassTaskId uuid.UUID `json:"classTaskId"`

	TestingVerdict        string `json:"testingVerdict" gorm:"default:NOT TESTED"`
	FinalTestingVerdict   string `json:"finalTestingVerdict" gorm:"default:NOT TESTED"`
	PostmoderationVerdict string `json:"postmoderationVerdict" gorm:"default:PENDING"`
	FinalVerdict          string `json:"finalVerdict" gorm:"default:PENDING"`

	Submissions      []generated.Submission `json:"submissions"`
	SolutionComments []SolutionComment      `json:"solutionComments"`

	Files []model2.File `json:"files" gorm:"-:all"`
}

type SolutionComment struct {
	model2.Base

	Content   string `json:"content"`
	IsPrivate *bool  `json:"isPrivate"`

	SolutionId uuid.UUID `json:"solutionId"`
	Solution   *Solution

	AuthorId uuid.UUID `json:"authorId"`
}
