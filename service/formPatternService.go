package service

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"hedgehog-forms/repository"
	"hedgehog-forms/util"
	"net/url"
	"slices"
	"strconv"
)

type FormPatternService struct {
	formPatternRepository *repository.FormPatternRepository
	formPatternFactory    *factory.FormPatternFactory
	formPatternMapper     *mapper.FormPatternMapper
	attachmentService     *AttachmentService
}

func NewFormPatternService() *FormPatternService {
	return &FormPatternService{
		formPatternRepository: repository.NewFormPatternRepository(),
		formPatternFactory:    factory.NewFormPatternFactory(),
		formPatternMapper:     mapper.NewFormPatternMapper(),
		attachmentService:     NewAttachmentService(),
	}
}

func (f *FormPatternService) CreatePattern(body create.FormPatternDto) (*get.FormPatternDto, error) {
	formPattern, err := f.formPatternFactory.BuildPattern(&body)
	if err != nil {
		return nil, err
	}

	attachmentIds, err := f.validatePatternAttachments(formPattern)
	if err != nil {
		return nil, err
	}

	if len(attachmentIds) > 0 {
		return nil, errs.New(fmt.Sprintf("incorrect attachment ids: %v", attachmentIds), 400)
	}

	if err = f.formPatternRepository.Create(formPattern); err != nil {
		return nil, err
	}

	dto, err := f.formPatternMapper.ToDto(formPattern)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (f *FormPatternService) GetForm(patternId string) (*get.FormPatternDto, error) {
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

func (f *FormPatternService) GetForms(query url.Values) (*get.PaginationResponse[get.FormPatternBaseDto], error) {
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

	patternBaseDtos := make([]get.FormPatternBaseDto, 0)
	for _, formPublished := range formsPublished {
		patternBaseDto := f.formPatternMapper.ToBaseDto(formPublished)
		patternBaseDtos = append(patternBaseDtos, *patternBaseDto)
	}

	return &get.PaginationResponse[get.FormPatternBaseDto]{
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

func (f *FormPatternService) extractQuestionsFromPattern(pattern *pattern.FormPattern) []question.IQuestion {
	questions := make([]question.IQuestion, 0)
	for _, currSection := range pattern.Sections {
		blocks := currSection.Blocks
		for _, iBlock := range blocks {
			switch assertedBlock := iBlock.(type) {
			case *block.DynamicBlock:
				questions = slices.Concat(questions, assertedBlock.Questions)
			case *block.StaticBlock:
				for _, variant := range assertedBlock.Variants {
					questions = slices.Concat(questions, variant.Questions)
				}
			}
		}
	}
	return questions
}

func (f *FormPatternService) validatePatternAttachments(pattern *pattern.FormPattern) ([]uuid.UUID, error) {
	questions := f.extractQuestionsFromPattern(pattern)
	attachmentIds := make([]uuid.UUID, 0)
	for _, iQuestion := range questions {
		for _, attachment := range iQuestion.GetAttachments() {
			attachmentIds = append(attachmentIds, attachment.Id)
		}
	}

	return f.attachmentService.validateAttachments(attachmentIds)
}
