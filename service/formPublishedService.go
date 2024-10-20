package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/repository"
)

type FormPublishedService struct {
	formPublishedRepository *repository.FormPublishedRepository
	formPatternService      *FormPatternService
	formPublishedMapper     *mapper.FormPublishedMapper
	formPublishedFactory    *factory.FormPublishedFactory
}

func NewFormPublishedService() *FormPublishedService {
	return &FormPublishedService{
		formPublishedRepository: repository.NewFormPublishedRepository(),
		formPatternService:      NewFormPatternService(),
		formPublishedMapper:     mapper.NewFormPublishedMapper(),
		formPublishedFactory:    factory.NewFormPublishedFactory(),
	}
}

func (f *FormPublishedService) PublishForm(publishDto create.FormPublishDto) (*get.FormPublishedBaseDto, error) {
	formPatternId, err := f.formPatternService.getFormId(publishDto.FormPatternId)
	if err != nil {
		return nil, err
	}

	publishDto.FormPatternId = formPatternId
	formPublished := f.formPublishedFactory.Build(publishDto)

	if err = f.formPublishedRepository.Create(&formPublished); err != nil {
		return nil, err
	}

	return f.formPublishedMapper.ToBaseDto(formPublished), nil
}

func (f *FormPublishedService) GetForm(formId string) (*get.FormPublishedDto, error) {
	if formId == "" {
		return nil, errs.New("formId is required", 400)
	}

	id, err := uuid.Parse(formId)
	if err != nil {
		return nil, errs.New(err.Error(), 500)
	}

	formPublished, err := f.formPublishedRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	publishedDto, err := f.formPublishedMapper.ToDto(formPublished)
	if err != nil {
		return nil, err
	}

	return publishedDto, nil
}
