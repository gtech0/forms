package create

type MultipleChoiceQuestionDto struct {
	NewQuestionDto
	Options        []string    `json:"options"`
	CorrectOptions []int       `json:"correctOptions"`
	Points         map[int]int `json:"points"`
}
