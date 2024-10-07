package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form"
)

type SubjectMapper struct {
}

func NewSubjectMapper() *SubjectMapper {
	return &SubjectMapper{}
}

func (s *SubjectMapper) toDto(subject form.Subject) get.SubjectDto {
	var subjectDto get.SubjectDto
	subjectDto.Id = subject.Id
	subjectDto.Name = subject.Name
	return subjectDto
}
