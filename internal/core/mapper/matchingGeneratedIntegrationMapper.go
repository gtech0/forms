package mapper

import (
	"hedgehog-forms/internal/core/dto/get"
	"hedgehog-forms/internal/core/model/form/generated"
	"hedgehog-forms/internal/core/model/form/pattern/question"
)

type MatchingGeneratedIntegrationMapper struct {
}

func NewMatchingGeneratedIntegrationMapper() *MatchingGeneratedIntegrationMapper {
	return &MatchingGeneratedIntegrationMapper{}
}

func (m *MatchingGeneratedIntegrationMapper) toDto(
	matching *generated.Matching,
	questionEntity *question.Question,
) (*get.IntegratedMatchingDto, error) {
	matchingDto := new(get.IntegratedMatchingDto)
	matchingDto.Id = matching.Id
	matchingDto.Description = matching.Description
	matchingDto.OwnerId = questionEntity.OwnerId
	matchingDto.Type = matching.Type
	matchingDto.Terms = matching.Terms
	matchingDto.Definitions = matching.Definitions
	matchingDto.TermsAndDefinitions = m.termsAndDefinitionsToDto(questionEntity.Matching.Definitions)
	matchingDto.EnteredAnswers = matching.EnteredAnswers
	return matchingDto, nil
}

func (m *MatchingGeneratedIntegrationMapper) termsAndDefinitionsToDto(
	definitions []question.MatchingDefinition,
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
