package repositories

import (
	"errors"

	"gorm.io/gorm"
)

type EntityID interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type IBaseRepository[T1 EntityID, T2 any] interface {
	Get(id T1) (*T2, error)
	GetAll() ([]*T2, error)
	GetByFilter(filter interface{}, args ...interface{}) (*T2, error)
	GetAllByFilter(filter interface{}, args ...interface{}) ([]*T2, error)
	Create(entity *T2) (*T2, error)
	Update(entity *T2) (*T2, error)
	Delete(id T1) error
	GetQueryable() *gorm.DB
	Count(query interface{}, args ...interface{}) (int64, error)
	Any(query interface{}, args ...interface{}) error
}

type BaseRepository[T1 EntityID, T2 any] struct {
	db        *gorm.DB
	tableName string
}

func NewBaseRepository[T1 EntityID, T2 any](db *gorm.DB, tableName string) IBaseRepository[T1, T2] {
	return &BaseRepository[T1, T2]{
		db:        db,
		tableName: tableName,
	}
}

func (repo *BaseRepository[T1, T2]) Get(id T1) (*T2, error) {
	var entity T2
	err := repo.db.Session(&gorm.Session{}).First(&entity, id).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *BaseRepository[T1, T2]) GetAll() ([]*T2, error) {
	var entities []*T2
	err := repo.db.Session(&gorm.Session{}).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (repo *BaseRepository[T1, T2]) GetByFilter(filter interface{}, args ...interface{}) (*T2, error) {
	var entity T2
	err := repo.db.Session(&gorm.Session{}).Where(filter, args...).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return &entity, nil
}

func (repo *BaseRepository[T1, T2]) GetAllByFilter(filter interface{}, args ...interface{}) ([]*T2, error) {
	var entities []*T2
	err := repo.db.Session(&gorm.Session{}).Where(filter, args...).Find(&entities).Error
	if err != nil {
		return nil, err
	}
	return entities, nil
}

func (repo *BaseRepository[T1, T2]) Create(entity *T2) (*T2, error) {
	if err := repo.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (repo *BaseRepository[T1, T2]) Update(entity *T2) (*T2, error) {
	if err := repo.db.Session(&gorm.Session{}).Save(&entity).Error; err != nil {
		return nil, err
	}

	return entity, nil
}

func (repo *BaseRepository[T1, T2]) Delete(id T1) error {
	if err := repo.db.Session(&gorm.Session{}).Delete(new(T2), id).Error; err != nil {
		return err
	}
	return nil
}

func (repo *BaseRepository[T1, T2]) Count(query interface{}, args ...interface{}) (int64, error) {
	var count int64
	err := repo.db.Session(&gorm.Session{}).Where(query, args...).Count(&count).Error
	if err != nil {
		return -1, err
	}

	return count, nil
}

func (repo *BaseRepository[T1, T2]) Any(query interface{}, args ...interface{}) error {
	count, err := repo.Count(query, args...)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("no record not found")
	}
	return nil
}

func (repo *BaseRepository[T1, T2]) GetQueryable() *gorm.DB {
	return repo.db
}
