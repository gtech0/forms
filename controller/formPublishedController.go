package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/mapper"
	"hedgehog-forms/model/form/published"
	"hedgehog-forms/service"
	"net/http"
)

type FormPublishedController struct {
	formPatternService   *service.FormPatternService
	formPublishedMapper  *mapper.FormPublishedMapper
	formPublishedService *service.FormPublishedService
}

func NewFormPublishedController() *FormPublishedController {
	return &FormPublishedController{
		formPatternService:   service.NewFormPatternService(),
		formPublishedMapper:  mapper.NewFormPublishedMapper(),
		formPublishedService: service.NewFormPublishedService(),
	}
}

func (f *FormPublishedController) PublishForm(ctx *gin.Context) {
	var publishDto create.FormPublishDto
	if err := ctx.Bind(&publishDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pattern, err := f.formPatternService.GetForm(publishDto.FormPatternId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var formPublished published.FormPublished
	formPublished.Id = uuid.New()
	formPublished.Deadline = publishDto.Deadline
	formPublished.Duration = publishDto.Duration

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

	formPublished.HideScore = publishDto.HideScore
	formPublished.PostModeration = publishDto.PostModeration
	formPublished.FormPatternId = pattern.Id
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
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, f.formPublishedMapper.ToBaseDto(formPublished))
}

func (f *FormPublishedController) GetForm(ctx *gin.Context) {
	formId := ctx.Param("formId")
	if formId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "formId is required",
		})
		return
	}

	id, err := uuid.Parse(formId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	formPublished, err := f.formPublishedService.GetForm(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	publishedDto, err := f.formPublishedMapper.ToDto(formPublished)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, publishedDto)
}
