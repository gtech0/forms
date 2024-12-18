package generated

import (
	"github.com/google/uuid"
	"hedgehog-forms/model"
)

type SolutionObject struct {
	model.Base
	NumberAttempts int        `json:"numberAttempts"`
	IsIndividual   bool       `json:"isIndividual"`
	TeamOwnerID    uuid.UUID  `json:"teamOwnerId"`
	UserOwnerID    uuid.UUID  `json:"userOwnerId"`
	Score          int        `json:"score"`
	Mark           string     `json:"mark"`
	Verdict        FormStatus `json:"verdict"`
}

//	SolutionObject struct {
//		ID                    uuid.UUID
//		CreatedAt             time.Time                  `json:"created_at"`
//		UpdatedAt             time.Time                  `json:"updated_at"`
//		NumberAttempts        int                        `json:"number_attempts"`
//		IsIndividual          bool                       `json:"is_individual"`
//		TeamOwnerID           uuid.UUID                  `json:"team_owner_id"`
//		UserOwnerID           uuid.UUID                  `json:"user_owner_id"`
//		Score                 int                        `json:"score"`
//		ClassTaskID           uuid.UUID                  `json:"class_task_id"`
//		TestingVerdict        enum.TestingVerdict        `json:"testing_verdict"`
//		FinalTestingVerdict   enum.FinalTestingVerdict   `json:"final_testing_verdict"`
//		PostmoderationVerdict enum.PostmoderationVerdict `json:"postmoderation_verdict"`
//		FinalVerdict          enum.FinalVerdict          `json:"final_verdict"`
//		Comment               []CommentObject            `json:"comment"`
//	}
