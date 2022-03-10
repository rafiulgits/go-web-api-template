package logger

import (
	"log"
	"os"
	"time"

	gormlog "gorm.io/gorm/logger"
)

func GormLogger() gormlog.Interface {
	return gormlog.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		gormlog.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  gormlog.Info, // Log level
			IgnoreRecordNotFoundError: true,         // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)
}
