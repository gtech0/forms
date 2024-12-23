package published

import (
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/generated"
)

type Submission struct {
	model.Base
	Sections []generated.Section
}
