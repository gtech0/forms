package get

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/model"
)

type AttemptObject struct {
	model.Base
	TaskId     uuid.UUID  `json:"taskId"`
	TaskName   string     `json:"taskName"`
	SectionId  uuid.UUID  `json:"sectionId"`
	TotalScore int        `json:"totalScore"`
	SolutionId *uuid.UUID `json:"solutionId"`
	SenderId   uuid.UUID  `json:"senderId"`
	SenderName string     `json:"senderName"`
	//ProgrammingLanguageId   *uuid.UUID   `json:"programmingLanguageId"`
	//ProgrammingLanguageName *string      `json:"programmingLanguageName"`
	FinalTestingVerdict   string       `json:"finalTestingVerdict"`
	TestingVerdict        string       `json:"testingVerdict"`
	PostmoderationVerdict string       `json:"postmoderationVerdict"`
	FileData              []model.File `json:"fileData"`
	ClassId               uuid.UUID    `json:"classId"`
}
