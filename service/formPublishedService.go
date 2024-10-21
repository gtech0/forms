package service

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/repository"
	"hedgehog-forms/util"
	"net/url"
	"strconv"
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
	if err := f.formPatternService.doesFormExist(publishDto.FormPatternId); err != nil {
		return nil, err
	}

	formPublished := f.formPublishedFactory.Build(publishDto)
	if err := f.formPublishedRepository.Create(&formPublished); err != nil {
		return nil, err
	}

	return f.formPublishedMapper.ToBaseDto(formPublished), nil
}

func (f *FormPublishedService) GetForm(formId string) (*get.FormPublishedDto, error) {
	id, err := util.IdCheckAndParse(formId)
	if err != nil {
		return nil, err
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

func (f *FormPublishedService) GetForms(query url.Values) (*get.PaginationResponse[get.FormPublishedBaseDto], error) {
	name := query.Get("name")
	page, _ := strconv.Atoi(query.Get("page"))
	if page <= 0 {
		page = 1
	}

	size, _ := strconv.Atoi(query.Get("size"))
	switch {
	case size > 50:
		size = 50
	case size <= 0:
		size = 5
	}

	formsPublished, err := f.formPublishedRepository.FindAndPaginate(name, page, size)
	if err != nil {
		return nil, err
	}

	publishedDtos := make([]get.FormPublishedBaseDto, 0)
	for _, formPublished := range formsPublished {
		publishedDto := f.formPublishedMapper.ToBaseDto(formPublished)
		publishedDtos = append(publishedDtos, *publishedDto)
	}

	return &get.PaginationResponse[get.FormPublishedBaseDto]{
		Page:     page,
		Size:     size,
		Elements: publishedDtos,
	}, nil
}
