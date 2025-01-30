package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hedgehog-forms/internal/core/dto/create"
	get2 "hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/factory"
	"hedgehog-forms/internal/core/mapper"
	"hedgehog-forms/internal/core/model/form/pattern"
	"hedgehog-forms/internal/core/model/form/pattern/section/block"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
	"hedgehog-forms/internal/core/repository"
	"hedgehog-forms/internal/core/util"
	"net/url"
	"slices"
	"strconv"
)

type FormPatternService struct {
	formPatternRepository *repository.FormPatternRepository
	formPatternFactory    *factory.FormPatternFactory
	formPatternMapper     *mapper.FormPatternMapper
	fileService           *FileService
}

func NewFormPatternService() *FormPatternService {
	return &FormPatternService{
		formPatternRepository: repository.NewFormPatternRepository(),
		formPatternFactory:    factory.NewFormPatternFactory(),
		formPatternMapper:     mapper.NewFormPatternMapper(),
		fileService:           NewFileService(),
	}
}

func (f *FormPatternService) CreatePattern(body create.FormPatternDto) (*get2.FormPatternDto, error) {
	formPattern, err := f.formPatternFactory.BuildPattern(&body)
	if err != nil {
		return nil, err
	}

	if err = f.validatePatternAttachments(formPattern); err != nil {
		return nil, err
	}

	if err = f.formPatternRepository.Create(formPattern); err != nil {
		return nil, err
	}

	formPatternEntity, err := f.formPatternRepository.FindById(formPattern.Id)
	if err != nil {
		return nil, err
	}

	dto, err := f.formPatternMapper.ToDto(formPatternEntity)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (f *FormPatternService) GetForm(patternId string) (*get2.FormPatternDto, error) {
	id, err := util.IdCheckAndParse(patternId)
	if err != nil {
		return nil, err
	}

	formPattern, err := f.formPatternRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	dto, err := f.formPatternMapper.ToDto(formPattern)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (f *FormPatternService) GetForms(query url.Values) (*get2.PaginationResponse[get2.FormPatternBaseDto], error) {
	name := query.Get("name")
	page, _ := strconv.Atoi(query.Get("page"))
	if page <= 0 {
		page = 1
	}

	size, _ := strconv.Atoi(query.Get("size"))
	switch {
	case size > 50:
		size = 50
	case size <= 0:
		size = 5
	}

	clauses := make([]clause.Expression, 0)
	subjectId := query.Get("subjectId")
	if subjectId != "" {
		parsedSubjectId, err := uuid.Parse(subjectId)
		if err != nil {
			return nil, errs.New(err.Error(), 400)
		}
		clauses = append(clauses, clause.Eq{Column: "subject_id", Value: parsedSubjectId})
	}

	formsPublished, err := f.formPatternRepository.FindAndPaginate(name, clauses, page, size)
	if err != nil {
		return nil, err
	}

	patternBaseDtos := make([]get2.FormPatternBaseDto, 0)
	for _, formPublished := range formsPublished {
		patternBaseDto := f.formPatternMapper.ToBaseDto(formPublished)
		patternBaseDtos = append(patternBaseDtos, *patternBaseDto)
	}

	return &get2.PaginationResponse[get2.FormPatternBaseDto]{
		Page:     page,
		Size:     size,
		Elements: patternBaseDtos,
	}, nil
}

func (f *FormPatternService) doesFormExist(id uuid.UUID) error {
	if _, err := f.formPatternRepository.FindById(id); err != nil {
		return err
	}

	return nil
}

func (f *FormPatternService) extractQuestionsFromPattern(pattern *pattern.FormPattern) []*question.Question {
	questions := make([]*question.Question, 0)
	for _, currSection := range pattern.Sections {
		blocks := currSection.Blocks
		for _, iBlock := range blocks {
			switch iBlock.Type {
			case block.DYNAMIC:
				questions = slices.Concat(questions, iBlock.DynamicBlock.Questions)
			case block.STATIC:
				for _, variant := range iBlock.StaticBlock.Variants {
					questions = slices.Concat(questions, variant.Questions)
				}
			}
		}
	}
	return questions
}

func (f *FormPatternService) validatePatternAttachments(pattern *pattern.FormPattern) error {
	questions := f.extractQuestionsFromPattern(pattern)
	return f.fileService.ValidateFiles(questions...)
}
