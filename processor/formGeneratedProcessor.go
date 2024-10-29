package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/dto/get"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern"
	"hedgehog-forms/model/form/pattern/section/block"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"hedgehog-forms/util"
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
	questions := f.extractQuestionsFromGeneratedForm(formGenerated)
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
	questions := f.extractQuestionsFromGeneratedForm(formGenerated)
	questionObjs := f.extractQuestionObjs(formPattern)

	for _, iQuestion := range questions {
		if err := f.checkQuestion(iQuestion.GetId(), iQuestion.GetType(), answers); err != nil {
			return err
		}
	}

	for _, iQuestion := range questionObjs {
		if err := f.checkQuestion(iQuestion.GetId(), iQuestion.GetType(), answers); err != nil {
			return err
		}
	}

	points := 0
	for _, iQuestion := range questions {
		for _, questionObj := range questionObjs {
			if iQuestion.GetId() == questionObj.GetId() {
				calculatedPoints, err := f.markAnswerAndCalculatePoints(iQuestion, questionObj, answers)
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

func (f *FormGeneratedProcessor) VerifyForm(formGenerated *generated.FormGenerated, checkDto create.CheckDto) {
	questions := f.extractQuestionsFromGeneratedForm(formGenerated)
	if checkDto.Points != nil {
		for questionId, points := range checkDto.Points {
			var generatedQuestion generated.IQuestion
			for _, iQuestion := range questions {
				if iQuestion.GetId() == questionId {
					generatedQuestion = iQuestion
					break
				}
			}
			f.applyNewPoints(formGenerated, generatedQuestion, points)
		}
	}
}

// FixMe: probably not a correct logic
func (f *FormGeneratedProcessor) applyNewPoints(formGenerated *generated.FormGenerated,
	questionGenerated generated.IQuestion,
	newPoints int,
) {
	difference := questionGenerated.GetPoints() - newPoints
	questionGenerated.SetPoints(formGenerated.Points - difference)
}

func (f *FormGeneratedProcessor) CalculateMark(formGenerated *generated.FormGenerated, marks map[string]int) error {
	pointsForForm := formGenerated.Points
	var requiredPointsForMark int
	for _, points := range marks {
		if points >= requiredPointsForMark && points <= pointsForForm {
			requiredPointsForMark = points
		}
	}

	mark, ok := util.FindKeyByValue(marks, requiredPointsForMark)
	if !ok {
		return errs.New(fmt.Sprintf("mark for %d points is not found", requiredPointsForMark), 500)
	}
	formGenerated.Mark = mark
	return nil
}

func (f *FormGeneratedProcessor) extractQuestionsFromGeneratedForm(formGenerated *generated.FormGenerated) []generated.IQuestion {
	questions := make([]generated.IQuestion, 0)
	for _, generatedSection := range formGenerated.Sections {
		for _, generatedBlock := range generatedSection.Blocks {
			if generatedBlock != nil {
				questions = slices.Concat(questions, generatedBlock.Questions)

				if generatedBlock.Variant != nil {
					questions = slices.Concat(questions, generatedBlock.Variant.Questions)
				}
			}
		}
	}
	return questions
}

func (f *FormGeneratedProcessor) extractQuestionObjs(formPattern pattern.FormPattern) []question.IQuestion {
	questions := make([]question.IQuestion, 0)
	for _, patternSection := range formPattern.Sections {
		for _, sectionBlock := range patternSection.Blocks {
			switch assertedBlock := sectionBlock.(type) {
			case *block.DynamicBlock:
				questions = slices.Concat(questions, assertedBlock.Questions)
			case *block.StaticBlock:
				variants := assertedBlock.Variants
				for _, variant := range variants {
					questions = slices.Concat(questions, variant.Questions)
				}
			}
		}
	}
	return questions
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
	questionObj question.IQuestion,
	answers get.AnswerDto,
) (int, error) {
	var points int
	var err error

	switch iQuestion.GetType() {
	case question.SINGLE_CHOICE:
		option := answers.SingleChoice[iQuestion.GetId()]
		points, err = f.singleChoiceProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.SingleChoice),
			questionObj.(*question.SingleChoice),
			option,
		)
	case question.MULTIPLE_CHOICE:
		options := answers.MultipleChoice[iQuestion.GetId()]
		points, err = f.multipleChoiceProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.MultipleChoice),
			questionObj.(*question.MultipleChoice),
			options,
		)
	case question.MATCHING:
		pairs := answers.Matching[iQuestion.GetId()]
		points, err = f.matchingProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.Matching),
			questionObj.(*question.Matching),
			pairs,
		)
	case question.TEXT_INPUT:
		answer := answers.TextInput[iQuestion.GetId()]
		points, err = f.textInputProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.TextInput),
			questionObj.(*question.TextInput),
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
