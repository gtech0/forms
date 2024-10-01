package form

type FormStatus string

const (
	NEW         FormStatus = "NEW"
	IN_PROGRESS FormStatus = "IN_PROGRESS"
	SUBMITTED   FormStatus = "SUBMITTED"
	COMPLETED   FormStatus = "COMPLETED"
)
