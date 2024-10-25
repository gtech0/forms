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
	"maps"
	"slices"
)

type QuestionProcessor struct {
	matchingProcessor       *MatchingProcessor
	singleChoiceProcessor   *SingleChoiceProcessor
	multipleChoiceProcessor *MultipleChoiceProcessor
	textInputProcessor      *TextInputProcessor
}

func NewQuestionProcessor() *QuestionProcessor {
	return &QuestionProcessor{
		matchingProcessor:       NewMatchingProcessor(),
		singleChoiceProcessor:   NewSingleChoiceProcessor(),
		multipleChoiceProcessor: NewMultipleChoiceProcessor(),
		textInputProcessor:      NewTextInputProcessor(),
	}
}

func (q *QuestionProcessor) markAnswers(formGenerated generated.FormGenerated, answers get.AnswerDto) error {
	questions := q.extractQuestionsFromGeneratedForm(formGenerated)
	for _, iQuestion := range questions {
		if err := q.checkQuestion(iQuestion.GetId(), iQuestion.GetType(), answers); err != nil {
			return err
		}

		if err := q.markAnswer(iQuestion, answers); err != nil {
			return err
		}
	}

	return nil
}

func (q *QuestionProcessor) markAnswersAndCalculatePoints(
	formGenerated generated.FormGenerated,
	formPattern pattern.FormPattern,
	answers get.AnswerDto,
) error {
	questions := q.extractQuestionsFromGeneratedForm(formGenerated)
	questionObjs := q.extractQuestionObjs(formPattern)

	for _, iQuestion := range questions {
		if err := q.checkQuestion(iQuestion.GetId(), iQuestion.GetType(), answers); err != nil {
			return err
		}
	}

	for _, iQuestion := range questionObjs {
		if err := q.checkQuestion(iQuestion.GetId(), iQuestion.GetType(), answers); err != nil {
			return err
		}
	}

	points := 0
	for _, iQuestion := range questions {
		for _, questionObj := range questionObjs {
			if iQuestion.GetId() == questionObj.GetId() {
				calculatedPoints, err := q.markAnswerAndCalculatePoints(iQuestion, questionObj, answers)
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

func (q *QuestionProcessor) verifyForm(formGenerated generated.FormGenerated, checkDto create.CheckDto) {
	questions := q.extractQuestionsFromGeneratedForm(formGenerated)
	if checkDto.Points != nil {
		for questionId, points := range checkDto.Points {
			var generatedQuestion generated.IQuestion
			for _, iQuestion := range questions {
				if iQuestion.GetId() == questionId {
					generatedQuestion = iQuestion
					break
				}
			}
			q.applyNewPoints(formGenerated, generatedQuestion, points)
		}
	}
}

// FixMe: probably not a correct logic
func (q *QuestionProcessor) applyNewPoints(formGenerated generated.FormGenerated,
	questionGenerated generated.IQuestion,
	newPoints int,
) {
	difference := questionGenerated.GetPoints() - newPoints
	questionGenerated.SetPoints(formGenerated.Points - difference)
}

func (q *QuestionProcessor) extractQuestionsFromGeneratedForm(formGenerated generated.FormGenerated) []generated.IQuestion {
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

func (q *QuestionProcessor) extractQuestionObjs(formPattern pattern.FormPattern) []question.IQuestion {
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

func (q *QuestionProcessor) markAnswer(iQuestion generated.IQuestion, answers get.AnswerDto) error {
	switch iQuestion.GetType() {
	case question.SINGLE_CHOICE:
		option := answers.SingleChoice[iQuestion.GetId()]
		if err := q.singleChoiceProcessor.markAnswer(iQuestion.(*generated.SingleChoice), option); err != nil {
			return err
		}
	case question.MULTIPLE_CHOICE:
		options := answers.MultipleChoice[iQuestion.GetId()]
		if err := q.multipleChoiceProcessor.markAnswer(iQuestion.(*generated.MultipleChoice), options); err != nil {
			return err
		}
	case question.MATCHING:
		pairs := answers.Matching[iQuestion.GetId()]
		if err := q.matchingProcessor.markAnswer(iQuestion.(*generated.Matching), pairs, iQuestion.GetId()); err != nil {
			return err
		}
	case question.TEXT_INPUT:
		answer := answers.TextInput[iQuestion.GetId()]
		if err := q.textInputProcessor.markAnswer(iQuestion.(*generated.TextInput), answer); err != nil {
			return err
		}
	default:
		return errs.New(fmt.Sprintf("invalid type %s", iQuestion.GetType()), 400)
	}
	return nil
}

func (q *QuestionProcessor) markAnswerAndCalculatePoints(
	iQuestion generated.IQuestion,
	questionObj question.IQuestion,
	answers get.AnswerDto) (int, error) {
	switch iQuestion.GetType() {
	case question.SINGLE_CHOICE:
		option := answers.SingleChoice[iQuestion.GetId()]
		points, err := q.singleChoiceProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.SingleChoice),
			questionObj.(*question.SingleChoice),
			option,
		)
		if err != nil {
			return 0, err
		}
		return points, nil
	case question.MULTIPLE_CHOICE:
		options := answers.MultipleChoice[iQuestion.GetId()]
		points, err := q.multipleChoiceProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.MultipleChoice),
			questionObj.(*question.MultipleChoice),
			options,
		)
		if err != nil {
			return 0, err
		}
		return points, nil
	case question.MATCHING:
		pairs := answers.Matching[iQuestion.GetId()]
		points, err := q.matchingProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.Matching),
			questionObj.(*question.Matching),
			pairs,
		)
		if err != nil {
			return 0, err
		}
		return points, nil
	case question.TEXT_INPUT:
		answer := answers.TextInput[iQuestion.GetId()]
		points, err := q.textInputProcessor.markAnswerAndCalculatePoints(
			iQuestion.(*generated.TextInput),
			questionObj.(*question.TextInput),
			answer,
		)
		if err != nil {
			return 0, err
		}
		return points, nil
	default:
		return 0, errs.New(fmt.Sprintf("invalid type %s", iQuestion.GetType()), 400)
	}
}

func (q *QuestionProcessor) checkQuestion(
	questionId uuid.UUID,
	questionType question.QuestionType,
	answers get.AnswerDto,
) error {
	switch questionType {
	case question.SINGLE_CHOICE:
		if err := q.checkQuestionId(slices.Collect(maps.Keys(answers.SingleChoice)), questionId); err != nil {
			return err
		}
	case question.MULTIPLE_CHOICE:
		if err := q.checkQuestionId(slices.Collect(maps.Keys(answers.MultipleChoice)), questionId); err != nil {
			return err
		}
	case question.MATCHING:
		if err := q.checkQuestionId(slices.Collect(maps.Keys(answers.Matching)), questionId); err != nil {
			return err
		}
	case question.TEXT_INPUT:
		if err := q.checkQuestionId(slices.Collect(maps.Keys(answers.TextInput)), questionId); err != nil {
			return err
		}
	default:
		return errs.New(fmt.Sprintf("invalid type %s", questionType), 400)
	}
	return nil
}

func (q *QuestionProcessor) checkQuestionId(questionIds []uuid.UUID, questionId uuid.UUID) error {
	if !slices.Contains(questionIds, questionId) {
		return errs.New(fmt.Sprintf("question with id %v not found", questionId), 404)
	}
	return nil
}
