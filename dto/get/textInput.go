package get

type TextInputDto struct {
	QuestionDto
	Points          int      `json:"points"`
	IsCaseSensitive bool     `json:"isCaseSensitive"`
	Answers         []string `json:"answers"`
}
