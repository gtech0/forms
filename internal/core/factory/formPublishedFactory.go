package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/model/form/published"
)

type FormPublishedFactory struct{}

func NewFormPublishedFactory() *FormPublishedFactory {
	return &FormPublishedFactory{}
}

func (f *FormPublishedFactory) Build(publishDto create.FormPublishDto) (*published.FormPublished, error) {
	formPublished := new(published.FormPublished)
	formPublished.Id = uuid.New()
	formPublished.Deadline = publishDto.Deadline
	formPublished.Duration = publishDto.Duration
	formPublished.HideScore = publishDto.HideScore
	formPublished.PostModeration = publishDto.PostModeration
	formPublished.FormPatternId = publishDto.FormPatternId
	formPublished.MaxAttempts = publishDto.MaxAttempts
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
	formPublished.MarkConfiguration = f.BuildMarkConfiguration(formPublishedDto.MarkConfiguration, formPublished.Id)
	return nil
}

func (f *FormPublishedFactory) BuildMarkConfiguration(
	marks map[int]int,
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
