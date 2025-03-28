package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/published"
)

type FormPublishedMapper struct {
	formPatternMapper *FormPatternMapper
}

func NewFormPublishedMapper() *FormPublishedMapper {
	return &FormPublishedMapper{
		formPatternMapper: NewFormPatternMapper(),
	}
}

func (f *FormPublishedMapper) ToBaseDto(publishedForm *published.FormPublished) *get.FormPublishedBaseDto {
	publishedBaseDto := new(get.FormPublishedBaseDto)
	publishedBaseDto.Id = publishedForm.Id
	publishedBaseDto.FormPatternId = publishedForm.FormPatternId
	publishedBaseDto.HideScore = publishedForm.HideScore
	publishedBaseDto.Deadline = publishedForm.Deadline
	publishedBaseDto.Duration = publishedForm.Duration
	publishedBaseDto.MaxAttempts = publishedForm.MaxAttempts
	groups := make([]uuid.UUID, 0)
	for _, publishedGroup := range publishedForm.Teams {
		groups = append(groups, publishedGroup.TeamId)
	}
	publishedBaseDto.TeamIds = groups

	users := make([]uuid.UUID, 0)
	for _, publishedUser := range publishedForm.Users {
		users = append(users, publishedUser.UserId)
	}
	publishedBaseDto.UserIds = users

	markConfig := make(map[int]int)
	for _, markConfiguration := range publishedForm.MarkConfiguration {
		markConfig[markConfiguration.Mark] = markConfiguration.MinPoints
	}
	publishedBaseDto.MarkConfiguration = markConfig
	return publishedBaseDto
}

func (f *FormPublishedMapper) ToDto(publishedForm *published.FormPublished) (*get.FormPublishedDto, error) {
	publishedDto := new(get.FormPublishedDto)
	formPattern, err := f.formPatternMapper.ToDto(&publishedForm.FormPattern)
	if err != nil {
		return nil, err
	}
	publishedDto.Id = publishedForm.Id
	publishedDto.FormPattern = *formPattern
	publishedDto.HideScore = publishedForm.HideScore
	publishedDto.Deadline = publishedForm.Deadline
	publishedDto.Duration = publishedForm.Duration
	groups := make([]uuid.UUID, 0)
	for _, publishedTeam := range publishedForm.Teams {
		groups = append(groups, publishedTeam.TeamId)
	}
	publishedDto.TeamIds = groups

	users := make([]uuid.UUID, 0)
	for _, publishedUser := range publishedForm.Users {
		users = append(users, publishedUser.UserId)
	}
	publishedDto.UserIds = users
	return publishedDto, nil
}
