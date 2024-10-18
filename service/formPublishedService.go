package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/mapper"
	"hedgehog-forms/model/form/published"
)

type FormPublishedService struct {
	formPatternService  *FormPatternService
	formPublishedMapper *mapper.FormPublishedMapper
}

func NewFormPublishedService() *FormPublishedService {
	return &FormPublishedService{
		formPatternService:  NewFormPatternService(),
		formPublishedMapper: mapper.NewFormPublishedMapper(),
	}
}

func (f *FormPublishedService) PublishForm(publishDto create.FormPublishDto) (get.FormPublishedBaseDto, error) {
	formPatternId, err := f.formPatternService.getFormId(publishDto.FormPatternId)
	if err != nil {
		return get.FormPublishedBaseDto{}, err
	}

	var formPublished published.FormPublished
	formPublished.Id = uuid.New()
	formPublished.Deadline = publishDto.Deadline
	formPublished.Duration = publishDto.Duration
	formPublished.HideScore = publishDto.HideScore
	formPublished.PostModeration = publishDto.PostModeration
	formPublished.FormPatternId = formPatternId

	groups := make([]published.FormPublishedGroup, 0)
	for _, groupId := range publishDto.GroupIds {
		var publishedGroup published.FormPublishedGroup
		publishedGroup.FormPublishedId = formPublished.Id
		publishedGroup.GroupId = groupId
		groups = append(groups, publishedGroup)
	}
	formPublished.Groups = groups

	users := make([]published.FormPublishedUser, 0)
	for _, userId := range publishDto.UserIds {
		var publishedUser published.FormPublishedUser
		publishedUser.FormPublishedId = formPublished.Id
		publishedUser.UserId = userId
		users = append(users, publishedUser)
	}
	formPublished.Users = users

	markConfigs := make([]published.MarkConfiguration, 0)
	for mark, points := range publishDto.MarkConfiguration {
		var markConfig published.MarkConfiguration
		markConfig.Mark = mark
		markConfig.MinPoints = points
		markConfig.FormPublishedId = formPublished.Id
		markConfigs = append(markConfigs, markConfig)
	}
	formPublished.MarkConfiguration = markConfigs

	if err = database.DB.Create(&formPublished).Error; err != nil {
		return get.FormPublishedBaseDto{}, errs.New(err.Error(), 500)
	}

	return f.formPublishedMapper.ToBaseDto(formPublished), nil
}

func (f *FormPublishedService) GetForm(formId string) (get.FormPublishedDto, error) {
	if formId == "" {
		return get.FormPublishedDto{}, errs.New("formId is required", 400)
	}

	id, err := uuid.Parse(formId)
	if err != nil {
		return get.FormPublishedDto{}, errs.New(err.Error(), 500)
	}

	var formPublished published.FormPublished
	if err = database.DB.Preload("~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~.~~~as~~~").
		Model(&published.FormPublished{}).
		Where("id = ?", id).
		First(&formPublished).
		Error; err != nil {
		return get.FormPublishedDto{}, errs.New(err.Error(), 500)
	}

	publishedDto, err := f.formPublishedMapper.ToDto(formPublished)
	if err != nil {
		return get.FormPublishedDto{}, err
	}

	return publishedDto, nil
}
