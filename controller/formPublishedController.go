package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/published"
	"hedgehog-forms/service"
	"net/http"
)

type FormPublishedController struct {
	formPatternService *service.FormPatternService
}

func NewFormPublishedController() *FormPublishedController {
	return &FormPublishedController{
		formPatternService: service.NewFormPatternService(),
	}
}

func (f *FormPublishedController) PublishForm(ctx *gin.Context) {
	var publishedDto create.FormPublishedDto
	if err := ctx.Bind(&publishedDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	pattern, err := f.formPatternService.GetForm(publishedDto.PatternId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	var formPublished published.FormPublished
	formPublished.Id = uuid.New()
	formPublished.Deadline = publishedDto.Deadline
	formPublished.Duration = publishedDto.Duration

	groups := make([]published.FormPublishedGroup, 0)
	for _, groupId := range publishedDto.GroupIds {
		var publishedGroup published.FormPublishedGroup
		publishedGroup.FormPublishedId = formPublished.Id
		publishedGroup.GroupId = groupId
		groups = append(groups, publishedGroup)
	}
	formPublished.Groups = groups

	users := make([]published.FormPublishedUser, 0)
	for _, userId := range publishedDto.UserIds {
		var publishedUser published.FormPublishedUser
		publishedUser.FormPublishedId = formPublished.Id
		publishedUser.UserId = userId
		users = append(users, publishedUser)
	}
	formPublished.Users = users

	formPublished.HideScore = publishedDto.HideScore
	formPublished.PostModeration = publishedDto.PostModeration
	formPublished.FormPattern = pattern
	markConfigs := make([]published.MarkConfiguration, 0)
	for mark, points := range publishedDto.MarkConfiguration {
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
	ctx.Status(http.StatusOK)
}
