package service

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/repository"
)

type SubjectService struct {
	subjectRepository *repository.SubjectRepository
	subjectFactory    *factory.SubjectFactory
	subjectMapper     *mapper.SubjectMapper
}

func NewSubjectService() *SubjectService {
	return &SubjectService{
		subjectRepository: repository.NewSubjectRepository(),
		subjectFactory:    factory.NewSubjectFactory(),
		subjectMapper:     mapper.NewSubjectMapper(),
	}
}

func (s *SubjectService) Create(dto create.SubjectDto) (*get.SubjectDto, error) {
	subject := s.subjectFactory.Build(dto)
	if err := s.subjectRepository.Create(subject); err != nil {
		return nil, err
	}
	return s.subjectMapper.ToDto(subject), nil
}
