package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/internal/core/dto/create"
	question2 "hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type MatchingFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewMatchingFactory() *MatchingFactory {
	return &MatchingFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (m *MatchingFactory) BuildFromDto(dto *create.MatchingQuestionDto) (*question2.Question, error) {
	questionEntity := new(question2.Question)
	questionEntity.Matching = new(question2.Matching)
	if err := m.commonMapper.MapCommonDtoFields(dto.NewQuestionDto, questionEntity); err != nil {
		return nil, err
	}

	terms, definitions := m.buildTermsAndDefinitions(dto.TermsAndDefinitions)
	questionEntity.Matching.Terms = terms
	questionEntity.Matching.Definitions = definitions
	for answer, value := range dto.Points {
		var pointEntity question2.MatchingPoints
		pointEntity.CorrectAnswer = answer
		pointEntity.Points = value
		questionEntity.Matching.Points = append(questionEntity.Matching.Points, pointEntity)
	}

	return questionEntity, nil
}

func (m *MatchingFactory) BuildFromEntity(questionEntity *question2.Question) (*question2.Question, error) {
	newQuestionEntity := new(question2.Question)
	newQuestionEntity.Matching = new(question2.Matching)
	terms := make([]question2.MatchingTerm, 0)
	definitions := make([]question2.MatchingDefinition, 0)

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
	term question2.MatchingTerm,
	questionId uuid.UUID,
	definition question2.MatchingDefinition,
) question2.MatchingTerm {
	var newTerm question2.MatchingTerm
	newTerm.Text = term.Text
	newTerm.MatchingId = questionId
	newTerm.MatchingDefinitionId = definition.Id
	return newTerm
}

func (m *MatchingFactory) buildDefinitionFromEntity(
	term question2.MatchingTerm,
	questionId uuid.UUID,
) question2.MatchingDefinition {
	var definition question2.MatchingDefinition
	definition.Text = term.Text
	definition.MatchingId = questionId
	return definition
}

func (m *MatchingFactory) buildTermsAndDefinitions(
	matchingMap map[string]string,
) ([]question2.MatchingTerm, []question2.MatchingDefinition) {
	terms := make([]question2.MatchingTerm, 0)
	definitions := make([]question2.MatchingDefinition, 0)
	for key, value := range matchingMap {
		var definition question2.MatchingDefinition
		definition.Id = uuid.New()
		definition.Text = value

		var term question2.MatchingTerm
		term.Text = key
		term.MatchingDefinitionId = definition.Id

		terms = append(terms, term)
		definitions = append(definitions, definition)
	}
	return terms, definitions
}
