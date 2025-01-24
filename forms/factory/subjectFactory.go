package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model"
)

type SubjectFactory struct{}

func NewSubjectFactory() *SubjectFactory {
	return &SubjectFactory{}
}

func (s *SubjectFactory) Build(subjectDto create.SubjectDto) model.Subject {
	var subject model.Subject
	subject.Id = uuid.New()
	subject.Name = subjectDto.Name
	return subject
}
