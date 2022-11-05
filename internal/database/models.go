package database

import "gorm.io/gorm"

type Configuration struct {
	gorm.Model

	UID   string
	Name  string
	Owner string
	Data  string
}
