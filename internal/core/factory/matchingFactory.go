package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/create"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type MatchingFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewMatchingFactory() *MatchingFactory {
	return &MatchingFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (m *MatchingFactory) BuildFromDto(dto *create.MatchingQuestionDto) (*question.Question, error) {
	questionEntity := new(question.Question)
	questionEntity.Matching = new(question.Matching)
	if err := m.commonMapper.MapCommonDtoFields(dto.NewQuestionDto, questionEntity); err != nil {
		return nil, err
	}

	terms, definitions := m.buildTermsAndDefinitions(dto.TermsAndDefinitions)
	questionEntity.Matching.Terms = terms
	questionEntity.Matching.Definitions = definitions
	for answer, value := range dto.Points {
		var pointEntity question.MatchingPoints
		pointEntity.CorrectAnswers = answer
		pointEntity.Points = value
		questionEntity.Matching.Points = append(questionEntity.Matching.Points, pointEntity)
	}

	return questionEntity, nil
}

func (m *MatchingFactory) BuildFromEntity(questionEntity *question.Question) (*question.Question, error) {
	newQuestionEntity := new(question.Question)
	newQuestionEntity.Matching = new(question.Matching)
	terms := make([]question.MatchingTerm, 0)
	definitions := make([]question.MatchingDefinition, 0)

	for _, term := range questionEntity.Matching.Terms {
		newDefinition := m.buildDefinitionFromEntity(term, newQuestionEntity.Id)
		newTerm := m.buildTermFromEntity(term, newQuestionEntity.Id, newDefinition)

		terms = append(terms, newTerm)
		definitions = append(definitions, newDefinition)
	}

	newQuestionEntity.Matching.Points = questionEntity.Matching.Points
	newQuestionEntity.Matching.Terms = terms
	newQuestionEntity.Matching.Definitions = definitions

	if err := m.commonMapper.MapCommonEntityFields(*questionEntity, newQuestionEntity); err != nil {
		return nil, err
	}

	return newQuestionEntity, nil
}

func (m *MatchingFactory) buildTermFromEntity(
	term question.MatchingTerm,
	questionId uuid.UUID,
	definition question.MatchingDefinition,
) question.MatchingTerm {
	var newTerm question.MatchingTerm
	newTerm.Text = term.Text
	newTerm.MatchingId = questionId
	newTerm.MatchingDefinitionId = definition.Id
	return newTerm
}

func (m *MatchingFactory) buildDefinitionFromEntity(
	term question.MatchingTerm,
	questionId uuid.UUID,
) question.MatchingDefinition {
	var definition question.MatchingDefinition
	definition.Text = term.Text
	definition.MatchingId = questionId
	return definition
}

func (m *MatchingFactory) buildTermsAndDefinitions(
	matchingMap map[string]string,
) ([]question.MatchingTerm, []question.MatchingDefinition) {
	terms := make([]question.MatchingTerm, 0)
	definitions := make([]question.MatchingDefinition, 0)
	for key, value := range matchingMap {
		var definition question.MatchingDefinition
		definition.Id = uuid.New()
		definition.Text = value

		var term question.MatchingTerm
		term.Text = key
		term.MatchingDefinitionId = definition.Id

		terms = append(terms, term)
		definitions = append(definitions, definition)
	}
	return terms, definitions
}
