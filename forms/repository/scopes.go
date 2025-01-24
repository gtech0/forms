package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}

func preload(d *gorm.DB) *gorm.DB {
	return d.Preload(clause.Associations, preload)
}
