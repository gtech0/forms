package service

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/published"
	"hedgehog-forms/processor"
	"hedgehog-forms/repository"
	"hedgehog-forms/util"
	"net/url"
	"strconv"
	"time"
)

type FormGeneratedService struct {
	formPublishedRepository *repository.FormPublishedRepository
	formGeneratedRepository *repository.FormGeneratedRepository
	formGeneratedFactory    *factory.FormGeneratedFactory
	formGeneratedMapper     *mapper.FormGeneratedMapper
	formGeneratedProcessor  *processor.FormGeneratedProcessor
}

func NewFormGeneratedService() *FormGeneratedService {
	return &FormGeneratedService{
		formPublishedRepository: repository.NewFormPublishedRepository(),
		formGeneratedRepository: repository.NewFormGeneratedRepository(),
		formGeneratedFactory:    factory.NewFormGeneratedFactory(),
		formGeneratedMapper:     mapper.NewFormGeneratedMapper(),
		formGeneratedProcessor:  processor.NewFormGeneratedProcessor(),
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

	return f.formGeneratedMapper.ToDto(formGenerated)
}

func (f *FormGeneratedService) buildAndCreate(
	formPublished published.FormPublished,
	// userId uuid.UUID,
) (*generated.FormGenerated, error) {
	generatedForm, err := f.formGeneratedFactory.BuildForm(formPublished)
	if err != nil {
		return nil, err
	}

	if err = f.formGeneratedRepository.Save(generatedForm); err != nil {
		return nil, err
	}

	return generatedForm, nil
}

func (f *FormGeneratedService) SaveAnswers(
	formGeneratedId uuid.UUID,
	answers get.AnswerDto,
) (*get.FormGeneratedDto, error) {
	formGenerated, err := f.formGeneratedRepository.FindById(formGeneratedId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedID)
	if err != nil {
		return nil, err
	}

	//TODO: check user access
	if err = f.checkTime(formPublished); err != nil {
		return nil, err
	}

	if err = f.checkDeadline(formPublished); err != nil {
		return nil, err
	}

	if err = f.checkStatus(formGenerated); err != nil {
		return nil, err
	}

	if err = f.formGeneratedProcessor.MarkAnswers(formGenerated, answers); err != nil {
		return nil, err
	}
	formGenerated.Status = generated.IN_PROGRESS
	if err = f.formGeneratedRepository.Save(formGenerated); err != nil {
		return nil, err
	}
	return f.formGeneratedMapper.ToDto(formGenerated)
}

func (f *FormGeneratedService) SubmitForm(
	userId uuid.UUID,
	generatedFormId uuid.UUID,
	answers get.AnswerDto,
) (*get.MyGeneratedDto, error) {
	formGenerated, err := f.formGeneratedRepository.FindById(generatedFormId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedID)
	if err != nil {
		return nil, err
	}

	if err = f.checkTime(formPublished); err != nil {
		return nil, err
	}

	if err = f.checkDeadline(formPublished); err != nil {
		return nil, err
	}

	if err = f.checkStatus(formGenerated); err != nil {
		return nil, err
	}

	if err = f.formGeneratedProcessor.MarkAnswersAndCalculatePoints(formGenerated, formPublished.FormPattern, answers); err != nil {
		return nil, err
	}

	if err = f.formGeneratedProcessor.CalculateMark(formGenerated, formPublished.GetMarkConfigMap()); err != nil {
		return nil, err
	}

	status := generated.COMPLETED
	if formPublished.PostModeration {
		status = generated.SUBMITTED
	}

	formGenerated.Status = status
	formGenerated.SubmitTime = time.Now()
	if err = f.formGeneratedRepository.Save(formGenerated); err != nil {
		return nil, err
	}

	return f.formGeneratedMapper.ToMyDto(formGenerated)
}

func (f *FormGeneratedService) GetMyForms(
	//userId,
	subjectId uuid.UUID,
	query url.Values,
) (*get.PaginationResponse[get.MyGeneratedDto], error) {
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

	formsGenerated, err := f.formGeneratedRepository.FindBySubjectIdAndPaginate(subjectId, page, size)
	if err != nil {
		return nil, err
	}

	myGeneratedDtos := make([]get.MyGeneratedDto, 0)
	for _, formGenerated := range formsGenerated {
		myGeneratedDto, err := f.formGeneratedMapper.ToMyDto(&formGenerated)
		if err != nil {
			return nil, err
		}

		myGeneratedDtos = append(myGeneratedDtos, *myGeneratedDto)
	}

	return &get.PaginationResponse[get.MyGeneratedDto]{
		Page:     page,
		Size:     size,
		Elements: myGeneratedDtos,
	}, nil
}

func (f *FormGeneratedService) GetSubmittedTests(
	//userId,
	formPublishedId uuid.UUID,
	status generated.FormStatus,
	query url.Values,
) (*get.PaginationResponse[get.SubmittedDto], error) {
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

	formsGenerated, err := f.formGeneratedRepository.FindByPublishedIdAndStatusAndPaginate(
		formPublishedId,
		status,
		page,
		size,
	)
	if err != nil {
		return nil, err
	}

	mySubmittedDtos := make([]get.SubmittedDto, 0)
	for _, formGenerated := range formsGenerated {
		myGeneratedDto := f.formGeneratedMapper.ToSubmittedDto(formGenerated)
		mySubmittedDtos = append(mySubmittedDtos, *myGeneratedDto)
	}

	return &get.PaginationResponse[get.SubmittedDto]{
		Page:     page,
		Size:     size,
		Elements: mySubmittedDtos,
	}, nil
}

func (f *FormGeneratedService) VerifyForm(generatedFormId uuid.UUID, checkDto create.CheckDto) (*get.FormGeneratedDto, error) {
	formGenerated, err := f.formGeneratedRepository.FindById(generatedFormId)
	if err != nil {
		return nil, err
	}

	if err = f.checkStatusForVerification(formGenerated.Status, checkDto.Status); err != nil {
		return nil, err
	}

	f.formGeneratedProcessor.VerifyForm(formGenerated, checkDto)
	if err = f.formGeneratedRepository.Save(formGenerated); err != nil {
		return nil, err
	}

	return f.formGeneratedMapper.ToDto(formGenerated)
}

func (f *FormGeneratedService) checkTime(formPublished *published.FormPublished) error {
	currentTime := time.Now()
	formCreatedAt := formPublished.CreatedAt
	endTime := formCreatedAt.Add(formPublished.Duration)
	if endTime.Before(currentTime) {
		return errs.New("Form duration is expired", 400)
	}

	return nil
}

func (f *FormGeneratedService) checkDeadline(formPublished *published.FormPublished) error {
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

func (f *FormGeneratedService) checkStatusForVerification(old, new generated.FormStatus) error {
	if old != generated.COMPLETED && old != generated.SUBMITTED {
		return errs.New(fmt.Sprintf("Old status %s of generated test is not suitable for verification", old), 400)
	}

	if new != generated.COMPLETED && new != generated.SUBMITTED {
		return errs.New(fmt.Sprintf("New status %s of generated test is not suitable for verification", new), 400)
	}
	return nil
}
