package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model"
)

type SolutionObject struct {
	model.Base
	NumberAttempts        int             `json:"numberAttempts"`
	IsIndividual          bool            `json:"isIndividual"`
	TeamOwnerId           uuid.UUID       `json:"teamOwnerId"`
	UserOwnerId           uuid.UUID       `json:"userOwnerId"`
	Score                 int             `json:"score"`
	Mark                  string          `json:"mark"`
	ClassTaskId           uuid.UUID       `json:"classTaskId"`
	TestingVerdict        string          `json:"testingVerdict"`
	FinalTestingVerdict   string          `json:"finalTestingVerdict"`
	PostmoderationVerdict string          `json:"postmoderationVerdict"`
	FinalVerdict          string          `json:"finalVerdict"`
	Comment               []CommentObject `json:"comment"`
}

type CommentObject struct {
	model.Base
	AuthorId   uuid.UUID `json:"authorId"`
	AuthorName string    `json:"authorName"`
	Content    string    `json:"content"`
	IsPrivate  bool      `json:"isPrivate"`
}
