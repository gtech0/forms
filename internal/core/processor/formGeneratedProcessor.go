package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/common"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern"
	"hedgehog-forms/internal/core/model/form/pattern/section/block/question"
	"hedgehog-forms/internal/core/util"
	"maps"
	"slices"
)

type FormGeneratedProcessor struct {
	matchingProcessor       *MatchingProcessor
	singleChoiceProcessor   *SingleChoiceProcessor
	multipleChoiceProcessor *MultipleChoiceProcessor
	textInputProcessor      *TextInputProcessor
}

func NewFormGeneratedProcessor() *FormGeneratedProcessor {
	return &FormGeneratedProcessor{
		matchingProcessor:       NewMatchingProcessor(),
		singleChoiceProcessor:   NewSingleChoiceProcessor(),
		multipleChoiceProcessor: NewMultipleChoiceProcessor(),
		textInputProcessor:      NewTextInputProcessor(),
	}
}

func (f *FormGeneratedProcessor) SaveAnswers(formGenerated *generated.FormGenerated, answers get.AnswerDto) error {
	questions := formGenerated.ExtractQuestions()
	for _, iQuestion := range questions {
		if err := f.checkQuestion(iQuestion.GetId(), iQuestion.GetType(), answers); err != nil {
			return err
		}

		if err := f.saveAnswer(iQuestion, answers); err != nil {
			return err
		}
	}

	return nil
}

func (f *FormGeneratedProcessor) CalculatePoints(
	formGenerated *generated.FormGenerated,
	formPattern pattern.FormPattern,
	answers get.AnswerDto,
) error {
	questions := formGenerated.ExtractQuestions()
	for _, iQuestion := range questions {
		if err := f.checkQuestion(iQuestion.GetId(), iQuestion.GetType(), answers); err != nil {
			return err
		}
	}

	questionEntities := formPattern.ExtractQuestionEntities()
	for _, questionEntity := range questionEntities {
		if err := f.checkQuestion(questionEntity.Id, questionEntity.Type, answers); err != nil {
			return err
		}
	}

	points := 0
	for _, iQuestion := range questions {
		for _, questionEntity := range questionEntities {
			if iQuestion.GetId() == questionEntity.Id {
				calculatedPoints, err := f.markAnswerAndCalculatePoints(iQuestion, questionEntity, answers)
				if err != nil {
					return err
				}
				points += calculatedPoints
			}
		}
	}

	formGenerated.Points = points
	return nil
}

func (f *FormGeneratedProcessor) ReapplyPoints(formGenerated *generated.FormGenerated, checkDto create.CheckDto) error {
	questions := formGenerated.ExtractQuestions()
	if checkDto.Points != nil {
		for questionId, points := range checkDto.Points {
			var questionGenerated generated.IQuestion
			for _, iQuestion := range questions {
				if iQuestion.GetId() == questionId {
					questionGenerated = iQuestion
					break
				}
			}

			if questionGenerated == nil {
				return errs.New(fmt.Sprintf("question with id %v not found", questionId), 500)
			}

			f.applyNewPoints(formGenerated, questionGenerated, points)
		}
	}

	formGenerated.Status = checkDto.Status
	return nil
}

func (f *FormGeneratedProcessor) applyNewPoints(formGenerated *generated.FormGenerated,
	questionGenerated generated.IQuestion,
	newPoints int,
) {
	difference := questionGenerated.GetPoints() - newPoints
	questionGenerated.SetPoints(newPoints)
	formGenerated.Points = formGenerated.Points - difference
}

func (f *FormGeneratedProcessor) CalculateMark(
	formGenerated *generated.FormGenerated,
	marks common.MarkConfiguration,
) error {
	pointsForForm := formGenerated.Points
	var requiredPointsForMark int
	for _, points := range marks {
		if points >= requiredPointsForMark && points <= pointsForForm {
			requiredPointsForMark = points
		}
	}

	mark, ok := util.FindKeyByValue(marks, requiredPointsForMark)
	if !ok && mark != 0 {
		return errs.New(fmt.Sprintf("mark for %d points is not found", requiredPointsForMark), 500)
	}

	formGenerated.Mark = mark
	return nil
}

func (f *FormGeneratedProcessor) saveAnswer(iQuestion generated.IQuestion, answers get.AnswerDto) error {
	switch iQuestion.GetType() {
	case question.SINGLE_CHOICE:
		option := answers.SingleChoice[iQuestion.GetId()]
		if err := f.singleChoiceProcessor.markAnswer(iQuestion.(*generated.SingleChoice), option); err != nil {
			return err
		}
	case question.MULTIPLE_CHOICE:
		options := answers.MultipleChoice[iQuestion.GetId()]
		if err := f.multipleChoiceProcessor.markAnswer(iQuestion.(*generated.MultipleChoice), options); err != nil {
			return err
		}
	case question.MATCHING:
		pairs := answers.Matching[iQuestion.GetId()]
		if err := f.matchingProcessor.markAnswer(iQuestion.(*generated.Matching), pairs, iQuestion.GetId()); err != nil {
			return err
		}
	case question.TEXT_INPUT:
		answer := answers.TextInput[iQuestion.GetId()]
		if err := f.textInputProcessor.markAnswer(iQuestion.(*generated.TextInput), answer); err != nil {
			return err
		}
	default:
		return errs.New(fmt.Sprintf("invalid type %s", iQuestion.GetType()), 400)
	}
	return nil
}

func (f *FormGeneratedProcessor) markAnswerAndCalculatePoints(
	iQuestion generated.IQuestion,
	questionEntity *question.Question,
	answers get.AnswerDto,
) (int, error) {
	var points int
	var err error

	switch iQuestion.GetType() {
	case question.SINGLE_CHOICE:
		option := answers.SingleChoice[iQuestion.GetId()]
		points, err = f.singleChoiceProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.SingleChoice),
			questionEntity.SingleChoice,
			option,
		)
	case question.MULTIPLE_CHOICE:
		options := answers.MultipleChoice[iQuestion.GetId()]
		points, err = f.multipleChoiceProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.MultipleChoice),
			questionEntity.MultipleChoice,
			options,
		)
	case question.MATCHING:
		pairs := answers.Matching[iQuestion.GetId()]
		points, err = f.matchingProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.Matching),
			questionEntity.Matching,
			pairs,
		)
	case question.TEXT_INPUT:
		answer := answers.TextInput[iQuestion.GetId()]
		points, err = f.textInputProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.TextInput),
			questionEntity.TextInput,
			answer,
		)
	default:
		return 0, errs.New(fmt.Sprintf("invalid type %s", iQuestion.GetType()), 400)
	}

	if err != nil {
		return 0, err
	}
	return points, nil
}

func (f *FormGeneratedProcessor) checkQuestion(
	questionId uuid.UUID,
	questionType question.QuestionType,
	answers get.AnswerDto,
) error {
	switch questionType {
	case question.SINGLE_CHOICE:
		if err := f.checkQuestionId(slices.Collect(maps.Keys(answers.SingleChoice)), questionId); err != nil {
			return err
		}
	case question.MULTIPLE_CHOICE:
		if err := f.checkQuestionId(slices.Collect(maps.Keys(answers.MultipleChoice)), questionId); err != nil {
			return err
		}
	case question.MATCHING:
		if err := f.checkQuestionId(slices.Collect(maps.Keys(answers.Matching)), questionId); err != nil {
			return err
		}
	case question.TEXT_INPUT:
		if err := f.checkQuestionId(slices.Collect(maps.Keys(answers.TextInput)), questionId); err != nil {
			return err
		}
	default:
		return errs.New(fmt.Sprintf("invalid type %s", questionType), 400)
	}
	return nil
}

func (f *FormGeneratedProcessor) checkQuestionId(questionIds []uuid.UUID, questionId uuid.UUID) error {
	if !slices.Contains(questionIds, questionId) {
		return errs.New(fmt.Sprintf("question with id %v not found", questionId), 404)
	}
	return nil
}
