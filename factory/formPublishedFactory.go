package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/published"
)

type FormPublishedFactory struct{}

func NewFormPublishedFactory() *FormPublishedFactory {
	return &FormPublishedFactory{}
}

func (f *FormPublishedFactory) Build(publishDto create.FormPublishDto) published.FormPublished {
	var formPublished published.FormPublished
	formPublished.Id = uuid.New()
	formPublished.Deadline = publishDto.Deadline
	formPublished.Duration = publishDto.Duration
	formPublished.HideScore = publishDto.HideScore
	formPublished.PostModeration = publishDto.PostModeration
	formPublished.FormPatternId = publishDto.FormPatternId

	formPublished.Groups = f.BuildGroups(publishDto.GroupIds, formPublished.Id)
	formPublished.Users = f.BuildUsers(publishDto.UserIds, formPublished.Id)
	formPublished.MarkConfiguration = f.BuildMarkConfiguration(publishDto.MarkConfiguration, formPublished.Id)

	return formPublished
}

func (f *FormPublishedFactory) Update(
	formPublished *published.FormPublished,
	formPublishedDto create.UpdateFormPublishedDto,
) {
	formPublished.Deadline = formPublishedDto.Deadline
	formPublished.Duration = formPublishedDto.Duration
	formPublished.HideScore = formPublishedDto.HideScore

	formPublished.Groups = f.BuildGroups(formPublishedDto.GroupIds, formPublished.Id)
	formPublished.Users = f.BuildUsers(formPublishedDto.UserIds, formPublished.Id)
	formPublished.MarkConfiguration = f.BuildMarkConfiguration(formPublishedDto.MarkConfiguration, formPublished.Id)
}

func (f *FormPublishedFactory) BuildGroups(groupIds []uuid.UUID, publishedId uuid.UUID) []published.FormPublishedGroup {
	groups := make([]published.FormPublishedGroup, 0)
	for _, groupId := range groupIds {
		var publishedGroup published.FormPublishedGroup
		publishedGroup.FormPublishedId = publishedId
		publishedGroup.GroupId = groupId
		groups = append(groups, publishedGroup)
	}
	return groups
}

func (f *FormPublishedFactory) BuildUsers(userIds []uuid.UUID, publishedId uuid.UUID) []published.FormPublishedUser {
	users := make([]published.FormPublishedUser, 0)
	for _, userId := range userIds {
		var publishedUser published.FormPublishedUser
		publishedUser.FormPublishedId = publishedId
		publishedUser.UserId = userId
		users = append(users, publishedUser)
	}
	return users
}

func (f *FormPublishedFactory) BuildMarkConfiguration(
	marks map[string]int,
	publishedId uuid.UUID,
) []published.MarkConfiguration {
	markConfigs := make([]published.MarkConfiguration, 0)
	for mark, points := range marks {
		var markConfig published.MarkConfiguration
		markConfig.Mark = mark
		markConfig.MinPoints = points
		markConfig.FormPublishedId = publishedId
		markConfigs = append(markConfigs, markConfig)
	}
	return markConfigs
}
