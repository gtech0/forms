package service

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
	"hedgehog-forms/repository"
	"hedgehog-forms/util"
	"time"
)

type FormGeneratedService struct {
	formPublishedRepository *repository.FormPublishedRepository
	formGeneratedRepository *repository.FormGeneratedRepository
	formGeneratedFactory    *factory.FormGeneratedFactory
	formGeneratedMapper     *mapper.FormGeneratedMapper
}

func NewFormGeneratedService() *FormGeneratedService {
	return &FormGeneratedService{
		formPublishedRepository: repository.NewFormPublishedRepository(),
		formGeneratedRepository: repository.NewFormGeneratedRepository(),
		formGeneratedFactory:    factory.NewFormGeneratedFactory(),
		formGeneratedMapper:     mapper.NewFormGeneratedMapper(),
	}
}

func (f *FormGeneratedService) GetMyForm(publishedId string) (*get.FormGeneratedDto, error) {
	parsedPublishedId, err := util.IdCheckAndParse(publishedId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(parsedPublishedId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindByPublishedId(parsedPublishedId)
	if err != nil {
		return nil, err
	}

	if formGenerated.Id == uuid.Nil {
		formGenerated, err = f.buildAndCreate(*formPublished)
		if err != nil {
			return nil, err
		}
	}

	return f.formGeneratedMapper.ToDto(*formGenerated)
}

func (f *FormGeneratedService) buildAndCreate(
	formPublished published.FormPublished,
	// userId uuid.UUID,
) (*generated.FormGenerated, error) {
	generatedForm, err := f.formGeneratedFactory.BuildForm(formPublished)
	if err != nil {
		return nil, err
	}

	if err = f.formGeneratedRepository.Create(generatedForm); err != nil {
		return nil, err
	}

	return generatedForm, nil
}

func (f *FormGeneratedService) checkTime(formGenerated *generated.FormGenerated) error {
	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedID)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	formCreatedAt := formPublished.CreatedAt
	endTime := formCreatedAt.Add(formPublished.Duration)
	if endTime.Before(currentTime) {
		return errs.New("Form duration is expired", 400)
	}

	return nil
}

func (f *FormGeneratedService) checkDeadline(formGenerated *generated.FormGenerated) error {
	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedID)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	if formPublished.Deadline.Before(currentTime) {
		return errs.New("Form deadline is expired", 400)
	}

	return nil
}

func (f *FormGeneratedService) checkStatus(formGenerated *generated.FormGenerated) error {
	if formGenerated.Status != generated.NEW && formGenerated.Status != generated.IN_PROGRESS {
		return errs.New("Form is already submitted", 400)
	}

	return nil
}
