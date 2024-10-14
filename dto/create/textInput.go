package create

type TextQuestionDto struct {
	NewQuestionDto
	IsCaseSensitive bool     `json:"isCaseSensitive"`
	Answers         []string `json:"answers"`
	Points          int      `json:"points"`
}
