package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model"
)

type SubjectMapper struct {
}

func NewSubjectMapper() *SubjectMapper {
	return &SubjectMapper{}
}

func (s *SubjectMapper) ToDto(subject model.Subject) *get.SubjectDto {
	subjectDto := new(get.SubjectDto)
	subjectDto.Id = subject.Id
	subjectDto.Name = subject.Name
	return subjectDto
}
