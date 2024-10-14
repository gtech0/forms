package create

type SingleChoiceQuestionDto struct {
	NewQuestionDto
	Options       []string `json:"options"`
	CorrectOption int      `json:"correctOption"`
	Points        int      `json:"points"`
}
