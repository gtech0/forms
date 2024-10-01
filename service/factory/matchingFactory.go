package factory

import (
	"hedgehog-forms/dto"
	"hedgehog-forms/model/form/section/block/question"
)

type MatchingFactory struct{}

func NewMatchingFactory() *MatchingFactory {
	return &MatchingFactory{}
}

func (m *MatchingFactory) buildFromDto(dto dto.CreateMatchingQuestionDto) question.Matching {
	var questionObj question.Matching
	NewCommonFieldQuestionMapper().MapCommonFieldsDto(dto.NewQuestionDto, questionObj.Question)

	listsSize := len(dto.TermsAndDefinitions)
	terms := make([]question.MatchingTerm, listsSize)
	definitions := make([]question.MatchingDefinition, listsSize)

	m.buildTermsAndDefinitions(dto.TermsAndDefinitions, terms, definitions, questionObj)

	questionObj.Terms = terms
	questionObj.Definitions = definitions
	for answer, value := range dto.Points {
		var pointObj question.MatchingPoint
		pointObj.CorrectAnswer = answer
		pointObj.Points = value
		questionObj.Points = append(questionObj.Points, pointObj)
	}

	return questionObj
}

func (m *MatchingFactory) buildFromObj(questionObj question.Matching) question.Matching {
	var newQuestionObj question.Matching
	terms := make([]question.MatchingTerm, len(questionObj.Terms))
	definitions := make([]question.MatchingDefinition, len(questionObj.Definitions))

	for _, term := range questionObj.Terms {
		newDefinition := m.buildDefinitionFromEntity(term, newQuestionObj)
		newTerm := m.buildTermFromEntity(term, newQuestionObj, newDefinition)

		terms = append(terms, newTerm)
		definitions = append(definitions, newDefinition)
	}

	newQuestionObj.Points = questionObj.Points
	newQuestionObj.Terms = terms
	newQuestionObj.Definitions = definitions

	NewCommonFieldQuestionMapper().MapCommonFieldsObj(questionObj.Question, newQuestionObj.Question)

	return newQuestionObj
}

func (m *MatchingFactory) buildTermFromEntity(
	term question.MatchingTerm,
	questionObj question.Matching,
	definition question.MatchingDefinition,
) question.MatchingTerm {
	var newTermObj question.MatchingTerm
	newTermObj.Text = term.Text
	newTermObj.MatchingId = questionObj.Id
	newTermObj.MatchingDefinitionId = definition.Id
	return newTermObj
}

func (m *MatchingFactory) buildDefinitionFromEntity(
	term question.MatchingTerm,
	questionObj question.Matching,
) question.MatchingDefinition {
	var newDefObj question.MatchingDefinition
	newDefObj.Text = term.Text
	newDefObj.MatchingId = questionObj.Id
	return newDefObj
}

func (m *MatchingFactory) buildTermsAndDefinitions(
	matchingMap map[string]string,
	terms []question.MatchingTerm,
	definitions []question.MatchingDefinition,
	questionObj question.Matching,
) {
	for key, value := range matchingMap {
		var definition question.MatchingDefinition
		definition.Text = value
		definition.MatchingId = questionObj.Id
		definitions = append(definitions, definition)

		var term question.MatchingTerm
		term.Text = key
		term.MatchingDefinitionId = definition.Id
		term.MatchingId = questionObj.Id
		terms = append(terms, term)
	}
}
