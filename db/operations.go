package db

import (
	"log"
	"webapi/models/domains"

	"gorm.io/gorm"
)

//Migration : auto migrate data models
func (database *DB) Migration() {
	log.Println("Database migrating...")
	database.AutoMigrate(
		domains.Demo{},
	)
}

func (database *DB) SetConstraints() {

}

func GetNewSession() *gorm.DB {
	return dbInstance.Session(&gorm.Session{})
}

func OpenTransection(callback func(tx *gorm.DB) error) error {
	return GetNewSession().Transaction(callback)
}
