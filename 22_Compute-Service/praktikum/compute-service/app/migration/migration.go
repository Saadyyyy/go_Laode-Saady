package migration

import (
	"praktikum_22/models"

	"gorm.io/gorm"
)

func InitMigrationMysql(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	// auto migrate untuk table book
}
