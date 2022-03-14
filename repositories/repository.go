package repositories

import (
	"errors"

	"gorm.io/gorm"
)

type IBaseRepository interface {
	Any(query interface{}, args ...interface{}) error
	GetQueryable() *gorm.DB
}

type BaseRepository struct {
	db *gorm.DB
}

func (repo *BaseRepository) Any(query interface{}, args ...interface{}) error {
	var count int64
	err := repo.db.Session(&gorm.Session{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("no record not found")
	}
	return nil
}

func (repo *BaseRepository) GetQueryable() *gorm.DB {
	return repo.db.Session(&gorm.Session{})
}
