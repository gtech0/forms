package service

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/dto/verify"
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
	formPublishedRepository          *repository.FormPublishedRepository
	formGeneratedRepository          *repository.FormGeneratedRepository
	formGeneratedFactory             *factory.FormGeneratedFactory
	formGeneratedVerificationFactory *mapper.FormGeneratedVerificationFactory
	formGeneratedMapper              *mapper.FormGeneratedMapper
	formGeneratedProcessor           *processor.FormGeneratedProcessor
}

func NewFormGeneratedService() *FormGeneratedService {
	return &FormGeneratedService{
		formPublishedRepository:          repository.NewFormPublishedRepository(),
		formGeneratedRepository:          repository.NewFormGeneratedRepository(),
		formGeneratedFactory:             factory.NewFormGeneratedFactory(),
		formGeneratedVerificationFactory: mapper.NewFormGeneratedVerificationFactory(),
		formGeneratedMapper:              mapper.NewFormGeneratedMapper(),
		formGeneratedProcessor:           processor.NewFormGeneratedProcessor(),
	}
}

//TODO: hide questions after retries

func (f *FormGeneratedService) GetMyForm(
	userId,
	publishedId string,
) (*get.FormGeneratedDto, error) {
	parsedPublishedId, err := util.IdCheckAndParse(publishedId)
	if err != nil {
		return nil, err
	}

	parsedUserId, err := util.IdCheckAndParse(userId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindByPublishedId(parsedPublishedId)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(parsedPublishedId)
	if err != nil {
		return nil, err
	}

	if formGenerated == nil {
		formGenerated, err = f.buildAndCreate(formPublished, parsedUserId)
		if err != nil {
			return nil, err
		}
	}

	if formGenerated.CurrentAttempts > 0 &&
		formGenerated.IsGenerated == false &&
		formGenerated.Status == generated.IN_PROGRESS {
		//TODO
		//questions := formGenerated.ExtractQuestionsFromGeneratedForm()

	}

	return f.formGeneratedMapper.ToDto(formGenerated)
}

func (f *FormGeneratedService) buildAndCreate(
	formPublished *published.FormPublished,
	userId uuid.UUID,
) (*generated.FormGenerated, error) {
	generatedForm, err := f.formGeneratedFactory.BuildForm(formPublished, userId)
	if err != nil {
		return nil, err
	}

	if err = f.formGeneratedRepository.Create(generatedForm); err != nil {
		return nil, err
	}

	return generatedForm, nil
}

func (f *FormGeneratedService) SaveAnswers(
	userId,
	generatedId string,
	answers get.AnswerDto,
) (*get.FormGeneratedDto, error) {
	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
	if err != nil {
		return nil, err
	}

	parsedUserId, err := util.IdCheckAndParse(userId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedID)
	if err != nil {
		return nil, err
	}

	if err = f.checkAccess(formPublished, parsedUserId); err != nil {
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

	if err = f.checkAttempts(formGenerated.CurrentAttempts, formPublished.MaxAttempts); err != nil {
		return nil, err
	}

	if err = f.formGeneratedProcessor.SaveAnswers(formGenerated, answers); err != nil {
		return nil, err
	}

	formGenerated.Status = generated.IN_PROGRESS
	if err = f.formGeneratedRepository.Create(formGenerated); err != nil {
		return nil, err
	}

	return f.formGeneratedMapper.ToDto(formGenerated)
}

func (f *FormGeneratedService) SubmitForm(
	userId,
	generatedId string,
	answers get.AnswerDto,
) (*get.MyGeneratedDto, error) {
	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
	if err != nil {
		return nil, err
	}

	parsedUserId, err := util.IdCheckAndParse(userId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedID)
	if err != nil {
		return nil, err
	}

	if err = f.checkAccess(formPublished, parsedUserId); err != nil {
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

	if err = f.checkAttempts(formGenerated.CurrentAttempts, formPublished.MaxAttempts); err != nil {
		return nil, err
	}

	if err = f.formGeneratedProcessor.CalculatePoints(formGenerated, formPublished.FormPattern, answers); err != nil {
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
	if err = f.formGeneratedRepository.Create(formGenerated); err != nil {
		return nil, err
	}

	formGenerated.CurrentAttempts += 1
	return f.formGeneratedMapper.ToMyDto(formGenerated)
}

func (f *FormGeneratedService) UnSubmitForm(generatedId string) (*get.MyGeneratedDto, error) {
	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
	if err != nil {
		return nil, err
	}

	formGenerated.Status = generated.IN_PROGRESS
	if err = f.formGeneratedRepository.Create(formGenerated); err != nil {
		return nil, err
	}

	return f.formGeneratedMapper.ToMyDto(formGenerated)
}

func (f *FormGeneratedService) GetMyForms(
	userId,
	subjectId string,
	query url.Values,
) (*get.PaginationResponse[get.MyGeneratedDto], error) {
	parsedSubjectId, err := util.IdCheckAndParse(subjectId)
	if err != nil {
		return nil, err
	}

	parsedUserId, err := util.IdCheckAndParse(userId)
	if err != nil {
		return nil, err
	}

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

	formsGenerated, err := f.formGeneratedRepository.FindBySubjectIdAndPaginate(parsedUserId, parsedSubjectId, page, size)
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

func (f *FormGeneratedService) GetSubmittedForms(
	userId,
	publishedId string,
	query url.Values,
) (*get.PaginationResponse[get.SubmittedDto], error) {
	parsedPublishedId, err := util.IdCheckAndParse(publishedId)
	if err != nil {
		return nil, err
	}

	parsedUserId, err := util.IdCheckAndParse(userId)
	if err != nil {
		return nil, err
	}

	status := query.Get("status")
	formStatus := generated.CheckStatusAndGet(status)

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
		parsedUserId,
		parsedPublishedId,
		formStatus,
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

func (f *FormGeneratedService) GetUsersWithUnsubmittedForm(publishedId string) ([]uuid.UUID, error) {
	parsedPublishedId, err := util.IdCheckAndParse(publishedId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(parsedPublishedId)
	if err != nil {
		return nil, err
	}

	userIdsWithAccess := make([]uuid.UUID, 0)
	for _, groupEntity := range formPublished.Groups {
		for _, userEntity := range groupEntity.Users {
			userIdsWithAccess = append(userIdsWithAccess, userEntity.Id)
		}
	}

	for _, userEntity := range formPublished.Users {
		userIdsWithAccess = append(userIdsWithAccess, userEntity.Id)
	}

	userIdsWithGeneratedForm := make([]uuid.UUID, 0)
	for _, formGenerated := range formPublished.FormsGenerated {
		userIdsWithGeneratedForm = append(userIdsWithGeneratedForm, formGenerated.UserId)
	}

	return util.Difference(userIdsWithAccess, userIdsWithGeneratedForm), nil
}

func (f *FormGeneratedService) GetSubmittedForm(generatedId string) (*verify.FormGenerated, error) {
	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedID)
	if err != nil {
		return nil, err
	}

	verifiedForm, err := f.formGeneratedVerificationFactory.Build(formGenerated, formPublished.FormPattern)
	if err != nil {
		return nil, err
	}

	return verifiedForm, nil
}

func (f *FormGeneratedService) VerifyForm(generatedId string, checkDto create.CheckDto) (*get.FormGeneratedDto, error) {
	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedID)
	if err != nil {
		return nil, err
	}

	if err = f.checkStatusForVerification(formGenerated.Status, checkDto.Status); err != nil {
		return nil, err
	}

	if err = f.formGeneratedProcessor.ReapplyPoints(formGenerated, checkDto); err != nil {
		return nil, err
	}

	if err = f.formGeneratedProcessor.CalculateMark(formGenerated, formPublished.GetMarkConfigMap()); err != nil {
		return nil, err
	}

	if err = f.formGeneratedRepository.Create(formGenerated); err != nil {
		return nil, err
	}

	return f.formGeneratedMapper.ToDto(formGenerated)
}

func (f *FormGeneratedService) checkAccess(formPublished *published.FormPublished, userId uuid.UUID) error {
	formHasUser := false
	for _, userEntity := range formPublished.Users {
		if userEntity.Id == userId {
			formHasUser = true
		}
	}

	groupsHasUser := false
	for _, groupEntity := range formPublished.Groups {
		for _, userEntity := range groupEntity.Users {
			if userEntity.Id == userId {
				groupsHasUser = true
			}
		}
	}

	if formHasUser || groupsHasUser {
		return nil
	}

	return errs.New(fmt.Sprintf("You don't have access to published test %v", formPublished.Id), 403)
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

func (f *FormGeneratedService) checkAttempts(current, max int) error {
	if current > max {
		return errs.New(fmt.Sprintf("Attempt limit exceeded"), 400)
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
