package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/dto/verify"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern"
)

type FormGeneratedVerificationFactory struct {
	questionMapper                      *QuestionMapper
	sectionGeneratedVerificationFactory *SectionGeneratedVerificationFactory
}

func NewFormGeneratedVerificationFactory() *FormGeneratedVerificationFactory {
	return &FormGeneratedVerificationFactory{
		questionMapper:                      NewQuestionMapper(),
		sectionGeneratedVerificationFactory: NewSectionGeneratedVerificationFactory(),
	}
}

func (f *FormGeneratedVerificationFactory) Build(
	formGenerated *generated.FormGenerated,
	formPattern pattern.FormPattern,
) (*verify.FormGenerated, error) {
	questionsWithCorrectAnswers, err := f.ExtractQuestionDtoMap(formPattern)
	if err != nil {
		return nil, err
	}

	verifiedSections, err := f.sectionGeneratedVerificationFactory.build(formGenerated.Sections, questionsWithCorrectAnswers)
	if err != nil {
		return nil, err
	}

	return &verify.FormGenerated{
		Id:       formGenerated.Id,
		Status:   formGenerated.Status,
		UserId:   formGenerated.Submission.UserId,
		Sections: verifiedSections,
	}, nil
}

func (f *FormGeneratedVerificationFactory) ExtractQuestionDtoMap(
	formPattern pattern.FormPattern,
) (map[uuid.UUID]get.IQuestionDto, error) {
	questionEntities := formPattern.ExtractQuestionEntities()

	questionDtoMap := make(map[uuid.UUID]get.IQuestionDto)
	for _, questionEntity := range questionEntities {
		questionDto, err := f.questionMapper.ToDto(questionEntity)
		if err != nil {
			return nil, err
		}

		questionDtoMap[questionEntity.Id] = questionDto
	}

	return questionDtoMap, nil
}
