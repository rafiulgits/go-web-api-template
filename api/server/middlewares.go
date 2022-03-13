package server

import (
	"flag"
	"webapi/configs"
	"webapi/db"
	"webapi/logger"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/rafiulgits/gotnet"
)

// middlewares
func ConfigureMiddlewares(app *gotnet.App) {
	preload(app.Service)
	parseConsoleArgs(app.Service)
	configureLogger(app)
}

// migrations and background services
func BeforeAppRun(app *gotnet.App) {
	dbMigration(app.Service.Container())
}

func preload(service *gotnet.Service) {
	err := service.Container().Invoke(func(
		cfg *configs.AppConfig, validator *validator.Validate) {
	})
	if err != nil {
		panic(err)
	}
}

func parseConsoleArgs(service *gotnet.Service) {
	dbMigration := flag.Bool("dbmigration", false, "Run with database migration")
	port := flag.Int("port", 8080, "Application port")
	debug := flag.Bool("debug", false, "Run app in debug mode")
	flag.Parse()
	appCfg := configs.GetAppConfig()
	if dbMigration != nil {
		appCfg.DBMigration = *dbMigration
	}
	if port != nil {
		appCfg.ListenPort = *port
	}
	if debug != nil {
		appCfg.DebugEnv = *debug
	}
}

func configureLogger(app *gotnet.App) {
	cfg := configs.GetAppConfig()
	logger.NewLogger(&cfg.LogConfig)
	if cfg.DebugEnv {
		app.Router.Use(middleware.Logger)
	} else {
		app.Router.Use(logger.ZapFileLogging(logger.Log))
	}
}

func dbMigration(container *gotnet.Container) {
	container.Invoke(func(database *db.DB, cfg *configs.AppConfig) {
		if cfg.DBMigration {
			database.Migration()
		}
	})
}
