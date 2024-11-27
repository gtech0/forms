package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/published"
	"hedgehog-forms/repository"
)

type FormPublishedFactory struct {
	groupRepository *repository.GroupRepository
	userRepository  *repository.UserRepository
}

func NewFormPublishedFactory() *FormPublishedFactory {
	return &FormPublishedFactory{
		userRepository:  repository.NewUserRepository(),
		groupRepository: repository.NewGroupRepository(),
	}
}

func (f *FormPublishedFactory) Build(publishDto create.FormPublishDto) (*published.FormPublished, error) {
	formPublished := new(published.FormPublished)
	formPublished.Id = uuid.New()
	formPublished.Deadline = publishDto.Deadline
	formPublished.Duration = publishDto.Duration
	formPublished.HideScore = publishDto.HideScore
	formPublished.PostModeration = publishDto.PostModeration
	formPublished.FormPatternId = publishDto.FormPatternId
	if err := f.BuildGroups(publishDto.GroupIds, formPublished); err != nil {
		return nil, err
	}

	if err := f.BuildUsers(publishDto.UserIds, formPublished); err != nil {
		return nil, err
	}

	formPublished.MarkConfiguration = f.BuildMarkConfiguration(publishDto.MarkConfiguration, formPublished.Id)
	return formPublished, nil
}

func (f *FormPublishedFactory) Update(
	formPublished *published.FormPublished,
	formPublishedDto create.UpdateFormPublishedDto,
) error {
	formPublished.Deadline = formPublishedDto.Deadline
	formPublished.Duration = formPublishedDto.Duration
	formPublished.HideScore = formPublishedDto.HideScore
	if err := f.BuildGroups(formPublishedDto.GroupIds, formPublished); err != nil {
		return err
	}

	if err := f.BuildUsers(formPublishedDto.UserIds, formPublished); err != nil {
		return err
	}

	formPublished.MarkConfiguration = f.BuildMarkConfiguration(formPublishedDto.MarkConfiguration, formPublished.Id)
	return nil
}

func (f *FormPublishedFactory) BuildGroups(
	groupIds []uuid.UUID,
	formPublished *published.FormPublished,
) error {
	groups := make([]published.Group, 0)
	for _, groupId := range groupIds {
		group, err := f.groupRepository.FindById(groupId)
		if err != nil {
			return err
		}
		groups = append(groups, *group)
	}
	formPublished.Groups = groups
	return nil
}

func (f *FormPublishedFactory) BuildUsers(
	userIds []uuid.UUID,
	formPublished *published.FormPublished,
) error {
	users := make([]published.User, 0)
	for _, userId := range userIds {
		user, err := f.userRepository.FindById(userId)
		if err != nil {
			return err
		}
		users = append(users, *user)
	}
	formPublished.Users = users
	return nil
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
