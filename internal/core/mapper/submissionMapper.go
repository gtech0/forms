package mapper

import (
	"hedgehog-forms/internal/core/repository"
)

type SubmissionMapper struct {
	formPublishedMapper     *FormPublishedMapper
	submissionRepository    *repository.SubmissionRepository
	formPublishedRepository *repository.FormPublishedRepository
}

func NewSubmissionMapper() *SubmissionMapper {
	return &SubmissionMapper{
		formPublishedMapper:     NewFormPublishedMapper(),
		submissionRepository:    repository.NewSubmissionRepository(),
		formPublishedRepository: repository.NewFormPublishedRepository(),
	}
}
