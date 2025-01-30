package service

import (
	"hedgehog-forms/internal/core/dto/create"
	get2 "hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/factory"
	"hedgehog-forms/internal/core/mapper"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/processor"
	repository2 "hedgehog-forms/internal/core/repository"
	"hedgehog-forms/internal/core/util"
	"maps"
	"net/url"
	"strconv"
)

type FormPublishedService struct {
	formPublishedRepository *repository2.FormPublishedRepository
	formPatternService      *FormPatternService
	formPublishedMapper     *mapper.FormPublishedMapper
	formPublishedFactory    *factory.FormPublishedFactory
	formGeneratedProcessor  *processor.FormGeneratedProcessor
	formGeneratedRepository *repository2.FormGeneratedRepository
}

func NewFormPublishedService() *FormPublishedService {
	return &FormPublishedService{
		formPublishedRepository: repository2.NewFormPublishedRepository(),
		formPatternService:      NewFormPatternService(),
		formPublishedMapper:     mapper.NewFormPublishedMapper(),
		formPublishedFactory:    factory.NewFormPublishedFactory(),
		formGeneratedProcessor:  processor.NewFormGeneratedProcessor(),
		formGeneratedRepository: repository2.NewFormGeneratedRepository(),
	}
}

func (f *FormPublishedService) PublishForm(publishDto create.FormPublishDto) (*get2.FormPublishedBaseDto, error) {
	if err := f.formPatternService.doesFormExist(publishDto.FormPatternId); err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedFactory.Build(publishDto)
	if err != nil {
		return nil, err
	}

	if err = f.formPublishedRepository.Create(formPublished); err != nil {
		return nil, err
	}

	return f.formPublishedMapper.ToBaseDto(formPublished), nil
}

func (f *FormPublishedService) GetForm(publishedId string) (*get2.FormPublishedDto, error) {
	id, err := util.IdCheckAndParse(publishedId)
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

func (f *FormPublishedService) GetForms(query url.Values) (*get2.PaginationResponse[get2.FormPublishedBaseDto], error) {
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

	formsPublished, err := f.formPublishedRepository.FindByNameAndPaginate(name, page, size)
	if err != nil {
		return nil, err
	}

	publishedDtos := make([]get2.FormPublishedBaseDto, 0)
	for _, formPublished := range formsPublished {
		publishedDto := f.formPublishedMapper.ToBaseDto(&formPublished)
		publishedDtos = append(publishedDtos, *publishedDto)
	}

	return &get2.PaginationResponse[get2.FormPublishedBaseDto]{
		Page:     page,
		Size:     size,
		Elements: publishedDtos,
	}, nil
}

func (f *FormPublishedService) UpdateForm(
	publishedId string,
	formPublishedDto create.UpdateFormPublishedDto,
) (*get2.FormPublishedBaseDto, error) {
	parsedPublishedId, err := util.IdCheckAndParse(publishedId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(parsedPublishedId)
	if err != nil {
		return nil, err
	}

	if err = f.formPublishedFactory.Update(formPublished, formPublishedDto); err != nil {
		return nil, err
	}

	if !maps.Equal(formPublished.GetMarkConfigMap(), formPublishedDto.MarkConfiguration) {
		if err = f.recalculateMarks(formPublished.FormsGenerated, formPublishedDto.MarkConfiguration); err != nil {
			return nil, err
		}
	}

	if err = f.formPublishedRepository.Save(formPublished); err != nil {
		return nil, err
	}

	return f.formPublishedMapper.ToBaseDto(formPublished), nil
}

func (f *FormPublishedService) recalculateMarks(formsGenerated []*generated.FormGenerated, marks map[int]int) error {
	for _, formGenerated := range formsGenerated {
		if err := f.formGeneratedProcessor.CalculateMark(formGenerated, marks); err != nil {
			return err
		}

		if err := f.formGeneratedRepository.Create(formGenerated); err != nil {
			return err
		}
	}
	return nil
}

func (f *FormPublishedService) DeleteForm(publishedId string) error {
	parsedPublishedId, err := util.IdCheckAndParse(publishedId)
	if err != nil {
		return err
	}

	if err = f.formPublishedRepository.DeleteById(parsedPublishedId); err != nil {
		return err
	}

	return nil
}
