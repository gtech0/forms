package service

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/factory"
	"hedgehog-forms/internal/core/mapper"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
	"hedgehog-forms/internal/core/repository"
	"hedgehog-forms/internal/core/util"
	"net/url"
	"strconv"
	"strings"
)

type QuestionService struct {
	questionRepository           *repository.QuestionRepository
	questionMapper               *mapper.QuestionMapper
	subjectRepository            *repository.SubjectRepository
	questionFactory              *factory.QuestionFactory
	subjectMapper                *mapper.SubjectMapper
	commonFieldQuestionDtoMapper *mapper.CommonFieldQuestionDtoMapper
	fileService                  *FileService
}

func NewQuestionService() *QuestionService {
	return &QuestionService{
		questionRepository:           repository.NewQuestionRepository(),
		questionMapper:               mapper.NewQuestionMapper(),
		subjectRepository:            repository.NewSubjectRepository(),
		questionFactory:              factory.NewQuestionFactory(),
		subjectMapper:                mapper.NewSubjectMapper(),
		commonFieldQuestionDtoMapper: mapper.NewCommonFieldQuestionDtoMapper(),
		fileService:                  NewFileService(),
	}
}

func (q *QuestionService) CreateQuestion(subjectId string, rawQuestionDto json.RawMessage) (*get.QuestionDto, error) {
	parsedSubjectId, err := util.IdCheckAndParse(subjectId)
	if err != nil {
		return nil, err
	}

	subject, err := q.subjectRepository.FindById(parsedSubjectId)
	if err != nil {
		return nil, err
	}

	questionDto, err := create.CommonQuestionDtoUnmarshal(rawQuestionDto)
	if err != nil {
		return nil, err
	}

	questionEntity, err := q.questionFactory.BuildQuestionFromDto(questionDto)
	if err != nil {
		return nil, err
	}

	if err = q.fileService.ValidateFiles(questionEntity); err != nil {
		return nil, err
	}

	questionEntity.Subject = *subject
	questionEntity.IsQuestionFromBank = true

	if err = q.questionRepository.Create(questionEntity); err != nil {
		return nil, err
	}

	getQuestionDto := new(get.QuestionDto)
	q.commonFieldQuestionDtoMapper.CommonFieldsToDto(questionEntity, getQuestionDto)
	return getQuestionDto, nil
}

func (q *QuestionService) GetQuestion(questionId string) (*get.QuestionDto, error) {
	parsedQuestionId, err := util.IdCheckAndParse(questionId)
	if err != nil {
		return nil, err
	}

	questionEntity, err := q.questionRepository.FindById(parsedQuestionId)
	if err != nil {
		return nil, err
	}

	getQuestionDto := new(get.QuestionDto)
	q.commonFieldQuestionDtoMapper.CommonFieldsToDto(questionEntity, getQuestionDto)
	return getQuestionDto, nil
}

//    public PaginationResponse<QuestionDto> getQuestions(UUID subjectId,
//                                                        String name,
//                                                        List<QuestionType> types,
//                                                        int page,
//                                                        int size
//    ) {
//        var specification = Specification.<QuestionEntity>where(null)
//                .and(hasSubject(subjectId))
//                .and(hasNameLike(name))
//                .and(hasTypeIn(types));
//
//        var pageable = PageRequest.of(page, size, Sort.Direction.ASC, QuestionEntity_.NAME);
//
//        var questions = questionRepository.findAll(specification, pageable).getContent()
//                .stream()
//                .map(questionMapper::toDto)
//                .toList();
//
//        return new PaginationResponse<>(page, size, questions);
//    }

func (q *QuestionService) GetQuestions(
	query url.Values,
) (*get.PaginationResponse[get.QuestionDto], error) {
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

	clauses := make([]clause.Expression, 0)
	subjectId := query.Get("subjectId")
	if subjectId != "" {
		parsedSubjectId, err := uuid.Parse(subjectId)
		if err != nil {
			return nil, errs.New(err.Error(), 400)
		}
		clauses = append(clauses, clause.Eq{Column: "subject_id", Value: parsedSubjectId})
	}

	typesSlice := strings.Split(query.Get("types"), ",")
	types := make([]question.QuestionType, 0)
	for _, questionType := range typesSlice {
		types = append(types, question.QuestionType(questionType))
	}

	questions, err := q.questionRepository.FindByParamsAndPaginate(clauses, name, page, size, types)
	if err != nil {
		return nil, err
	}

	questionDtos := make([]get.QuestionDto, 0)
	for _, questionEntity := range questions {
		questionDto, err := q.questionMapper.ToDto(&questionEntity)
		if err != nil {
			return nil, err
		}

		q.commonFieldQuestionDtoMapper.CommonFieldsToDto(&questionEntity, questionDto)
	}

	return &get.PaginationResponse[get.QuestionDto]{
		Page:     page,
		Size:     size,
		Elements: questionDtos,
	}, nil
}

func (q *QuestionService) DeleteQuestion(questionId string) error {
	parsedQuestionId, err := util.IdCheckAndParse(questionId)
	if err != nil {
		return err
	}

	questionEntity, err := q.questionRepository.FindById(parsedQuestionId)
	if err != nil {
		return err
	}

	if !questionEntity.IsQuestionFromBank {
		return errs.New(fmt.Sprintf("question %v is not from bank", questionId), 400)
	}

	return q.questionRepository.DeleteById(parsedQuestionId)
}
