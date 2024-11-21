package factory

import (
	"github.com/google/uuid"
	"hedgehog-forms/dto/create"
	"hedgehog-forms/model/form/pattern/section/block/question"
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
	questionObj := new(question.Question)
	questionObj.Matching = new(question.Matching)
	if err := m.commonMapper.MapCommonFieldsDto(dto.NewQuestionDto, questionObj); err != nil {
		return nil, err
	}

	terms, definitions := m.buildTermsAndDefinitions(dto.TermsAndDefinitions)
	questionObj.Matching.Terms = terms
	questionObj.Matching.Definitions = definitions
	for answer, value := range dto.Points {
		var pointObj question.MatchingPoint
		pointObj.CorrectAnswers = answer
		pointObj.Points = value
		questionObj.Matching.Points = append(questionObj.Matching.Points, pointObj)
	}

	return questionObj, nil
}

func (m *MatchingFactory) BuildFromObj(questionObj *question.Question) (*question.Question, error) {
	newQuestionObj := new(question.Question)
	newQuestionObj.Matching = new(question.Matching)
	terms := make([]question.MatchingTerm, 0)
	definitions := make([]question.MatchingDefinition, 0)

	for _, term := range questionObj.Matching.Terms {
		newDefinition := m.buildDefinitionFromEntity(term, newQuestionObj.Id)
		newTerm := m.buildTermFromEntity(term, newQuestionObj.Id, newDefinition)

		terms = append(terms, newTerm)
		definitions = append(definitions, newDefinition)
	}

	newQuestionObj.Matching.Points = questionObj.Matching.Points
	newQuestionObj.Matching.Terms = terms
	newQuestionObj.Matching.Definitions = definitions

	if err := m.commonMapper.MapCommonFieldsObj(*questionObj, newQuestionObj); err != nil {
		return nil, err
	}

	return newQuestionObj, nil
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

		//definition.MatchingTerm = term
		terms = append(terms, term)
		definitions = append(definitions, definition)
	}
	return terms, definitions
}
