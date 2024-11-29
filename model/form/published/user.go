package published

import (
	"hedgehog-forms/model"
	"hedgehog-forms/model/form/generated"
)

type User struct {
	model.Base
	Name          string
	Groups        []Group `gorm:"many2many:user_groups"`
	FormGenerated generated.FormGenerated
}
