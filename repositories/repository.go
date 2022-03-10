package repositories

import "gorm.io/gorm"

type BaseRepository struct {
	db *gorm.DB
}
