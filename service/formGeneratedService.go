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

func (f *FormGeneratedService) GetMyForm(
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

	solution, err := f.solutionRepository.FindByTaskIdAndUserId(parsedPublishedId, parsedUserId)
	if err != nil {
		return nil, err
	}

	currAttempt := new(generated.FormGenerated)
	if solution.Id == uuid.Nil {
		currAttempt, err = f.buildAndCreate(formPublished, parsedUserId)
		if err != nil {
			return nil, err
		}

		solution = f.solutionFactory.BuildFromPublished(formPublished, &parsedUserId, currAttempt)
		if err = f.solutionRepository.Create(solution); err != nil {
			return nil, err
		}
	} else {
		attempts, err := f.formGeneratedRepository.FindAttemptsByUserAndPublished(parsedUserId, parsedPublishedId)
		if err != nil {
			return nil, err
		}

		//TODO: possibly more changes needed (?)
		currAttempt, _ = f.findActiveGeneratedForm(attempts)
		//if err != nil {
		//	return nil, err
		//}

		if currAttempt == nil && len(attempts) < formPublished.MaxAttempts {
			currAttempt, err = f.buildAndCreate(formPublished, parsedUserId)
			if err != nil {
				return nil, err
			}

			solution.Submissions = append(solution.Submissions, *currAttempt)
			solution.NumberAttempts = len(attempts)
		} else {
			solution.NumberAttempts = formPublished.MaxAttempts
		}

		if err = f.solutionRepository.Save(solution); err != nil {
			return nil, err
		}
	}

	return f.formGeneratedMapper.ToDto(currAttempt)
}

func (f *FormGeneratedService) buildAndCreate(
	formPublished *published.FormPublished,
	userId uuid.UUID,
) (*generated.FormGenerated, error) {
	formGenerated, err := f.formGeneratedFactory.BuildForm(formPublished, userId)
	if err != nil {
		return nil, err
	}

	questions := formGenerated.ExtractQuestionsFromGeneratedForm()
	questionIds := f.getAllQuestionIds(userId, questions, formPublished.Id)
	formPublished.ExcludedQuestions = append(formPublished.ExcludedQuestions, questionIds...)
	formPublished.FormsGenerated = append(formPublished.FormsGenerated, formGenerated)

	if err = f.formPublishedRepository.Save(formPublished); err != nil {
		return nil, err
	}

	return formGenerated, nil
}

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

func (f *FormGeneratedService) SaveAnswers(
	generatedId string,
	answers get.AnswerDto,
) (*get.FormGeneratedDto, error) {
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

	attempts, err := f.formGeneratedRepository.FindAttemptsByUserAndPublished(formGenerated.UserId, formGenerated.FormPublishedId)
	if err != nil {
		return nil, err
	}

	currForm, err := f.findActiveGeneratedForm(attempts)
	if err != nil {
		return nil, err
	}

	if err = f.checkTime(currForm, formPublished.Duration); err != nil {
		return nil, err
	}

	if err = f.checkDeadline(formPublished); err != nil {
		return nil, err
	}

	if err = f.checkStatus(formGenerated); err != nil {
		return nil, err
	}

	if err = f.checkAttempts(len(attempts), formPublished.MaxAttempts); err != nil {
		return nil, err
	}

	if err = f.formGeneratedProcessor.SaveAnswers(formGenerated, answers); err != nil {
		return nil, err
	}

	formGenerated.Status = generated.IN_PROGRESS
	if err = f.formGeneratedRepository.Save(formGenerated); err != nil {
		return nil, err
	}

	return f.formGeneratedMapper.ToDto(formGenerated)
}

//TODO: add solution logic

func (f *FormGeneratedService) SubmitForm(
	generatedId string,
	answers get.AnswerDto,
) (*get.MyGeneratedDto, error) {
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

	attempts, err := f.formGeneratedRepository.FindAttemptsByUserAndPublished(formGenerated.UserId, formGenerated.FormPublishedId)
	if err != nil {
		return nil, err
	}

	currForm, err := f.findActiveGeneratedForm(attempts)
	if err != nil {
		return nil, err
	}

	if err = f.checkTime(currForm, formPublished.Duration); err != nil {
		return nil, err
	}

	if err = f.checkDeadline(formPublished); err != nil {
		return nil, err
	}

	if err = f.checkStatus(formGenerated); err != nil {
		return nil, err
	}

	if err = f.checkAttempts(len(attempts), formPublished.MaxAttempts); err != nil {
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
	if err = f.formGeneratedRepository.Save(formGenerated); err != nil {
		return nil, err
	}

	solution, err := f.solutionRepository.FindByTaskIdAndUserId(formGenerated.FormPublishedId, formGenerated.UserId)
	if err != nil {
		return nil, err
	}

	solution.Score = formGenerated.Points
	solution.NumberAttempts = len(attempts)
	if err = f.solutionRepository.Save(solution); err != nil {
		return nil, err
	}

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
	//TODO: user group ids
	//for _, groupEntity := range formPublished.Teams {
	//	for _, userEntity := range groupEntity.Users {
	//		userIdsWithAccess = append(userIdsWithAccess, userEntity.Id)
	//	}
	//}

	for _, userEntity := range formPublished.Users {
		userIdsWithAccess = append(userIdsWithAccess, userEntity.UserId)
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

	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedId)
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

func (f *FormGeneratedService) checkTime(formGenerated *generated.FormGenerated, duration time.Duration) error {
	startTime := formGenerated.StartTime
	endTime := startTime.Add(duration)
	if endTime.Before(time.Now()) {
		formGenerated.IsCompleted = true
		if err := f.formGeneratedRepository.Save(formGenerated); err != nil {
			return err
		}
		return errs.New("Generated form duration is expired", 400)
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

func (f *FormGeneratedService) findActiveGeneratedForm(forms []*generated.FormGenerated) (*generated.FormGenerated, error) {
	for _, formGenerated := range forms {
		if formGenerated.IsCompleted == false {
			return formGenerated, nil
		}
	}
	return nil, errs.New("Generated form is completed", 400)
}
