package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/repository"
	"time"
)

type FormGeneratedMapper struct {
	formPublishedMapper     *FormPublishedMapper
	submissionRepository    *repository.SubmissionRepository
	formPublishedRepository *repository.FormPublishedRepository
}

func NewFormGeneratedMapper() *FormGeneratedMapper {
	return &FormGeneratedMapper{
		formPublishedMapper:     NewFormPublishedMapper(),
		submissionRepository:    repository.NewSubmissionRepository(),
		formPublishedRepository: repository.NewFormPublishedRepository(),
	}
}

func (f *FormGeneratedMapper) ToDto(formGenerated *generated.FormGenerated) (*get.FormGeneratedDto, error) {
	formGeneratedDto := new(get.FormGeneratedDto)
	formGeneratedDto.Id = formGenerated.Id
	formGeneratedDto.Status = formGenerated.Status
	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedId)
	if err != nil {
		return nil, err
	}

	formGeneratedDto.FormPublished = *f.formPublishedMapper.ToBaseDto(formPublished)
	submission, err := f.submissionRepository.FindById(formGenerated.SubmissionId)
	if err != nil {
		return nil, err
	}

	formGeneratedDto.UserId = submission.UserId
	formGeneratedDto.Sections = formGenerated.Sections
	return formGeneratedDto, nil
}

func (f *FormGeneratedMapper) ToMyDto(formGenerated *generated.FormGenerated) (*get.MyGeneratedDto, error) {
	myGeneratedDto := new(get.MyGeneratedDto)
	myGeneratedDto.Id = formGenerated.Id
	myGeneratedDto.Status = formGenerated.Status
	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedId)
	if err != nil {
		return nil, err
	}

	myGeneratedDto.FormPublished = *f.formPublishedMapper.ToBaseDto(formPublished)
	myGeneratedDto.Points = formGenerated.Points
	myGeneratedDto.Mark = formGenerated.Mark

	hideScore := myGeneratedDto.FormPublished.HideScore
	isAfterDeadline := myGeneratedDto.FormPublished.Deadline.After(time.Now())
	if hideScore && isAfterDeadline {
		myGeneratedDto.Points = 0
		myGeneratedDto.Mark = 0
	}

	return myGeneratedDto, nil
}

func (f *FormGeneratedMapper) ToSubmittedDto(formGenerated generated.FormGenerated) *get.SubmittedDto {
	submittedDto := new(get.SubmittedDto)
	submittedDto.Status = formGenerated.Status
	submittedDto.Points = formGenerated.Points
	submittedDto.Mark = formGenerated.Mark
	//submittedDto.SubmitTime = formGenerated.SubmitTime
	return submittedDto
}
