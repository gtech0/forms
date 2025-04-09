package mapper

import (
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
	publishedDto.MaxAttempts = publishedForm.MaxAttempts

	markConfig := make(map[int]int)
	for _, markConfiguration := range publishedForm.MarkConfiguration {
		markConfig[markConfiguration.Mark] = markConfiguration.MinPoints
	}
	publishedDto.MarkConfiguration = markConfig
	return publishedDto, nil
}

func (f *FormPublishedMapper) ToTaskDescription(publishedForm *published.FormPublished) (*get.TaskDescription, error) {
	formPattern := publishedForm.FormPattern

	taskDescription := new(get.TaskDescription)
	taskDescription.Name = formPattern.Title
	taskDescription.Description = formPattern.Description
	return taskDescription, nil
}
