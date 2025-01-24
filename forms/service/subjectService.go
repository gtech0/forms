package service

import (
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/repository"
	"hedgehog-forms/util"
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

func (s *SubjectService) CreateSubject(dto create.SubjectDto) (*get.SubjectDto, error) {
	subject := s.subjectFactory.Build(dto)
	if err := s.subjectRepository.Save(subject); err != nil {
		return nil, err
	}
	return s.subjectMapper.ToDto(subject), nil
}

func (s *SubjectService) GetSubject(subjectId string) (*get.SubjectDto, error) {
	parsedSubjectId, err := util.IdCheckAndParse(subjectId)
	if err != nil {
		return nil, err
	}

	subject, err := s.subjectRepository.FindById(parsedSubjectId)
	if err != nil {
		return nil, err
	}

	return s.subjectMapper.ToDto(*subject), nil
}

func (s *SubjectService) GetSubjects(name string) ([]get.SubjectDto, error) {
	subjects, err := s.subjectRepository.FindByName(name)
	if err != nil {
		return nil, err
	}

	subjectDtos := make([]get.SubjectDto, 0)
	for _, subject := range subjects {
		subjectDto := s.subjectMapper.ToDto(subject)
		subjectDtos = append(subjectDtos, *subjectDto)
	}

	return subjectDtos, nil
}

func (s *SubjectService) UpdateSubject(subjectId string, name string) (*get.SubjectDto, error) {
	parsedSubjectId, err := util.IdCheckAndParse(subjectId)
	if err != nil {
		return nil, err
	}

	if err = s.subjectRepository.Update(parsedSubjectId, name); err != nil {
		return nil, err
	}

	return &get.SubjectDto{
		Id:   parsedSubjectId,
		Name: name,
	}, err
}

func (s *SubjectService) DeleteSubject(subjectId string) error {
	parsedSubjectId, err := util.IdCheckAndParse(subjectId)
	if err != nil {
		return err
	}

	if err = s.subjectRepository.Delete(parsedSubjectId); err != nil {
		return err
	}
	return nil
}
