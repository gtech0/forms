package question

type QuestionType string

const (
	MULTIPLE_CHOICE QuestionType = "MULTIPLE_CHOICE"
	SINGLE_CHOICE   QuestionType = "SINGLE_CHOICE"
	MATCHING        QuestionType = "MATCHING"
	TEXT_INPUT      QuestionType = "TEXT_INPUT"
	EXISTING        QuestionType = "EXISTING"
)
