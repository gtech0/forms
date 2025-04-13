package get

type IntegratedTextInputDto struct {
	IntegratedQuestionDto
	IsCaseSensitive bool     `json:"isCaseSensitive"`
	Answers         []string `json:"answers,omitempty"`
	EnteredAnswer   string   `json:"enteredAnswer"`
}
