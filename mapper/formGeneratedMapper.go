package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/repository"
)

type FormGeneratedMapper struct {
	formPublishedMapper     *FormPublishedMapper
	formPublishedRepository *repository.FormPublishedRepository
}

func NewFormGeneratedMapper() *FormGeneratedMapper {
	return &FormGeneratedMapper{
		formPublishedMapper:     NewFormPublishedMapper(),
		formPublishedRepository: repository.NewFormPublishedRepository(),
	}
}

func (f *FormGeneratedMapper) ToDto(formGenerated generated.FormGenerated) (*get.FormGeneratedDto, error) {
	formGeneratedDto := new(get.FormGeneratedDto)
	formGeneratedDto.Id = formGenerated.Id
	formGeneratedDto.Status = formGenerated.Status
	formPublished, err := f.formPublishedRepository.FindById(formGenerated.FormPublishedID)
	if err != nil {
		return nil, err
	}

	formGeneratedDto.FormPublished = *f.formPublishedMapper.ToBaseDto(*formPublished)
	formGeneratedDto.Sections = formGenerated.Sections
	return formGeneratedDto, nil
}
