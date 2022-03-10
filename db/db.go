package db

import (
	"fmt"
	"log"
	"sync"
	"webapi/configs"
	"webapi/logger"
	"webapi/models/domains"

	"go.uber.org/zap"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

//DB database structure
type DB struct {
	*gorm.DB
}

var dbInstance *DB

func connectDB(config *configs.DBConfig) error {
	gormCfg := &gorm.Config{
		Logger:                   logger.GormLogger(),
		DisableNestedTransaction: true,
		SkipDefaultTransaction:   true,
	}

	connString := fmt.Sprintf("server=%s; port=%d; user id=%s; password=%s; database=%s;", config.Server, config.Port, config.User, config.Password, config.DbName)
	conn, err := gorm.Open(sqlserver.Open(connString), gormCfg)
	if err != nil {
		return err
	}
	dbInstance = &DB{conn}
	return nil
}

func OpenConnection(cfg *configs.AppConfig) *DB {
	var connDBOnce sync.Once
	connDBOnce.Do(func() {
		err := connectDB(&cfg.DBConfig)
		if err != nil {
			logger.Log.Error("failed to connect database", zap.Any("err", err))
			panic(err)
		}
	})
	log.Println("Database connected")
	return dbInstance
}

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
