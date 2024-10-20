package service

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/database"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/factory"
	"hedgehog-forms/mapper"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"hedgehog-forms/repository"
	"slices"
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
		return nil, errs.New(
			fmt.Sprintf("incorrect attachment ids: %v", attachmentIds), 400,
		)
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
	if patternId == "" {
		return nil, errs.New("patternId is required", 400)
	}

	parsedPatternId, err := uuid.Parse(patternId)
	if err != nil {
		return nil, errs.New(err.Error(), 400)
	}

	formPattern, err := f.formPatternRepository.FindById(parsedPatternId)
	if err != nil {
		return nil, err
	}

	dto, err := f.formPatternMapper.ToDto(formPattern)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (f *FormPatternService) getFormId(id uuid.UUID) (uuid.UUID, error) {
	var formPattern pattern.FormPattern
	if err := database.DB.Model(&pattern.FormPattern{}).
		Where("id = ?", id).
		First(&formPattern).
		Error; err != nil {
		return uuid.Nil, errs.New(err.Error(), 500)
	}
	return formPattern.Id, nil
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
