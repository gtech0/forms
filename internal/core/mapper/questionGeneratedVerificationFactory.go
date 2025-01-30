package mapper

import (
	"fmt"
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/dto/verify"
	"hedgehog-forms/internal/core/errs"
	"hedgehog-forms/internal/core/model/form/generated"
)

type QuestionGeneratedVerificationFactory struct{}

func NewQuestionGeneratedVerificationFactory() *QuestionGeneratedVerificationFactory {
	return &QuestionGeneratedVerificationFactory{}
}

func (q *QuestionGeneratedVerificationFactory) build(
	questions []generated.IQuestion,
	questionsWithCorrectAnswers map[uuid.UUID]get.IQuestionDto,
) ([]verify.Question, error) {
	verifiedQuestions := make([]verify.Question, 0)
	for _, currQuestion := range questions {
		newQuestion, err := q.buildQuestion(currQuestion, questionsWithCorrectAnswers)
		if err != nil {
			return nil, err
		}

		verifiedQuestions = append(verifiedQuestions, *newQuestion)
	}
	return verifiedQuestions, nil
}

func (q *QuestionGeneratedVerificationFactory) buildQuestion(
	iQuestion generated.IQuestion,
	questionsWithCorrectAnswers map[uuid.UUID]get.IQuestionDto,
) (*verify.Question, error) {
	correctAnswer, err := q.getCorrectAnswer(iQuestion.GetId(), questionsWithCorrectAnswers)
	if err != nil {
		return nil, err
	}

	return &verify.Question{
		QuestionWithCorrectAnswer: correctAnswer,
		QuestionWithStudentAnswer: iQuestion,
	}, nil
}

func (q *QuestionGeneratedVerificationFactory) getCorrectAnswer(
	questionId uuid.UUID,
	questionsWithCorrectAnswers map[uuid.UUID]get.IQuestionDto,
) (get.IQuestionDto, error) {
	if _, ok := questionsWithCorrectAnswers[questionId]; ok {
		return questionsWithCorrectAnswers[questionId], nil
	}

	return nil, errs.New(
		fmt.Sprintf("question with id %v not found in questions with correct answer", questionId),
		400,
	)
}
