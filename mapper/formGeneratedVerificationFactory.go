package mapper

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/dto/verify"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/util"
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
		UserId:   formGenerated.UserId,
		Sections: verifiedSections,
	}, nil
}

func (f *FormGeneratedVerificationFactory) ExtractQuestionDtoMap(
	formPattern pattern.FormPattern,
) (map[uuid.UUID]get.IQuestionDto, error) {
	questionObjs := util.ExtractQuestionObjs(formPattern)

	questionDtoMap := make(map[uuid.UUID]get.IQuestionDto)
	for _, questionObj := range questionObjs {
		questionDto, err := f.questionMapper.toDto(questionObj)
		if err != nil {
			return nil, err
		}

		questionDtoMap[questionObj.GetId()] = questionDto
	}

	return questionDtoMap, nil
}
