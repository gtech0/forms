package mapper

import (
	"hedgehog-forms/dto/get"
	"hedgehog-forms/model/form/pattern/section/block/question"
)

type MatchingMapper struct {
	commonMapper *CommonFieldQuestionDtoMapper
}

func NewMatchingMapper() *MatchingMapper {
	return &MatchingMapper{
		commonMapper: NewCommonFieldQuestionDtoMapper(),
	}
}

func (m *MatchingMapper) toDto(matchingObj *question.Matching) (*get.MatchingDto, error) {
	matchingDto := new(get.MatchingDto)
	m.commonMapper.commonFieldsToDto(matchingObj.Question, matchingDto)
	matchingDto.TermsAndDefinitions = m.termsAndDefinitionsToDto(matchingObj.Definitions)
	matchingDto.Points = m.pointsToDto(matchingObj.Points)
	return matchingDto, nil
}

func (m *MatchingMapper) termsAndDefinitionsToDto(
	definitionsObj []question.MatchingDefinition,
) []get.MatchingTermDefinitionDto {
	termDefinitionDtos := make([]get.MatchingTermDefinitionDto, 0)
	for _, definitionObj := range definitionsObj {
		var termAndDefinition get.MatchingTermDefinitionDto

		var definition get.MatchingOptionDto
		definition.Id = definitionObj.Id
		definition.Text = definitionObj.Text
		termAndDefinition.Definition = definition

		var term get.MatchingOptionDto
		term.Id = definitionObj.Term.Id
		term.Text = definitionObj.Term.Text
		termAndDefinition.Term = term

		termDefinitionDtos = append(termDefinitionDtos, termAndDefinition)
	}
	return termDefinitionDtos
}

func (m *MatchingMapper) pointsToDto(pointsObj []question.MatchingPoint) map[int]int {
	points := make(map[int]int)
	for _, pointObj := range pointsObj {
		points[pointObj.CorrectAnswers] = pointObj.Points
	}
	return points
}
