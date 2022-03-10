package repositories

import (
	"webapi/db"
	"webapi/logger"
	"webapi/models/domains"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IDemoRepository interface {
	Get(id int) (*domains.Demo, error)
	GetByFilter(filter interface{}, args ...interface{}) (*domains.Demo, error)
	GetAll() ([]*domains.Demo, error)
	GetAllByFilter(filter interface{}, args ...interface{}) ([]*domains.Demo, error)
	Create(demo *domains.Demo) (*domains.Demo, error)
	Update(demo *domains.Demo) (*domains.Demo, error)
}

type DemoRepository struct {
	*BaseRepository
}

func NewDemoRepository(db *db.DB) IDemoRepository {
	return &DemoRepository{
		&BaseRepository{
			db: db.Table(domains.DemoTableName),
		},
	}
}

func (repo *DemoRepository) Get(id int) (*domains.Demo, error) {
	var demo domains.Demo
	err := repo.db.Session(&gorm.Session{}).Where("id=?", id).First(&demo).Error
	if err != nil {
		logger.Log.Error("demo repository failed to get demo by id", zap.Int("id", id), zap.String("err", err.Error()))
		return nil, err
	}
	return &demo, nil
}

func (repo *DemoRepository) GetByFilter(filter interface{}, args ...interface{}) (*domains.Demo, error) {
	var demo domains.Demo
	err := repo.db.Session(&gorm.Session{}).Where(filter, args...).First(&demo).Error
	if err != nil {
		logger.Log.Error("demo repository failed to filter a single demo", zap.Any("filter", filter), zap.Any("args", args), zap.String("err", err.Error()))
		return nil, err
	}
	return &demo, nil
}

func (repo *DemoRepository) GetAll() ([]*domains.Demo, error) {
	var demos []*domains.Demo
	err := repo.db.Session(&gorm.Session{}).Find(&demos).Error
	if err != nil {
		logger.Log.Error("demo repository failed to get all demos")
		return nil, err
	}

	return demos, nil
}

func (repo *DemoRepository) GetAllByFilter(filter interface{}, args ...interface{}) ([]*domains.Demo, error) {
	var demos []*domains.Demo
	err := repo.db.Session(&gorm.Session{}).Where(filter, args...).Find(&demos).Error
	if err != nil {
		logger.Log.Error("demo repository failed to filter demos", zap.Any("filter", filter), zap.Any("args", args), zap.String("err", err.Error()))
		return nil, err
	}
	return demos, nil
}

func (repo *DemoRepository) Create(demo *domains.Demo) (*domains.Demo, error) {
	demo.ID = 0
	if err := repo.db.Session(&gorm.Session{}).Create(&demo).Error; err != nil {
		logger.Log.Error("demo repository failed to create a demo", zap.Any("data", demo), zap.Any("err", err))
		return nil, err
	}
	logger.Log.Info("a new demo created", zap.Any("data", demo))
	return demo, nil
}

func (repo *DemoRepository) Update(demo *domains.Demo) (*domains.Demo, error) {
	if err := repo.db.Session(&gorm.Session{}).Save(demo).Error; err != nil {
		logger.Log.Error("demo repository failed to update a demo", zap.Any("data", demo), zap.Any("err", err))
		return nil, err
	}
	return demo, nil
}
