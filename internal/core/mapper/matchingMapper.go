package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	question2 "hedgehog-forms/internal/core/model/form/pattern/section/block/question"
)

type MatchingMapper struct {
	commonMapper *CommonFieldQuestionDtoMapper
}

func NewMatchingMapper() *MatchingMapper {
	return &MatchingMapper{
		commonMapper: NewCommonFieldQuestionDtoMapper(),
	}
}

func (m *MatchingMapper) toDto(questionEntity *question2.Question) (*get.MatchingDto, error) {
	matchingDto := new(get.MatchingDto)
	m.commonMapper.CommonFieldsToDto(questionEntity, matchingDto)
	matchingDto.TermsAndDefinitions = m.termsAndDefinitionsToDto(questionEntity.Matching.Definitions)
	matchingDto.Points = m.pointsToDto(questionEntity.Matching.Points)
	return matchingDto, nil
}

func (m *MatchingMapper) termsAndDefinitionsToDto(
	definitions []question2.MatchingDefinition,
) []get.MatchingTermDefinitionDto {
	termDefinitionDtos := make([]get.MatchingTermDefinitionDto, 0)
	for _, definitionEntity := range definitions {
		var termAndDefinition get.MatchingTermDefinitionDto

		var definition get.MatchingOptionDto
		definition.Id = definitionEntity.Id
		definition.Text = definitionEntity.Text
		termAndDefinition.Definition = definition

		var term get.MatchingOptionDto
		term.Id = definitionEntity.MatchingTerm.Id
		term.Text = definitionEntity.MatchingTerm.Text
		termAndDefinition.Term = term

		termDefinitionDtos = append(termDefinitionDtos, termAndDefinition)
	}
	return termDefinitionDtos
}

func (m *MatchingMapper) pointsToDto(matchingPoints []question2.MatchingPoints) map[int]int {
	points := make(map[int]int)
	for _, matchingPoint := range matchingPoints {
		points[matchingPoint.CorrectAnswer] = matchingPoint.Points
	}
	return points
}
