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

	formPublished.Groups = f.buildGroups(publishDto)
	formPublished.Users = f.buildUsers(publishDto)
	formPublished.MarkConfiguration = f.buildMarkConfiguration(publishDto)

	return formPublished
}

func (f *FormPublishedFactory) buildGroups(publishDto create.FormPublishDto) []published.FormPublishedGroup {
	groups := make([]published.FormPublishedGroup, 0)
	for _, groupId := range publishDto.GroupIds {
		var publishedGroup published.FormPublishedGroup
		publishedGroup.FormPublishedId = publishDto.FormPatternId
		publishedGroup.GroupId = groupId
		groups = append(groups, publishedGroup)
	}
	return groups
}

func (f *FormPublishedFactory) buildUsers(publishDto create.FormPublishDto) []published.FormPublishedUser {
	users := make([]published.FormPublishedUser, 0)
	for _, userId := range publishDto.UserIds {
		var publishedUser published.FormPublishedUser
		publishedUser.FormPublishedId = publishDto.FormPatternId
		publishedUser.UserId = userId
		users = append(users, publishedUser)
	}
	return users
}

func (f *FormPublishedFactory) buildMarkConfiguration(publishDto create.FormPublishDto) []published.MarkConfiguration {
	markConfigs := make([]published.MarkConfiguration, 0)
	for mark, points := range publishDto.MarkConfiguration {
		var markConfig published.MarkConfiguration
		markConfig.Mark = mark
		markConfig.MinPoints = points
		markConfig.FormPublishedId = publishDto.FormPatternId
		markConfigs = append(markConfigs, markConfig)
	}
	return markConfigs
}
