package config

import (
	"fmt"
	"saady/app/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDbMysql(cfg *config.AppConfig) *gorm.DB {
	DNS := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DBUSER, cfg.DBPASS, cfg.DBHOST, cfg.DBPORT, cfg.DBNAME)

	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}
