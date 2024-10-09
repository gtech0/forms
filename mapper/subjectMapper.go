package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model"
)

type SubjectMapper struct {
}

func NewSubjectMapper() *SubjectMapper {
	return &SubjectMapper{}
}

func (s *SubjectMapper) toDto(subject model.Subject) get.SubjectDto {
	var subjectDto get.SubjectDto
	subjectDto.Id = subject.Id
	subjectDto.Name = subject.Name
	return subjectDto
}
