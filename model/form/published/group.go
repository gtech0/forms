package published

import "hedgehog-forms/model"

type Group struct {
	model.Base
	Name  string
	Users []User `gorm:"many2many:user_groups"`
}
