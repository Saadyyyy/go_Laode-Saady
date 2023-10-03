package migration

import (
	r "Laode_Saady/20_Clean-Architecture/praktikum/belajar-go-echo/feature/users/repository"

	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		r.User{},
	)
}
