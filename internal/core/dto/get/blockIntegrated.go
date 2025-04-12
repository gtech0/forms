package get

type IntegrationBlockDto struct {
	BlockDto
	Variant   *IntegratedVariantDto    `json:"variant"`
	Questions []IntegratedIQuestionDto `json:"questions"`
}
