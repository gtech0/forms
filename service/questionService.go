package service

import (
	"encoding/json"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/repository"
	"hedgehog-forms/util"
)

type QuestionService struct {
	questionRepository           *repository.QuestionRepository
	subjectRepository            *repository.SubjectRepository
	questionFactory              *factory.QuestionFactory
	subjectMapper                *mapper.SubjectMapper
	commonFieldQuestionDtoMapper *mapper.CommonFieldQuestionDtoMapper
	attachmentService            *AttachmentService
}

func NewQuestionService() *QuestionService {
	return &QuestionService{
		questionRepository:           repository.NewQuestionRepository(),
		subjectRepository:            repository.NewSubjectRepository(),
		questionFactory:              factory.NewQuestionFactory(),
		subjectMapper:                mapper.NewSubjectMapper(),
		commonFieldQuestionDtoMapper: mapper.NewCommonFieldQuestionDtoMapper(),
		attachmentService:            NewAttachmentService(),
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

	iQuestion, err := q.questionFactory.BuildQuestionFromDto(questionDto)
	if err != nil {
		return nil, err
	}

	if err = q.attachmentService.ValidateAttachments(iQuestion); err != nil {
		return nil, err
	}

	iQuestion.SetSubject(*subject)
	iQuestion.SetIsQuestionFromBank(true)

	if err = q.questionRepository.Create(iQuestion); err != nil {
		return nil, err
	}

	getQuestionDto := new(get.QuestionDto)
	q.commonFieldQuestionDtoMapper.CommonFieldsToDto(iQuestion, getQuestionDto)
	return getQuestionDto, nil
}
