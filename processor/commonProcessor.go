package processor

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/errs"
	"hedgehog-forms/model/form/generated"
	"hedgehog-forms/model/form/pattern/section/block/question"
	"hedgehog-forms/util"
)

func filterQuestions[T generated.IQuestion](questions []T, questionType question.QuestionType) []T {
	filteredQuestions := make([]T, 0)
	for _, iQuestion := range questions {
		if iQuestion.GetType() == questionType {
			filteredQuestions = append(filteredQuestions, iQuestion)
		}
	}
	return filteredQuestions
}

func filterQuestionObjs[T question.IQuestion](questions []T, questionType question.QuestionType) []T {
	filteredQuestions := make([]T, 0)
	for _, iQuestion := range questions {
		if iQuestion.GetType() == questionType {
			filteredQuestions = append(filteredQuestions, iQuestion)
		}
	}
	return filteredQuestions
}

func findQuestion[T generated.IQuestion](questions []T, id uuid.UUID) (T, error) {
	for _, iQuestion := range questions {
		if iQuestion.GetId() == id {
			return iQuestion, nil
		}
	}

	return util.Zero[T](), errs.New(fmt.Sprintf("question with id %v not found", id), 404)
}

func findQuestionObj[T question.IQuestion](questions []T, id uuid.UUID) (T, error) {
	for _, iQuestion := range questions {
		if iQuestion.GetId() == id {
			return iQuestion, nil
		}
	}

	return util.Zero[T](), errs.New(fmt.Sprintf("question with id %v not found", id), 404)
}
