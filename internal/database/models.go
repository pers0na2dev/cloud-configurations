package database

import "gorm.io/gorm"

type Configuration struct {
	gorm.Model

	UID   string `gorm:"index"`
	Name  string
	Owner string
	Data  string
}
