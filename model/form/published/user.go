package published

import "hedgehog-forms/model"

type User struct {
	model.Base
	Name   string
	Groups []Group `gorm:"many2many:user_groups"`
}
