package database

import (
	"cloud-configurations/internal/environment"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Database() *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(environment.Config.DSN),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func Migration() {
	err := DB.AutoMigrate(
		&Configuration{},
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}
