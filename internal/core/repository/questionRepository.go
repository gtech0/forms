package repository

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
	"hedgehog-forms/internal/core/errs"
	question2 "hedgehog-forms/internal/core/model/form/pattern/section/block/question"
	"hedgehog-forms/pkg/database"
)

type QuestionRepository struct{}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}

func (q *QuestionRepository) Create(questionEntity *question2.Question) error {
	if err := database.DB.Create(questionEntity).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}

func (q *QuestionRepository) FindById(id uuid.UUID) (*question2.Question, error) {
	questionEntity := new(question2.Question)
	if err := database.DB.Model(&question2.Question{}).
		Where("id = ?", id).
		First(questionEntity).Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return questionEntity, nil
}

func (q *QuestionRepository) FindByParamsAndPaginate(
	clauses []clause.Expression,
	name string,
	page, size int,
	types []question2.QuestionType,
) ([]question2.Question, error) {
	questions := make([]question2.Question, 0)
	if err := database.DB.
		Preload(clause.Associations, preload).
		Model(&question2.Question{}).
		Clauses(clauses...).
		Where("title LIKE ?", fmt.Sprintf("%%%s%%", name)).
		Where("type IN ?", types).
		Scopes(paginate(page, size)).
		Find(&questions).
		Error; err != nil {
		return nil, errs.New(err.Error(), 500)
	}
	return questions, nil
}

func (q *QuestionRepository) DeleteById(id uuid.UUID) error {
	if err := database.DB.Delete(&question2.Question{}, id).Error; err != nil {
		return errs.New(err.Error(), 500)
	}
	return nil
}
