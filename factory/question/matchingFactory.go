package question

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/section/block/question"
)

type MatchingFactory struct {
	commonMapper *CommonFieldQuestionMapper
}

func NewMatchingFactory() *MatchingFactory {
	return &MatchingFactory{
		commonMapper: NewCommonFieldQuestionMapper(),
	}
}

func (m *MatchingFactory) BuildFromDto(dto *create.MatchingQuestionDto) (*question.Matching, error) {
	questionObj := new(question.Matching)
	m.commonMapper.MapCommonFieldsDto(dto.NewQuestionDto, questionObj)

	terms := make([]question.MatchingTerm, 0)
	definitions := make([]question.MatchingDefinition, 0)

	m.buildTermsAndDefinitions(dto.TermsAndDefinitions, terms, definitions, questionObj.Id)

	questionObj.Terms = terms
	questionObj.Definitions = definitions
	for answer, value := range dto.Points {
		var pointObj question.MatchingPoint
		pointObj.CorrectAnswer = answer
		pointObj.Points = value
		questionObj.Points = append(questionObj.Points, pointObj)
	}

	return questionObj, nil
}

func (m *MatchingFactory) BuildFromObj(questionObj *question.Matching) *question.Matching {
	newQuestionObj := new(question.Matching)
	terms := make([]question.MatchingTerm, 0)
	definitions := make([]question.MatchingDefinition, 0)

	for _, term := range questionObj.Terms {
		newDefinition := m.buildDefinitionFromEntity(term, newQuestionObj.Id)
		newTerm := m.buildTermFromEntity(term, newQuestionObj.Id, newDefinition)

		terms = append(terms, newTerm)
		definitions = append(definitions, newDefinition)
	}

	newQuestionObj.Points = questionObj.Points
	newQuestionObj.Terms = terms
	newQuestionObj.Definitions = definitions

	m.commonMapper.MapCommonFieldsObj(questionObj.Question, newQuestionObj)

	return newQuestionObj
}

func (m *MatchingFactory) buildTermFromEntity(
	term question.MatchingTerm,
	questionId uuid.UUID,
	definition question.MatchingDefinition,
) question.MatchingTerm {
	var newTermObj question.MatchingTerm
	newTermObj.Text = term.Text
	newTermObj.MatchingId = questionId
	newTermObj.MatchingDefinitionId = definition.Id
	return newTermObj
}

func (m *MatchingFactory) buildDefinitionFromEntity(
	term question.MatchingTerm,
	questionId uuid.UUID,
) question.MatchingDefinition {
	var newDefObj question.MatchingDefinition
	newDefObj.Text = term.Text
	newDefObj.MatchingId = questionId
	return newDefObj
}

func (m *MatchingFactory) buildTermsAndDefinitions(
	matchingMap map[string]string,
	terms []question.MatchingTerm,
	definitions []question.MatchingDefinition,
	questionId uuid.UUID,
) {
	for key, value := range matchingMap {
		var definition question.MatchingDefinition
		definition.Text = value
		definition.MatchingId = questionId
		definitions = append(definitions, definition)

		var term question.MatchingTerm
		term.Text = key
		term.MatchingDefinitionId = definition.Id
		term.MatchingId = questionId
		terms = append(terms, term)
	}
}
