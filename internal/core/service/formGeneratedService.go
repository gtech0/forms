package service

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/factory"
	"hedgehog-forms/internal/core/mapper"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/published"
	"hedgehog-forms/internal/core/processor"
	"hedgehog-forms/internal/core/repository"
	"hedgehog-forms/internal/core/util"
	"net/url"
	"strconv"
	"time"
)

type FormGeneratedService struct {
	submissionRepository             *repository.SubmissionRepository
	submissionFactory                *factory.SubmissionFactory
	solutionRepository               *repository.SolutionRepository
	solutionFactory                  *factory.SolutionFactory
	formPublishedRepository          *repository.FormPublishedRepository
	formGeneratedRepository          *repository.FormGeneratedRepository
	formGeneratedFactory             *factory.FormGeneratedFactory
	formGeneratedVerificationFactory *mapper.FormGeneratedVerificationFactory
	formGeneratedMapper              *mapper.FormGeneratedMapper
	formGeneratedProcessor           *processor.FormGeneratedProcessor
}

func NewFormGeneratedService() *FormGeneratedService {
	return &FormGeneratedService{
		submissionRepository:             repository.NewSubmissionRepository(),
		submissionFactory:                factory.NewSubmissionFactory(),
		solutionRepository:               repository.NewSolutionRepository(),
		solutionFactory:                  factory.NewSolutionFactory(),
		formPublishedRepository:          repository.NewFormPublishedRepository(),
		formGeneratedRepository:          repository.NewFormGeneratedRepository(),
		formGeneratedFactory:             factory.NewFormGeneratedFactory(),
		formGeneratedVerificationFactory: mapper.NewFormGeneratedVerificationFactory(),
		formGeneratedMapper:              mapper.NewFormGeneratedMapper(),
		formGeneratedProcessor:           processor.NewFormGeneratedProcessor(),
	}
}

func (f *FormGeneratedService) Create(
	publishedId,
	userId string,
) (*get.FormGeneratedDto, error) {
	parsedPublishedId, err := util.IdCheckAndParse(publishedId)
	if err != nil {
		return nil, err
	}

	parsedUserId, err := util.IdCheckAndParse(userId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(parsedPublishedId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedFactory.BuildForm(formPublished)
	if err != nil {
		return nil, err
	}

	//-should be optional?-
	questionIds := f.getAllQuestionIds(parsedUserId, formGenerated.ExtractQuestions(), formPublished.Id)
	formPublished.ExcludedQuestions = append(formPublished.ExcludedQuestions, questionIds...)
	//---------------------
	formPublished.FormsGenerated = append(formPublished.FormsGenerated, formGenerated)
	if err = f.formPublishedRepository.Save(formPublished); err != nil {
		return nil, err
	}

	return f.formGeneratedMapper.ToDto(formGenerated)
}

//func (f *FormGeneratedService) savePublished(
//	formPublished *published.FormPublished,
//	submission *generated.Submission,
//) error {
//	formGenerated := submission.FormGenerated
//
//	questions := formGenerated.ExtractQuestions()
//	questionIds := f.getAllQuestionIds(*submission.UserId, questions, formPublished.Id)
//
//	formPublished.ExcludedQuestions = append(formPublished.ExcludedQuestions, questionIds...)
//	formPublished.FormsGenerated = append(formPublished.FormsGenerated, formGenerated)
//
//	if err := f.formPublishedRepository.Save(formPublished); err != nil {
//		return err
//	}
//
//	return nil
//}

func (f *FormGeneratedService) Get(generatedId string) (*get.IntegrationGeneratedFormDto, error) {
	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
	if err != nil {
		return nil, err
	}

	return f.formGeneratedMapper.ToIntegrationDto(formGenerated)
}

//func (f *FormGeneratedService) GetMyForm(
//	publishedId,
//	userId string,
//) (*get.FormGeneratedDto, error) {
//	parsedPublishedId, err := util.IdCheckAndParse(publishedId)
//	if err != nil {
//		return nil, err
//	}
//
//	parsedUserId, err := util.IdCheckAndParse(userId)
//	if err != nil {
//		return nil, err
//	}
//
//	formPublished, err := f.formPublishedRepository.FindById(parsedPublishedId)
//	if err != nil {
//		return nil, err
//	}
//
//	solution, err := f.solutionRepository.FindByTaskIdAndUserId(parsedPublishedId, parsedUserId)
//	if err != nil {
//		return nil, err
//	}
//
//	submission := new(generated.Submission)
//	if solution.Id == uuid.Nil {
//		solution, err = f.solutionFactory.BuildFromPublished(formPublished, &parsedUserId)
//		if err = f.solutionRepository.Create(solution); err != nil {
//			return nil, err
//		}
//
//		if err = f.savePublished(formPublished, &solution.Submissions[0]); err != nil {
//			return nil, err
//		}
//	} else {
//		submission, err = f.findActiveGeneratedForm(solution.Submissions, formPublished, &parsedUserId)
//		if err != nil {
//			return nil, err
//		}
//
//		if len(solution.Submissions) < formPublished.MaxAttempts {
//			solution.Submissions = append(solution.Submissions, *submission)
//			if err = f.savePublished(formPublished, submission); err != nil {
//				return nil, err
//			}
//		} else {
//			return nil, errors.New("max attempts reached")
//		}
//
//		if err = f.solutionRepository.Save(solution); err != nil {
//			return nil, err
//		}
//	}
//
//	return f.formGeneratedMapper.ToDto(submission.FormGenerated)
//}

func (f *FormGeneratedService) getAllQuestionIds(
	userId uuid.UUID,
	questions []generated.IQuestion,
	formPublishedId uuid.UUID,
) []published.ExcludedQuestion {
	questionIds := make([]published.ExcludedQuestion, 0)
	for _, currQuestion := range questions {
		var excludedQuestion published.ExcludedQuestion
		excludedQuestion.UserId = userId
		excludedQuestion.QuestionId = currQuestion.GetId()
		excludedQuestion.FormPublishedId = formPublishedId
		questionIds = append(questionIds, excludedQuestion)
	}
	return questionIds
}

//func (f *FormGeneratedService) SaveAnswers(
//	generatedId string,
//	answers get.AnswerDto,
//) (*get.FormGeneratedDto, error) {
//	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
//	if err != nil {
//		return nil, err
//	}
//
//	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
//	if err != nil {
//		return nil, err
//	}
//
//	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedId)
//	if err != nil {
//		return nil, err
//	}
//
//	submission, err := f.submissionRepository.FindById(formGenerated.SubmissionId)
//	if err != nil {
//		return nil, err
//	}
//
//	submissions, err := f.submissionRepository.FindAttemptsByUserAndPublished(*submission.UserId, formGenerated.FormPublishedId)
//	if err != nil {
//		return nil, err
//	}
//
//	if err = f.checkTime(submission, formPublished.Duration); err != nil {
//		return nil, err
//	}
//
//	if err = f.checkDeadline(formPublished); err != nil {
//		return nil, err
//	}
//
//	if err = f.checkStatus(formGenerated); err != nil {
//		return nil, err
//	}
//
//	if err = f.checkAttempts(len(submissions), formPublished.MaxAttempts); err != nil {
//		return nil, err
//	}
//
//	if err = f.formGeneratedProcessor.SaveAnswers(formGenerated, answers); err != nil {
//		return nil, err
//	}
//
//	formGenerated.Status = generated.IN_PROGRESS
//	if err = f.formGeneratedRepository.Save(formGenerated); err != nil {
//		return nil, err
//	}
//
//	return f.formGeneratedMapper.ToDto(formGenerated)
//}
//
//func (f *FormGeneratedService) SubmitForm(
//	generatedId string,
//	answers get.AnswerDto,
//) (*get.MyGeneratedDto, error) {
//	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
//	if err != nil {
//		return nil, err
//	}
//
//	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
//	if err != nil {
//		return nil, err
//	}
//
//	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedId)
//	if err != nil {
//		return nil, err
//	}
//
//	submission, err := f.submissionRepository.FindById(formGenerated.SubmissionId)
//	if err != nil {
//		return nil, err
//	}
//
//	submissions, err := f.submissionRepository.FindAttemptsByUserAndPublished(*submission.UserId, formGenerated.FormPublishedId)
//	if err != nil {
//		return nil, err
//	}
//
//	if err = f.checkTime(submission, formPublished.Duration); err != nil {
//		return nil, err
//	}
//
//	if err = f.checkDeadline(formPublished); err != nil {
//		return nil, err
//	}
//
//	if err = f.checkStatus(formGenerated); err != nil {
//		return nil, err
//	}
//
//	if err = f.checkAttempts(len(submissions), formPublished.MaxAttempts); err != nil {
//		return nil, err
//	}
//
//	if err = f.formGeneratedProcessor.CalculatePoints(formGenerated, formPublished.FormPattern, answers); err != nil {
//		return nil, err
//	}
//
//	if err = f.formGeneratedProcessor.CalculateMark(formGenerated, formPublished.GetMarkConfigMap()); err != nil {
//		return nil, err
//	}
//
//	status := generated.COMPLETED
//	if formPublished.PostModeration {
//		status = generated.SUBMITTED
//	}
//
//	formGenerated.Status = status
//	if err = f.formGeneratedRepository.Save(formGenerated); err != nil {
//		return nil, err
//	}
//
//	submission.SubmitTime = util.Pointer(time.Now())
//	if err = f.submissionRepository.Save(submission); err != nil {
//		return nil, err
//	}
//
//	solution, err := f.solutionRepository.FindByTaskIdAndUserId(formGenerated.FormPublishedId, *submission.UserId)
//	if err != nil {
//		return nil, err
//	}
//
//	solution.Score = formGenerated.Points
//	if err = f.solutionRepository.Save(solution); err != nil {
//		return nil, err
//	}
//
//	return f.formGeneratedMapper.ToMyDto(formGenerated)
//}

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

func (f *FormGeneratedService) paginationConfig(
	minSize, maxSize int,
	query url.Values,
) (int, int, error) {
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		return 0, 0, err
	}

	if page <= 0 {
		page = 1
	}

	size, err := strconv.Atoi(query.Get("size"))
	if err != nil {
		return 0, 0, err
	}

	switch {
	case size > maxSize:
		size = maxSize
	case size <= 0:
		size = minSize
	}

	return page, size, nil
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

	page, size, err := f.paginationConfig(5, 50, query)
	if err != nil {
		return nil, err
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
		mySubmittedDto := f.formGeneratedMapper.ToSubmittedDto(formGenerated)
		mySubmittedDtos = append(mySubmittedDtos, *mySubmittedDto)
	}

	return &get.PaginationResponse[get.SubmittedDto]{
		Page:     page,
		Size:     size,
		Elements: mySubmittedDtos,
	}, nil
}

//func (f *FormGeneratedService) GetSubmittedForm(generatedId string) (*verify.FormGenerated, error) {
//	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
//	if err != nil {
//		return nil, err
//	}
//
//	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
//	if err != nil {
//		return nil, err
//	}
//
//	submission, err := f.submissionRepository.FindById(formGenerated.SubmissionId)
//	if err != nil {
//		return nil, err
//	}
//
//	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedId)
//	if err != nil {
//		return nil, err
//	}
//
//	verifiedForm, err := f.formGeneratedVerificationFactory.Build(formGenerated, formPublished.FormPattern, submission.UserId)
//	if err != nil {
//		return nil, err
//	}
//
//	return verifiedForm, nil
//}

func (f *FormGeneratedService) VerifyForm(generatedId string, checkDto create.CheckDto) (*get.FormGeneratedDto, error) {
	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
	if err != nil {
		return nil, err
	}

	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedId)
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

func (f *FormGeneratedService) ReturnForm(generatedId string) (*get.MyGeneratedDto, error) {
	parsedGeneratedId, err := util.IdCheckAndParse(generatedId)
	if err != nil {
		return nil, err
	}

	formGenerated, err := f.formGeneratedRepository.FindById(parsedGeneratedId)
	if err != nil {
		return nil, err
	}

	formGenerated.Status = generated.RETURNED
	if err = f.formGeneratedRepository.Create(formGenerated); err != nil {
		return nil, err
	}

	return f.formGeneratedMapper.ToMyDto(formGenerated)
}

//func (f *FormGeneratedService) checkTime(
//	submission *generated.Submission,
//	duration time.Duration,
//) error {
//	endTime := submission.StartTime.Add(duration)
//	if endTime.Before(time.Now()) {
//		submission.FormGenerated.Status = generated.COMPLETED
//		if err := f.submissionRepository.Save(submission); err != nil {
//			return err
//		}
//		return errs.New("Generated form duration is expired", 400)
//	}
//
//	return nil
//}

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

//func (f *FormGeneratedService) findActiveGeneratedForm(
//	forms []generated.Submission,
//	formPublished *published.FormPublished,
//	userId *uuid.UUID,
//) (*generated.Submission, error) {
//	for _, submission := range forms {
//		if submission.FormGenerated.Status != generated.COMPLETED {
//			return &submission, nil
//		}
//	}
//
//	submission, err := f.submissionFactory.Build(userId, formPublished)
//	if err != nil {
//		return nil, err
//	}
//
//	return submission, nil
//}
