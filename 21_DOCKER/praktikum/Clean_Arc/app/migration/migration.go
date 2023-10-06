package migration

import (
	r "saady/feature/users/repository"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		r.User{},
	)
}
